package kong

import (
	"testing"

	"github.com/kr/pretty"
)

func getClient() *Client {

	return NewClient(nil, "http://10.13.172.10:8001/")
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

func TestAddAPI(t *testing.T) {
	client := getClient()
	addAPI, _, _ := client.APIService.Add("ApiTest", "", "/test", "http://localhost:8080/testing", false, false)

	getAPI, _, _ := client.APIService.Get("ApiTest")

	client.APIService.Delete("ApiTest")
	if addAPI.Id != getAPI.Id {
		t.Error("ApiTest added with ID ", addAPI.Id, "Get returned ", getAPI.Id)
	}
}

func TestUpdateAPI(t *testing.T) {
	client := getClient()
	addAPI, _, _ := client.APIService.Add("ApiTest", "", "/test", "http://localhost:8080/testing", false, false)
	pretty.Println(addAPI)

	updateAPI, _, _ := client.APIService.AddOrUpdate(addAPI.Id, "ApiTest", "", "/testupdate", "http://localhost:8080/testing", false, false, addAPI.CreatedAt)
	pretty.Println(updateAPI)

	getAPI, _, _ := client.APIService.Get("ApiTest")
	pretty.Println(getAPI)

	client.APIService.Delete("ApiTest")
	if addAPI.Uris == getAPI.Uris {
		t.Error("ApiTest added with RequestPath ", addAPI.Uris, "Get returned ", getAPI.Uris)
	}
}

func TestDeleteAPI(t *testing.T) {
	client := getClient()
	addAPI, _, _ := client.APIService.Add("ApiTest", "", "/test", "http://localhost:8080/testing", false, false)

	getAPI, _, _ := client.APIService.Get("ApiTest")

	client.APIService.Delete("ApiTest")

	getAPIAfterDelete, _, _ := client.APIService.Get("ApiTest")

	if addAPI.Id != getAPI.Id {
		t.Error("Error including API")
	}

	if len(getAPIAfterDelete.Name) > 0 {
		t.Error("Error deleting API")
	}
}

func TestAddPlugin(t *testing.T) {
	client := getClient()
	client.APIService.Add("ApiTest", "", "/test", "http://localhost:8080/testing", false, false)

	body := &RequestTransformer{
		Name:           "request-transformer",
		ReplaceHeaders: "Host:www.test.org",
	}

	resp, err := client.APIService.AddPlugin("ApiTest", body)

	if err != nil {
		t.Error("Error adding a plugin: " + err.Error())
	}

	if resp.StatusCode != 201 {
		t.Error("StatusCode different from 201: ", resp.StatusCode)
	}

	client.APIService.Delete("ApiTest")
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
