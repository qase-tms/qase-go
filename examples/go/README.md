# How to upload test results to Qase from the Go tests

## Automatic Test Run Completion

Starting from version X.X.X, Qase Go reporter now supports **automatic test run completion**. This means you no longer need to call `defer qase.CompleteTestRun()` in every test function.

### How it works

- Add `TestMain` function to your test file
- The test run is automatically completed **after all tests finish**
- `sync.Once` ensures completion happens only once
- You can still call `qase.CompleteTestRun()` manually if immediate completion is needed

### Example

```go
package yourpackage

import (
    "testing"
    "github.com/qase-tms/qase-go/pkg/qase-go/qase"
)

// TestMain automatically completes the test run after all tests finish
func TestMain(m *testing.M) {
    qase.TestMain(m)
}

func TestExample(t *testing.T) {
    // No defer qase.CompleteTestRun() needed!
    
    qase.Test(t, qase.TestMetadata{...}, func() {
        // test logic
    })
    // Test run will be automatically completed after ALL tests finish
}
```

### Benefits

- **Cleaner code**: No need for `defer qase.CompleteTestRun()` in every test
- **No duplicate results**: Test run is completed only once, after all tests
- **Backward compatible**: Manual completion still works if needed
- **Automatic cleanup**: Test run completion happens automatically at the right time
- **Proper sequencing**: Results are sent before test run completion
- **Robust error handling**: Individual result failures don't crash the entire test run

### How it's implemented

1. **TestMain**: Uses Go's built-in `testing.M` to run tests
2. **Auto-completion**: Automatically calls `CompleteTestRun()` after all tests finish
3. **Single execution**: `sync.Once` ensures completion happens only once
4. **Graceful degradation**: Problematic results are skipped instead of failing the entire run

### Error Handling

The reporter now handles errors more gracefully:
- **Individual result failures** are logged as warnings and skipped
- **Batch uploads continue** even if some results fail to convert
- **Test run completion** proceeds as long as at least some results were uploaded
- **Detailed logging** shows which results succeeded and which failed

---

You can upload test results to Qase from the Go tests. 

First, you need to install the `gotestsum` package. This package is used to run the tests and save the results in the JUnit format. 

Then you can upload the test results to Qase using the `qasectl` tool.

Step-by-step guide:

1. Clone the repository

    ```bash
    git clone https://github.com/qase-tms/qase-go.git
    ```

2. Move to the directory with the examples

    ```bash
    cd qase-go/examples/go
    ```

3. Install the required packages

    ```bash
    go mod tidy
    ```

4. Install the `gotestsum` package

    ```bash
    go install gotest.tools/gotestsum@latest
    ```

5. Run the tests and save the results to the `test-results.xml` file

    ```bash
    gotestsum --junitfile results/test-results.xml
    ```

6. Upload the test results to Qase using the [qasectl](https://github.com/qase-tms/qasectl/blob/main/docs/command.md) tool

    ```bash 
    qli testops result upload --project DEMO --token token --id 1 --format junit --path /results/test-results.xml --verbose
    ```  
