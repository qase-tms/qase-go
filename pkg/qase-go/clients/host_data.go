package clients

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/qase-tms/qase-go/pkg/qase-go/domain"
	"golang.org/x/mod/modfile"
)

// HostData contains information about the host system and package versions.
// Field names are unified across all Qase reporter languages.
type HostData struct {
	System         string `json:"system"`         // OS name: "linux", "darwin", "windows"
	MachineName    string `json:"machineName"`    // Machine/hostname
	Release        string `json:"release"`        // OS kernel/release version (uname -r)
	Version        string `json:"version"`        // Detailed OS version (e.g., macOS 15.4, Ubuntu 22.04)
	Arch           string `json:"arch"`           // CPU architecture (e.g., "amd64", "arm64")
	Language       string `json:"language"`       // Runtime version (Go version without "go" prefix)
	PackageManager string `json:"packageManager"` // Package manager version (empty for Go)
	Framework      string `json:"framework"`      // Test framework version (empty for Go stdlib)
	Reporter       string `json:"reporter"`       // Reporter version
	Commons        string `json:"commons"`        // Commons/core version
	ApiClientV1    string `json:"apiClientV1"`    // API client v1 version
	ApiClientV2    string `json:"apiClientV2"`    // API client v2 version
}

// GetHostInfo collects system information and package versions
func GetHostInfo() *HostData {
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "unknown"
	}

	// Get Go version and remove "go" prefix
	goVersion := strings.TrimPrefix(runtime.Version(), "go")

	// Try to get module versions from go.mod
	apiClientV1Version, apiClientV2Version := getModuleVersionsFromGoMod()

	return &HostData{
		System:         runtime.GOOS,
		MachineName:    hostname,
		Release:        getOSRelease(),
		Version:        getOSVersion(),
		Arch:           runtime.GOARCH,
		Language:       goVersion,
		PackageManager: "",
		Framework:      "",
		Reporter:       domain.Version,
		Commons:        domain.Version,
		ApiClientV1:    apiClientV1Version,
		ApiClientV2:    apiClientV2Version,
	}
}

// getOSRelease returns the OS kernel/release version (uname -r on Unix systems)
func getOSRelease() string {
	switch runtime.GOOS {
	case "linux", "darwin":
		out, err := exec.Command("uname", "-r").Output()
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(out))
	default:
		return ""
	}
}

// getOSVersion returns a detailed OS version string
func getOSVersion() string {
	switch runtime.GOOS {
	case "darwin":
		out, err := exec.Command("sw_vers", "-productVersion").Output()
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(out))
	case "linux":
		data, err := os.ReadFile("/etc/os-release")
		if err != nil {
			return ""
		}
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				val := strings.TrimPrefix(line, "PRETTY_NAME=")
				return strings.Trim(val, "\"")
			}
		}
		return ""
	default:
		return ""
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
// Format: reporter=qase-go;reporter_version=<version>;framework=go;framework_version=<version>;client_version_v1=<version>;client_version_v2=<version>;core_version=<version>
func buildXClientHeader(hostData *HostData) string {
	if hostData == nil {
		return ""
	}

	var parts []string

	parts = append(parts, "reporter=qase-go")

	if hostData.Reporter != "" {
		parts = append(parts, "reporter_version="+hostData.Reporter)
	}

	parts = append(parts, "framework=go")

	if hostData.Framework != "" {
		parts = append(parts, "framework_version="+hostData.Framework)
	}

	if hostData.ApiClientV1 != "" {
		parts = append(parts, "client_version_v1="+hostData.ApiClientV1)
	}

	if hostData.ApiClientV2 != "" {
		parts = append(parts, "client_version_v2="+hostData.ApiClientV2)
	}

	if hostData.Commons != "" {
		parts = append(parts, "core_version="+hostData.Commons)
	}

	return strings.Join(parts, ";")
}

// buildXPlatformHeader builds the X-Platform header value from HostData
// Format: os={os_name};arch={arch};go={language_version}
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

	if hostData.Language != "" {
		parts = append(parts, "go="+hostData.Language)
	}

	return strings.Join(parts, ";")
}
