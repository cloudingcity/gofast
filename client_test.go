package gofast

import (
	"io/ioutil"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

const testURL = "http://example.com/"

func TestClient_Get(t *testing.T) {
	t.Run("200", func(t *testing.T) {
		c := New()
		c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, testURL, ctx.Request.URI().String())
			ctx.SetBodyString(`{"foo":"bar"}`)
		})

		var out struct{ Foo string }
		err := c.Get(testURL, &out, nil)
		assert.NoError(t, err)
		assert.Equal(t, "bar", out.Foo)
	})

	t.Run("error handle when status code not 2xx", func(t *testing.T) {
		c := New()
		c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBodyString("something wrong")
		})

		err := c.Get(testURL, nil, nil)
		assert.Error(t, err)
		assert.Equal(t, "code: 500, body: something wrong", err.Error())
	})

	t.Run("get with header", func(t *testing.T) {
		c := New()
		c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {
			assert.Equal(t, "bar", string(ctx.Request.Header.Peek("foo")))
		})

		err := c.Get(testURL, nil, Header{"foo": "bar"})
		assert.NoError(t, err)
	})
}

func TestClient_Post(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := New()
		c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {
			assert.JSONEq(t, `{"foo":"bar"}`, string(ctx.Request.Body()))
		})

		in := Body{"foo": "bar"}
		err := c.Post(testURL, &in, nil, nil)
		assert.NoError(t, err)
	})

	t.Run("request encode fail", func(t *testing.T) {
		c := New()
		c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {})

		in := make(chan struct{})
		err := c.Post(testURL, in, nil, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "encode request:")
	})

	t.Run("response decode fail", func(t *testing.T) {
		log.SetOutput(ioutil.Discard)

		c := New()
		c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {
			ctx.SetBodyString("wrong format")
		})

		var out struct{ Foo string }
		err := c.Post(testURL, nil, &out, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "decode response:")
	})
}

func TestClient_Put(t *testing.T) {
	c := New()
	c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {})

	err := c.Put(testURL, nil, nil, nil)
	assert.NoError(t, err)
}

func TestClient_Patch(t *testing.T) {
	c := New()
	c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {})
	err := c.Patch(testURL, nil, nil, nil)
	assert.NoError(t, err)
}

func TestClient_Delete(t *testing.T) {
	c := New()
	c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {})
	err := c.Delete(testURL, nil, nil, nil)
	assert.NoError(t, err)
}

func BenchmarkPostJSON(b *testing.B) {
	c := New()
	c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {
		ctx.SetBodyString(`{"hello": "world"}`)
	})

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		in := struct{ Foo string }{
			Foo: "bar",
		}
		for pb.Next() {
			err := c.Post(testURL, &in, nil, Header{"foo": "bar"})
			if err != nil {
				b.Fatalf("unexpected error: %s", err)
			}
		}
	})
}

func BenchmarkPostURLEncode(b *testing.B) {
	c := New(Config{
		RequestEncoder:  URLEncoder,
		ResponseDecoder: TextDecoder,
	})
	c.fastClient = mockFastHTTPClient(func(ctx *fasthttp.RequestCtx) {
		ctx.SetBodyString(`{"hello": "world"}`)
	})

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		in := Body{
			"foo": "bar",
		}
		for pb.Next() {
			err := c.Post(testURL, in, nil, Header{"foo": "bar"})
			if err != nil {
				b.Fatalf("unexpected error: %s", err)
			}
		}
	})
}

func mockFastHTTPClient(handler fasthttp.RequestHandler) *fasthttp.Client {
	ln := fasthttputil.NewInmemoryListener()
	srv := &fasthttp.Server{
		Handler: handler,
	}
	go srv.Serve(ln) //nolint:errcheck

	return &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}
}
