package handler

import (
	"net/http"

	"github.com/quarkcloudio/quark-go/v3/examples/zeroadmin/internal/logic"
	"github.com/quarkcloudio/quark-go/v3/examples/zeroadmin/internal/svc"
	"github.com/quarkcloudio/quark-go/v3/examples/zeroadmin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ZeroadminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewZeroadminLogic(r.Context(), svcCtx)
		resp, err := l.Zeroadmin(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
