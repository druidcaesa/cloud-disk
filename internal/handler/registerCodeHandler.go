package handler

import (
	"net/http"

	"cloud-disk/internal/logic"
	"cloud-disk/internal/svc"
	"cloud-disk/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterCodeLogic(r.Context(), svcCtx)
		resp, err := l.RegisterCode(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
