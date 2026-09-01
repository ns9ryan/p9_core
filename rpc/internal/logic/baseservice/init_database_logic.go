package baseservicelogic

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/bsm/redislock"
	"github.com/ns9ryan/common/enum/common"
	"github.com/ns9ryan/common/i18n"
	"github.com/ns9ryan/common/password"
	"github.com/ns9ryan/common/rpcerror"
	"github.com/ns9ryan/p9_core/rpc/ent"
	"github.com/ns9ryan/p9_core/rpc/ent/role"
	"github.com/ns9ryan/p9_core/rpc/internal/svc"
	"github.com/ns9ryan/p9_core/rpc/pb/core/base"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	initDatabaseLockKey  = "INIT:DATABASE:LOCK"
	initDatabaseErrorKey = "INIT:DATABASE:ERROR"
	initDatabaseStateKey = "INIT:DATABASE:STATE"

	initDatabaseStateRunning = "0"
	initDatabaseStateSuccess = "1"

	initDatabaseLockTTL  = 10 * time.Minute
	initDatabaseErrorTTL = 5 * time.Minute
	initDatabaseStateTTL = 24 * time.Hour
)

// InitDatabaseLogic 数据库初始化逻辑
type InitDatabaseLogic struct {
	ctx    context.Context     // 请求上下文
	svcCtx *svc.ServiceContext // 服务上下文
	logx.Logger
}

// NewInitDatabaseLogic 创建数据库初始化逻辑
func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// handleInitError 统一处理数据库初始化错误
func (l *InitDatabaseLogic) handleInitError(err error) (*base.InitDatabaseResponse, error) {
	logx.Errorw("数据库初始化失败", logx.Field("detail", err))

	_ = l.svcCtx.Redis.Set(
		l.ctx,
		initDatabaseErrorKey,
		err.Error(),
		initDatabaseErrorTTL,
	).Err()

	return nil, rpcerror.NewInternal(err.Error())
}

// InitDatabase 初始化数据库表结构及基础数据
func (l *InitDatabaseLogic) InitDatabase(in *base.InitDatabaseRequest) (*base.InitDatabaseResponse, error) {
	// 初始化任务可能耗时较长，移除请求取消和超时限制
	l.ctx = context.Background()

	// 获取初始化锁，避免重复执行
	locker := redislock.New(l.svcCtx.Redis)
	lock, err := locker.Obtain(
		l.ctx,
		initDatabaseLockKey,
		initDatabaseLockTTL,
		nil,
	)
	if errors.Is(err, redislock.ErrNotObtained) {
		logx.Error("数据库初始化任务正在执行")
		return nil, rpcerror.NewInternal("数据库初始化任务正在执行")
	}

	if err != nil {
		logx.Errorw("获取 Redis 锁失败", logx.Field("detail", err.Error()))
		return nil, rpcerror.NewInternal("获取 Redis 锁失败")
	}

	// 初始化完成后释放锁
	defer func() {
		if err := lock.Release(l.ctx); err != nil {
			logx.Errorw("释放 Redis 初始化锁失败", logx.Field("detail", err))
		}
	}()

	// 初始化数据库表结构
	if err := l.svcCtx.DB.Schema.Create(
		l.ctx,
		schema.WithForeignKeys(false),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	); err != nil {
		return l.handleInitError(err)
	}

	// 检查数据库是否已经初始化
	apiCount, err := l.svcCtx.DB.API.Query().Count(l.ctx)
	if err != nil {
		return l.handleInitError(err)
	}

	if apiCount > 0 {
		// 已初始化时重新同步超级管理员 Casbin 策略
		if err := l.insertCasbinPoliciesData(); err != nil {
			return l.handleInitError(err)
		}

		if err := l.svcCtx.Redis.Set(l.ctx, initDatabaseStateKey, initDatabaseStateSuccess, initDatabaseStateTTL).Err(); err != nil {
			logx.Errorw("更新数据库初始化状态失败", logx.Field("detail", err))
			return nil, rpcerror.NewInternal(i18n.RedisError)
		}

		return &base.InitDatabaseResponse{
			Msg: i18n.AlreadyInit,
		}, nil
	}

	// 初始化 Redis 状态
	_ = l.svcCtx.Redis.Set(l.ctx, initDatabaseErrorKey, "", initDatabaseErrorTTL).Err()
	_ = l.svcCtx.Redis.Set(l.ctx, initDatabaseStateKey, initDatabaseStateRunning, initDatabaseErrorTTL).Err()

	// 初始化角色
	if err := l.insertRoleData(); err != nil {
		return l.handleInitError(err)
	}

	// 初始化部门
	if err := l.insertDepartmentData(); err != nil {
		return l.handleInitError(err)
	}

	// 初始化职位
	if err := l.insertPositionData(); err != nil {
		return l.handleInitError(err)
	}

	// 初始化用户
	if err := l.insertUserData(); err != nil {
		return l.handleInitError(err)
	}

	// 初始化菜单
	if err := l.insertMenuData(); err != nil {
		return l.handleInitError(err)
	}

	// 初始化 API
	if err := l.insertApiData(); err != nil {
		return l.handleInitError(err)
	}

	// 初始化超级管理员菜单权限
	if err := l.insertRoleMenuAuthorityData(); err != nil {
		return l.handleInitError(err)
	}

	// 初始化第三方登录配置
	if err := l.insertProviderData(); err != nil {
		return l.handleInitError(err)
	}

	// 初始化 Casbin 策略
	if err := l.insertCasbinPoliciesData(); err != nil {
		return l.handleInitError(err)
	}

	// 标记数据库初始化完成
	if err := l.svcCtx.Redis.Set(l.ctx, initDatabaseStateKey, initDatabaseStateSuccess, initDatabaseStateTTL).Err(); err != nil {
		logx.Errorw("更新数据库初始化状态失败", logx.Field("detail", err))
		return nil, rpcerror.NewInternal(i18n.RedisError)
	}

	return &base.InitDatabaseResponse{}, nil
}

