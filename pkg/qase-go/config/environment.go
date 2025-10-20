package config

const (
	// Main Configuration
	QaseModeEnvVar        = "QASE_MODE"
	QaseFallbackEnvVar    = "QASE_FALLBACK"
	QaseDebugEnvVar       = "QASE_DEBUG"
	QaseEnvironmentEnvVar = "QASE_ENVIRONMENT"
	QaseCaptureLogsEnvVar = "QASE_CAPTURE_LOGS"

	// TestOps Configuration
	QaseTestOpsAPITokenEnvVar     = "QASE_TESTOPS_API_TOKEN"
	QaseTestOpsProjectEnvVar      = "QASE_TESTOPS_PROJECT"
	QaseTestOpsRunIDEnvVar        = "QASE_TESTOPS_RUN_ID"
	QaseTestOpsAPIHostEnvVar      = "QASE_TESTOPS_API_HOST"
	QaseTestOpsDefectEnvVar       = "QASE_TESTOPS_DEFECT"
	QaseTestOpsBatchSizeEnvVar    = "QASE_TESTOPS_BATCH_SIZE"
	QaseTestOpsStatusFilterEnvVar = "QASE_TESTOPS_STATUS_FILTER"
	QaseStatusMappingEnvVar       = "QASE_STATUS_MAPPING"

	// Report Configuration
	QaseReportDriverEnvVar           = "QASE_REPORT_DRIVER"
	QaseReportConnectionPathEnvVar   = "QASE_REPORT_CONNECTION_PATH"
	QaseReportConnectionFormatEnvVar = "QASE_REPORT_CONNECTION_FORMAT"

	// Logging Configuration
	QaseLoggingConsoleEnvVar = "QASE_LOGGING_CONSOLE"
	QaseLoggingFileEnvVar    = "QASE_LOGGING_FILE"
)

const (
	MODE_TESTOPS = "testops"
	MODE_REPORT  = "report"
	MODE_OFF     = "off"
)

var VALID_MODES [3]string = [3]string{MODE_TESTOPS, MODE_REPORT, MODE_OFF}
var VALID_FALLBACKS [2]string = [2]string{MODE_REPORT, MODE_OFF}

const (
	STATUS_PASSED      = "passed"
	STATUS_FAILED      = "failed"
	STATUS_BLOCKED     = "blocked"
	STATUS_SKIPPED     = "skipped"
	STATUS_IN_PROGRESS = "in_progress"
	STATUS_INVALID     = "invalid"
)

var VALID_STATUSES [6]string = [6]string{STATUS_PASSED, STATUS_FAILED,
	STATUS_BLOCKED, STATUS_SKIPPED, STATUS_IN_PROGRESS, STATUS_INVALID}
