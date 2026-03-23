# ZSphere SDK for Go



## ZSphere SDK for Go

This repository provides a Go SDK for interacting with the ZSphere platform API. It allows developers to manage ZSphere resources programmatically using Go.

### Features

- Complete API coverage for ZSphere platform
- Type-safe Go client for ZSphere API operations
- Support for authentication and session management
- Structured request and response handling
- Error handling and logging capabilities
- Compatible with Terraform provider development

### Installation

```bash
go get github.com/terraform-zstack-modules/zsphere-sdk-go
```

### Usage

#### Initialize Client with Account Login

```go
import (
    "github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/client"
    "github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// Create a client with account login authentication
accountLoginClient := client.NewZSClient(
    client.DefaultZSConfig("ZStack MN HOST IP").
        LoginAccount("admin", "password").
        ReadOnly(false).
        Debug(true),
)
ctx := context.Background()
// Login to ZStack
_, err := accountLoginClient.Login(ctx)
if err != nil {
    // Handle error
}
defer accountLoginClient.Logout(ctx)

// Query clusters
clusters, err := accountLoginClient.QueryCluster(ctx, param.NewQueryParam())
if err != nil {
    // Handle error
}
// Use clusters data

```

#### Initialize Client with Access Key
```go
import (
    "github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/client"
    "github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// Create a client with access key authentication
accessKeyClient := client.NewZSClient(
    client.DefaultZSConfig("ZStack MN HOST IP").
        AccessKey("your-access-key-id", "your-access-key-secret").
        ReadOnly(false).
        Debug(false),
)
ctx := context.Background()
// Query clusters
clusters, err := accessKeyClient.QueryCluster(ctx, param.NewQueryParam())
if err != nil {
    // Handle error
}
// Use clusters data

```
### Authentication

The SDK supports the following authentication methods:
- Username and password authentication
- Session-based authentication with automatic token renewal
- API key authentication

### API Coverage

This SDK provides Go bindings for all major ZSphere API operations, including but not limited to:

- Compute resources (instances, volumes)
- Network resources (L2/L3 networks, security groups)
- Storage resources (primary storage, backup storage)
- Identity and access management
- System management operations


### Testing

To run tests and generate an HTML report:

**Windows (PowerShell):**
```powershell
./scripts/run_tests_with_report.ps1
```

**Linux (Bash):**
```bash
./scripts/run_tests_with_report.sh
```

### Contributing


Contributions to the ZSphere Go SDK are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

### Related Projects

- [ZSphere Documentation](https://www.zstack.io/help/zstack_zsphere/dev_manual/)

### Support

For issues, questions and discussions please use the [GitHub Issues](https://github.com/terraform-zstack-modules/zsphere-sdk-go/issues).

