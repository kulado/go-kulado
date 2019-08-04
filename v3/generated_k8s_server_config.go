package client

const (
	K8S_SERVER_CONFIG_TYPE = "k8sServerConfig"
)

type K8sServerConfig struct {
	Resource `yaml:"-"`

	AdmissionControllers []string `json:"admissionControllers,omitempty" yaml:"admission_controllers,omitempty"`

	ServiceNetCidr string `json:"serviceNetCidr,omitempty" yaml:"service_net_cidr,omitempty"`
}

type K8sServerConfigCollection struct {
	Collection
	Data   []K8sServerConfig `json:"data,omitempty"`
	client *K8sServerConfigClient
}

type K8sServerConfigClient struct {
	kuladoClient *KuladoClient
}

type K8sServerConfigOperations interface {
	List(opts *ListOpts) (*K8sServerConfigCollection, error)
	Create(opts *K8sServerConfig) (*K8sServerConfig, error)
	Update(existing *K8sServerConfig, updates interface{}) (*K8sServerConfig, error)
	ById(id string) (*K8sServerConfig, error)
	Delete(container *K8sServerConfig) error
}

func newK8sServerConfigClient(kuladoClient *KuladoClient) *K8sServerConfigClient {
	return &K8sServerConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *K8sServerConfigClient) Create(container *K8sServerConfig) (*K8sServerConfig, error) {
	resp := &K8sServerConfig{}
	err := c.kuladoClient.doCreate(K8S_SERVER_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *K8sServerConfigClient) Update(existing *K8sServerConfig, updates interface{}) (*K8sServerConfig, error) {
	resp := &K8sServerConfig{}
	err := c.kuladoClient.doUpdate(K8S_SERVER_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *K8sServerConfigClient) List(opts *ListOpts) (*K8sServerConfigCollection, error) {
	resp := &K8sServerConfigCollection{}
	err := c.kuladoClient.doList(K8S_SERVER_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *K8sServerConfigCollection) Next() (*K8sServerConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &K8sServerConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *K8sServerConfigClient) ById(id string) (*K8sServerConfig, error) {
	resp := &K8sServerConfig{}
	err := c.kuladoClient.doById(K8S_SERVER_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *K8sServerConfigClient) Delete(container *K8sServerConfig) error {
	return c.kuladoClient.doResourceDelete(K8S_SERVER_CONFIG_TYPE, &container.Resource)
}
