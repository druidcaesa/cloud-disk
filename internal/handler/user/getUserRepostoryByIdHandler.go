package user

import (
	"net/http"
	"strconv"

	"cloud-disk/internal/logic/user"
	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserRepostoryByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserRepostoryByIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetUserRepostoryByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUserRepostoryById(&req, r.Header.Get("UserIdentity"))
		atoi, _ := strconv.Atoi(r.URL.Query().Get("id"))
		req.Id = atoi
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
