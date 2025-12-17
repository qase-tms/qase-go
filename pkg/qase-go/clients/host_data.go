package clients

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"golang.org/x/mod/modfile"
)

// HostData contains information about the host system and package versions
type HostData struct {
	System           string // OS name (e.g., "Linux", "Darwin", "Windows")
	MachineName      string // Machine/hostname
	Release          string // OS release version
	Version          string // OS version
	Arch             string // Architecture (e.g., "amd64", "arm64")
	Framework        string // Testing framework name (e.g., "testing", "testify")
	FrameworkVersion string // Testing framework version
	Reporter         string // Reporter name (e.g., "qase-go")
	ReporterVersion  string // Reporter version
	Commons          string // Commons/core version
	APIClientV1      string // API client v1 version
	APIClientV2      string // API client v2 version
}

// GetHostInfo collects system information and package versions
// frameworkPackage: name of the testing framework (e.g., "testing", "testify")
// frameworkVersion: version of the testing framework
// reporterName: name of the reporter (e.g., "qase-go")
// reporterVersion: version of the reporter
func GetHostInfo() *HostData {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "unknown"
	}

	// Get OS information
	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Get Go version and remove "go" prefix
	goVersion := strings.TrimPrefix(runtime.Version(), "go")

	// Try to get module versions from go.mod
	apiClientV1Version, apiClientV2Version := getModuleVersionsFromGoMod()

	return &HostData{
		System:           osName,
		MachineName:      hostname,
		Release:          "",
		Version:          "",
		Arch:             arch,
		Framework:        "go",
		FrameworkVersion: goVersion,
		Reporter:         "qase-go",
		ReporterVersion:  domain.Version,
		Commons:          domain.Version,
		APIClientV1:      apiClientV1Version,
		APIClientV2:      apiClientV2Version,
	}
}

// getModuleVersionsFromGoMod reads go.mod file and extracts versions of API client modules
// Returns versions for qase-api-client and qase-api-v2-client
func getModuleVersionsFromGoMod() (string, string) {
	var apiClientV1Version, apiClientV2Version string

	var goModPath string

	// First, try to find go.mod relative to this package file
	// This is more reliable than using working directory
	_, filename, _, ok := runtime.Caller(1) // Get caller (GetHostInfo) location
	if ok {
		packageDir := filepath.Dir(filename)
		// Go up to find go.mod (pkg/qase-go/clients -> pkg/qase-go -> go.mod)
		for i := 0; i < 2; i++ {
			packageDir = filepath.Dir(packageDir)
			candidate := filepath.Join(packageDir, "go.mod")
			if _, err := os.Stat(candidate); err == nil {
				goModPath = candidate
				break
			}
		}
	}

	// Fallback: try to find go.mod file starting from current directory and going up
	if goModPath == "" {
		dir, err := os.Getwd()
		if err == nil {
			for {
				candidate := filepath.Join(dir, "go.mod")
				if _, err := os.Stat(candidate); err == nil {
					goModPath = candidate
					break
				}

				parent := filepath.Dir(dir)
				if parent == dir {
					// Reached root directory
					break
				}
				dir = parent
			}
		}
	}

	if goModPath == "" {
		return "", ""
	}

	// Read go.mod file
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return "", ""
	}

	// Parse go.mod file
	file, err := modfile.Parse(goModPath, data, nil)
	if err != nil {
		return "", ""
	}

	// Find versions of API client modules
	for _, req := range file.Require {
		modulePath := req.Mod.Path
		version := req.Mod.Version

		if strings.Contains(modulePath, "qase-api-client") && !strings.Contains(modulePath, "qase-api-v2-client") {
			apiClientV1Version = strings.TrimPrefix(version, "v")
		} else if strings.Contains(modulePath, "qase-api-v2-client") {
			apiClientV2Version = strings.TrimPrefix(version, "v")
		}
	}

	return apiClientV1Version, apiClientV2Version
}

// buildXClientHeader builds the X-Client header value from HostData
// Format: reporter={reporter_name};reporter_version=v{reporter_version};framework={framework};framework_version={framework_version};client_version_v1=v{api_client_v1_version};client_version_v2=v{api_client_v2_version};core_version=v{commons_version}
func buildXClientHeader(hostData *HostData) string {
	if hostData == nil {
		return ""
	}

	var parts []string

	if hostData.Reporter != "" {
		parts = append(parts, "reporter="+hostData.Reporter)
	}

	if hostData.ReporterVersion != "" {
		version := strings.TrimPrefix(hostData.ReporterVersion, "v")
		parts = append(parts, "reporter_version="+version)
	}

	if hostData.Framework != "" {
		parts = append(parts, "framework="+hostData.Framework)
	}

	if hostData.FrameworkVersion != "" {
		version := strings.TrimPrefix(hostData.FrameworkVersion, "v")
		parts = append(parts, "framework_version="+version)
	}

	if hostData.APIClientV1 != "" {
		version := strings.TrimPrefix(hostData.APIClientV1, "v")
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}
		parts = append(parts, "client_version_v1="+version)
	}

	if hostData.APIClientV2 != "" {
		version := strings.TrimPrefix(hostData.APIClientV2, "v")
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}
		parts = append(parts, "client_version_v2="+version)
	}

	if hostData.Commons != "" {
		version := strings.TrimPrefix(hostData.Commons, "v")
		if !strings.HasPrefix(version, "v") {
			version = "v" + version
		}
		parts = append(parts, "core_version="+version)
	}

	return strings.Join(parts, ";")
}

// buildXPlatformHeader builds the X-Platform header value from HostData
// Format: os={os_name};arch={arch};{language}={language_version};{package_manager}={package_manager_version}
func buildXPlatformHeader(hostData *HostData) string {
	if hostData == nil {
		return ""
	}

	var parts []string

	if hostData.System != "" {
		parts = append(parts, "os="+hostData.System)
	}

	if hostData.Arch != "" {
		parts = append(parts, "arch="+hostData.Arch)
	}

	// For Go, we use "go" as the language
	// Go version from runtime.Version() includes "go" prefix (e.g., "go1.21.0")
	goVersion := runtime.Version()
	if goVersion != "" {
		parts = append(parts, "go="+goVersion)
	}

	// Package manager for Go would be "go" with version, but we can't easily get it at runtime
	// So we'll skip it for now

	return strings.Join(parts, ";")
}
