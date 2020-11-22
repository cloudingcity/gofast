package gofast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestEncoding_JSON(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		req := fasthttp.AcquireRequest()
		in := struct {
			Foo string `json:"foo"`
		}{
			Foo: "bar",
		}
		err := JSONEncoder(req, &in)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"foo": "bar"}`, string(req.Body()))
	})

	t.Run("decode", func(t *testing.T) {
		var out struct {
			Foo string `json:"foo"`
		}
		resp := fasthttp.AcquireResponse()
		resp.SetBodyString(`{"foo": "bar"}`)
		err := JSONDecoder(resp, &out)
		assert.NoError(t, err)
		assert.Equal(t, "bar", out.Foo)
	})
}

func TestEncoding_URL(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		req := fasthttp.AcquireRequest()
		in := Body{
			"foo": "bar",
		}
		err := URLEncoder(req, in)
		assert.NoError(t, err)
		assert.Equal(t, "foo=bar", string(req.Body()))
	})
}

func TestEncoding_Text(t *testing.T) {
	t.Run("decode", func(t *testing.T) {
		var out string
		resp := fasthttp.AcquireResponse()
		resp.SetBodyString("hello string")
		err := TextDecoder(resp, &out)
		assert.NoError(t, err)
		assert.Equal(t, "hello string", out)
	})
}
