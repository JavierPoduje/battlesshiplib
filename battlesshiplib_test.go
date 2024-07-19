package battlesshiplib

import (
	"testing"
)

func TestBattlesshiplib_Connect(t *testing.T) {
	r := NewRedis()
	r.Ping()
	r.Set("yo", "meo")
	value := r.Get("yo")
	if value != "meo" {
		t.Errorf("Expected value: meo, got: %v", value)
	}

	defer r.rdb.Close()
}
