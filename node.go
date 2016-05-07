package kong

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Node struct {
	Configuration struct {
		AdminAPIListen string `json:"admin_api_listen"`
		Cassandra      struct {
			ContactPoints       []string `json:"contact_points"`
			DataCenters         struct{} `json:"data_centers"`
			Keyspace            string   `json:"keyspace"`
			ReplicationFactor   int      `json:"replication_factor"`
			ReplicationStrategy string   `json:"replication_strategy"`
			Ssl                 struct {
				Enabled bool `json:"enabled"`
				Verify  bool `json:"verify"`
			} `json:"ssl"`
			Timeout int `json:"timeout"`
		} `json:"cassandra"`
		Cluster struct {
			Auto_join bool   `json:"auto-join"`
			Profile   string `json:"profile"`
		} `json:"cluster"`
		ClusterListen    string   `json:"cluster_listen"`
		ClusterListenRPC string   `json:"cluster_listen_rpc"`
		CustomPlugins    struct{} `json:"custom_plugins"`
		DaoConfig        struct {
			ContactPoints       []string `json:"contact_points"`
			DataCenters         struct{} `json:"data_centers"`
			Keyspace            string   `json:"keyspace"`
			ReplicationFactor   int      `json:"replication_factor"`
			ReplicationStrategy string   `json:"replication_strategy"`
			Ssl                 struct {
				Enabled bool `json:"enabled"`
				Verify  bool `json:"verify"`
			} `json:"ssl"`
			Timeout int `json:"timeout"`
		} `json:"dao_config"`
		Database    string `json:"database"`
		DNSResolver struct {
			Address string `json:"address"`
			Dnsmasq bool   `json:"dnsmasq"`
			Port    int    `json:"port"`
		} `json:"dns_resolver"`
		DNSResolversAvailable struct {
			Dnsmasq struct {
				Port int `json:"port"`
			} `json:"dnsmasq"`
			Server struct {
				Address string `json:"address"`
			} `json:"server"`
		} `json:"dns_resolvers_available"`
		MemoryCacheSize      int      `json:"memory_cache_size"`
		Nginx                string   `json:"nginx"`
		NginxWorkingDir      string   `json:"nginx_working_dir"`
		PidFile              string   `json:"pid_file"`
		Plugins              []string `json:"plugins"`
		ProxyListen          string   `json:"proxy_listen"`
		ProxyListenSsl       string   `json:"proxy_listen_ssl"`
		SendAnonymousReports bool     `json:"send_anonymous_reports"`
	} `json:"configuration"`
	Hostname   string `json:"hostname"`
	LuaVersion string `json:"lua_version"`
	Plugins    struct {
		AvailableOnServer []string `json:"available_on_server"`
		EnabledInCluster  []string `json:"enabled_in_cluster"`
	} `json:"plugins"`
	Tagline string `json:"tagline"`
	Version string `json:"version"`
}

type NodeStatus struct {
	Database struct {
		Acls                        int `json:"acls"`
		Apis                        int `json:"apis"`
		BasicauthCredentials        int `json:"basicauth_credentials"`
		Consumers                   int `json:"consumers"`
		HmacauthCredentials         int `json:"hmacauth_credentials"`
		JwtSecrets                  int `json:"jwt_secrets"`
		KeyauthCredentials          int `json:"keyauth_credentials"`
		Nodes                       int `json:"nodes"`
		Oauth2AuthorizationCodes    int `json:"oauth2_authorization_codes"`
		Oauth2Credentials           int `json:"oauth2_credentials"`
		Oauth2Tokens                int `json:"oauth2_tokens"`
		Plugins                     int `json:"plugins"`
		RatelimitingMetrics         int `json:"ratelimiting_metrics"`
		ResponseRatelimitingMetrics int `json:"response_ratelimiting_metrics"`
	} `json:"database"`
	Server struct {
		ConnectionsAccepted int `json:"connections_accepted"`
		ConnectionsActive   int `json:"connections_active"`
		ConnectionsHandled  int `json:"connections_handled"`
		ConnectionsReading  int `json:"connections_reading"`
		ConnectionsWaiting  int `json:"connections_waiting"`
		ConnectionsWriting  int `json:"connections_writing"`
		TotalRequests       int `json:"total_requests"`
	} `json:"server"`
}

// IssueService provides methods for creating and reading issues.
type NodeService struct {
	sling *sling.Sling
}

// NewIssueService returns a new IssueService.
func NewNodeService(httpClient *http.Client, baseUrl string) *NodeService {
	return &NodeService{
		sling: sling.New().Client(httpClient).Base(baseUrl),
	}
}

func (s *NodeService) Information() (*Node, *http.Response, error) {
	node := new(Node)
	resp, err := s.sling.New().Path("/").ReceiveSuccess(node)
	return node, resp, err
}

func (s *NodeService) Status() (*NodeStatus, *http.Response, error) {
	nodestatus := new(NodeStatus)
	resp, err := s.sling.New().Path("/status").ReceiveSuccess(nodestatus)
	return nodestatus, resp, err
}
