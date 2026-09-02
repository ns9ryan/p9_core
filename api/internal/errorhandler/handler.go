package errorhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/ns9ryan/p9_core/api/internal/rpcerror"
	"github.com/ns9ryan/p9_core/api/internal/types"
	"github.com/ns9ryan/p9_core/pkg/i18n"
	"github.com/ns9ryan/p9_core/pkg/i18nkey"
	"github.com/ns9ryan/p9_core/pkg/validate"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler 错误处理器
type Handler struct {
	trans *i18n.Translator // 翻译器
	debug bool             // 是否开启调试
}

// New 创建错误处理器
func New(trans *i18n.Translator, debug bool) *Handler {
	return &Handler{
		trans: trans,
		debug: debug,
	}
}

// Handle 处理API错误
func (h *Handler) Handle(ctx context.Context, err error) (int, any) {
	// 处理参数校验错误
	var validationError *validate.Error
	if errors.As(err, &validationError) {
		return h.response(
			http.StatusBadRequest,
			validationError.Error(),
			h.debugDetail("api", "", err),
		)
	}

	// 获取RPC调用来源
	source := "api"
	rpcMethod := ""

	// 判断错误是否来自RPC调用，并记录具体RPC方法
	if rpcErr, ok := rpcerror.FromError(err); ok {
		source = "rpc"
		rpcMethod = rpcErr.Method
	}

	// 解析gRPC错误并转换为统一HTTP错误响应
	if grpcStatus, ok := status.FromError(err); ok {
		return h.handleGRPC(ctx, err, grpcStatus, source, rpcMethod)
	}

	// 未分类错误
	h.logInternal(ctx, source, rpcMethod, err)

	return h.response(
		http.StatusInternalServerError,
		h.trans.Trans(ctx, i18nkey.InternalError),
		h.debugDetail(source, rpcMethod, err),
	)
}

// handleGRPC 处理gRPC错误
func (h *Handler) handleGRPC(
	ctx context.Context,
	err error,
	grpcStatus *status.Status,
	source string,
	rpcMethod string,
) (int, any) {
	// 将gRPC状态码转换为对应的HTTP状态码
	httpStatus := grpcCodeToHTTPStatus(grpcStatus.Code())

	// 根据gRPC状态获取最终返回给前端的提示信息
	message := h.grpcMessage(ctx, grpcStatus)

	// HTTP 500及以上属于服务端错误，需要记录内部错误日志
	if httpStatus >= http.StatusInternalServerError {
		h.logInternal(ctx, source, rpcMethod, err)
	}

	// 创建统一错误响应
	return h.response(
		httpStatus,
		message,
		h.debugDetail(source, rpcMethod, err),
	)
}

// grpcMessage 获取gRPC错误提示
func (h *Handler) grpcMessage(ctx context.Context, grpcStatus *status.Status) string {
	switch grpcStatus.Code() {
	// 资源耗尽，例如限流或配额不足
	case codes.ResourceExhausted:
		return h.trans.Trans(ctx, i18nkey.TooManyRequests)

	// 服务不可用，例如RPC服务未启动或连接失败
	case codes.Unavailable:
		return h.trans.Trans(ctx, i18nkey.ServiceUnavailable)

	// 请求超时
	case codes.DeadlineExceeded:
		return h.trans.Trans(ctx, i18nkey.RequestTimeout)

	// 请求被取消
	case codes.Canceled:
		return h.trans.Trans(ctx, i18nkey.RequestTimeout)

	// 未明确分类的gRPC错误
	case codes.Unknown:
		return h.trans.Trans(ctx, i18nkey.InternalError)
	}

	// 将RPC返回的消息Key翻译成当前请求语言
	message := h.trans.Trans(ctx, grpcStatus.Message())

	// Internal代表服务内部错误，翻译不到时不直接暴露原始错误信息
	if grpcStatus.Code() == codes.Internal && message == grpcStatus.Message() {
		return h.trans.Trans(ctx, i18nkey.InternalError)
	}

	return message
}

// debugDetail 创建调试信息
func (h *Handler) debugDetail(source string, rpcMethod string, err error) *types.ErrorDetail {
	if !h.debug {
		return nil
	}

	return &types.ErrorDetail{
		Source: source,
		Rpc:    rpcMethod,
		Error:  err.Error(),
	}
}

// logInternal 记录内部错误
func (h *Handler) logInternal(
	ctx context.Context,
	source string,
	rpcMethod string,
	err error,
) {
	logger := logx.WithContext(ctx)

	if rpcMethod != "" {
		logger.Errorw(
			"RPC调用失败",
			logx.Field("rpc", rpcMethod),
			logx.Field("error", err.Error()),
		)
		return
	}

	logger.Errorw(
		"接口处理失败",
		logx.Field("source", source),
		logx.Field("error", err.Error()),
	)
}

// response 创建错误响应
func (h *Handler) response(code int, message string, detail *types.ErrorDetail) (int, any) {
	return code, &types.BaseResponse{
		Code:   code,
		Msg:    message,
		Detail: detail,
	}
}

// grpcCodeToHTTPStatus 将gRPC状态码转换为HTTP状态码
func grpcCodeToHTTPStatus(code codes.Code) int {
	switch code {
	// 参数错误、资源已存在、前置条件不满足、参数超出范围
	case codes.InvalidArgument,
		codes.AlreadyExists,
		codes.FailedPrecondition,
		codes.OutOfRange:
		return http.StatusBadRequest

	// 未认证，例如Token无效或未登录
	case codes.Unauthenticated:
		return http.StatusUnauthorized

	// 已认证但没有操作权限
	case codes.PermissionDenied:
		return http.StatusForbidden

	// 请求的资源不存在
	case codes.NotFound:
		return http.StatusNotFound

	// 请求过多或资源配额耗尽
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests

	// 请求被客户端或上游主动取消
	case codes.Canceled:
		return http.StatusRequestTimeout

	// 操作冲突，例如并发事务被中止
	case codes.Aborted:
		return http.StatusConflict

	// RPC方法未实现
	case codes.Unimplemented:
		return http.StatusNotImplemented

	// RPC服务当前不可用
	case codes.Unavailable:
		return http.StatusServiceUnavailable

	// RPC调用超过截止时间
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout

	// 未明确映射的错误统一按服务器内部错误处理
	default:
		return http.StatusInternalServerError
	}
}
