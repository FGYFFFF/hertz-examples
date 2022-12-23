// this is my custom handler.
package example

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
	"hertz-examples/all_in_one/biz/service"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethodService(ctx, c).Run(&req)
	c.JSON(200, resp)
}

// HelloMethod2 .
// @router /hello2 [GET]
func HelloMethod2(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethod2Service(ctx, c).Run(&req)
	c.JSON(200, resp)
}

// HelloMethod3 .
// @router /hello3 [GET]
func HelloMethod3(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethod3Service(ctx, c).Run(&req)
	c.JSON(200, resp)
}

// HelloMethod4 .
// @router /hello4 [GET]
func HelloMethod4(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethod4Service(ctx, c).Run(&req)
	c.JSON(200, resp)
}

// HelloMethod5 .
// @router /hello5 [GET]
func HelloMethod5(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethod5Service(ctx, c).Run(&req)
	c.JSON(200, resp)
}
