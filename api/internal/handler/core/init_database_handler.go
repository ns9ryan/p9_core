// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package core

import (
	"net/http"

	"github.com/ns9ryan/p9_core/api/internal/logic/core"
	"github.com/ns9ryan/p9_core/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InitDatabaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := core.NewInitDatabaseLogic(r.Context(), svcCtx)
		resp, err := l.InitDatabase()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
