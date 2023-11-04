package pip3

import (
	"aws-buildpacks/src/common"
	"aws-buildpacks/src/python"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/fs"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func NewPipDetectFunc(logs scribe.Emitter) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		exists, err := fs.Exists(python.RequirementsTxt)
		if err != nil {
			return packit.DetectResult{}, err
		}

		if !exists {
			return packit.DetectResult{}, packit.Fail.WithMessage("requirements.txt not found")
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: Pip},
				},
				Requires: []packit.BuildPlanRequirement{
					{
						Name: Pip,
						Metadata: common.BuildPlanRequirementMetadata{
							Build:  true,
							Launch: true,
						},
					},
				},
			},
		}, nil
	}
}
