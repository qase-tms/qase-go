# Qase TestOps Reporters for Go

Monorepo containing Qase TestOps integration tools for Go testing frameworks. This repository includes reporters and API clients designed to streamline the process of reporting test results to Qase TestOps.

## Projects Overview

### Core SDK

* **[qase-go](pkg/qase-go)**  
Main Go SDK package providing comprehensive tools for integrating Go applications with Qase TestOps. Includes domain models, reporters, and testing utilities.

### API Clients

* **[qase-api-client](/qase-api-client)**  
Official client for interacting with the Qase TestOps API (v1). Recommended for most use cases.
* **[qase-api-v2-client](/qase-api-v2-client)**  
API client supporting the new Qase TestOps API (v2). Use this client for projects leveraging the latest API features.

### Examples

* **[examples/go](/examples/go)**  
Comprehensive examples demonstrating SDK usage, including basic integration, multi-package testing, and advanced configuration scenarios.

