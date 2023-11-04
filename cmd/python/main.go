package main

import (
	"aws-buildpacks/src/python"
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/cargo"
	"github.com/paketo-buildpacks/packit/v2/chronos"
	"github.com/paketo-buildpacks/packit/v2/draft"
	"github.com/paketo-buildpacks/packit/v2/postal"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func main() {
	// Setup log level based on environment variable
	logger := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	entryResolver := draft.NewPlanner()
	dependencyManager := postal.NewService(cargo.NewTransport())
	// SbomGeneratorImpl implements SbomGenerator interface

	// Run pack
	packit.Run(
		python.NewPythonRuntimeDetectFunc(logger),
		python.NewPythonBuildFunc(
			logger,
			entryResolver,
			dependencyManager,
			chronos.DefaultClock,
		),
	)
}
