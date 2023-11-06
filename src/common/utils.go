package common

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func RecursiveSearch(logs scribe.Emitter, rootDir string, regex regexp.Regexp) int {
	found := 0
	entries, _ := os.ReadDir(rootDir)
	for _, e := range entries {
		if e.IsDir() {
			found += RecursiveSearch(logs, filepath.Join(rootDir, e.Name()), regex)
		} else {
			if regex.MatchString(e.Name()) {
				found += 1
				logs.Subprocess("Matched file: %s", e.Name())
			}
		}
	}
	return found
}
