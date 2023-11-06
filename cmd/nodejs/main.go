package main

import (
	"aws-buildpacks/src/nodejs"
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func main() {
	// Setup log level based on environment variable
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	// SbomGeneratorImpl implements SbomGenerator interface

	// Run pack
	packit.Run(
		nodejs.NewNodejsDetectFunc(logger),
		nodejs.NewNodejsBuildFunc(logger),
	)
}
