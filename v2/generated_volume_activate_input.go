package client

const (
	VOLUME_ACTIVATE_INPUT_TYPE = "volumeActivateInput"
)

type VolumeActivateInput struct {
	Resource

	HostId string `json:"hostId,omitempty" yaml:"host_id,omitempty"`
}

type VolumeActivateInputCollection struct {
	Collection
	Data   []VolumeActivateInput `json:"data,omitempty"`
	client *VolumeActivateInputClient
}

type VolumeActivateInputClient struct {
	kuladoClient *KuladoClient
}

type VolumeActivateInputOperations interface {
	List(opts *ListOpts) (*VolumeActivateInputCollection, error)
	Create(opts *VolumeActivateInput) (*VolumeActivateInput, error)
	Update(existing *VolumeActivateInput, updates interface{}) (*VolumeActivateInput, error)
	ById(id string) (*VolumeActivateInput, error)
	Delete(container *VolumeActivateInput) error
}

func newVolumeActivateInputClient(kuladoClient *KuladoClient) *VolumeActivateInputClient {
	return &VolumeActivateInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *VolumeActivateInputClient) Create(container *VolumeActivateInput) (*VolumeActivateInput, error) {
	resp := &VolumeActivateInput{}
	err := c.kuladoClient.doCreate(VOLUME_ACTIVATE_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *VolumeActivateInputClient) Update(existing *VolumeActivateInput, updates interface{}) (*VolumeActivateInput, error) {
	resp := &VolumeActivateInput{}
	err := c.kuladoClient.doUpdate(VOLUME_ACTIVATE_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VolumeActivateInputClient) List(opts *ListOpts) (*VolumeActivateInputCollection, error) {
	resp := &VolumeActivateInputCollection{}
	err := c.kuladoClient.doList(VOLUME_ACTIVATE_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *VolumeActivateInputCollection) Next() (*VolumeActivateInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &VolumeActivateInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *VolumeActivateInputClient) ById(id string) (*VolumeActivateInput, error) {
	resp := &VolumeActivateInput{}
	err := c.kuladoClient.doById(VOLUME_ACTIVATE_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *VolumeActivateInputClient) Delete(container *VolumeActivateInput) error {
	return c.kuladoClient.doResourceDelete(VOLUME_ACTIVATE_INPUT_TYPE, &container.Resource)
}
