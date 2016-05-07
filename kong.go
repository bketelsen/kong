package kong

import "net/http"

// Client is a tiny Kong client
type Client struct {
	BaseURL         string
	NodeService     *NodeService
	ClusterService  *ClusterService
	APIService      *APIService
	ConsumerService *ConsumerService
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, baseUrl string) *Client {
	return &Client{
		BaseURL:         baseUrl,
		NodeService:     NewNodeService(httpClient, baseUrl),
		ClusterService:  NewClusterService(httpClient, baseUrl),
		APIService:      NewAPIService(httpClient, baseUrl),
		ConsumerService: NewConsumerService(httpClient, baseUrl),
	}
}
