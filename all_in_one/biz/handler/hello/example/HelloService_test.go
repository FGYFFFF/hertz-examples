package example

import (
	"bytes"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"testing"
)

func TestHelloMethod(t *testing.T) {
	h := server.Default()
	h.GET("/hello", HelloMethod)
	w := ut.PerformRequest(h.Engine, "GET", "/hello", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestHelloMethod2(t *testing.T) {
	h := server.Default()
	h.GET("/hello2", HelloMethod2)
	w := ut.PerformRequest(h.Engine, "GET", "/hello2", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestHelloMethod3(t *testing.T) {
	h := server.Default()
	h.GET("/hello3", HelloMethod3)
	w := ut.PerformRequest(h.Engine, "GET", "/hello3", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestHelloMethod4(t *testing.T) {
	h := server.Default()
	h.GET("/hello4", HelloMethod4)
	w := ut.PerformRequest(h.Engine, "GET", "/hello4", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestHelloMethod5(t *testing.T) {
	h := server.Default()
	h.GET("/hello5", HelloMethod5)
	w := ut.PerformRequest(h.Engine, "GET", "/hello5", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
