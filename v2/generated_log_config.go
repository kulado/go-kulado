package client

const (
	LOG_CONFIG_TYPE = "logConfig"
)

type LogConfig struct {
	Resource

	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`

	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`
}

type LogConfigCollection struct {
	Collection
	Data   []LogConfig `json:"data,omitempty"`
	client *LogConfigClient
}

type LogConfigClient struct {
	kuladoClient *KuladoClient
}

type LogConfigOperations interface {
	List(opts *ListOpts) (*LogConfigCollection, error)
	Create(opts *LogConfig) (*LogConfig, error)
	Update(existing *LogConfig, updates interface{}) (*LogConfig, error)
	ById(id string) (*LogConfig, error)
	Delete(container *LogConfig) error
}

func newLogConfigClient(kuladoClient *KuladoClient) *LogConfigClient {
	return &LogConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *LogConfigClient) Create(container *LogConfig) (*LogConfig, error) {
	resp := &LogConfig{}
	err := c.kuladoClient.doCreate(LOG_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *LogConfigClient) Update(existing *LogConfig, updates interface{}) (*LogConfig, error) {
	resp := &LogConfig{}
	err := c.kuladoClient.doUpdate(LOG_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LogConfigClient) List(opts *ListOpts) (*LogConfigCollection, error) {
	resp := &LogConfigCollection{}
	err := c.kuladoClient.doList(LOG_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LogConfigCollection) Next() (*LogConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LogConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LogConfigClient) ById(id string) (*LogConfig, error) {
	resp := &LogConfig{}
	err := c.kuladoClient.doById(LOG_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *LogConfigClient) Delete(container *LogConfig) error {
	return c.kuladoClient.doResourceDelete(LOG_CONFIG_TYPE, &container.Resource)
}
