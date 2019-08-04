package client

const (
	SECRET_REFERENCE_TYPE = "secretReference"
)

type SecretReference struct {
	Resource `yaml:"-"`

	Gid string `json:"gid,omitempty" yaml:"gid,omitempty"`

	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	SecretId string `json:"secretId,omitempty" yaml:"secret_id,omitempty"`

	Uid string `json:"uid,omitempty" yaml:"uid,omitempty"`
}

type SecretReferenceCollection struct {
	Collection
	Data   []SecretReference `json:"data,omitempty"`
	client *SecretReferenceClient
}

type SecretReferenceClient struct {
	kuladoClient *KuladoClient
}

type SecretReferenceOperations interface {
	List(opts *ListOpts) (*SecretReferenceCollection, error)
	Create(opts *SecretReference) (*SecretReference, error)
	Update(existing *SecretReference, updates interface{}) (*SecretReference, error)
	ById(id string) (*SecretReference, error)
	Delete(container *SecretReference) error
}

func newSecretReferenceClient(kuladoClient *KuladoClient) *SecretReferenceClient {
	return &SecretReferenceClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SecretReferenceClient) Create(container *SecretReference) (*SecretReference, error) {
	resp := &SecretReference{}
	err := c.kuladoClient.doCreate(SECRET_REFERENCE_TYPE, container, resp)
	return resp, err
}

func (c *SecretReferenceClient) Update(existing *SecretReference, updates interface{}) (*SecretReference, error) {
	resp := &SecretReference{}
	err := c.kuladoClient.doUpdate(SECRET_REFERENCE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SecretReferenceClient) List(opts *ListOpts) (*SecretReferenceCollection, error) {
	resp := &SecretReferenceCollection{}
	err := c.kuladoClient.doList(SECRET_REFERENCE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SecretReferenceCollection) Next() (*SecretReferenceCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SecretReferenceCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SecretReferenceClient) ById(id string) (*SecretReference, error) {
	resp := &SecretReference{}
	err := c.kuladoClient.doById(SECRET_REFERENCE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SecretReferenceClient) Delete(container *SecretReference) error {
	return c.kuladoClient.doResourceDelete(SECRET_REFERENCE_TYPE, &container.Resource)
}
