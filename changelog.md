# qase-go@1.2.0

## What's new

Added Tags support for test cases:

- Added `Tags` field to `TestMetadata` struct for specifying tags on tests
- Added `Tags` field to `TestResult` domain model with `AddTags()` method
- Tags are sent to Qase API as comma-separated string in `ResultCreateFields.Tags`
- Support for tags via `Fields` map (`Fields: map[string]string{"tags": "smoke,regression"}`)
- Duplicate tags are automatically removed
- Updated API v2 client dependency to v1.1.7

# qase-go@1.1.0

## What's new

Unified HostData model to align field names across all Qase reporter languages:

- Added `Language`, `PackageManager` fields
- Consolidated `Framework`/`FrameworkVersion` into single `Framework` field (version only)
- Consolidated `Reporter`/`ReporterVersion` into single `Reporter` field (version only)
- Renamed `APIClientV1`/`APIClientV2` to `ApiClientV1`/`ApiClientV2` for cross-language consistency
- Added `Release` (OS kernel version via `uname -r`) and `Version` (detailed OS version) fields
- X-Platform header now uses `hostData.Language` instead of `runtime.Version()`
- Reporter and framework names are hardcoded in header builders, not stored in HostData
