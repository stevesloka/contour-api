package main

import (
	"fmt"
	"log"

	"github.com/projectcontour/contour-api/internal/examples"

	"gopkg.in/yaml.v2"
)

func main() {
	d, err := yaml.Marshal(examples.Basic())
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}
