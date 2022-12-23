package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
)

type HelloMethod5Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod5Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod5Service {
	return &HelloMethod5Service{RequestContext: RequestContext, Context: Context}
}
func (h *HelloMethod5Service) Run(req *example.HelloReq) (resp *example.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
