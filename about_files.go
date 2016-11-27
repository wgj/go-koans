package go_koans

import "io/ioutil"
import "strings"

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	assert(lines[5] == "func aboutFiles() {") // handling files is too trivial

	//TODO add an exercise for writing files too.
	// Execise suggestion: assert input, write to a file, read file, assert output.
}
