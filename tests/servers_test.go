package tests

import (
	"github.com/mailosaur/mailosaur-go/mailosaur"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServersList(t *testing.T) {
	client := createTestClient()
	result, err := client.Servers().List()
	assert.Nil(t, err)
	assert.True(t, len(result.Items) > 0)
}

func TestServersGetNotFound(t *testing.T) {
	client := createTestClient()

	// Should fail if server is not found
	result, err := client.Servers().Get("efe907e9-74ed-4113-a3e0-a3d41d914765")
	assert.NotNil(t, err)
	_, ok := err.(mailosaur.MailosaurException)
	assert.True(t, ok)
	assert.Nil(t, result)
}

func TestServersCRUD(t *testing.T) {
	client := createTestClient()
	serverName := "My test"

	// Create a new server
	options := mailosaur.ServerCreateOptions{Name: serverName}
	createdServer, err := client.Servers().Create(options)
	assert.Nil(t, err)
	assert.NotNil(t, createdServer)
	assert.NotEmpty(t, createdServer.Id)
	assert.Equal(t, serverName, createdServer.Name)
	assert.NotEmpty(t, createdServer.Password)
	assert.NotNil(t, createdServer.Users)
	assert.Equal(t, 0, createdServer.Messages)
	assert.NotNil(t, createdServer.ForwardingRules)

	// Retrieve a server and confirm it has expected content
	retrievedServer, err := client.Servers().Get(createdServer.Id)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedServer)
	assert.Equal(t, createdServer.Id, retrievedServer.Id)
	assert.Equal(t, createdServer.Name, retrievedServer.Name)
	assert.Equal(t, createdServer.Password, retrievedServer.Password)
	assert.NotNil(t, retrievedServer.Users)
	assert.Equal(t, 0, retrievedServer.Messages)
	assert.NotNil(t, retrievedServer.ForwardingRules)

	// Update a server and confirm it has changed
	retrievedServer.Name += " EDITED"
	updatedServer, err := client.Servers().Update(retrievedServer.Id, retrievedServer)
	assert.Nil(t, err)
	assert.NotNil(t, updatedServer)
	assert.Equal(t, retrievedServer.Id, updatedServer.Id)
	assert.Equal(t, retrievedServer.Name, updatedServer.Name)
	assert.Equal(t, retrievedServer.Password, updatedServer.Password)
	assert.Equal(t, retrievedServer.Users, updatedServer.Users)
	assert.Equal(t, retrievedServer.Messages, updatedServer.Messages)
	assert.Equal(t, retrievedServer.ForwardingRules, updatedServer.ForwardingRules)

	err = client.Servers().Delete(retrievedServer.Id)
	assert.Nil(t, err)

	// Attempting to delete again should fail
	err = client.Servers().Delete(retrievedServer.Id)
	assert.NotNil(t, err)
	_, ok := err.(mailosaur.MailosaurException)
	assert.True(t, ok)
}

func TestServersFailedCreateTest(t *testing.T) {
	client := createTestClient()
	options := mailosaur.ServerCreateOptions{}
	_, err := client.Servers().Create(options)
	assert.NotNil(t, err)
	ex, ok := err.(mailosaur.MailosaurException)
	assert.True(t, ok)
	assert.Equal(t, "Operation returned an invalid status code '400 Bad Request'", ex.Message)
	assert.Equal(t, "ValidationError", ex.MailosaurError.Type)
	assert.Equal(t, 1, len(ex.MailosaurError.Messages))
	assert.NotEmpty(t, ex.MailosaurError.Messages["name"])
}
