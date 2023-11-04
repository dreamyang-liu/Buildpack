package python

import (
	"aws-buildpacks/src/common"
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func NewPythonRuntimeDetectFunc(logs scribe.Emitter) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		var requirements []packit.BuildPlanRequirement

		logs.Detail("Setting up default python version")
		requirements = append(requirements, packit.BuildPlanRequirement{
			Name: Python,
			Metadata: common.BuildPlanRequirementMetadata{
				VersionSource: BpDefaultPythonEnv,
				Version:       DefaultPythonVersion,
				Build:         true,
				Launch:        true,
			},
		})

		logs.Detail("Checking for user specified python version")
		pythonVersion := os.Getenv(BpUserPythonEnv)
		if pythonVersion != "" {
			requirements = append(requirements, packit.BuildPlanRequirement{
				Name: Python,
				Metadata: common.BuildPlanRequirementMetadata{
					VersionSource: BpUserPythonEnv,
					Version:       pythonVersion,
					Build:         true,
					Launch:        true,
				},
			})
		}

		if len(requirements) == 0 {
			return packit.DetectResult{}, packit.Fail.WithMessage("failed to detect python")
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: Python},
				},
				Requires: requirements,
			},
		}, nil
	}
}
