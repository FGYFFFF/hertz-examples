package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
)

type HelloMethod4Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod4Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod4Service {
	return &HelloMethod4Service{RequestContext: RequestContext, Context: Context}
}
func (h *HelloMethod4Service) Run(req *example.HelloReq) (resp *example.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
