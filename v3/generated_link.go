package client

const (
	LINK_TYPE = "link"
)

type Link struct {
	Resource `yaml:"-"`

	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

type LinkCollection struct {
	Collection
	Data   []Link `json:"data,omitempty"`
	client *LinkClient
}

type LinkClient struct {
	kuladoClient *KuladoClient
}

type LinkOperations interface {
	List(opts *ListOpts) (*LinkCollection, error)
	Create(opts *Link) (*Link, error)
	Update(existing *Link, updates interface{}) (*Link, error)
	ById(id string) (*Link, error)
	Delete(container *Link) error
}

func newLinkClient(kuladoClient *KuladoClient) *LinkClient {
	return &LinkClient{
		kuladoClient: kuladoClient,
	}
}

func (c *LinkClient) Create(container *Link) (*Link, error) {
	resp := &Link{}
	err := c.kuladoClient.doCreate(LINK_TYPE, container, resp)
	return resp, err
}

func (c *LinkClient) Update(existing *Link, updates interface{}) (*Link, error) {
	resp := &Link{}
	err := c.kuladoClient.doUpdate(LINK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LinkClient) List(opts *ListOpts) (*LinkCollection, error) {
	resp := &LinkCollection{}
	err := c.kuladoClient.doList(LINK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LinkCollection) Next() (*LinkCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LinkCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LinkClient) ById(id string) (*Link, error) {
	resp := &Link{}
	err := c.kuladoClient.doById(LINK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *LinkClient) Delete(container *Link) error {
	return c.kuladoClient.doResourceDelete(LINK_TYPE, &container.Resource)
}
