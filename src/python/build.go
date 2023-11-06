package python

import (
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/chronos"
	"github.com/paketo-buildpacks/packit/v2/draft"
	"github.com/paketo-buildpacks/packit/v2/postal"
	"github.com/paketo-buildpacks/packit/v2/sbom"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func NewPythonBuildFunc(logs scribe.Emitter, entryResolver draft.Planner, dependencyManager postal.Service, clock chronos.Clock) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {

		priorities := []interface{}{
			BpUserPythonEnv,
			BpDefaultPythonEnv,
		}
		entry, _ := entryResolver.Resolve(Python, context.Plan.Entries, priorities)

		pythonLayer, err := context.Layers.Get(PythonLayer)
		if err != nil {
			return packit.BuildResult{}, err
		}

		version, ok := entry.Metadata["version"].(string)
		if !ok {
			version = "default"
		}

		pythonLayer.Launch, pythonLayer.Build, pythonLayer.Cache = true, true, true

		dependency, err := dependencyManager.Resolve(filepath.Join(context.CNBPath, "buildpack.toml"), entry.Name, version, context.Stack)

		if err != nil {
			return packit.BuildResult{}, err
		}

		logs.SelectedDependency(entry, dependency, clock.Now())

		dependencyManager.Deliver(dependency, context.CNBPath, pythonLayer.Path, context.Platform.Path)

		logs.GeneratingSBOM(pythonLayer.Path)

		sbomContent, err := sbom.GenerateFromDependency(dependency, pythonLayer.Path)

		if err != nil {
			return packit.BuildResult{}, err
		}

		pythonLayer.SBOM, err = sbomContent.InFormats(context.BuildpackInfo.SBOMFormats...)

		if err != nil {
			return packit.BuildResult{}, err
		}

		os.Setenv(PythonPath, pythonLayer.Path)
		pythonLayer.SharedEnv.Default(PythonPath, pythonLayer.Path)

		launchMetadata := packit.LaunchMetadata{
			Processes: []packit.Process{
				{
					Type:    "web",
					Command: "python3 app.py",
					Default: true,
				},
			},
		}

		entrypoint := os.Getenv("BP_PYTHON_ENTRYPOINT")
		if entrypoint != "" {
			launchMetadata.Processes[0].Command = entrypoint
		}

		return packit.BuildResult{
			Layers: []packit.Layer{pythonLayer},
			Launch: launchMetadata,
			Build:  packit.BuildMetadata{},
		}, nil

	}
}
