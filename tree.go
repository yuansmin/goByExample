// tools like linux tree
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	path := os.Args[1]
	err := listPath(path)
	if err != nil {
		fmt.Printf("tree dir err: %v", err)
	}
}

func listPath(path string) (err error) {
	var depth int
	err = nil
	var listPathHandler func(curPath string)
	listPathHandler = func(curPath string) {
		fileInfo, err := os.Lstat(curPath)
		if err != nil {
			return
		}
		baseName := filepath.Base(curPath)
		if fileInfo.IsDir() {
			fmt.Printf("%*s%s:\n", depth*4, "", baseName)
		} else {
			fmt.Printf("%*s%s\n", depth*4, "", baseName)
		}

		if fileInfo.IsDir() {
			depth++
			fileInfos, err := ioutil.ReadDir(curPath)
			if err != nil {
				return
			}
			for _, info := range fileInfos {
				listPathHandler(filepath.Join(curPath, info.Name()))
			}
			depth--
		}
	}
	listPathHandler(path)
	return
}