// insertRoleData 插入初始角色数据
func (l *InitDatabaseLogic) insertRoleData() error {
	roles := []*ent.RoleCreate{
		l.svcCtx.DB.Role.Create().
			SetName("role.admin").
			SetCode("001").
			SetRemark("超级管理员").
			SetSort(1),

		l.svcCtx.DB.Role.Create().
			SetName("role.stuff").
			SetCode("002").
			SetRemark("普通员工").
			SetSort(2),
	}

	return l.svcCtx.DB.Role.CreateBulk(roles...).Exec(l.ctx)
}

// insertDepartmentData 插入初始部门数据
func (l *InitDatabaseLogic) insertDepartmentData() error {
	departments := []*ent.DepartmentCreate{
		l.svcCtx.DB.Department.Create().
			SetName("department.managementDepartment").
			SetAncestors("").
			SetLeader("admin").
			SetEmail("simpleadmin@gmail.com").
			SetPhone("18888888888").
			SetRemark("Super Administrator").
			SetSort(1).
			SetParentID(common.DefaultParentID),
	}

	return l.svcCtx.DB.Department.CreateBulk(departments...).Exec(l.ctx)
}

// insertPositionData 插入初始职位数据
func (l *InitDatabaseLogic) insertPositionData() error {
	positions := []*ent.PositionCreate{
		l.svcCtx.DB.Position.Create().
			SetName("position.ceo").
			SetCode("001").
			SetRemark("CEO").
			SetSort(1),
	}

	return l.svcCtx.DB.Position.CreateBulk(positions...).Exec(l.ctx)
}

// insertUserData 插入初始用户数据
func (l *InitDatabaseLogic) insertUserData() error {
	passwordHash, err := password.Hash("simple-admin")
	if err != nil {
		return err
	}

	users := []*ent.UserCreate{
		l.svcCtx.DB.User.Create().
			SetUsername("admin").
			SetNickname("admin").
			SetPassword(passwordHash).
			SetEmail("simple_admin@gmail.com").
			AddRoleIDs(1).
			SetDepartmentID(1).
			AddPositionIDs(1),
	}

	return l.svcCtx.DB.User.CreateBulk(users...).Exec(l.ctx)
}

