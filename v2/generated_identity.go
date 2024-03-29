package client

const (
	IDENTITY_TYPE = "identity"
)

type Identity struct {
	Resource

	All string `json:"all,omitempty" yaml:"all,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	ExternalIdType string `json:"externalIdType,omitempty" yaml:"external_id_type,omitempty"`

	Login string `json:"login,omitempty" yaml:"login,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	ProfilePicture string `json:"profilePicture,omitempty" yaml:"profile_picture,omitempty"`

	ProfileUrl string `json:"profileUrl,omitempty" yaml:"profile_url,omitempty"`

	ProjectId string `json:"projectId,omitempty" yaml:"project_id,omitempty"`

	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	User bool `json:"user,omitempty" yaml:"user,omitempty"`
}

type IdentityCollection struct {
	Collection
	Data   []Identity `json:"data,omitempty"`
	client *IdentityClient
}

type IdentityClient struct {
	kuladoClient *KuladoClient
}

type IdentityOperations interface {
	List(opts *ListOpts) (*IdentityCollection, error)
	Create(opts *Identity) (*Identity, error)
	Update(existing *Identity, updates interface{}) (*Identity, error)
	ById(id string) (*Identity, error)
	Delete(container *Identity) error
}

func newIdentityClient(kuladoClient *KuladoClient) *IdentityClient {
	return &IdentityClient{
		kuladoClient: kuladoClient,
	}
}

func (c *IdentityClient) Create(container *Identity) (*Identity, error) {
	resp := &Identity{}
	err := c.kuladoClient.doCreate(IDENTITY_TYPE, container, resp)
	return resp, err
}

func (c *IdentityClient) Update(existing *Identity, updates interface{}) (*Identity, error) {
	resp := &Identity{}
	err := c.kuladoClient.doUpdate(IDENTITY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *IdentityClient) List(opts *ListOpts) (*IdentityCollection, error) {
	resp := &IdentityCollection{}
	err := c.kuladoClient.doList(IDENTITY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *IdentityCollection) Next() (*IdentityCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &IdentityCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *IdentityClient) ById(id string) (*Identity, error) {
	resp := &Identity{}
	err := c.kuladoClient.doById(IDENTITY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *IdentityClient) Delete(container *Identity) error {
	return c.kuladoClient.doResourceDelete(IDENTITY_TYPE, &container.Resource)
}
