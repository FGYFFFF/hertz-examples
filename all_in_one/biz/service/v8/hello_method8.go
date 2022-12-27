package v8

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
)

type HelloMethod8Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod8Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod8Service {
	return &HelloMethod8Service{RequestContext: RequestContext, Context: Context}
}
func (h *HelloMethod8Service) Run(req *example.HelloReq) (resp *example.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
