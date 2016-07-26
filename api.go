package kong

import (
	"net/http"

	"github.com/dghubble/sling"
)

type API struct {
	Id               string `json:"id,omitempty"`
	CreatedAt        int    `json:"created_at,omitempty"`
	Name             string `json:"name,omitempty"`
	PreserveHost     bool   `json:"preserve_host,omitempty"`
	RequestHost      string `json:"request_host,omitempty"`
	RequestPath      string `json:"request_path,omitempty"`
	StripRequestPath bool   `json:"strip_request_path,omitempty"`
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
	resp, err := s.sling.New().Path("/apis/").Path(name).ReceiveSuccess(api)
	return api, resp, err
}

func (s *APIService) List() (*APIList, *http.Response, error) {
	apis := new(APIList)
	resp, err := s.sling.New().Path("/apis/").ReceiveSuccess(apis)
	return apis, resp, err
}

func (s *APIService) Add(name, requestHost, requestPath, upstreamURL string, stripRequestPath, preserveHost bool) (*API, *http.Response, error) {
	api := new(API)
	if len(name) > 0 {
		api.Name = name
	}
	if len(requestHost) > 0 {
		api.RequestHost = requestHost
	}
	if len(requestPath) > 0 {
		api.RequestPath = requestPath
	}
	api.UpstreamURL = upstreamURL
	api.StripRequestPath = stripRequestPath
	api.PreserveHost = preserveHost

	resp, err := s.sling.New().Post("/apis/").BodyJSON(api).ReceiveSuccess(api)

	return api, resp, err
}

func (s *APIService) AddOrUpdate(id, name, requestHost, requestPath, upstreamURL string, stripRequestPath, preserveHost bool, createdAt int) (*API, *http.Response, error) {
	api := new(API)
	if len(id) > 0 {
		api.Id = id
	}
	if len(name) > 0 {
		api.Name = name
	}
	if len(requestHost) > 0 {
		api.RequestHost = requestHost
	}
	if len(requestPath) > 0 {
		api.RequestPath = requestPath
	}
	api.UpstreamURL = upstreamURL
	api.StripRequestPath = stripRequestPath
	api.PreserveHost = preserveHost

	if createdAt > 0 {
		api.CreatedAt = createdAt
	}

	resp, err := s.sling.New().Put("/apis").BodyJSON(api).ReceiveSuccess(api)

	return api, resp, err
}

func (s *APIService) Delete(name string) (*http.Response, error) {
	resp, err := s.sling.New().Delete("/apis/").Path(name).Receive(nil, nil)
	return resp, err
}
