package client

const (
	K8S_CLIENT_CONFIG_TYPE = "k8sClientConfig"
)

type K8sClientConfig struct {
	Resource `yaml:"-"`

	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	BearerToken string `json:"bearerToken,omitempty" yaml:"bearer_token,omitempty"`

	CaCert string `json:"caCert,omitempty" yaml:"ca_cert,omitempty"`
}

type K8sClientConfigCollection struct {
	Collection
	Data   []K8sClientConfig `json:"data,omitempty"`
	client *K8sClientConfigClient
}

type K8sClientConfigClient struct {
	kuladoClient *KuladoClient
}

type K8sClientConfigOperations interface {
	List(opts *ListOpts) (*K8sClientConfigCollection, error)
	Create(opts *K8sClientConfig) (*K8sClientConfig, error)
	Update(existing *K8sClientConfig, updates interface{}) (*K8sClientConfig, error)
	ById(id string) (*K8sClientConfig, error)
	Delete(container *K8sClientConfig) error
}

func newK8sClientConfigClient(kuladoClient *KuladoClient) *K8sClientConfigClient {
	return &K8sClientConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *K8sClientConfigClient) Create(container *K8sClientConfig) (*K8sClientConfig, error) {
	resp := &K8sClientConfig{}
	err := c.kuladoClient.doCreate(K8S_CLIENT_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *K8sClientConfigClient) Update(existing *K8sClientConfig, updates interface{}) (*K8sClientConfig, error) {
	resp := &K8sClientConfig{}
	err := c.kuladoClient.doUpdate(K8S_CLIENT_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *K8sClientConfigClient) List(opts *ListOpts) (*K8sClientConfigCollection, error) {
	resp := &K8sClientConfigCollection{}
	err := c.kuladoClient.doList(K8S_CLIENT_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *K8sClientConfigCollection) Next() (*K8sClientConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &K8sClientConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *K8sClientConfigClient) ById(id string) (*K8sClientConfig, error) {
	resp := &K8sClientConfig{}
	err := c.kuladoClient.doById(K8S_CLIENT_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *K8sClientConfigClient) Delete(container *K8sClientConfig) error {
	return c.kuladoClient.doResourceDelete(K8S_CLIENT_CONFIG_TYPE, &container.Resource)
}
