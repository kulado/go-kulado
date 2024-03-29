package client

const (
	IP_ADDRESS_ASSOCIATE_INPUT_TYPE = "ipAddressAssociateInput"
)

type IpAddressAssociateInput struct {
	Resource

	IpAddressId string `json:"ipAddressId,omitempty" yaml:"ip_address_id,omitempty"`
}

type IpAddressAssociateInputCollection struct {
	Collection
	Data   []IpAddressAssociateInput `json:"data,omitempty"`
	client *IpAddressAssociateInputClient
}

type IpAddressAssociateInputClient struct {
	kuladoClient *KuladoClient
}

type IpAddressAssociateInputOperations interface {
	List(opts *ListOpts) (*IpAddressAssociateInputCollection, error)
	Create(opts *IpAddressAssociateInput) (*IpAddressAssociateInput, error)
	Update(existing *IpAddressAssociateInput, updates interface{}) (*IpAddressAssociateInput, error)
	ById(id string) (*IpAddressAssociateInput, error)
	Delete(container *IpAddressAssociateInput) error
}

func newIpAddressAssociateInputClient(kuladoClient *KuladoClient) *IpAddressAssociateInputClient {
	return &IpAddressAssociateInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *IpAddressAssociateInputClient) Create(container *IpAddressAssociateInput) (*IpAddressAssociateInput, error) {
	resp := &IpAddressAssociateInput{}
	err := c.kuladoClient.doCreate(IP_ADDRESS_ASSOCIATE_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *IpAddressAssociateInputClient) Update(existing *IpAddressAssociateInput, updates interface{}) (*IpAddressAssociateInput, error) {
	resp := &IpAddressAssociateInput{}
	err := c.kuladoClient.doUpdate(IP_ADDRESS_ASSOCIATE_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *IpAddressAssociateInputClient) List(opts *ListOpts) (*IpAddressAssociateInputCollection, error) {
	resp := &IpAddressAssociateInputCollection{}
	err := c.kuladoClient.doList(IP_ADDRESS_ASSOCIATE_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *IpAddressAssociateInputCollection) Next() (*IpAddressAssociateInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &IpAddressAssociateInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *IpAddressAssociateInputClient) ById(id string) (*IpAddressAssociateInput, error) {
	resp := &IpAddressAssociateInput{}
	err := c.kuladoClient.doById(IP_ADDRESS_ASSOCIATE_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *IpAddressAssociateInputClient) Delete(container *IpAddressAssociateInput) error {
	return c.kuladoClient.doResourceDelete(IP_ADDRESS_ASSOCIATE_INPUT_TYPE, &container.Resource)
}
