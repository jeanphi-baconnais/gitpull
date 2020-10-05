package utils

import (
	"testing"
)

func TestIsValidDir(t *testing.T) {
	bool := IsValidGitDir("test")
	if bool {
		t.Errorf("test is not a valid git directory")
	}
}

func TestIsValidDirGit(t *testing.T) {
	bool := IsValidGitDir(".git")
	if !bool {
		t.Errorf(".git is not a directory")
	}
}

func TestIsRealDir(t *testing.T) {
	bool := IsRealDir(".test")
	if bool {
		t.Errorf(".git is not a real directory")
	}
}
func TestIsRealDirTest(t *testing.T) {
	bool := IsRealDir("target")
	if bool {
		t.Errorf("target is not a real directory")
	}
}

func TestIsRealDirIdea(t *testing.T) {
	bool := IsRealDir(".idea")
	if bool {
		t.Errorf(".idea is not a real directory")
	}
}

func TestIsRealDirNodeModules(t *testing.T) {
	bool := IsRealDir("node_modules")
	if bool {
		t.Errorf("node_modules is not a real directory")
	}
}

func TestIsRealDirCache(t *testing.T) {
	bool := IsRealDir("cache")
	if bool {
		t.Errorf("cache is not a real directory")
	}
}

func TestIsRealDirPkg(t *testing.T) {
	bool := IsRealDir("pkg")
	if bool {
		t.Errorf("pkg is not a real directory")
	}
}