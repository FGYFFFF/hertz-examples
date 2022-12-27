package v7

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
)

type HelloMethod7Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod7Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod7Service {
	return &HelloMethod7Service{RequestContext: RequestContext, Context: Context}
}
func (h *HelloMethod7Service) Run(req *example.HelloReq) (resp *example.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
