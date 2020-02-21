package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	"github.com/stretchr/testify/require"
)

// GetStorageAccount is a helper function that will check the existence of a storage account
func GetStorageAccount(t *testing.T, resGroupName string, storageAccountName string, subscriptionID string) *storage.Account {
	storageAccount, err := GetStorageAccountE(t, resGroupName, storageAccountName, subscriptionID)
	require.NoError(t, err)

	return storageAccount
}

// GetStorageAccountE is a helper function that will check the existence of a storage account
func GetStorageAccountE(t *testing.T, resGroupName string, storageAccountName string, subscriptionID string) (*storage.Account, error) {

	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}
	resGroupName, err = getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	accountsClient := storage.NewAccountsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	accountsClient.Authorizer = *authorizer
	var accountExpand storage.AccountExpand
	storageAccount, err := accountsClient.GetProperties(context.Background(), resGroupName, storageAccountName, accountExpand)
	if err != nil {
		return nil, err
	}
	return &storageAccount, nil
}
