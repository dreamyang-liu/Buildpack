package nodejs

import (
	"os"
	"path/filepath"
	"regexp"

	"aws-buildpacks/src/common"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/fs"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func NewNodejsDetectFunc(logs scribe.Emitter) packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		workingDir := context.WorkingDir
		var requirements []packit.BuildPlanRequirement

		requirements = append(requirements, packit.BuildPlanRequirement{
			Name: Nodejs,
			Metadata: common.BuildPlanRequirementMetadata{
				VersionSource: BpNodejsDefaultVersion,
				Version:       NodejsDefaultVersion,
				Build:         true,
				Launch:        true,
			},
		})

		nodejsUserVersion := os.Getenv(BpNodejsUserVersion)
		if nodejsUserVersion != "" {
			requirements = append(requirements, packit.BuildPlanRequirement{
				Name: Nodejs,
				Metadata: common.BuildPlanRequirementMetadata{
					VersionSource: BpNodejsUserVersion,
					Version:       nodejsUserVersion,
					Build:         true,
					Launch:        true,
				},
			})
		}

		exists, err := fs.Exists(filepath.Join(workingDir, PackageJSON))
		if err != nil {
			return packit.DetectResult{}, err
		}
		if exists {
			requirements = append(requirements, packit.BuildPlanRequirement{
				Name: Nodejs,
				Metadata: common.BuildPlanRequirementMetadata{
					VersionSource: PackageJSON,
					Version:       "18.17.1",
					Build:         true,
					Launch:        true,
				},
			})
		}

		regexExpression, _ := regexp.Compile(`.+\.js`)
		match := common.RecursiveSearch(logs, context.WorkingDir, *regexExpression)
		if match == 0 {
			return packit.DetectResult{}, packit.Fail.WithMessage("No js file found !")
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Provides: []packit.BuildPlanProvision{
					{Name: Nodejs},
				},
				Requires: requirements,
			},
		}, nil
	}
}
