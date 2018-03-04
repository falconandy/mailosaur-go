package examples

import (
	"github.com/mailosaur/mailosaur-go/mailosaur"
)

func serversList() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Servers().List()
}

func serversCreate() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	options := mailosaur.ServerCreateOptions{Name: "My Server"}
	client.Servers().Create(options)
}

func serversGet() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Servers().Get("SERVER_ID")
}

func serversUpdate() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	server, _ := client.Servers().Get("SERVER_ID")
	server.Name = "New server name"
	client.Servers().Update(server.Id, server)
}

func serversDelete() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Servers().Delete("SERVER_ID")
}
