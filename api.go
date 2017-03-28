package kong

import (
	"net/http"

	"github.com/dghubble/sling"
)

type API struct {
	Id           string `json:"id,omitempty"`
	CreatedAt    int    `json:"created_at,omitempty"`
	Name         string `json:"name,omitempty"`
	PreserveHost bool   `json:"preserve_host,omitempty"`
	Hosts        string `json:"hosts,omitempty"`
	Uris         string `json:"uris,omitempty"`
	StripUri     bool   `json:"strip_uri,omitempty"`
	UpstreamURL  string `json:"upstream_url"`
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

func (s *APIService) Add(name, Hosts, Uris, upstreamURL string, stripUri, preserveHost bool) (*API, *http.Response, error) {
	api := new(API)
	if len(name) > 0 {
		api.Name = name
	}
	if len(Hosts) > 0 {
		api.Hosts = Hosts
	}
	if len(Uris) > 0 {
		api.Uris = Uris
	}
	api.UpstreamURL = upstreamURL
	api.StripUri = stripUri
	api.PreserveHost = preserveHost

	resp, err := s.sling.New().Post("/apis/").BodyJSON(api).ReceiveSuccess(api)

	return api, resp, err
}

func (s *APIService) AddOrUpdate(id, name, Hosts, Uris, upstreamURL string, stripUri, preserveHost bool, createdAt int) (*API, *http.Response, error) {
	api := new(API)
	if len(id) > 0 {
		api.Id = id
	}
	if len(name) > 0 {
		api.Name = name
	}
	if len(Hosts) > 0 {
		api.Hosts = Hosts
	}
	if len(Uris) > 0 {
		api.Uris = Uris
	}
	api.UpstreamURL = upstreamURL
	api.StripUri = stripUri
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

func (s *APIService) AddPlugin(name string, plugin interface{}) (*http.Response, error) {
	resp, err := s.sling.New().Post("/apis/"+name+"/plugins").BodyForm(plugin).Receive(nil, nil)
	return resp, err
}
