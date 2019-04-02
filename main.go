package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var zoneDirs = []string{
	"/usr/share/zoneinfo/",
	"/usr/share/lib/zoneinfo/",
	"/usr/lib/locale/TZ/",
}

var zoneDir string
var all map[string]bool

func main() {
	all = map[string]bool{}
	for _, zoneDir = range zoneDirs {
		ReadFile("")
	}

	location := []string{}
	for name, _ := range all {
		location = append(location, name)
	}
	b, _ := json.MarshalIndent(&location, "", " ")
	fmt.Println(string(b))
}

func ReadFile(path string) {
	files, _ := ioutil.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			ReadFile(path + "/" + f.Name())
		} else {
			p := path + "/" + f.Name()
			all[p[1:]] = true
		}
	}
}
