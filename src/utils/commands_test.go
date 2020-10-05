package utils

import (
	"testing"
)

func TestInitTool(t *testing.T) {
	bool := IsValidGitDir("test")
	if bool {
		t.Errorf("test is not a valid git directory")
	}
}
