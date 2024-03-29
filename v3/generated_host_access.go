package client

const (
	HOST_ACCESS_TYPE = "hostAccess"
)

type HostAccess struct {
	Resource `yaml:"-"`

	Token string `json:"token,omitempty" yaml:"token,omitempty"`

	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}

type HostAccessCollection struct {
	Collection
	Data   []HostAccess `json:"data,omitempty"`
	client *HostAccessClient
}

type HostAccessClient struct {
	kuladoClient *KuladoClient
}

type HostAccessOperations interface {
	List(opts *ListOpts) (*HostAccessCollection, error)
	Create(opts *HostAccess) (*HostAccess, error)
	Update(existing *HostAccess, updates interface{}) (*HostAccess, error)
	ById(id string) (*HostAccess, error)
	Delete(container *HostAccess) error
}

func newHostAccessClient(kuladoClient *KuladoClient) *HostAccessClient {
	return &HostAccessClient{
		kuladoClient: kuladoClient,
	}
}

func (c *HostAccessClient) Create(container *HostAccess) (*HostAccess, error) {
	resp := &HostAccess{}
	err := c.kuladoClient.doCreate(HOST_ACCESS_TYPE, container, resp)
	return resp, err
}

func (c *HostAccessClient) Update(existing *HostAccess, updates interface{}) (*HostAccess, error) {
	resp := &HostAccess{}
	err := c.kuladoClient.doUpdate(HOST_ACCESS_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HostAccessClient) List(opts *ListOpts) (*HostAccessCollection, error) {
	resp := &HostAccessCollection{}
	err := c.kuladoClient.doList(HOST_ACCESS_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HostAccessCollection) Next() (*HostAccessCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HostAccessCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HostAccessClient) ById(id string) (*HostAccess, error) {
	resp := &HostAccess{}
	err := c.kuladoClient.doById(HOST_ACCESS_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *HostAccessClient) Delete(container *HostAccess) error {
	return c.kuladoClient.doResourceDelete(HOST_ACCESS_TYPE, &container.Resource)
}
