package nodejs

import (
	"os"
	"path/filepath"
	"strings"

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

		entries, err := os.ReadDir("./")
		if err != nil {
			return packit.DetectResult{}, err
		}
		match := 0
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), ".js") {
				match += 1
			}
		}
		if match == 0 {
			return packit.DetectResult{}, packit.Fail.WithMessage("No Js file found !")
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
