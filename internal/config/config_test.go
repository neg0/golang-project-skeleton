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
		if len(sut.API.Host) < 1 {
			t.Log(sut.API.Host)
			t.Log(len(sut.API.Host))
			t.Fail()
		}
	})

	t.Run("should have a port number", func(t *testing.T) {
		if sut.API.Port < 1 {
			t.Log(sut.API.Port)
			t.Fail()
		}
	})
}
