# qase-go@1.2.1

## What's fixed

Result uploads are no longer rejected or lost silently:

- Batch size is clamped to 200, the maximum the bulk results endpoint accepts — a larger configured value used to make every upload fail with a non-retryable HTTP 413
- Added `DefaultBatchSize` and `MaxBatchSize` constants and a `Config.GetBatchSize()` accessor that applies the limit
- Clamping of an oversized `testops.batch.size` / `QASE_TESTOPS_BATCH_SIZE` is reported in the log
- A partial upload failure now logs how many results were lost and returns an error carrying the same counts, instead of returning `nil` and reporting the number only at debug level
- Removed the dead fallback to a batch size of 50 that contradicted the documented default of 100
- Updated API v1 client dependency to v1.2.13 and API v2 client dependency to v1.1.8

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
