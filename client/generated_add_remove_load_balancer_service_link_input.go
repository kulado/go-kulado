package client

const (
	ADD_REMOVE_LOAD_BALANCER_SERVICE_LINK_INPUT_TYPE = "addRemoveLoadBalancerServiceLinkInput"
)

type AddRemoveLoadBalancerServiceLinkInput struct {
	Resource

	ServiceLink LoadBalancerServiceLink `json:"serviceLink,omitempty" yaml:"service_link,omitempty"`
}

type AddRemoveLoadBalancerServiceLinkInputCollection struct {
	Collection
	Data   []AddRemoveLoadBalancerServiceLinkInput `json:"data,omitempty"`
	client *AddRemoveLoadBalancerServiceLinkInputClient
}

type AddRemoveLoadBalancerServiceLinkInputClient struct {
	kuladoClient *KuladoClient
}

type AddRemoveLoadBalancerServiceLinkInputOperations interface {
	List(opts *ListOpts) (*AddRemoveLoadBalancerServiceLinkInputCollection, error)
	Create(opts *AddRemoveLoadBalancerServiceLinkInput) (*AddRemoveLoadBalancerServiceLinkInput, error)
	Update(existing *AddRemoveLoadBalancerServiceLinkInput, updates interface{}) (*AddRemoveLoadBalancerServiceLinkInput, error)
	ById(id string) (*AddRemoveLoadBalancerServiceLinkInput, error)
	Delete(container *AddRemoveLoadBalancerServiceLinkInput) error
}

func newAddRemoveLoadBalancerServiceLinkInputClient(kuladoClient *KuladoClient) *AddRemoveLoadBalancerServiceLinkInputClient {
	return &AddRemoveLoadBalancerServiceLinkInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *AddRemoveLoadBalancerServiceLinkInputClient) Create(container *AddRemoveLoadBalancerServiceLinkInput) (*AddRemoveLoadBalancerServiceLinkInput, error) {
	resp := &AddRemoveLoadBalancerServiceLinkInput{}
	err := c.kuladoClient.doCreate(ADD_REMOVE_LOAD_BALANCER_SERVICE_LINK_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *AddRemoveLoadBalancerServiceLinkInputClient) Update(existing *AddRemoveLoadBalancerServiceLinkInput, updates interface{}) (*AddRemoveLoadBalancerServiceLinkInput, error) {
	resp := &AddRemoveLoadBalancerServiceLinkInput{}
	err := c.kuladoClient.doUpdate(ADD_REMOVE_LOAD_BALANCER_SERVICE_LINK_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AddRemoveLoadBalancerServiceLinkInputClient) List(opts *ListOpts) (*AddRemoveLoadBalancerServiceLinkInputCollection, error) {
	resp := &AddRemoveLoadBalancerServiceLinkInputCollection{}
	err := c.kuladoClient.doList(ADD_REMOVE_LOAD_BALANCER_SERVICE_LINK_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AddRemoveLoadBalancerServiceLinkInputCollection) Next() (*AddRemoveLoadBalancerServiceLinkInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AddRemoveLoadBalancerServiceLinkInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AddRemoveLoadBalancerServiceLinkInputClient) ById(id string) (*AddRemoveLoadBalancerServiceLinkInput, error) {
	resp := &AddRemoveLoadBalancerServiceLinkInput{}
	err := c.kuladoClient.doById(ADD_REMOVE_LOAD_BALANCER_SERVICE_LINK_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *AddRemoveLoadBalancerServiceLinkInputClient) Delete(container *AddRemoveLoadBalancerServiceLinkInput) error {
	return c.kuladoClient.doResourceDelete(ADD_REMOVE_LOAD_BALANCER_SERVICE_LINK_INPUT_TYPE, &container.Resource)
}
