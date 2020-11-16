package gofast

import "time"

const (
	defaultUserAgent = "gofast"
	defaultTimeout   = 6 * time.Second
)

type Config struct {
	// fasthttp client configurations
	Name                     string
	NoDefaultUserAgentHeader bool
	ReadTimeout              time.Duration
	WriteTimeout             time.Duration

	// ErrorHandler handle the status code without 2xx
	ErrorHandler ErrorHandler

	// RequestEncoder encode request before send
	RequestEncoder RequestEncoder

	// ResponseDecoder decode response after send
	ResponseDecoder ResponseDecoder
}
