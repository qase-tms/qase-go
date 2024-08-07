# Go API client for api_v1_client

Qase TestOps API v1 Specification.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.1
- Generator version: 7.4.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://qase.io](https://qase.io)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import api_v1_client "github.com/qase-tms/qase-go/qase-api-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `api_v1_client.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), api_v1_client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `api_v1_client.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), api_v1_client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `api_v1_client.ContextOperationServerIndices` and `api_v1_client.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), api_v1_client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), api_v1_client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

### Example

```go
package main

import (
	"context"
	"fmt"
	
	api_v1_client "github.com/qase-tms/qase-go/qase-api-client"
)

func main() {
	ctx := context.WithValue(context.Background(), api_v1_client.ContextAPIKeys, map[string]api_v1_client.APIKey{
		"TokenAuth": {
			Key: "<TOKEN>",
		},
	})

	configuration := api_v1_client.NewConfiguration()
	apiClient := api_v1_client.NewAPIClient(configuration)

	proj, _, err := apiClient.ProjectsAPI.GetProject(ctx, "<PROJECT_CODE>").Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(proj)
}

```
## Documentation for API Endpoints

All URIs are relative to *https://api.qase.io/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AttachmentsAPI* | [**DeleteAttachment**](docs/AttachmentsAPI.md#deleteattachment) | **Delete** /attachment/{hash} | Remove attachment by Hash
*AttachmentsAPI* | [**GetAttachment**](docs/AttachmentsAPI.md#getattachment) | **Get** /attachment/{hash} | Get attachment by Hash
*AttachmentsAPI* | [**GetAttachments**](docs/AttachmentsAPI.md#getattachments) | **Get** /attachment | Get all attachments
*AttachmentsAPI* | [**UploadAttachment**](docs/AttachmentsAPI.md#uploadattachment) | **Post** /attachment/{code} | Upload attachment
*AuthorsAPI* | [**GetAuthor**](docs/AuthorsAPI.md#getauthor) | **Get** /author/{id} | Get a specific author
*AuthorsAPI* | [**GetAuthors**](docs/AuthorsAPI.md#getauthors) | **Get** /author | Get all authors
*CasesAPI* | [**Bulk**](docs/CasesAPI.md#bulk) | **Post** /case/{code}/bulk | Create test cases in bulk
*CasesAPI* | [**CaseAttachExternalIssue**](docs/CasesAPI.md#caseattachexternalissue) | **Post** /case/{code}/external-issue/attach | Attach the external issues to the test cases
*CasesAPI* | [**CaseDetachExternalIssue**](docs/CasesAPI.md#casedetachexternalissue) | **Post** /case/{code}/external-issue/detach | Detach the external issues from the test cases
*CasesAPI* | [**CreateCase**](docs/CasesAPI.md#createcase) | **Post** /case/{code} | Create a new test case
*CasesAPI* | [**DeleteCase**](docs/CasesAPI.md#deletecase) | **Delete** /case/{code}/{id} | Delete test case
*CasesAPI* | [**GetCase**](docs/CasesAPI.md#getcase) | **Get** /case/{code}/{id} | Get a specific test case
*CasesAPI* | [**GetCases**](docs/CasesAPI.md#getcases) | **Get** /case/{code} | Get all test cases
*CasesAPI* | [**UpdateCase**](docs/CasesAPI.md#updatecase) | **Patch** /case/{code}/{id} | Update test case
*ConfigurationsAPI* | [**CreateConfiguration**](docs/ConfigurationsAPI.md#createconfiguration) | **Post** /configuration/{code} | Create a new configuration in a particular group.
*ConfigurationsAPI* | [**CreateConfigurationGroup**](docs/ConfigurationsAPI.md#createconfigurationgroup) | **Post** /configuration/{code}/group | Create a new configuration group.
*ConfigurationsAPI* | [**GetConfigurations**](docs/ConfigurationsAPI.md#getconfigurations) | **Get** /configuration/{code} | Get all configuration groups with configurations.
*CustomFieldsAPI* | [**CreateCustomField**](docs/CustomFieldsAPI.md#createcustomfield) | **Post** /custom_field | Create new Custom Field
*CustomFieldsAPI* | [**DeleteCustomField**](docs/CustomFieldsAPI.md#deletecustomfield) | **Delete** /custom_field/{id} | Delete Custom Field by id
*CustomFieldsAPI* | [**GetCustomField**](docs/CustomFieldsAPI.md#getcustomfield) | **Get** /custom_field/{id} | Get Custom Field by id
*CustomFieldsAPI* | [**GetCustomFields**](docs/CustomFieldsAPI.md#getcustomfields) | **Get** /custom_field | Get all Custom Fields
*CustomFieldsAPI* | [**UpdateCustomField**](docs/CustomFieldsAPI.md#updatecustomfield) | **Patch** /custom_field/{id} | Update Custom Field by id
*DefectsAPI* | [**CreateDefect**](docs/DefectsAPI.md#createdefect) | **Post** /defect/{code} | Create a new defect
*DefectsAPI* | [**DeleteDefect**](docs/DefectsAPI.md#deletedefect) | **Delete** /defect/{code}/{id} | Delete defect
*DefectsAPI* | [**GetDefect**](docs/DefectsAPI.md#getdefect) | **Get** /defect/{code}/{id} | Get a specific defect
*DefectsAPI* | [**GetDefects**](docs/DefectsAPI.md#getdefects) | **Get** /defect/{code} | Get all defects
*DefectsAPI* | [**ResolveDefect**](docs/DefectsAPI.md#resolvedefect) | **Patch** /defect/{code}/resolve/{id} | Resolve a specific defect
*DefectsAPI* | [**UpdateDefect**](docs/DefectsAPI.md#updatedefect) | **Patch** /defect/{code}/{id} | Update defect
*DefectsAPI* | [**UpdateDefectStatus**](docs/DefectsAPI.md#updatedefectstatus) | **Patch** /defect/{code}/status/{id} | Update a specific defect status
*EnvironmentsAPI* | [**CreateEnvironment**](docs/EnvironmentsAPI.md#createenvironment) | **Post** /environment/{code} | Create a new environment
*EnvironmentsAPI* | [**DeleteEnvironment**](docs/EnvironmentsAPI.md#deleteenvironment) | **Delete** /environment/{code}/{id} | Delete environment
*EnvironmentsAPI* | [**GetEnvironment**](docs/EnvironmentsAPI.md#getenvironment) | **Get** /environment/{code}/{id} | Get a specific environment
*EnvironmentsAPI* | [**GetEnvironments**](docs/EnvironmentsAPI.md#getenvironments) | **Get** /environment/{code} | Get all environments
*EnvironmentsAPI* | [**UpdateEnvironment**](docs/EnvironmentsAPI.md#updateenvironment) | **Patch** /environment/{code}/{id} | Update environment
*MilestonesAPI* | [**CreateMilestone**](docs/MilestonesAPI.md#createmilestone) | **Post** /milestone/{code} | Create a new milestone
*MilestonesAPI* | [**DeleteMilestone**](docs/MilestonesAPI.md#deletemilestone) | **Delete** /milestone/{code}/{id} | Delete milestone
*MilestonesAPI* | [**GetMilestone**](docs/MilestonesAPI.md#getmilestone) | **Get** /milestone/{code}/{id} | Get a specific milestone
*MilestonesAPI* | [**GetMilestones**](docs/MilestonesAPI.md#getmilestones) | **Get** /milestone/{code} | Get all milestones
*MilestonesAPI* | [**UpdateMilestone**](docs/MilestonesAPI.md#updatemilestone) | **Patch** /milestone/{code}/{id} | Update milestone
*PlansAPI* | [**CreatePlan**](docs/PlansAPI.md#createplan) | **Post** /plan/{code} | Create a new plan
*PlansAPI* | [**DeletePlan**](docs/PlansAPI.md#deleteplan) | **Delete** /plan/{code}/{id} | Delete plan
*PlansAPI* | [**GetPlan**](docs/PlansAPI.md#getplan) | **Get** /plan/{code}/{id} | Get a specific plan
*PlansAPI* | [**GetPlans**](docs/PlansAPI.md#getplans) | **Get** /plan/{code} | Get all plans
*PlansAPI* | [**UpdatePlan**](docs/PlansAPI.md#updateplan) | **Patch** /plan/{code}/{id} | Update plan
*ProjectsAPI* | [**CreateProject**](docs/ProjectsAPI.md#createproject) | **Post** /project | Create new project
*ProjectsAPI* | [**DeleteProject**](docs/ProjectsAPI.md#deleteproject) | **Delete** /project/{code} | Delete Project by code
*ProjectsAPI* | [**GetProject**](docs/ProjectsAPI.md#getproject) | **Get** /project/{code} | Get Project by code
*ProjectsAPI* | [**GetProjects**](docs/ProjectsAPI.md#getprojects) | **Get** /project | Get All Projects
*ProjectsAPI* | [**GrantAccessToProject**](docs/ProjectsAPI.md#grantaccesstoproject) | **Post** /project/{code}/access | Grant access to project by code
*ProjectsAPI* | [**RevokeAccessToProject**](docs/ProjectsAPI.md#revokeaccesstoproject) | **Delete** /project/{code}/access | Revoke access to project by code
*ResultsAPI* | [**CreateResult**](docs/ResultsAPI.md#createresult) | **Post** /result/{code}/{id} | Create test run result
*ResultsAPI* | [**CreateResultBulk**](docs/ResultsAPI.md#createresultbulk) | **Post** /result/{code}/{id}/bulk | Bulk create test run result
*ResultsAPI* | [**DeleteResult**](docs/ResultsAPI.md#deleteresult) | **Delete** /result/{code}/{id}/{hash} | Delete test run result
*ResultsAPI* | [**GetResult**](docs/ResultsAPI.md#getresult) | **Get** /result/{code}/{hash} | Get test run result by code
*ResultsAPI* | [**GetResults**](docs/ResultsAPI.md#getresults) | **Get** /result/{code} | Get all test run results
*ResultsAPI* | [**UpdateResult**](docs/ResultsAPI.md#updateresult) | **Patch** /result/{code}/{id}/{hash} | Update test run result
*RunsAPI* | [**CompleteRun**](docs/RunsAPI.md#completerun) | **Post** /run/{code}/{id}/complete | Complete a specific run
*RunsAPI* | [**CreateRun**](docs/RunsAPI.md#createrun) | **Post** /run/{code} | Create a new run
*RunsAPI* | [**DeleteRun**](docs/RunsAPI.md#deleterun) | **Delete** /run/{code}/{id} | Delete run
*RunsAPI* | [**GetRun**](docs/RunsAPI.md#getrun) | **Get** /run/{code}/{id} | Get a specific run
*RunsAPI* | [**GetRuns**](docs/RunsAPI.md#getruns) | **Get** /run/{code} | Get all runs
*RunsAPI* | [**UpdateRunPublicity**](docs/RunsAPI.md#updaterunpublicity) | **Patch** /run/{code}/{id}/public | Update publicity of a specific run
*SearchAPI* | [**Search**](docs/SearchAPI.md#search) | **Get** /search | Search entities by Qase Query Language (QQL)
*SharedStepsAPI* | [**CreateSharedStep**](docs/SharedStepsAPI.md#createsharedstep) | **Post** /shared_step/{code} | Create a new shared step
*SharedStepsAPI* | [**DeleteSharedStep**](docs/SharedStepsAPI.md#deletesharedstep) | **Delete** /shared_step/{code}/{hash} | Delete shared step
*SharedStepsAPI* | [**GetSharedStep**](docs/SharedStepsAPI.md#getsharedstep) | **Get** /shared_step/{code}/{hash} | Get a specific shared step
*SharedStepsAPI* | [**GetSharedSteps**](docs/SharedStepsAPI.md#getsharedsteps) | **Get** /shared_step/{code} | Get all shared steps
*SharedStepsAPI* | [**UpdateSharedStep**](docs/SharedStepsAPI.md#updatesharedstep) | **Patch** /shared_step/{code}/{hash} | Update shared step
*SuitesAPI* | [**CreateSuite**](docs/SuitesAPI.md#createsuite) | **Post** /suite/{code} | Create a new test suite
*SuitesAPI* | [**DeleteSuite**](docs/SuitesAPI.md#deletesuite) | **Delete** /suite/{code}/{id} | Delete test suite
*SuitesAPI* | [**GetSuite**](docs/SuitesAPI.md#getsuite) | **Get** /suite/{code}/{id} | Get a specific test suite
*SuitesAPI* | [**GetSuites**](docs/SuitesAPI.md#getsuites) | **Get** /suite/{code} | Get all test suites
*SuitesAPI* | [**UpdateSuite**](docs/SuitesAPI.md#updatesuite) | **Patch** /suite/{code}/{id} | Update test suite
*SystemFieldsAPI* | [**GetSystemFields**](docs/SystemFieldsAPI.md#getsystemfields) | **Get** /system_field | Get all System Fields


## Documentation For Models

 - [Attachment](docs/Attachment.md)
 - [AttachmentGet](docs/AttachmentGet.md)
 - [AttachmentHash](docs/AttachmentHash.md)
 - [AttachmentListResponse](docs/AttachmentListResponse.md)
 - [AttachmentListResponseAllOfResult](docs/AttachmentListResponseAllOfResult.md)
 - [AttachmentResponse](docs/AttachmentResponse.md)
 - [AttachmentUploadsResponse](docs/AttachmentUploadsResponse.md)
 - [Attachmentupload](docs/Attachmentupload.md)
 - [Author](docs/Author.md)
 - [AuthorListResponse](docs/AuthorListResponse.md)
 - [AuthorListResponseAllOfResult](docs/AuthorListResponseAllOfResult.md)
 - [AuthorResponse](docs/AuthorResponse.md)
 - [BaseResponse](docs/BaseResponse.md)
 - [Bulk200Response](docs/Bulk200Response.md)
 - [Bulk200ResponseAllOfResult](docs/Bulk200ResponseAllOfResult.md)
 - [Configuration](docs/Configuration.md)
 - [ConfigurationCreate](docs/ConfigurationCreate.md)
 - [ConfigurationGroup](docs/ConfigurationGroup.md)
 - [ConfigurationGroupCreate](docs/ConfigurationGroupCreate.md)
 - [ConfigurationListResponse](docs/ConfigurationListResponse.md)
 - [ConfigurationListResponseAllOfResult](docs/ConfigurationListResponseAllOfResult.md)
 - [CustomField](docs/CustomField.md)
 - [CustomFieldCreate](docs/CustomFieldCreate.md)
 - [CustomFieldCreateValueInner](docs/CustomFieldCreateValueInner.md)
 - [CustomFieldListResponse](docs/CustomFieldListResponse.md)
 - [CustomFieldResponse](docs/CustomFieldResponse.md)
 - [CustomFieldUpdate](docs/CustomFieldUpdate.md)
 - [CustomFieldValue](docs/CustomFieldValue.md)
 - [CustomFieldsResponse](docs/CustomFieldsResponse.md)
 - [CustomFieldsResponseAllOfResult](docs/CustomFieldsResponseAllOfResult.md)
 - [Defect](docs/Defect.md)
 - [DefectCreate](docs/DefectCreate.md)
 - [DefectListResponse](docs/DefectListResponse.md)
 - [DefectListResponseAllOfResult](docs/DefectListResponseAllOfResult.md)
 - [DefectQuery](docs/DefectQuery.md)
 - [DefectResponse](docs/DefectResponse.md)
 - [DefectStatus](docs/DefectStatus.md)
 - [DefectUpdate](docs/DefectUpdate.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentCreate](docs/EnvironmentCreate.md)
 - [EnvironmentListResponse](docs/EnvironmentListResponse.md)
 - [EnvironmentListResponseAllOfResult](docs/EnvironmentListResponseAllOfResult.md)
 - [EnvironmentResponse](docs/EnvironmentResponse.md)
 - [EnvironmentUpdate](docs/EnvironmentUpdate.md)
 - [ExternalIssue](docs/ExternalIssue.md)
 - [ExternalIssueIssuesInner](docs/ExternalIssueIssuesInner.md)
 - [HashResponse](docs/HashResponse.md)
 - [HashResponseAllOfResult](docs/HashResponseAllOfResult.md)
 - [IdResponse](docs/IdResponse.md)
 - [IdResponseAllOfResult](docs/IdResponseAllOfResult.md)
 - [Milestone](docs/Milestone.md)
 - [MilestoneCreate](docs/MilestoneCreate.md)
 - [MilestoneListResponse](docs/MilestoneListResponse.md)
 - [MilestoneListResponseAllOfResult](docs/MilestoneListResponseAllOfResult.md)
 - [MilestoneResponse](docs/MilestoneResponse.md)
 - [MilestoneUpdate](docs/MilestoneUpdate.md)
 - [Plan](docs/Plan.md)
 - [PlanCreate](docs/PlanCreate.md)
 - [PlanDetailed](docs/PlanDetailed.md)
 - [PlanDetailedAllOfCases](docs/PlanDetailedAllOfCases.md)
 - [PlanListResponse](docs/PlanListResponse.md)
 - [PlanListResponseAllOfResult](docs/PlanListResponseAllOfResult.md)
 - [PlanQuery](docs/PlanQuery.md)
 - [PlanResponse](docs/PlanResponse.md)
 - [PlanUpdate](docs/PlanUpdate.md)
 - [Project](docs/Project.md)
 - [ProjectAccess](docs/ProjectAccess.md)
 - [ProjectCodeResponse](docs/ProjectCodeResponse.md)
 - [ProjectCodeResponseAllOfResult](docs/ProjectCodeResponseAllOfResult.md)
 - [ProjectCounts](docs/ProjectCounts.md)
 - [ProjectCountsDefects](docs/ProjectCountsDefects.md)
 - [ProjectCountsRuns](docs/ProjectCountsRuns.md)
 - [ProjectCreate](docs/ProjectCreate.md)
 - [ProjectListResponse](docs/ProjectListResponse.md)
 - [ProjectListResponseAllOfResult](docs/ProjectListResponseAllOfResult.md)
 - [ProjectResponse](docs/ProjectResponse.md)
 - [QqlDefect](docs/QqlDefect.md)
 - [QqlPlan](docs/QqlPlan.md)
 - [QqlTestCase](docs/QqlTestCase.md)
 - [Requirement](docs/Requirement.md)
 - [Response](docs/Response.md)
 - [Result](docs/Result.md)
 - [ResultCreate](docs/ResultCreate.md)
 - [ResultCreateBulk](docs/ResultCreateBulk.md)
 - [ResultCreateCase](docs/ResultCreateCase.md)
 - [ResultCreateResponse](docs/ResultCreateResponse.md)
 - [ResultCreateResponseAllOfResult](docs/ResultCreateResponseAllOfResult.md)
 - [ResultListResponse](docs/ResultListResponse.md)
 - [ResultListResponseAllOfResult](docs/ResultListResponseAllOfResult.md)
 - [ResultResponse](docs/ResultResponse.md)
 - [ResultUpdate](docs/ResultUpdate.md)
 - [ResultcreateBulk](docs/ResultcreateBulk.md)
 - [Run](docs/Run.md)
 - [RunCreate](docs/RunCreate.md)
 - [RunEnvironment](docs/RunEnvironment.md)
 - [RunListResponse](docs/RunListResponse.md)
 - [RunListResponseAllOfResult](docs/RunListResponseAllOfResult.md)
 - [RunMilestone](docs/RunMilestone.md)
 - [RunPublic](docs/RunPublic.md)
 - [RunPublicResponse](docs/RunPublicResponse.md)
 - [RunPublicResponseAllOfResult](docs/RunPublicResponseAllOfResult.md)
 - [RunResponse](docs/RunResponse.md)
 - [RunStats](docs/RunStats.md)
 - [SearchResponse](docs/SearchResponse.md)
 - [SearchResponseAllOfResult](docs/SearchResponseAllOfResult.md)
 - [SearchResponseAllOfResultEntities](docs/SearchResponseAllOfResultEntities.md)
 - [SharedStep](docs/SharedStep.md)
 - [SharedStepContent](docs/SharedStepContent.md)
 - [SharedStepContentCreate](docs/SharedStepContentCreate.md)
 - [SharedStepCreate](docs/SharedStepCreate.md)
 - [SharedStepListResponse](docs/SharedStepListResponse.md)
 - [SharedStepListResponseAllOfResult](docs/SharedStepListResponseAllOfResult.md)
 - [SharedStepResponse](docs/SharedStepResponse.md)
 - [SharedStepUpdate](docs/SharedStepUpdate.md)
 - [Suite](docs/Suite.md)
 - [SuiteCreate](docs/SuiteCreate.md)
 - [SuiteDelete](docs/SuiteDelete.md)
 - [SuiteListResponse](docs/SuiteListResponse.md)
 - [SuiteListResponseAllOfResult](docs/SuiteListResponseAllOfResult.md)
 - [SuiteResponse](docs/SuiteResponse.md)
 - [SuiteUpdate](docs/SuiteUpdate.md)
 - [SystemField](docs/SystemField.md)
 - [SystemFieldListResponse](docs/SystemFieldListResponse.md)
 - [SystemFieldOption](docs/SystemFieldOption.md)
 - [TagValue](docs/TagValue.md)
 - [TestCase](docs/TestCase.md)
 - [TestCaseCreate](docs/TestCaseCreate.md)
 - [TestCaseExternalIssues](docs/TestCaseExternalIssues.md)
 - [TestCaseExternalIssuesLinksInner](docs/TestCaseExternalIssuesLinksInner.md)
 - [TestCaseListResponse](docs/TestCaseListResponse.md)
 - [TestCaseListResponseAllOfResult](docs/TestCaseListResponseAllOfResult.md)
 - [TestCaseParams](docs/TestCaseParams.md)
 - [TestCaseQuery](docs/TestCaseQuery.md)
 - [TestCaseResponse](docs/TestCaseResponse.md)
 - [TestCaseUpdate](docs/TestCaseUpdate.md)
 - [TestCasebulk](docs/TestCasebulk.md)
 - [TestCasebulkCasesInner](docs/TestCasebulkCasesInner.md)
 - [TestCaseexternalIssues](docs/TestCaseexternalIssues.md)
 - [TestStep](docs/TestStep.md)
 - [TestStepCreate](docs/TestStepCreate.md)
 - [TestStepResult](docs/TestStepResult.md)
 - [TestStepResultCreate](docs/TestStepResultCreate.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### TokenAuth

- **Type**: API key
- **API key parameter name**: Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Token and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		api_v1_client.ContextAPIKeys,
		map[string]api_v1_client.APIKey{
			"Token": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@qase.io

