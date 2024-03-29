package client

const (
	DYNAMIC_SCHEMA_TYPE = "dynamicSchema"
)

type DynamicSchema struct {
	Resource `yaml:"-"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Definition string `json:"definition,omitempty" yaml:"definition,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type DynamicSchemaCollection struct {
	Collection
	Data   []DynamicSchema `json:"data,omitempty"`
	client *DynamicSchemaClient
}

type DynamicSchemaClient struct {
	kuladoClient *KuladoClient
}

type DynamicSchemaOperations interface {
	List(opts *ListOpts) (*DynamicSchemaCollection, error)
	Create(opts *DynamicSchema) (*DynamicSchema, error)
	Update(existing *DynamicSchema, updates interface{}) (*DynamicSchema, error)
	ById(id string) (*DynamicSchema, error)
	Delete(container *DynamicSchema) error

	ActionCreate(*DynamicSchema) (*DynamicSchema, error)

	ActionRemove(*DynamicSchema) (*DynamicSchema, error)
}

func newDynamicSchemaClient(kuladoClient *KuladoClient) *DynamicSchemaClient {
	return &DynamicSchemaClient{
		kuladoClient: kuladoClient,
	}
}

func (c *DynamicSchemaClient) Create(container *DynamicSchema) (*DynamicSchema, error) {
	resp := &DynamicSchema{}
	err := c.kuladoClient.doCreate(DYNAMIC_SCHEMA_TYPE, container, resp)
	return resp, err
}

func (c *DynamicSchemaClient) Update(existing *DynamicSchema, updates interface{}) (*DynamicSchema, error) {
	resp := &DynamicSchema{}
	err := c.kuladoClient.doUpdate(DYNAMIC_SCHEMA_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DynamicSchemaClient) List(opts *ListOpts) (*DynamicSchemaCollection, error) {
	resp := &DynamicSchemaCollection{}
	err := c.kuladoClient.doList(DYNAMIC_SCHEMA_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DynamicSchemaCollection) Next() (*DynamicSchemaCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DynamicSchemaCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DynamicSchemaClient) ById(id string) (*DynamicSchema, error) {
	resp := &DynamicSchema{}
	err := c.kuladoClient.doById(DYNAMIC_SCHEMA_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *DynamicSchemaClient) Delete(container *DynamicSchema) error {
	return c.kuladoClient.doResourceDelete(DYNAMIC_SCHEMA_TYPE, &container.Resource)
}

func (c *DynamicSchemaClient) ActionCreate(resource *DynamicSchema) (*DynamicSchema, error) {

	resp := &DynamicSchema{}

	err := c.kuladoClient.doAction(DYNAMIC_SCHEMA_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *DynamicSchemaClient) ActionRemove(resource *DynamicSchema) (*DynamicSchema, error) {

	resp := &DynamicSchema{}

	err := c.kuladoClient.doAction(DYNAMIC_SCHEMA_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
