package catalog

const (
	API_VERSION_TYPE = "apiVersion"
)

type ApiVersion struct {
	Resource `yaml:"-"`

	Actions map[string]string `json:"actions,omitempty" yaml:"actions,omitempty"`

	Links map[string]string `json:"links,omitempty" yaml:"links,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type ApiVersionCollection struct {
	Collection
	Data   []ApiVersion `json:"data,omitempty"`
	client *ApiVersionClient
}

type ApiVersionClient struct {
	kuladoClient *KuladoClient
}

type ApiVersionOperations interface {
	List(opts *ListOpts) (*ApiVersionCollection, error)
	Create(opts *ApiVersion) (*ApiVersion, error)
	Update(existing *ApiVersion, updates interface{}) (*ApiVersion, error)
	ById(id string) (*ApiVersion, error)
	Delete(container *ApiVersion) error
}

func newApiVersionClient(kuladoClient *KuladoClient) *ApiVersionClient {
	return &ApiVersionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ApiVersionClient) Create(container *ApiVersion) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := c.kuladoClient.doCreate(API_VERSION_TYPE, container, resp)
	return resp, err
}

func (c *ApiVersionClient) Update(existing *ApiVersion, updates interface{}) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := c.kuladoClient.doUpdate(API_VERSION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ApiVersionClient) List(opts *ListOpts) (*ApiVersionCollection, error) {
	resp := &ApiVersionCollection{}
	err := c.kuladoClient.doList(API_VERSION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ApiVersionCollection) Next() (*ApiVersionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ApiVersionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ApiVersionClient) ById(id string) (*ApiVersion, error) {
	resp := &ApiVersion{}
	err := c.kuladoClient.doById(API_VERSION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ApiVersionClient) Delete(container *ApiVersion) error {
	return c.kuladoClient.doResourceDelete(API_VERSION_TYPE, &container.Resource)
}
