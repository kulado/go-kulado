package client

const (
	SERVICE_EXPOSE_MAP_TYPE = "serviceExposeMap"
)

type ServiceExposeMap struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	InstanceId string `json:"instanceId,omitempty" yaml:"instance_id,omitempty"`

	IpAddress string `json:"ipAddress,omitempty" yaml:"ip_address,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Managed bool `json:"managed,omitempty" yaml:"managed,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ServiceExposeMapCollection struct {
	Collection
	Data   []ServiceExposeMap `json:"data,omitempty"`
	client *ServiceExposeMapClient
}

type ServiceExposeMapClient struct {
	kuladoClient *KuladoClient
}

type ServiceExposeMapOperations interface {
	List(opts *ListOpts) (*ServiceExposeMapCollection, error)
	Create(opts *ServiceExposeMap) (*ServiceExposeMap, error)
	Update(existing *ServiceExposeMap, updates interface{}) (*ServiceExposeMap, error)
	ById(id string) (*ServiceExposeMap, error)
	Delete(container *ServiceExposeMap) error

	ActionCreate(*ServiceExposeMap) (*ServiceExposeMap, error)

	ActionRemove(*ServiceExposeMap) (*ServiceExposeMap, error)
}

func newServiceExposeMapClient(kuladoClient *KuladoClient) *ServiceExposeMapClient {
	return &ServiceExposeMapClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ServiceExposeMapClient) Create(container *ServiceExposeMap) (*ServiceExposeMap, error) {
	resp := &ServiceExposeMap{}
	err := c.kuladoClient.doCreate(SERVICE_EXPOSE_MAP_TYPE, container, resp)
	return resp, err
}

func (c *ServiceExposeMapClient) Update(existing *ServiceExposeMap, updates interface{}) (*ServiceExposeMap, error) {
	resp := &ServiceExposeMap{}
	err := c.kuladoClient.doUpdate(SERVICE_EXPOSE_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ServiceExposeMapClient) List(opts *ListOpts) (*ServiceExposeMapCollection, error) {
	resp := &ServiceExposeMapCollection{}
	err := c.kuladoClient.doList(SERVICE_EXPOSE_MAP_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ServiceExposeMapCollection) Next() (*ServiceExposeMapCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ServiceExposeMapCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ServiceExposeMapClient) ById(id string) (*ServiceExposeMap, error) {
	resp := &ServiceExposeMap{}
	err := c.kuladoClient.doById(SERVICE_EXPOSE_MAP_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ServiceExposeMapClient) Delete(container *ServiceExposeMap) error {
	return c.kuladoClient.doResourceDelete(SERVICE_EXPOSE_MAP_TYPE, &container.Resource)
}

func (c *ServiceExposeMapClient) ActionCreate(resource *ServiceExposeMap) (*ServiceExposeMap, error) {

	resp := &ServiceExposeMap{}

	err := c.kuladoClient.doAction(SERVICE_EXPOSE_MAP_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ServiceExposeMapClient) ActionRemove(resource *ServiceExposeMap) (*ServiceExposeMap, error) {

	resp := &ServiceExposeMap{}

	err := c.kuladoClient.doAction(SERVICE_EXPOSE_MAP_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
