package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
)

type HelloMethod12Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod12Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod12Service {
	return &HelloMethod12Service{RequestContext: RequestContext, Context: Context}
}
func (h *HelloMethod12Service) Run(req *example.HelloReq) (resp *example.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
