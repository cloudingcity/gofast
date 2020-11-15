package gofast

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestDefaultErrorHandler(t *testing.T) {
	resp := fasthttp.AcquireResponse()
	resp.SetStatusCode(400)
	resp.SetBodyString("hi there")

	want := "code: 400, body: hi there"
	got := defaultErrorHandler(resp).Error()
	assert.Equal(t, want, got)
}
