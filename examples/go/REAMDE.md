# How to upload test results to Qase from the Go tests

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
