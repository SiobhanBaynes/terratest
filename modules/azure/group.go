package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	"github.com/stretchr/testify/require"
)

// GetResourceGroup is a helper function that will check the existence of a resource group
func GetResourceGroup(t *testing.T, resGroupName string, subscriptionID string) *resources.Group {
	resGroup, err := GetResourceGroupE(t, resGroupName, subscriptionID)
	require.NoError(t, err)

	return resGroup
}

// GetResourceGroupE is a helper function that will check the existence of a resource group
func GetResourceGroupE(t *testing.T, resGroupName string, subscriptionID string) (*resources.Group, error) {

	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}
	resGroupName, err = getTargetAzureResourceGroupName(resGroupName)
	if err != nil {
		return nil, err
	}

	groupsClient := resources.NewGroupsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	groupsClient.Authorizer = *authorizer
	resourceGroup, err := groupsClient.Get(context.Background(), resGroupName)
	if err != nil {
		return nil, err
	}
	return &resourceGroup, nil
}
