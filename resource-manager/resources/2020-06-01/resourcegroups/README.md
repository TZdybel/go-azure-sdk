
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/resourcegroups` Documentation

The `resourcegroups` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2020-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2020-06-01/resourcegroups"
```


### Client Initialization

```go
client := resourcegroups.NewResourceGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ResourceGroupsClient.CheckExistence`

```go
ctx := context.TODO()
id := resourcegroups.NewResourceGroupID()
read, err := client.CheckExistence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := resourcegroups.NewResourceGroupID()

payload := resourcegroups.ResourceGroup{
	// ...
}

read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGroupsClient.Delete`

```go
ctx := context.TODO()
id := resourcegroups.NewResourceGroupID()
future, err := client.Delete(ctx, id, resourcegroups.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ResourceGroupsClient.ExportTemplate`

```go
ctx := context.TODO()
id := resourcegroups.NewResourceGroupID()

payload := resourcegroups.ExportTemplateRequest{
	// ...
}

future, err := client.ExportTemplate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ResourceGroupsClient.Get`

```go
ctx := context.TODO()
id := resourcegroups.NewResourceGroupID()
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGroupsClient.List`

```go
ctx := context.TODO()
id := resourcegroups.NewSubscriptionID()
// alternatively `client.List(ctx, id, resourcegroups.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, resourcegroups.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGroupsClient.ResourcesListByResourceGroup`

```go
ctx := context.TODO()
id := resourcegroups.NewResourceGroupID()
// alternatively `client.ResourcesListByResourceGroup(ctx, id, resourcegroups.DefaultResourcesListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ResourcesListByResourceGroupComplete(ctx, id, resourcegroups.DefaultResourcesListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGroupsClient.Update`

```go
ctx := context.TODO()
id := resourcegroups.NewResourceGroupID()

payload := resourcegroups.ResourceGroupPatchable{
	// ...
}

read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```