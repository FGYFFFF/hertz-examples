package v6

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
)

type HelloMethod6Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod6Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod6Service {
	return &HelloMethod6Service{RequestContext: RequestContext, Context: Context}
}
func (h *HelloMethod6Service) Run(req *example.HelloReq) (resp *example.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
