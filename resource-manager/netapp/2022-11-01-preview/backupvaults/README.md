
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/backupvaults` Documentation

The `backupvaults` SDK allows for interaction with the Azure Resource Manager Service `netapp` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-11-01-preview/backupvaults"
```


### Client Initialization

```go
client := backupvaults.NewBackupVaultsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupVaultsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := backupvaults.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "backupVaultValue")

payload := backupvaults.BackupVault{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BackupVaultsClient.Delete`

```go
ctx := context.TODO()
id := backupvaults.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "backupVaultValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BackupVaultsClient.Get`

```go
ctx := context.TODO()
id := backupvaults.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "backupVaultValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BackupVaultsClient.ListByNetAppAccount`

```go
ctx := context.TODO()
id := backupvaults.NewNetAppAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue")

// alternatively `client.ListByNetAppAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByNetAppAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BackupVaultsClient.Update`

```go
ctx := context.TODO()
id := backupvaults.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "backupVaultValue")

payload := backupvaults.BackupVaultPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
