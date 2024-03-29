package client

const (
	BINDING_TYPE = "binding"
)

type Binding struct {
	Resource

	Services map[string]interface{} `json:"services,omitempty" yaml:"services,omitempty"`
}

type BindingCollection struct {
	Collection
	Data   []Binding `json:"data,omitempty"`
	client *BindingClient
}

type BindingClient struct {
	kuladoClient *KuladoClient
}

type BindingOperations interface {
	List(opts *ListOpts) (*BindingCollection, error)
	Create(opts *Binding) (*Binding, error)
	Update(existing *Binding, updates interface{}) (*Binding, error)
	ById(id string) (*Binding, error)
	Delete(container *Binding) error
}

func newBindingClient(kuladoClient *KuladoClient) *BindingClient {
	return &BindingClient{
		kuladoClient: kuladoClient,
	}
}

func (c *BindingClient) Create(container *Binding) (*Binding, error) {
	resp := &Binding{}
	err := c.kuladoClient.doCreate(BINDING_TYPE, container, resp)
	return resp, err
}

func (c *BindingClient) Update(existing *Binding, updates interface{}) (*Binding, error) {
	resp := &Binding{}
	err := c.kuladoClient.doUpdate(BINDING_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *BindingClient) List(opts *ListOpts) (*BindingCollection, error) {
	resp := &BindingCollection{}
	err := c.kuladoClient.doList(BINDING_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *BindingCollection) Next() (*BindingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &BindingCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *BindingClient) ById(id string) (*Binding, error) {
	resp := &Binding{}
	err := c.kuladoClient.doById(BINDING_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *BindingClient) Delete(container *Binding) error {
	return c.kuladoClient.doResourceDelete(BINDING_TYPE, &container.Resource)
}
