package client

const (
	CONTAINER_PROXY_TYPE = "containerProxy"
)

type ContainerProxy struct {
	Resource `yaml:"-"`

	Port int64 `json:"port,omitempty" yaml:"port,omitempty"`

	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}

type ContainerProxyCollection struct {
	Collection
	Data   []ContainerProxy `json:"data,omitempty"`
	client *ContainerProxyClient
}

type ContainerProxyClient struct {
	kuladoClient *KuladoClient
}

type ContainerProxyOperations interface {
	List(opts *ListOpts) (*ContainerProxyCollection, error)
	Create(opts *ContainerProxy) (*ContainerProxy, error)
	Update(existing *ContainerProxy, updates interface{}) (*ContainerProxy, error)
	ById(id string) (*ContainerProxy, error)
	Delete(container *ContainerProxy) error
}

func newContainerProxyClient(kuladoClient *KuladoClient) *ContainerProxyClient {
	return &ContainerProxyClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ContainerProxyClient) Create(container *ContainerProxy) (*ContainerProxy, error) {
	resp := &ContainerProxy{}
	err := c.kuladoClient.doCreate(CONTAINER_PROXY_TYPE, container, resp)
	return resp, err
}

func (c *ContainerProxyClient) Update(existing *ContainerProxy, updates interface{}) (*ContainerProxy, error) {
	resp := &ContainerProxy{}
	err := c.kuladoClient.doUpdate(CONTAINER_PROXY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ContainerProxyClient) List(opts *ListOpts) (*ContainerProxyCollection, error) {
	resp := &ContainerProxyCollection{}
	err := c.kuladoClient.doList(CONTAINER_PROXY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ContainerProxyCollection) Next() (*ContainerProxyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ContainerProxyCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ContainerProxyClient) ById(id string) (*ContainerProxy, error) {
	resp := &ContainerProxy{}
	err := c.kuladoClient.doById(CONTAINER_PROXY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ContainerProxyClient) Delete(container *ContainerProxy) error {
	return c.kuladoClient.doResourceDelete(CONTAINER_PROXY_TYPE, &container.Resource)
}
