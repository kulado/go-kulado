package client

const (
	VIRTUAL_MACHINE_DISK_TYPE = "virtualMachineDisk"
)

type VirtualMachineDisk struct {
	Resource

	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Opts map[string]interface{} `json:"opts,omitempty" yaml:"opts,omitempty"`

	ReadIops int64 `json:"readIops,omitempty" yaml:"read_iops,omitempty"`

	Root bool `json:"root,omitempty" yaml:"root,omitempty"`

	Size string `json:"size,omitempty" yaml:"size,omitempty"`

	WriteIops int64 `json:"writeIops,omitempty" yaml:"write_iops,omitempty"`
}

type VirtualMachineDiskCollection struct {
	Collection
	Data   []VirtualMachineDisk `json:"data,omitempty"`
	client *VirtualMachineDiskClient
}

type VirtualMachineDiskClient struct {
	kuladoClient *KuladoClient
}

type VirtualMachineDiskOperations interface {
	List(opts *ListOpts) (*VirtualMachineDiskCollection, error)
	Create(opts *VirtualMachineDisk) (*VirtualMachineDisk, error)
	Update(existing *VirtualMachineDisk, updates interface{}) (*VirtualMachineDisk, error)
	ById(id string) (*VirtualMachineDisk, error)
	Delete(container *VirtualMachineDisk) error
}

func newVirtualMachineDiskClient(kuladoClient *KuladoClient) *VirtualMachineDiskClient {
	return &VirtualMachineDiskClient{
		kuladoClient: kuladoClient,
	}
}

func (c *VirtualMachineDiskClient) Create(container *VirtualMachineDisk) (*VirtualMachineDisk, error) {
	resp := &VirtualMachineDisk{}
	err := c.kuladoClient.doCreate(VIRTUAL_MACHINE_DISK_TYPE, container, resp)
	return resp, err
}

func (c *VirtualMachineDiskClient) Update(existing *VirtualMachineDisk, updates interface{}) (*VirtualMachineDisk, error) {
	resp := &VirtualMachineDisk{}
	err := c.kuladoClient.doUpdate(VIRTUAL_MACHINE_DISK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *VirtualMachineDiskClient) List(opts *ListOpts) (*VirtualMachineDiskCollection, error) {
	resp := &VirtualMachineDiskCollection{}
	err := c.kuladoClient.doList(VIRTUAL_MACHINE_DISK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *VirtualMachineDiskCollection) Next() (*VirtualMachineDiskCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &VirtualMachineDiskCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *VirtualMachineDiskClient) ById(id string) (*VirtualMachineDisk, error) {
	resp := &VirtualMachineDisk{}
	err := c.kuladoClient.doById(VIRTUAL_MACHINE_DISK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *VirtualMachineDiskClient) Delete(container *VirtualMachineDisk) error {
	return c.kuladoClient.doResourceDelete(VIRTUAL_MACHINE_DISK_TYPE, &container.Resource)
}
