
## `github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/deletedconfigurationstores` Documentation

The `deletedconfigurationstores` SDK allows for interaction with the Azure Resource Manager Service `appconfiguration` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/deletedconfigurationstores"
```


### Client Initialization

```go
client := deletedconfigurationstores.NewDeletedConfigurationStoresClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `DeletedConfigurationStoresClient.ConfigurationStoresGetDeleted`

```go
ctx := context.TODO()
id := deletedconfigurationstores.NewDeletedConfigurationStoreID("12345678-1234-9876-4563-123456789012", "locationValue", "configStoreValue")
read, err := client.ConfigurationStoresGetDeleted(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedConfigurationStoresClient.ConfigurationStoresListDeleted`

```go
ctx := context.TODO()
id := deletedconfigurationstores.NewSubscriptionID()
// alternatively `client.ConfigurationStoresListDeleted(ctx, id)` can be used to do batched pagination
items, err := client.ConfigurationStoresListDeletedComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedConfigurationStoresClient.ConfigurationStoresPurgeDeleted`

```go
ctx := context.TODO()
id := deletedconfigurationstores.NewDeletedConfigurationStoreID("12345678-1234-9876-4563-123456789012", "locationValue", "configStoreValue")
future, err := client.ConfigurationStoresPurgeDeleted(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```