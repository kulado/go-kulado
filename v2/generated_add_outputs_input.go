package client

const (
	ADD_OUTPUTS_INPUT_TYPE = "addOutputsInput"
)

type AddOutputsInput struct {
	Resource

	Outputs map[string]interface{} `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}

type AddOutputsInputCollection struct {
	Collection
	Data   []AddOutputsInput `json:"data,omitempty"`
	client *AddOutputsInputClient
}

type AddOutputsInputClient struct {
	kuladoClient *KuladoClient
}

type AddOutputsInputOperations interface {
	List(opts *ListOpts) (*AddOutputsInputCollection, error)
	Create(opts *AddOutputsInput) (*AddOutputsInput, error)
	Update(existing *AddOutputsInput, updates interface{}) (*AddOutputsInput, error)
	ById(id string) (*AddOutputsInput, error)
	Delete(container *AddOutputsInput) error
}

func newAddOutputsInputClient(kuladoClient *KuladoClient) *AddOutputsInputClient {
	return &AddOutputsInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *AddOutputsInputClient) Create(container *AddOutputsInput) (*AddOutputsInput, error) {
	resp := &AddOutputsInput{}
	err := c.kuladoClient.doCreate(ADD_OUTPUTS_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *AddOutputsInputClient) Update(existing *AddOutputsInput, updates interface{}) (*AddOutputsInput, error) {
	resp := &AddOutputsInput{}
	err := c.kuladoClient.doUpdate(ADD_OUTPUTS_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AddOutputsInputClient) List(opts *ListOpts) (*AddOutputsInputCollection, error) {
	resp := &AddOutputsInputCollection{}
	err := c.kuladoClient.doList(ADD_OUTPUTS_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AddOutputsInputCollection) Next() (*AddOutputsInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AddOutputsInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AddOutputsInputClient) ById(id string) (*AddOutputsInput, error) {
	resp := &AddOutputsInput{}
	err := c.kuladoClient.doById(ADD_OUTPUTS_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *AddOutputsInputClient) Delete(container *AddOutputsInput) error {
	return c.kuladoClient.doResourceDelete(ADD_OUTPUTS_INPUT_TYPE, &container.Resource)
}
