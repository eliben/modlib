package auth

import (
	"strings"
	"testing"
)

func TestGetAuth(t *testing.T) {
	msg := GetAuth()
	if strings.Index(msg, "not") >= 0 {
		t.Errorf("expected not to find 'not', got %v", msg)
	}
}
