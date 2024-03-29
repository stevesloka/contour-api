package main

import (
	"fmt"
	"os"

	"github.com/stevesloka/contour-api/internal/examples"

	"gopkg.in/yaml.v2"
)

const (
	outputPath = "examples"
	apiVersion = "projectcontour.io/v1alpha1"
	kind       = "HTTPProxy"
)

func main() {
	ex := examples.Example{}

	err := os.RemoveAll(outputPath)
	check(err)
	err = os.Mkdir(outputPath, os.ModePerm)
	check(err)

	for _, v := range ex.Get() {
		currentDir := fmt.Sprintf("%s/%s", outputPath, v.DirName)
		err := os.Mkdir(currentDir, os.ModePerm)
		check(err)

		file, err := os.Create(fmt.Sprintf("%s/%s.yaml", currentDir, v.Name))
		check(err)

		_, err = file.WriteString(fmt.Sprintf("# Name: %s\n# Description: %s\n\n", v.Name, v.Description))
		check(err)

		for _, HTTPProxy := range v.HTTPProxy {
			d, err := yaml.Marshal(HTTPProxy)
			check(err)

			_, err = file.WriteString(fmt.Sprintf("apiVersion: %s\nkind: %s\n%s---\n", apiVersion, kind, string(d)))
			check(err)
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
