package main

import (
	"fmt"
	"path"
)

func main() {
	// Split file and directory
	dir, file := path.Split("/usr/bin/ls")
	fmt.Printf("dir = %+v\n", dir)
	fmt.Printf("file = %+v\n", file)

	// Join paths
	fmt.Println(path.Join("usr", "bin", "ls"))
	fmt.Println(path.Join("usr/", "bin", "ls"))
	fmt.Println(path.Join("usr", "/bin", "ls"))
	fmt.Println(path.Join("usr", "/bin", "/ls"))
	fmt.Println(path.Join("usr/bin", "/ls"))

	// verify if path is absolute
	fmt.Println("Is absolute? ", path.IsAbs("/dev/null"))
	fmt.Println("Is absolute? ", path.IsAbs("./dev/null"))

	// Get the extension
	fmt.Println("Extension: static.css -> ", path.Ext("static.css"))
	fmt.Println("Extension: noextension ->", path.Ext("noextension"))

	// Get the path of a file
	fmt.Println("Dir: usr/bin/ls -> ", path.Dir("usr/bin/ls"))
	fmt.Println("Dir: ->  ", path.Dir(""))

	// See more on docs: https://golang.org/pkg/path

}
