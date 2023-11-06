package python

import (
	"aws-buildpacks/src/common"
	"os"
	"strings"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func NewPythonRuntimeDetectFunc(logs scribe.Emitter) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		var requirements []packit.BuildPlanRequirement

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

		entries, err := os.ReadDir("./")
		if err != nil {
			return packit.DetectResult{}, err
		}
		match := 0
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".py") || e.Name() == RequirementsTxt {
				match += 1
			}
		}
		if match == 0 {
			return packit.DetectResult{}, packit.Fail.WithMessage("No py or requirements.txt file found !")
		}

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
