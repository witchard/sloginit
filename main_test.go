package sloginit

import (
	"log/slog"
	"testing"
)

func TestOptionsEnv(t *testing.T) {

}

func TestOptionsFunctional(t *testing.T) {
	o := opts{}

	Console(&o)
	if !o.console {
		t.Error("Console() should set console")
	}

	JSON(&o)
	if o.console {
		t.Error("JSON() should unset console")
	}

	Level(slog.LevelError)(&o)
	if o.level != slog.LevelError {
		t.Error("Level should set level correctly")
	}

	EnvPrefix("TEST")(&o)
	if o.prefix != "TEST" {
		t.Error("EnvPrefix should set prefix correctly")
	}

	o = setOpts(Console, Level(slog.LevelWarn), EnvPrefix("test"))
	if !o.console {
		t.Error("setOpts(Console) should set console")
	}
	if o.level != slog.LevelWarn {
		t.Error("setOpts(Level) should set level correctly")
	}
	if o.prefix != "test" {
		t.Error("setOpts(EnvPrefix) should set prefix correctly")
	}
}
