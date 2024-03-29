package client

const (
	COMPOSE_CONFIG_TYPE = "composeConfig"
)

type ComposeConfig struct {
	Resource `yaml:"-"`

	Templates map[string]string `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type ComposeConfigCollection struct {
	Collection
	Data   []ComposeConfig `json:"data,omitempty"`
	client *ComposeConfigClient
}

type ComposeConfigClient struct {
	kuladoClient *KuladoClient
}

type ComposeConfigOperations interface {
	List(opts *ListOpts) (*ComposeConfigCollection, error)
	Create(opts *ComposeConfig) (*ComposeConfig, error)
	Update(existing *ComposeConfig, updates interface{}) (*ComposeConfig, error)
	ById(id string) (*ComposeConfig, error)
	Delete(container *ComposeConfig) error
}

func newComposeConfigClient(kuladoClient *KuladoClient) *ComposeConfigClient {
	return &ComposeConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ComposeConfigClient) Create(container *ComposeConfig) (*ComposeConfig, error) {
	resp := &ComposeConfig{}
	err := c.kuladoClient.doCreate(COMPOSE_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *ComposeConfigClient) Update(existing *ComposeConfig, updates interface{}) (*ComposeConfig, error) {
	resp := &ComposeConfig{}
	err := c.kuladoClient.doUpdate(COMPOSE_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ComposeConfigClient) List(opts *ListOpts) (*ComposeConfigCollection, error) {
	resp := &ComposeConfigCollection{}
	err := c.kuladoClient.doList(COMPOSE_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ComposeConfigCollection) Next() (*ComposeConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ComposeConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ComposeConfigClient) ById(id string) (*ComposeConfig, error) {
	resp := &ComposeConfig{}
	err := c.kuladoClient.doById(COMPOSE_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ComposeConfigClient) Delete(container *ComposeConfig) error {
	return c.kuladoClient.doResourceDelete(COMPOSE_CONFIG_TYPE, &container.Resource)
}
