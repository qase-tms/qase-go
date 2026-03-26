package api_v1_client

import (
	"runtime/debug"
	"strings"
	"sync"
)

const modulePath = "github.com/qase-tms/qase-go/qase-api-client"

var (
	moduleVersion     string
	moduleVersionOnce sync.Once
)

// getModuleVersion returns the version of this module resolved from build info.
func getModuleVersion() string {
	moduleVersionOnce.Do(func() {
		moduleVersion = resolveModuleVersion()
	})
	return moduleVersion
}

func resolveModuleVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	for _, dep := range info.Deps {
		if dep.Path == modulePath {
			return strings.TrimPrefix(dep.Version, "v")
		}
	}

	if info.Main.Path == modulePath && info.Main.Version != "" && info.Main.Version != "(devel)" {
		return strings.TrimPrefix(info.Main.Version, "v")
	}

	return "unknown"
}
