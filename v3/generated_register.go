package client

const (
	REGISTER_TYPE = "register"
)

type Register struct {
	Resource `yaml:"-"`

	AccessKey string `json:"accessKey,omitempty" yaml:"access_key,omitempty"`

	K8sClientConfig K8sClientConfig `json:"k8sClientConfig,omitempty" yaml:"k8s_client_config,omitempty"`

	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	Orchestration string `json:"orchestration,omitempty" yaml:"orchestration,omitempty"`

	SecretKey string `json:"secretKey,omitempty" yaml:"secret_key,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`
}

type RegisterCollection struct {
	Collection
	Data   []Register `json:"data,omitempty"`
	client *RegisterClient
}

type RegisterClient struct {
	kuladoClient *KuladoClient
}

type RegisterOperations interface {
	List(opts *ListOpts) (*RegisterCollection, error)
	Create(opts *Register) (*Register, error)
	Update(existing *Register, updates interface{}) (*Register, error)
	ById(id string) (*Register, error)
	Delete(container *Register) error

	ActionCreate(*Register) (*Register, error)

	ActionRemove(*Register) (*Register, error)
}

func newRegisterClient(kuladoClient *KuladoClient) *RegisterClient {
	return &RegisterClient{
		kuladoClient: kuladoClient,
	}
}

func (c *RegisterClient) Create(container *Register) (*Register, error) {
	resp := &Register{}
	err := c.kuladoClient.doCreate(REGISTER_TYPE, container, resp)
	return resp, err
}

func (c *RegisterClient) Update(existing *Register, updates interface{}) (*Register, error) {
	resp := &Register{}
	err := c.kuladoClient.doUpdate(REGISTER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RegisterClient) List(opts *ListOpts) (*RegisterCollection, error) {
	resp := &RegisterCollection{}
	err := c.kuladoClient.doList(REGISTER_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RegisterCollection) Next() (*RegisterCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RegisterCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RegisterClient) ById(id string) (*Register, error) {
	resp := &Register{}
	err := c.kuladoClient.doById(REGISTER_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RegisterClient) Delete(container *Register) error {
	return c.kuladoClient.doResourceDelete(REGISTER_TYPE, &container.Resource)
}

func (c *RegisterClient) ActionCreate(resource *Register) (*Register, error) {

	resp := &Register{}

	err := c.kuladoClient.doAction(REGISTER_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *RegisterClient) ActionRemove(resource *Register) (*Register, error) {

	resp := &Register{}

	err := c.kuladoClient.doAction(REGISTER_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
