package nodejs

import (
	"os"
	"path/filepath"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/cargo"
	"github.com/paketo-buildpacks/packit/v2/draft"
	"github.com/paketo-buildpacks/packit/v2/postal"
	"github.com/paketo-buildpacks/packit/v2/sbom"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func NewNodejsBuildFunc(logs scribe.Emitter) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		priority := []interface{}{
			BpNodejsUserVersion,
			PackageJSON,
			BpNodejsDefaultVersion,
		}
		resolver := draft.NewPlanner()
		entry, entries := resolver.Resolve(Nodejs, context.Plan.Entries, priority)
		logs.Candidates(entries)
		version, ok := entry.Metadata["version"].(string)
		if !ok {
			version = "default"
		}
		nodejsLayer, err := context.Layers.Get(NodejsLayer)
		if err != nil {
			return packit.BuildResult{}, err
		}

		nodejsLayer.Build, nodejsLayer.Launch, nodejsLayer.Cache = true, true, true
		dependencyManager := postal.NewService(cargo.NewTransport())

		dependency, err := dependencyManager.Resolve(filepath.Join(context.CNBPath, "buildpack.toml"), entry.Name, version, context.Stack)
		if err != nil {
			return packit.BuildResult{}, err
		}

		dependencyManager.Deliver(dependency, context.CNBPath, nodejsLayer.Path, context.Platform.Path)

		logs.GeneratingSBOM(nodejsLayer.Path)
		sbomContent, err := sbom.GenerateFromDependency(dependency, nodejsLayer.Path)
		if err != nil {
			return packit.BuildResult{}, err
		}

		logs.FormattingSBOM(context.BuildpackInfo.SBOMFormats...)
		nodejsLayer.SBOM, err = sbomContent.InFormats(context.BuildpackInfo.SBOMFormats...)
		if err != nil {
			return packit.BuildResult{}, err
		}
		nodejsLayer.SharedEnv.Default(NODE_ENV, DEFAULT_NODE_ENV_VALUE)

		launchMetadata := packit.LaunchMetadata{
			Processes: []packit.Process{
				{
					Type:    "web",
					Command: "node server.js",
					Default: true,
				},
			},
		}

		entrypoint := os.Getenv("BP_PYTHON_ENTRYPOINT")
		if entrypoint != "" {
			launchMetadata.Processes[0].Command = entrypoint
		}

		return packit.BuildResult{
			Layers: []packit.Layer{nodejsLayer},
			Build:  packit.BuildMetadata{},
			Launch: launchMetadata,
		}, nil

	}
}
