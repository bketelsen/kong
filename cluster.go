package kong

import (
	"net/http"

	"github.com/dghubble/sling"
)

type ClusterStatus struct {
	Data []struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Status  string `json:"status"`
	} `json:"data"`
	Total int `json:"total"`
}
type ClusterService struct {
	sling *sling.Sling
}

// NewIssueService returns a new IssueService.
func NewClusterService(httpClient *http.Client, baseUrl string) *ClusterService {
	return &ClusterService{
		sling: sling.New().Client(httpClient).Base(baseUrl),
	}
}

func (s *ClusterService) Status() (*ClusterStatus, *http.Response, error) {
	status := new(ClusterStatus)
	resp, err := s.sling.New().Path("/cluster").ReceiveSuccess(status)
	return status, resp, err
}

/*
func (s *ClusterService) Remove() (*http.Response, error) {
	req, _ := s.sling.New().Delete("/cluster").Request()
	return s.sling.Do(req, nil, nil)
}
*/
