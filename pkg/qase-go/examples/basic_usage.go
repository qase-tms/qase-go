package main

import (
	"fmt"
	"log"

	"github.com/qase-tms/qase-go/config"
)

func main() {
	fmt.Println("=== Qase Go Configuration Examples ===\n")

	// Example 1: Basic loading with defaults
	fmt.Println("1. Loading default configuration...")
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Error loading config: %v", err)
	} else {
		fmt.Printf("   Mode: %s\n", cfg.Mode)
		fmt.Printf("   Debug: %t\n", cfg.Debug)
		fmt.Printf("   Report Path: %s\n", cfg.Report.Connection.Local.Path)
		fmt.Printf("   Report Format: %s\n", cfg.Report.Connection.Local.Format)
	}

	// Example 2: Using Builder Pattern
	fmt.Println("\n2. Using Builder Pattern...")
	builderCfg, err := config.NewConfigBuilder().
		TestOpsMode().
		WithAPIToken("demo-token-12345").
		WithProject("DEMO").
		WithDebug(true).
		AddRunTag("smoke").
		AddRunTag("regression").
		AddRunConfiguration("browser", "chrome").
		AddRunConfiguration("environment", "staging").
		WithBatchSize(50).
		Build()

	if err != nil {
		log.Printf("Error building config: %v", err)
	} else {
		fmt.Printf("   Mode: %s\n", builderCfg.Mode)
		fmt.Printf("   Project: %s\n", builderCfg.TestOps.Project)
		fmt.Printf("   API Token: %s...\n", builderCfg.TestOps.API.Token[:10])
		fmt.Printf("   Tags: %v\n", builderCfg.TestOps.Run.Tags)
		fmt.Printf("   Batch Size: %d\n", builderCfg.TestOps.Batch.Size)
		fmt.Printf("   Configurations: %d\n", len(builderCfg.TestOps.Run.Configurations.Values))
	}

	// Example 3: Creating default config file
	fmt.Println("\n3. Creating example configuration file...")
	err = config.CreateDefaultConfigAt("./qase.config.example.json")
	if err != nil {
		log.Printf("Error creating config file: %v", err)
	} else {
		fmt.Println("   ✓ Created qase.config.example.json")
	}

	// Example 4: Loading from file
	fmt.Println("\n4. Loading from configuration file...")
	fileCfg, err := config.LoadFrom("./qase.config.example.json")
	if err != nil {
		log.Printf("Error loading from file: %v", err)
	} else {
		fmt.Printf("   Mode: %s\n", fileCfg.Mode)
		fmt.Printf("   Project: %s\n", fileCfg.TestOps.Project)
		fmt.Printf("   Host: %s\n", fileCfg.TestOps.API.Host)
	}

	// Example 5: Environment variable demonstration
	fmt.Println("\n5. Environment variable override example...")
	// Simulate environment variables
	envCfg := config.NewConfig()
	envCfg.Mode = "report" // Default from file/code
	
	fmt.Printf("   Before env loading - Mode: %s\n", envCfg.Mode)
	
	// In real usage, set environment variable like:
	// export QASE_MODE=testops
	// export QASE_DEBUG=true
	fmt.Println("   (Set QASE_MODE=testops and QASE_DEBUG=true to see override effect)")

	// Example 6: Configuration modes
	fmt.Println("\n6. Configuration modes demonstration...")
	
	modes := []string{"testops", "report", "off"}
	for _, mode := range modes {
		fmt.Printf("   Mode '%s':\n", mode)
		switch mode {
		case "testops":
			fmt.Println("     - Results sent to Qase TestOps")
			fmt.Println("     - Requires API token and project code")
		case "report":
			fmt.Println("     - Results saved to local files")
			fmt.Println("     - JSON or YAML format supported")
		case "off":
			fmt.Println("     - Reporter disabled")
			fmt.Println("     - No results collected or sent")
		}
	}

	fmt.Println("\n=== Configuration Examples Complete ===")
}