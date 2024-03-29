package client

const (
	API_KEY_TYPE = "apiKey"
)

type ApiKey struct {
	Resource `yaml:"-"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	PublicValue string `json:"publicValue,omitempty" yaml:"public_value,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	SecretValue string `json:"secretValue,omitempty" yaml:"secret_value,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ApiKeyCollection struct {
	Collection
	Data   []ApiKey `json:"data,omitempty"`
	client *ApiKeyClient
}

type ApiKeyClient struct {
	kuladoClient *KuladoClient
}

type ApiKeyOperations interface {
	List(opts *ListOpts) (*ApiKeyCollection, error)
	Create(opts *ApiKey) (*ApiKey, error)
	Update(existing *ApiKey, updates interface{}) (*ApiKey, error)
	ById(id string) (*ApiKey, error)
	Delete(container *ApiKey) error

	ActionActivate(*ApiKey) (*Credential, error)

	ActionCreate(*ApiKey) (*Credential, error)

	ActionDeactivate(*ApiKey) (*Credential, error)

	ActionRemove(*ApiKey) (*Credential, error)
}

func newApiKeyClient(kuladoClient *KuladoClient) *ApiKeyClient {
	return &ApiKeyClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ApiKeyClient) Create(container *ApiKey) (*ApiKey, error) {
	resp := &ApiKey{}
	err := c.kuladoClient.doCreate(API_KEY_TYPE, container, resp)
	return resp, err
}

func (c *ApiKeyClient) Update(existing *ApiKey, updates interface{}) (*ApiKey, error) {
	resp := &ApiKey{}
	err := c.kuladoClient.doUpdate(API_KEY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ApiKeyClient) List(opts *ListOpts) (*ApiKeyCollection, error) {
	resp := &ApiKeyCollection{}
	err := c.kuladoClient.doList(API_KEY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ApiKeyCollection) Next() (*ApiKeyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ApiKeyCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ApiKeyClient) ById(id string) (*ApiKey, error) {
	resp := &ApiKey{}
	err := c.kuladoClient.doById(API_KEY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ApiKeyClient) Delete(container *ApiKey) error {
	return c.kuladoClient.doResourceDelete(API_KEY_TYPE, &container.Resource)
}

func (c *ApiKeyClient) ActionActivate(resource *ApiKey) (*Credential, error) {

	resp := &Credential{}

	err := c.kuladoClient.doAction(API_KEY_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ApiKeyClient) ActionCreate(resource *ApiKey) (*Credential, error) {

	resp := &Credential{}

	err := c.kuladoClient.doAction(API_KEY_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ApiKeyClient) ActionDeactivate(resource *ApiKey) (*Credential, error) {

	resp := &Credential{}

	err := c.kuladoClient.doAction(API_KEY_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ApiKeyClient) ActionRemove(resource *ApiKey) (*Credential, error) {

	resp := &Credential{}

	err := c.kuladoClient.doAction(API_KEY_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
