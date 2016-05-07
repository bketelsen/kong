package kong

import (
	"net/http"

	"github.com/dghubble/sling"
)

type API struct {
	CreatedAt        int    `json:"created_at"`
	Name             string `json:"name"`
	PreserveHost     bool   `json:"preserve_host"`
	RequestHost      string `json:"request_host"`
	RequestPath      string `json:"request_path"`
	StripRequestPath bool   `json:"strip_request_path"`
	UpstreamURL      string `json:"upstream_url"`
}
type APIList struct {
	Data  []API  `json:"data"`
	Next  string `json:"next"`
	Total int    `json:"total"`
}
type APIService struct {
	sling *sling.Sling
}

func NewAPIService(httpClient *http.Client, baseUrl string) *APIService {
	return &APIService{
		sling: sling.New().Client(httpClient).Base(baseUrl),
	}
}

func (s *APIService) Get(name string) (*API, *http.Response, error) {
	api := new(API)
	resp, err := s.sling.New().Path("/apis").Path(name).ReceiveSuccess(api)
	return api, resp, err
}

func (s *APIService) List() (*APIList, *http.Response, error) {
	apis := new(APIList)
	resp, err := s.sling.New().Path("/apis/").ReceiveSuccess(apis)
	return apis, resp, err
}
