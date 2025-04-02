// golang gin framework mvc and clean code project
// Licensed under the Apache License 2.0
// @author Selman TUNÇ <selmantunc@gmail.com>
// @link: https://github.com/stnc/go-mvc-blog-clean-code
// @license: Apache License 2.0
package main

import (
	"fmt"
	"os"

	// "path"

	"github.com/go-git/go-git/v5"
	// "stncCms/internal/commons"
	// "github.com/go-liquor/liquor/internal/constants"
	// "github.com/go-liquor/liquor/internal/message"
	// cp "github.com/otiai10/copy"
	// "io"
	// "path"
	// "path/filepath"
)

// IsExists return true if path or file exists
func IsExist(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func main() {

	url := "https://github.com/git-fixtures/basic.git"
	directory := "git_/r"

	// Clone the given repository to the given directory
	Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)

}
