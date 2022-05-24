package user

import (
	"net/http"

	"cloud-disk/internal/logic/user"
	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileMovedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileMovedRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserFileMovedLogic(r.Context(), svcCtx)
		resp, err := l.UserFileMoved(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
