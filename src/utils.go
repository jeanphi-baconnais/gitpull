package main

import (
	"strings"
)

func IsValidDir(dirName string) bool {
	result := 0

	result &= strings.Compare(dirName, ".git")

	return result == 0
}

func isRealDir(dir string) bool {
	result := false
	/*

	result = result || !strings.HasPrefix(dir, ".")
	result = result || strings.Compare(dir,"target") != 0
	result = result || strings.Compare(dir,".idea") != 0
	result = result || strings.Compare(dir,"node_modules") != 0
	result = result || strings.Compare(dir,"cache") != 0

			result = result && dir.IsDir()
			result = result && !strings.HasPrefix(dir, ".")
			result = result && strings.Compare(dir,"target") != 0
			result = result && strings.Compare(dir,".idea") != 0
			result = result && strings.Compare(dir,"node_modules") != 0
			result = result && strings.Compare(dir,"cache") != 0


	*/

	if !strings.HasPrefix(dir, ".") && strings.Compare(dir,"target") != 0 && strings.Compare(dir,".idea") != 0 && strings.Compare(dir,"node_modules") != 0 && strings.Compare(dir,"cache") != 0 && strings.Compare(dir,"pkg") != 0 {
		result = true
	}

	return result
}

func TestMethod() bool {
	r := false

	r = r && true

	return r
}