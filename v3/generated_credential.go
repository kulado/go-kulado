package client

const (
	CREDENTIAL_TYPE = "credential"
)

type Credential struct {
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

type CredentialCollection struct {
	Collection
	Data   []Credential `json:"data,omitempty"`
	client *CredentialClient
}

type CredentialClient struct {
	kuladoClient *KuladoClient
}

type CredentialOperations interface {
	List(opts *ListOpts) (*CredentialCollection, error)
	Create(opts *Credential) (*Credential, error)
	Update(existing *Credential, updates interface{}) (*Credential, error)
	ById(id string) (*Credential, error)
	Delete(container *Credential) error

	ActionActivate(*Credential) (*Credential, error)

	ActionCreate(*Credential) (*Credential, error)

	ActionDeactivate(*Credential) (*Credential, error)

	ActionRemove(*Credential) (*Credential, error)
}

func newCredentialClient(kuladoClient *KuladoClient) *CredentialClient {
	return &CredentialClient{
		kuladoClient: kuladoClient,
	}
}

func (c *CredentialClient) Create(container *Credential) (*Credential, error) {
	resp := &Credential{}
	err := c.kuladoClient.doCreate(CREDENTIAL_TYPE, container, resp)
	return resp, err
}

func (c *CredentialClient) Update(existing *Credential, updates interface{}) (*Credential, error) {
	resp := &Credential{}
	err := c.kuladoClient.doUpdate(CREDENTIAL_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *CredentialClient) List(opts *ListOpts) (*CredentialCollection, error) {
	resp := &CredentialCollection{}
	err := c.kuladoClient.doList(CREDENTIAL_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *CredentialCollection) Next() (*CredentialCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &CredentialCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *CredentialClient) ById(id string) (*Credential, error) {
	resp := &Credential{}
	err := c.kuladoClient.doById(CREDENTIAL_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *CredentialClient) Delete(container *Credential) error {
	return c.kuladoClient.doResourceDelete(CREDENTIAL_TYPE, &container.Resource)
}

func (c *CredentialClient) ActionActivate(resource *Credential) (*Credential, error) {

	resp := &Credential{}

	err := c.kuladoClient.doAction(CREDENTIAL_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *CredentialClient) ActionCreate(resource *Credential) (*Credential, error) {

	resp := &Credential{}

	err := c.kuladoClient.doAction(CREDENTIAL_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *CredentialClient) ActionDeactivate(resource *Credential) (*Credential, error) {

	resp := &Credential{}

	err := c.kuladoClient.doAction(CREDENTIAL_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *CredentialClient) ActionRemove(resource *Credential) (*Credential, error) {

	resp := &Credential{}

	err := c.kuladoClient.doAction(CREDENTIAL_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
