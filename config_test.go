package gofast

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("set empty config get default", func(t *testing.T) {
		cfg := configDefault(Config{})

		assert.Equal(t, ConfigDefault.Name, cfg.Name)
		assert.Equal(t, ConfigDefault.NoDefaultUserAgentHeader, cfg.NoDefaultUserAgentHeader)
		assert.Equal(t, ConfigDefault.ReadTimeout, cfg.ReadTimeout)
		assert.Equal(t, ConfigDefault.WriteTimeout, cfg.WriteTimeout)

		want := runtime.FuncForPC(reflect.ValueOf(ConfigDefault.RequestEncoder).Pointer()).Name()
		got := runtime.FuncForPC(reflect.ValueOf(cfg.RequestEncoder).Pointer()).Name()
		assert.Equal(t, want, got)

		want2 := runtime.FuncForPC(reflect.ValueOf(ConfigDefault.ResponseDecoder).Pointer()).Name()
		got2 := runtime.FuncForPC(reflect.ValueOf(cfg.ResponseDecoder).Pointer()).Name()
		assert.Equal(t, want2, got2)

		want3 := runtime.FuncForPC(reflect.ValueOf(ConfigDefault.ErrorHandler).Pointer()).Name()
		got3 := runtime.FuncForPC(reflect.ValueOf(cfg.ErrorHandler).Pointer()).Name()
		assert.Equal(t, want3, got3)
	})

	t.Run("customize config", func(t *testing.T) {
		cfg := configDefault(Config{Name: "foo"})

		assert.Equal(t, "foo", cfg.Name)
	})
}
