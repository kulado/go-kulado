package client

const (
	SERVICE_PROXY_TYPE = "serviceProxy"
)

type ServiceProxy struct {
	Resource

	Port int64 `json:"port,omitempty" yaml:"port,omitempty"`

	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`

	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	Token string `json:"token,omitempty" yaml:"token,omitempty"`

	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}

type ServiceProxyCollection struct {
	Collection
	Data   []ServiceProxy `json:"data,omitempty"`
	client *ServiceProxyClient
}

type ServiceProxyClient struct {
	kuladoClient *KuladoClient
}

type ServiceProxyOperations interface {
	List(opts *ListOpts) (*ServiceProxyCollection, error)
	Create(opts *ServiceProxy) (*ServiceProxy, error)
	Update(existing *ServiceProxy, updates interface{}) (*ServiceProxy, error)
	ById(id string) (*ServiceProxy, error)
	Delete(container *ServiceProxy) error
}

func newServiceProxyClient(kuladoClient *KuladoClient) *ServiceProxyClient {
	return &ServiceProxyClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServiceProxyClient) Create(container *ServiceProxy) (*ServiceProxy, error) {
	resp := &ServiceProxy{}
	err := c.kuladoClient.doCreate(SERVICE_PROXY_TYPE, container, resp)
	return resp, err
}

func (c *ServiceProxyClient) Update(existing *ServiceProxy, updates interface{}) (*ServiceProxy, error) {
	resp := &ServiceProxy{}
	err := c.kuladoClient.doUpdate(SERVICE_PROXY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceProxyClient) List(opts *ListOpts) (*ServiceProxyCollection, error) {
	resp := &ServiceProxyCollection{}
	err := c.kuladoClient.doList(SERVICE_PROXY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceProxyCollection) Next() (*ServiceProxyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceProxyCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceProxyClient) ById(id string) (*ServiceProxy, error) {
	resp := &ServiceProxy{}
	err := c.kuladoClient.doById(SERVICE_PROXY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServiceProxyClient) Delete(container *ServiceProxy) error {
	return c.kuladoClient.doResourceDelete(SERVICE_PROXY_TYPE, &container.Resource)
}
