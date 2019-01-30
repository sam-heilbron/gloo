package main

import (
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/solo-kit/pkg/code-generator/cmd"
	"github.com/solo-io/solo-kit/pkg/utils/log"
)

//go:generate go run generate.go

func main() {
	err := version.CheckVersions()
	if err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
	log.Printf("starting generate")
	if err := cmd.Run("projects", true, true, nil, nil); err != nil {
		log.Fatalf("generate failed!: %v", err)
	}
}
