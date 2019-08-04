package client

const (
	VOLUME_TYPE = "volume"
)

type Volume struct {
	Resource `yaml:"-"`

	AccessMode string `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`

	DriverOpts map[string]string `json:"driverOpts,omitempty" yaml:"driver_opts,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	HostId string `json:"hostId,omitempty" yaml:"host_id,omitempty"`

	IsHostPath bool `json:"isHostPath,omitempty" yaml:"is_host_path,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Mounts []MountEntry `json:"mounts,omitempty" yaml:"mounts,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	SizeMb int64 `json:"sizeMb,omitempty" yaml:"size_mb,omitempty"`

	StackId string `json:"stackId,omitempty" yaml:"stack_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	StorageDriverId string `json:"storageDriverId,omitempty" yaml:"storage_driver_id,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	VolumeTemplateId string `json:"volumeTemplateId,omitempty" yaml:"volume_template_id,omitempty"`
}

type VolumeCollection struct {
	Collection
	Data   []Volume `json:"data,omitempty"`
	client *VolumeClient
}

type VolumeClient struct {
	kuladoClient *KuladoClient
}

type VolumeOperations interface {
	List(opts *ListOpts) (*VolumeCollection, error)
	Create(opts *Volume) (*Volume, error)
	Update(existing *Volume, updates interface{}) (*Volume, error)
	ById(id string) (*Volume, error)
	Delete(container *Volume) error

	ActionCreate(*Volume) (*Volume, error)

	ActionRemove(*Volume) (*Volume, error)

	ActionUpdate(*Volume) (*Volume, error)
}

func newVolumeClient(kuladoClient *KuladoClient) *VolumeClient {
	return &VolumeClient{
		kuladoClient: kuladoClient,
	}
}

func (c *VolumeClient) Create(container *Volume) (*Volume, error) {
	resp := &Volume{}
	err := c.kuladoClient.doCreate(VOLUME_TYPE, container, resp)
	return resp, err
}

func (c *VolumeClient) Update(existing *Volume, updates interface{}) (*Volume, error) {
	resp := &Volume{}
	err := c.kuladoClient.doUpdate(VOLUME_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VolumeClient) List(opts *ListOpts) (*VolumeCollection, error) {
	resp := &VolumeCollection{}
	err := c.kuladoClient.doList(VOLUME_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *VolumeCollection) Next() (*VolumeCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &VolumeCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *VolumeClient) ById(id string) (*Volume, error) {
	resp := &Volume{}
	err := c.kuladoClient.doById(VOLUME_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *VolumeClient) Delete(container *Volume) error {
	return c.kuladoClient.doResourceDelete(VOLUME_TYPE, &container.Resource)
}

func (c *VolumeClient) ActionCreate(resource *Volume) (*Volume, error) {

	resp := &Volume{}

	err := c.kuladoClient.doAction(VOLUME_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *VolumeClient) ActionRemove(resource *Volume) (*Volume, error) {

	resp := &Volume{}

	err := c.kuladoClient.doAction(VOLUME_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *VolumeClient) ActionUpdate(resource *Volume) (*Volume, error) {

	resp := &Volume{}

	err := c.kuladoClient.doAction(VOLUME_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}