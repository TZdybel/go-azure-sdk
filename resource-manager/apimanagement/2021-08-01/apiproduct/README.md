
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2021-08-01/apiproduct` Documentation

The `apiproduct` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2021-08-01/apiproduct"
```


### Client Initialization

```go
client := apiproduct.NewApiProductClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApiProductClient.ListByApis`

```go
ctx := context.TODO()
id := apiproduct.NewApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue")

// alternatively `client.ListByApis(ctx, id, apiproduct.DefaultListByApisOperationOptions())` can be used to do batched pagination
items, err := client.ListByApisComplete(ctx, id, apiproduct.DefaultListByApisOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
