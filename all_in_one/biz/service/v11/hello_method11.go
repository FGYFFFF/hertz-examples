package v11

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
)

type HelloMethod11Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod11Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod11Service {
	return &HelloMethod11Service{RequestContext: RequestContext, Context: Context}
}
func (h *HelloMethod11Service) Run(req *example.HelloReq) (resp *example.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