// insertRoleMenuAuthorityData 插入超级管理员菜单权限
func (l *InitDatabaseLogic) insertRoleMenuAuthorityData() error {
	menus, err := l.svcCtx.DB.Menu.Query().All(l.ctx)
	if err != nil {
		return err
	}

	menuIDs := make([]uint64, 0, len(menus))
	for _, menu := range menus {
		menuIDs = append(menuIDs, menu.ID)
	}

	if len(menuIDs) == 0 {
		return nil
	}

	// 仅为超级管理员分配全部菜单
	_, err = l.svcCtx.DB.Role.Update().
		Where(role.CodeEQ("001")).
		AddMenuIDs(menuIDs...).
		Save(l.ctx)

	return err
}

// insertProviderData 插入初始第三方登录配置
func (l *InitDatabaseLogic) insertProviderData() error {
	providers := []*ent.OauthProviderCreate{
		l.svcCtx.DB.OauthProvider.Create().
			SetName("google").
			SetClientID("your client id").
			SetClientSecret("your client secret").
			SetRedirectURL("http://localhost:3100/oauth/login/callback").
			SetScopes("email openid").
			SetAuthURL("https://accounts.google.com/o/oauth2/auth").
			SetTokenURL("https://oauth2.googleapis.com/token").
			SetAuthStyle(1).
			SetInfoURL("https://www.googleapis.com/oauth2/v2/userinfo?access_token=TOKEN"),

		l.svcCtx.DB.OauthProvider.Create().
			SetName("github").
			SetClientID("your client id").
			SetClientSecret("your client secret").
			SetRedirectURL("http://localhost:3100/oauth/login/callback").
			SetScopes("email openid").
			SetAuthURL("https://github.com/login/oauth/authorize").
			SetTokenURL("https://github.com/login/oauth/access_token").
			SetAuthStyle(2).
			SetInfoURL("https://api.github.com/user"),
	}

	return l.svcCtx.DB.OauthProvider.CreateBulk(providers...).Exec(l.ctx)
}

// insertCasbinPoliciesData 插入超级管理员 Casbin 策略
func (l *InitDatabaseLogic) insertCasbinPoliciesData() error {
	// 查询全部 API
	apis, err := l.svcCtx.DB.API.Query().All(l.ctx)
	if err != nil {
		return err
	}

	// 没有 API 时无需同步策略
	if len(apis) == 0 {
		return nil
	}

	// 获取超级管理员角色编码
	roleCode := "001"

	adminRole, err := l.svcCtx.DB.Role.Query().Where(role.NameEQ("role.admin")).First(l.ctx)
	if err == nil {
		roleCode = adminRole.Code
	}

	// 构建超级管理员权限策略
	policies := make([][]string, 0, len(apis))
	for _, api := range apis {
		policies = append(policies, []string{
			roleCode,
			api.Path,
			api.Method,
		})
	}

	// 创建 Casbin 权限执行器
	enforcer, err := l.svcCtx.Config.CasbinConf.NewCasbin(
		l.svcCtx.Config.DatabaseConf.Type,
		l.svcCtx.Config.DatabaseConf.GetDSN(),
	)
	if err != nil {
		logx.Errorw("初始化 Casbin 权限执行器失败", logx.Field("detail", err))
		return err
	}

	// 查询超级管理员原有策略
	oldPolicies, err := enforcer.GetFilteredPolicy(0, roleCode)
	if err != nil {
		logx.Errorw("获取超级管理员 Casbin 策略失败", logx.Field("role_code", roleCode), logx.Field("detail", err))
		return err
	}

	// 删除超级管理员原有策略
	if len(oldPolicies) > 0 {
		removed, err := enforcer.RemoveFilteredPolicy(0, roleCode)
		if err != nil {
			logx.Errorw("删除超级管理员 Casbin 策略失败", logx.Field("role_code", roleCode), logx.Field("detail", err))
			return err
		}

		if !removed {
			return errors.New("删除超级管理员 Casbin 策略失败")
		}
	}

	// 添加最新超级管理员策略
	added, err := enforcer.AddPolicies(policies)
	if err != nil {
		return err
	}

	if !added {
		return errors.New("添加超级管理员 Casbin 策略失败")
	}

	return nil
}
