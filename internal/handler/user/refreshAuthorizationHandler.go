package user

import (
	"net/http"

	"cloud-disk/internal/logic/user"
	"cloud-disk/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RefreshAuthorizationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewRefreshAuthorizationLogic(r.Context(), svcCtx)
		resp, err := l.RefreshAuthorization(r.Header.Get("Authorization"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
