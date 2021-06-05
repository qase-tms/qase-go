package examples

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.qase.io/client"
)

func Example_projects() {
	cfg := client.NewConfiguration()
	cl := client.NewAPIClient(cfg)

	auth := client.APIKey{Key: os.Getenv("API_TOKEN")}
	ctx, cancel := context.WithTimeout(
		context.WithValue(context.Background(), client.ContextAPIKey, auth),
		2*time.Second, // TODO:performance regress
	)
	defer cancel()

	opt := new(client.ProjectsApiGetProjectsOpts)
	list, resp, err := cl.ProjectsApi.GetProjects(ctx, opt)
	if err != nil || resp == nil || !list.Status {
		return
	}

	fmt.Println("success")
	// output: success
}
