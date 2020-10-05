package utils

import (
	"strings"
)

func IsValidGitDir(dirName string) bool {
	result := strings.Compare(dirName, ".git")

	return result == 0
}

func IsRealDir(dir string) bool {
	result := false

	if !strings.HasPrefix(dir, ".") && strings.Compare(dir, "target") != 0 && strings.Compare(dir, ".idea") != 0 && strings.Compare(dir, "node_modules") != 0 && strings.Compare(dir, "cache") != 0 && strings.Compare(dir, "pkg") != 0 {
		result = true
	}

	return result
}

