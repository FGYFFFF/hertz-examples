// this is my custom handler.
package example

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	example "hertz-examples/all_in_one/biz/model/hello/example"
	"hertz-examples/all_in_one/biz/service"
	"hertz-examples/all_in_one/biz/service/v10"
	"hertz-examples/all_in_one/biz/service/v11"
	"hertz-examples/all_in_one/biz/service/v5"
	"hertz-examples/all_in_one/biz/service/v7"
	"hertz-examples/all_in_one/biz/service/v8"
	"hertz-examples/all_in_one/biz/service/v9"

	"hertz-examples/all_in_one/biz/service/v6"
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

	resp := v5.NewHelloMethod5Service(ctx, c).Run(&req)

	c.JSON(200, resp)
}

// HelloMethod6 .
// @router /hello6 [GET]
func HelloMethod6(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := v6.NewHelloMethod6Service(ctx, c).Run(&req)

	c.JSON(200, resp)
}

// HelloMethod7 .
// @router /hello7 [GET]
func HelloMethod7(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := v7.NewHelloMethod7Service(ctx, c).Run(&req)

	c.JSON(200, resp)
}

// HelloMethod8 .
// @router /hello8 [GET]
func HelloMethod8(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := v8.NewHelloMethod8Service(ctx, c).Run(&req)

	c.JSON(200, resp)
}

// HelloMethod9 .
// @router /hello9 [GET]
func HelloMethod9(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := v9.NewHelloMethod9Service(ctx, c).Run(&req)

	c.JSON(200, resp)
}

// HelloMethod10 .
// @router /hello10 [GET]
func HelloMethod10(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := v10.NewHelloMethod10Service(ctx, c).Run(&req)

	c.JSON(200, resp)
}

// HelloMethod11 .
// @router /hello11 [GET]
func HelloMethod11(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := v11.NewHelloMethod11Service(ctx, c).Run(&req)

	c.JSON(200, resp)
}

// HelloMethod12 .
// @router /hello12 [GET]
func HelloMethod12(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethod12Service(ctx, c).Run(&req)

	c.JSON(200, resp)
}
