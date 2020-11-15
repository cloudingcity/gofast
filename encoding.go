package gofast

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type RequestEncoder func(req *fasthttp.Request, in interface{}) error

type ResponseDecoder func(resp *fasthttp.Response, out interface{}) error

var jsonRequestEncoder = func(req *fasthttp.Request, in interface{}) error {
	req.Header.SetContentType("application/json")
	return json.NewEncoder(req.BodyWriter()).Encode(in)
}

var jsonResponseDecoder = func(resp *fasthttp.Response, out interface{}) error {
	return json.Unmarshal(resp.Body(), out)
}
