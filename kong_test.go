package kong

import (
	"testing"

	"github.com/kr/pretty"
)

func getClient() *Client {

	return NewClient(nil, "http://localhost:38001/")
}
func TestNodeInformation(t *testing.T) {
	client := getClient()
	nodes, _, _ := client.NodeService.Information()
	pretty.Println(nodes)

}

func TestNodeStatus(t *testing.T) {
	client := getClient()
	status, _, _ := client.NodeService.Status()
	pretty.Println(status)

}

func TestClusterStatus(t *testing.T) {
	client := getClient()
	status, _, _ := client.ClusterService.Status()
	pretty.Println(status)

}

func TestListAPI(t *testing.T) {
	client := getClient()
	apis, _, _ := client.APIService.List()
	pretty.Println(apis)

}

func TestListConsumers(t *testing.T) {
	client := getClient()
	consumers, _, _ := client.ConsumerService.List()
	pretty.Println(consumers)

}

func TestCreateConsumer(t *testing.T) {
	client := getClient()
	consumer, _, _ := client.ConsumerService.Create("2", "paul@atreides.com")
	pretty.Println(consumer)

	consumer, resp, err := client.ConsumerService.BasicAuth(consumer.ID, "newusername", "newpass")
	pretty.Println(consumer)
	pretty.Println(resp)
	pretty.Println(err)

}
