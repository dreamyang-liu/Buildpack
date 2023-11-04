package pip3

import (
	"aws-buildpacks/src/python"
	"fmt"
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/chronos"
	"github.com/paketo-buildpacks/packit/v2/pexec"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func NewPipBuildFunc(logs scribe.Emitter, clock chronos.Clock) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		// Fetch environment variables
		pythonPath := os.Getenv(python.PythonPath)

		execution := pexec.Execution{
			Args:   []string{"-m", "pip", "install", "-r", "requirements.txt"},
			Dir:    context.WorkingDir,
			Env:    append(os.Environ(), fmt.Sprintf("PYTHONPATH=%s", pythonPath)),
			Stdout: logs.ActionWriter,
			Stderr: logs.ActionWriter,
		}

		duration, err := clock.Measure(func() error {
			return pexec.NewExecutable(fmt.Sprintf("%s/bin/python3", pythonPath)).Execute(execution)
		})

		logs.Detail("pip install completed in %s", duration.Seconds())

		if err != nil {
			return packit.BuildResult{}, err
		}

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
			Layers: []packit.Layer{},
			Launch: launchMetadata,
			Build:  packit.BuildMetadata{},
		}, nil

	}
}
