package client

const (
	INSTANCE_STOP_TYPE = "instanceStop"
)

type InstanceStop struct {
	Resource `yaml:"-"`

	Remove bool `json:"remove,omitempty" yaml:"remove,omitempty"`

	StopSource string `json:"stopSource,omitempty" yaml:"stop_source,omitempty"`

	Timeout int64 `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}

type InstanceStopCollection struct {
	Collection
	Data   []InstanceStop `json:"data,omitempty"`
	client *InstanceStopClient
}

type InstanceStopClient struct {
	kuladoClient *KuladoClient
}

type InstanceStopOperations interface {
	List(opts *ListOpts) (*InstanceStopCollection, error)
	Create(opts *InstanceStop) (*InstanceStop, error)
	Update(existing *InstanceStop, updates interface{}) (*InstanceStop, error)
	ById(id string) (*InstanceStop, error)
	Delete(container *InstanceStop) error
}

func newInstanceStopClient(kuladoClient *KuladoClient) *InstanceStopClient {
	return &InstanceStopClient{
		kuladoClient: kuladoClient,
	}
}

func (c *InstanceStopClient) Create(container *InstanceStop) (*InstanceStop, error) {
	resp := &InstanceStop{}
	err := c.kuladoClient.doCreate(INSTANCE_STOP_TYPE, container, resp)
	return resp, err
}

func (c *InstanceStopClient) Update(existing *InstanceStop, updates interface{}) (*InstanceStop, error) {
	resp := &InstanceStop{}
	err := c.kuladoClient.doUpdate(INSTANCE_STOP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceStopClient) List(opts *ListOpts) (*InstanceStopCollection, error) {
	resp := &InstanceStopCollection{}
	err := c.kuladoClient.doList(INSTANCE_STOP_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *InstanceStopCollection) Next() (*InstanceStopCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &InstanceStopCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *InstanceStopClient) ById(id string) (*InstanceStop, error) {
	resp := &InstanceStop{}
	err := c.kuladoClient.doById(INSTANCE_STOP_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *InstanceStopClient) Delete(container *InstanceStop) error {
	return c.kuladoClient.doResourceDelete(INSTANCE_STOP_TYPE, &container.Resource)
}
