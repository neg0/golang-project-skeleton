package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	sut, err := newConfig().fromFile(nil)
	if err != nil {
		t.Error(err)
	}

	t.Run("should have host address", func(t *testing.T) {
		if len(sut.Api.Host) <  1 {
			t.Log(sut.Api.Host)
			t.Log(len(sut.Api.Host))
			t.Fail()
		}
	})

	t.Run("should have a port number", func(t *testing.T) {
		if sut.Api.Port < 1 {
			t.Log(sut.Api.Port)
			t.Fail()
		}
	})
}
