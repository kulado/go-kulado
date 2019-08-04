package client

const (
	INSTANCE_STATUS_TYPE = "instanceStatus"
)

type InstanceStatus struct {
	Resource `yaml:"-"`

	DockerInspect map[string]interface{} `json:"dockerInspect,omitempty" yaml:"docker_inspect,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	InstanceUuid string `json:"instanceUuid,omitempty" yaml:"instance_uuid,omitempty"`

	PrimaryIpAddress string `json:"primaryIpAddress,omitempty" yaml:"primary_ip_address,omitempty"`
}

type InstanceStatusCollection struct {
	Collection
	Data   []InstanceStatus `json:"data,omitempty"`
	client *InstanceStatusClient
}

type InstanceStatusClient struct {
	kuladoClient *KuladoClient
}

type InstanceStatusOperations interface {
	List(opts *ListOpts) (*InstanceStatusCollection, error)
	Create(opts *InstanceStatus) (*InstanceStatus, error)
	Update(existing *InstanceStatus, updates interface{}) (*InstanceStatus, error)
	ById(id string) (*InstanceStatus, error)
	Delete(container *InstanceStatus) error
}

func newInstanceStatusClient(kuladoClient *KuladoClient) *InstanceStatusClient {
	return &InstanceStatusClient{
		kuladoClient: kuladoClient,
	}
}

func (c *InstanceStatusClient) Create(container *InstanceStatus) (*InstanceStatus, error) {
	resp := &InstanceStatus{}
	err := c.kuladoClient.doCreate(INSTANCE_STATUS_TYPE, container, resp)
	return resp, err
}

func (c *InstanceStatusClient) Update(existing *InstanceStatus, updates interface{}) (*InstanceStatus, error) {
	resp := &InstanceStatus{}
	err := c.kuladoClient.doUpdate(INSTANCE_STATUS_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InstanceStatusClient) List(opts *ListOpts) (*InstanceStatusCollection, error) {
	resp := &InstanceStatusCollection{}
	err := c.kuladoClient.doList(INSTANCE_STATUS_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *InstanceStatusCollection) Next() (*InstanceStatusCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &InstanceStatusCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *InstanceStatusClient) ById(id string) (*InstanceStatus, error) {
	resp := &InstanceStatus{}
	err := c.kuladoClient.doById(INSTANCE_STATUS_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *InstanceStatusClient) Delete(container *InstanceStatus) error {
	return c.kuladoClient.doResourceDelete(INSTANCE_STATUS_TYPE, &container.Resource)
}
