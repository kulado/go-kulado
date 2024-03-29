package client

const (
	STORAGE_DRIVER_TYPE = "storageDriver"
)

type StorageDriver struct {
	Resource `yaml:"-"`

	BlockDevicePath string `json:"blockDevicePath,omitempty" yaml:"block_device_path,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	VolumeAccessMode string `json:"volumeAccessMode,omitempty" yaml:"volume_access_mode,omitempty"`

	VolumeCapabilities []string `json:"volumeCapabilities,omitempty" yaml:"volume_capabilities,omitempty"`
}

type StorageDriverCollection struct {
	Collection
	Data   []StorageDriver `json:"data,omitempty"`
	client *StorageDriverClient
}

type StorageDriverClient struct {
	kuladoClient *KuladoClient
}

type StorageDriverOperations interface {
	List(opts *ListOpts) (*StorageDriverCollection, error)
	Create(opts *StorageDriver) (*StorageDriver, error)
	Update(existing *StorageDriver, updates interface{}) (*StorageDriver, error)
	ById(id string) (*StorageDriver, error)
	Delete(container *StorageDriver) error

	ActionCreate(*StorageDriver) (*StorageDriver, error)

	ActionRemove(*StorageDriver) (*StorageDriver, error)
}

func newStorageDriverClient(kuladoClient *KuladoClient) *StorageDriverClient {
	return &StorageDriverClient{
		kuladoClient: kuladoClient,
	}
}

func (c *StorageDriverClient) Create(container *StorageDriver) (*StorageDriver, error) {
	resp := &StorageDriver{}
	err := c.kuladoClient.doCreate(STORAGE_DRIVER_TYPE, container, resp)
	return resp, err
}

func (c *StorageDriverClient) Update(existing *StorageDriver, updates interface{}) (*StorageDriver, error) {
	resp := &StorageDriver{}
	err := c.kuladoClient.doUpdate(STORAGE_DRIVER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StorageDriverClient) List(opts *ListOpts) (*StorageDriverCollection, error) {
	resp := &StorageDriverCollection{}
	err := c.kuladoClient.doList(STORAGE_DRIVER_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *StorageDriverCollection) Next() (*StorageDriverCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &StorageDriverCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *StorageDriverClient) ById(id string) (*StorageDriver, error) {
	resp := &StorageDriver{}
	err := c.kuladoClient.doById(STORAGE_DRIVER_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *StorageDriverClient) Delete(container *StorageDriver) error {
	return c.kuladoClient.doResourceDelete(STORAGE_DRIVER_TYPE, &container.Resource)
}

func (c *StorageDriverClient) ActionCreate(resource *StorageDriver) (*StorageDriver, error) {

	resp := &StorageDriver{}

	err := c.kuladoClient.doAction(STORAGE_DRIVER_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *StorageDriverClient) ActionRemove(resource *StorageDriver) (*StorageDriver, error) {

	resp := &StorageDriver{}

	err := c.kuladoClient.doAction(STORAGE_DRIVER_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
