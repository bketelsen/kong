package kong

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type Consumer struct {
	CreatedAt int    `json:"created_at"`
	ID        string `json:"id"`
	CustomID  string `json:"custom_id"`
	Username  string `json:"username"`
}
type ConsumerList struct {
	Data  []Consumer `json:"data"`
	Next  string     `json:"next"`
	Total int        `json:"total"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ConsumerService struct {
	sling *sling.Sling
}

func NewConsumerService(httpClient *http.Client, baseUrl string) *ConsumerService {
	return &ConsumerService{
		sling: sling.New().Client(httpClient).Base(baseUrl),
	}
}

func (s *ConsumerService) Get(name string) (*Consumer, *http.Response, error) {
	consumer := new(Consumer)
	resp, err := s.sling.New().Path("/consumers").Path(name).ReceiveSuccess(consumer)
	return consumer, resp, err
}

func (s *ConsumerService) List() (*ConsumerList, *http.Response, error) {
	consumers := new(ConsumerList)
	resp, err := s.sling.New().Path("/consumers/").ReceiveSuccess(consumers)
	return consumers, resp, err
}

func (s *ConsumerService) Create(customID, username string) (*Consumer, *http.Response, error) {
	consumer := new(Consumer)
	consumer.CustomID = customID
	consumer.Username = username
	resp, err := s.sling.New().Post("/consumers/").BodyJSON(consumer).ReceiveSuccess(consumer)
	return consumer, resp, err
}

func (s *ConsumerService) BasicAuth(cons, username, password string) (*Consumer, *http.Response, error) {
	creds := new(Credentials)
	creds.Username = username
	creds.Password = password
	consumer := new(Consumer)
	path := fmt.Sprintf("%s/basic-auth", cons)
	resp, err := s.sling.New().Post("/consumers/").Path(path).BodyJSON(creds).ReceiveSuccess(consumer)
	return consumer, resp, err
}
