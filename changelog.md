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
