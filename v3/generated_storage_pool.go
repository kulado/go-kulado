package client

const (
	STORAGE_POOL_TYPE = "storagePool"
)

type StoragePool struct {
	Resource `yaml:"-"`

	BlockDevicePath string `json:"blockDevicePath,omitempty" yaml:"block_device_path,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	DriverName string `json:"driverName,omitempty" yaml:"driver_name,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	HostIds []string `json:"hostIds,omitempty" yaml:"host_ids,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	StorageDriverId string `json:"storageDriverId,omitempty" yaml:"storage_driver_id,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	VolumeAccessMode string `json:"volumeAccessMode,omitempty" yaml:"volume_access_mode,omitempty"`

	VolumeCapabilities []string `json:"volumeCapabilities,omitempty" yaml:"volume_capabilities,omitempty"`

	VolumeIds []string `json:"volumeIds,omitempty" yaml:"volume_ids,omitempty"`
}

type StoragePoolCollection struct {
	Collection
	Data   []StoragePool `json:"data,omitempty"`
	client *StoragePoolClient
}

type StoragePoolClient struct {
	kuladoClient *KuladoClient
}

type StoragePoolOperations interface {
	List(opts *ListOpts) (*StoragePoolCollection, error)
	Create(opts *StoragePool) (*StoragePool, error)
	Update(existing *StoragePool, updates interface{}) (*StoragePool, error)
	ById(id string) (*StoragePool, error)
	Delete(container *StoragePool) error

	ActionActivate(*StoragePool) (*StoragePool, error)

	ActionCreate(*StoragePool) (*StoragePool, error)

	ActionDeactivate(*StoragePool) (*StoragePool, error)

	ActionRemove(*StoragePool) (*StoragePool, error)
}

func newStoragePoolClient(kuladoClient *KuladoClient) *StoragePoolClient {
	return &StoragePoolClient{
		kuladoClient: kuladoClient,
	}
}

func (c *StoragePoolClient) Create(container *StoragePool) (*StoragePool, error) {
	resp := &StoragePool{}
	err := c.kuladoClient.doCreate(STORAGE_POOL_TYPE, container, resp)
	return resp, err
}

func (c *StoragePoolClient) Update(existing *StoragePool, updates interface{}) (*StoragePool, error) {
	resp := &StoragePool{}
	err := c.kuladoClient.doUpdate(STORAGE_POOL_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *StoragePoolClient) List(opts *ListOpts) (*StoragePoolCollection, error) {
	resp := &StoragePoolCollection{}
	err := c.kuladoClient.doList(STORAGE_POOL_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *StoragePoolCollection) Next() (*StoragePoolCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &StoragePoolCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *StoragePoolClient) ById(id string) (*StoragePool, error) {
	resp := &StoragePool{}
	err := c.kuladoClient.doById(STORAGE_POOL_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *StoragePoolClient) Delete(container *StoragePool) error {
	return c.kuladoClient.doResourceDelete(STORAGE_POOL_TYPE, &container.Resource)
}

func (c *StoragePoolClient) ActionActivate(resource *StoragePool) (*StoragePool, error) {

	resp := &StoragePool{}

	err := c.kuladoClient.doAction(STORAGE_POOL_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *StoragePoolClient) ActionCreate(resource *StoragePool) (*StoragePool, error) {

	resp := &StoragePool{}

	err := c.kuladoClient.doAction(STORAGE_POOL_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *StoragePoolClient) ActionDeactivate(resource *StoragePool) (*StoragePool, error) {

	resp := &StoragePool{}

	err := c.kuladoClient.doAction(STORAGE_POOL_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *StoragePoolClient) ActionRemove(resource *StoragePool) (*StoragePool, error) {

	resp := &StoragePool{}

	err := c.kuladoClient.doAction(STORAGE_POOL_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
