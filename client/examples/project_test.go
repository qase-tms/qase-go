package examples

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.qase.io/client"
)

func Example_projects() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cfg := client.NewConfiguration()
	cl := client.NewAPIClient(cfg)
	opt := new(client.ProjectsApiGetProjectsOpts)

	cfg.AddDefaultHeader("Token", os.Getenv("API_TOKEN"))

	list, resp, err := cl.ProjectsApi.GetProjects(ctx, opt)
	if err != nil || resp == nil || !list.Status {
		return
	}

	fmt.Println("success")
	// output: success
}
