package client

const (
	ADD_REMOVE_SERVICE_LINK_INPUT_TYPE = "addRemoveServiceLinkInput"
)

type AddRemoveServiceLinkInput struct {
	Resource

	ServiceLink ServiceLink `json:"serviceLink,omitempty" yaml:"service_link,omitempty"`
}

type AddRemoveServiceLinkInputCollection struct {
	Collection
	Data   []AddRemoveServiceLinkInput `json:"data,omitempty"`
	client *AddRemoveServiceLinkInputClient
}

type AddRemoveServiceLinkInputClient struct {
	kuladoClient *KuladoClient
}

type AddRemoveServiceLinkInputOperations interface {
	List(opts *ListOpts) (*AddRemoveServiceLinkInputCollection, error)
	Create(opts *AddRemoveServiceLinkInput) (*AddRemoveServiceLinkInput, error)
	Update(existing *AddRemoveServiceLinkInput, updates interface{}) (*AddRemoveServiceLinkInput, error)
	ById(id string) (*AddRemoveServiceLinkInput, error)
	Delete(container *AddRemoveServiceLinkInput) error
}

func newAddRemoveServiceLinkInputClient(kuladoClient *KuladoClient) *AddRemoveServiceLinkInputClient {
	return &AddRemoveServiceLinkInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *AddRemoveServiceLinkInputClient) Create(container *AddRemoveServiceLinkInput) (*AddRemoveServiceLinkInput, error) {
	resp := &AddRemoveServiceLinkInput{}
	err := c.kuladoClient.doCreate(ADD_REMOVE_SERVICE_LINK_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *AddRemoveServiceLinkInputClient) Update(existing *AddRemoveServiceLinkInput, updates interface{}) (*AddRemoveServiceLinkInput, error) {
	resp := &AddRemoveServiceLinkInput{}
	err := c.kuladoClient.doUpdate(ADD_REMOVE_SERVICE_LINK_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AddRemoveServiceLinkInputClient) List(opts *ListOpts) (*AddRemoveServiceLinkInputCollection, error) {
	resp := &AddRemoveServiceLinkInputCollection{}
	err := c.kuladoClient.doList(ADD_REMOVE_SERVICE_LINK_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AddRemoveServiceLinkInputCollection) Next() (*AddRemoveServiceLinkInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AddRemoveServiceLinkInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AddRemoveServiceLinkInputClient) ById(id string) (*AddRemoveServiceLinkInput, error) {
	resp := &AddRemoveServiceLinkInput{}
	err := c.kuladoClient.doById(ADD_REMOVE_SERVICE_LINK_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *AddRemoveServiceLinkInputClient) Delete(container *AddRemoveServiceLinkInput) error {
	return c.kuladoClient.doResourceDelete(ADD_REMOVE_SERVICE_LINK_INPUT_TYPE, &container.Resource)
}
