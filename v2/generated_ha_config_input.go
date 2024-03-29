package client

const (
	HA_CONFIG_INPUT_TYPE = "haConfigInput"
)

type HaConfigInput struct {
	Resource

	Cert string `json:"cert,omitempty" yaml:"cert,omitempty"`

	CertChain string `json:"certChain,omitempty" yaml:"cert_chain,omitempty"`

	ClusterSize int64 `json:"clusterSize,omitempty" yaml:"cluster_size,omitempty"`

	HostRegistrationUrl string `json:"hostRegistrationUrl,omitempty" yaml:"host_registration_url,omitempty"`

	HttpEnabled bool `json:"httpEnabled,omitempty" yaml:"http_enabled,omitempty"`

	HttpPort int64 `json:"httpPort,omitempty" yaml:"http_port,omitempty"`

	HttpsPort int64 `json:"httpsPort,omitempty" yaml:"https_port,omitempty"`

	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	PpHttpPort int64 `json:"ppHttpPort,omitempty" yaml:"pp_http_port,omitempty"`

	PpHttpsPort int64 `json:"ppHttpsPort,omitempty" yaml:"pp_https_port,omitempty"`

	RedisPort int64 `json:"redisPort,omitempty" yaml:"redis_port,omitempty"`

	SwarmEnabled bool `json:"swarmEnabled,omitempty" yaml:"swarm_enabled,omitempty"`

	SwarmPort int64 `json:"swarmPort,omitempty" yaml:"swarm_port,omitempty"`

	ZookeeperClientPort int64 `json:"zookeeperClientPort,omitempty" yaml:"zookeeper_client_port,omitempty"`

	ZookeeperLeaderPort int64 `json:"zookeeperLeaderPort,omitempty" yaml:"zookeeper_leader_port,omitempty"`

	ZookeeperQuorumPort int64 `json:"zookeeperQuorumPort,omitempty" yaml:"zookeeper_quorum_port,omitempty"`
}

type HaConfigInputCollection struct {
	Collection
	Data   []HaConfigInput `json:"data,omitempty"`
	client *HaConfigInputClient
}

type HaConfigInputClient struct {
	kuladoClient *KuladoClient
}

type HaConfigInputOperations interface {
	List(opts *ListOpts) (*HaConfigInputCollection, error)
	Create(opts *HaConfigInput) (*HaConfigInput, error)
	Update(existing *HaConfigInput, updates interface{}) (*HaConfigInput, error)
	ById(id string) (*HaConfigInput, error)
	Delete(container *HaConfigInput) error
}

func newHaConfigInputClient(kuladoClient *KuladoClient) *HaConfigInputClient {
	return &HaConfigInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *HaConfigInputClient) Create(container *HaConfigInput) (*HaConfigInput, error) {
	resp := &HaConfigInput{}
	err := c.kuladoClient.doCreate(HA_CONFIG_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *HaConfigInputClient) Update(existing *HaConfigInput, updates interface{}) (*HaConfigInput, error) {
	resp := &HaConfigInput{}
	err := c.kuladoClient.doUpdate(HA_CONFIG_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HaConfigInputClient) List(opts *ListOpts) (*HaConfigInputCollection, error) {
	resp := &HaConfigInputCollection{}
	err := c.kuladoClient.doList(HA_CONFIG_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HaConfigInputCollection) Next() (*HaConfigInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HaConfigInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HaConfigInputClient) ById(id string) (*HaConfigInput, error) {
	resp := &HaConfigInput{}
	err := c.kuladoClient.doById(HA_CONFIG_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *HaConfigInputClient) Delete(container *HaConfigInput) error {
	return c.kuladoClient.doResourceDelete(HA_CONFIG_INPUT_TYPE, &container.Resource)
}
