package client

const (
	STACK_CONFIGURATION_TYPE = "stackConfiguration"
)

type StackConfiguration struct {
	Resource `yaml:"-"`

	Answers map[string]interface{} `json:"answers,omitempty" yaml:"answers,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	Templates map[string]string `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type StackConfigurationCollection struct {
	Collection
	Data   []StackConfiguration `json:"data,omitempty"`
	client *StackConfigurationClient
}

type StackConfigurationClient struct {
	kuladoClient *KuladoClient
}

type StackConfigurationOperations interface {
	List(opts *ListOpts) (*StackConfigurationCollection, error)
	Create(opts *StackConfiguration) (*StackConfiguration, error)
	Update(existing *StackConfiguration, updates interface{}) (*StackConfiguration, error)
	ById(id string) (*StackConfiguration, error)
	Delete(container *StackConfiguration) error
}

func newStackConfigurationClient(kuladoClient *KuladoClient) *StackConfigurationClient {
	return &StackConfigurationClient{
		kuladoClient: kuladoClient,
	}
}

func (c *StackConfigurationClient) Create(container *StackConfiguration) (*StackConfiguration, error) {
	resp := &StackConfiguration{}
	err := c.kuladoClient.doCreate(STACK_CONFIGURATION_TYPE, container, resp)
	return resp, err
}

func (c *StackConfigurationClient) Update(existing *StackConfiguration, updates interface{}) (*StackConfiguration, error) {
	resp := &StackConfiguration{}
	err := c.kuladoClient.doUpdate(STACK_CONFIGURATION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StackConfigurationClient) List(opts *ListOpts) (*StackConfigurationCollection, error) {
	resp := &StackConfigurationCollection{}
	err := c.kuladoClient.doList(STACK_CONFIGURATION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *StackConfigurationCollection) Next() (*StackConfigurationCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &StackConfigurationCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *StackConfigurationClient) ById(id string) (*StackConfiguration, error) {
	resp := &StackConfiguration{}
	err := c.kuladoClient.doById(STACK_CONFIGURATION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *StackConfigurationClient) Delete(container *StackConfiguration) error {
	return c.kuladoClient.doResourceDelete(STACK_CONFIGURATION_TYPE, &container.Resource)
}
