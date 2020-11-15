package gofast

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

type RequestEncoder func(req *fasthttp.Request, in interface{}) error

type ResponseDecoder func(resp *fasthttp.Response, out interface{}) error

var jsonRequestEncoder = func(req *fasthttp.Request, in interface{}) error {
	req.Header.SetContentType("application/json")
	return json.NewEncoder(req.BodyWriter()).Encode(in)
}

var jsonResponseDecoder = func(resp *fasthttp.Response, out interface{}) error {
	if err := json.Unmarshal(resp.Body(), out); err != nil {
		log.Printf("[gofast] response decode failed - code: %v, body: %v", resp.StatusCode(), string(resp.Body()))
		return err
	}
	return nil
}
