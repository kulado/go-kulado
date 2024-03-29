package client

const (
	BLKIO_DEVICE_OPTION_TYPE = "blkioDeviceOption"
)

type BlkioDeviceOption struct {
	Resource `yaml:"-"`

	ReadBps int64 `json:"readBps,omitempty" yaml:"read_bps,omitempty"`

	ReadIops int64 `json:"readIops,omitempty" yaml:"read_iops,omitempty"`

	Weight int64 `json:"weight,omitempty" yaml:"weight,omitempty"`

	WriteBps int64 `json:"writeBps,omitempty" yaml:"write_bps,omitempty"`

	WriteIops int64 `json:"writeIops,omitempty" yaml:"write_iops,omitempty"`
}

type BlkioDeviceOptionCollection struct {
	Collection
	Data   []BlkioDeviceOption `json:"data,omitempty"`
	client *BlkioDeviceOptionClient
}

type BlkioDeviceOptionClient struct {
	kuladoClient *KuladoClient
}

type BlkioDeviceOptionOperations interface {
	List(opts *ListOpts) (*BlkioDeviceOptionCollection, error)
	Create(opts *BlkioDeviceOption) (*BlkioDeviceOption, error)
	Update(existing *BlkioDeviceOption, updates interface{}) (*BlkioDeviceOption, error)
	ById(id string) (*BlkioDeviceOption, error)
	Delete(container *BlkioDeviceOption) error
}

func newBlkioDeviceOptionClient(kuladoClient *KuladoClient) *BlkioDeviceOptionClient {
	return &BlkioDeviceOptionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *BlkioDeviceOptionClient) Create(container *BlkioDeviceOption) (*BlkioDeviceOption, error) {
	resp := &BlkioDeviceOption{}
	err := c.kuladoClient.doCreate(BLKIO_DEVICE_OPTION_TYPE, container, resp)
	return resp, err
}

func (c *BlkioDeviceOptionClient) Update(existing *BlkioDeviceOption, updates interface{}) (*BlkioDeviceOption, error) {
	resp := &BlkioDeviceOption{}
	err := c.kuladoClient.doUpdate(BLKIO_DEVICE_OPTION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *BlkioDeviceOptionClient) List(opts *ListOpts) (*BlkioDeviceOptionCollection, error) {
	resp := &BlkioDeviceOptionCollection{}
	err := c.kuladoClient.doList(BLKIO_DEVICE_OPTION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *BlkioDeviceOptionCollection) Next() (*BlkioDeviceOptionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &BlkioDeviceOptionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *BlkioDeviceOptionClient) ById(id string) (*BlkioDeviceOption, error) {
	resp := &BlkioDeviceOption{}
	err := c.kuladoClient.doById(BLKIO_DEVICE_OPTION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *BlkioDeviceOptionClient) Delete(container *BlkioDeviceOption) error {
	return c.kuladoClient.doResourceDelete(BLKIO_DEVICE_OPTION_TYPE, &container.Resource)
}
