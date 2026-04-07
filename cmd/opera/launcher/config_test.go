package launcher

import (
	"testing"
)

func TestGCModeFlagDefaultValue(t *testing.T) {
	if GCModeFlag.Value != "full" {
		t.Errorf("GCModeFlag.Value = %q, want %q", GCModeFlag.Value, "full")
	}
}
