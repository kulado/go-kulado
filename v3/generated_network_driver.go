package client

const (
	NETWORK_DRIVER_TYPE = "networkDriver"
)

type NetworkDriver struct {
	Resource `yaml:"-"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	CniConfig map[string]interface{} `json:"cniConfig,omitempty" yaml:"cni_config,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	DefaultNetwork DefaultNetwork `json:"defaultNetwork,omitempty" yaml:"default_network,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	NetworkMetadata map[string]interface{} `json:"networkMetadata,omitempty" yaml:"network_metadata,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type NetworkDriverCollection struct {
	Collection
	Data   []NetworkDriver `json:"data,omitempty"`
	client *NetworkDriverClient
}

type NetworkDriverClient struct {
	kuladoClient *KuladoClient
}

type NetworkDriverOperations interface {
	List(opts *ListOpts) (*NetworkDriverCollection, error)
	Create(opts *NetworkDriver) (*NetworkDriver, error)
	Update(existing *NetworkDriver, updates interface{}) (*NetworkDriver, error)
	ById(id string) (*NetworkDriver, error)
	Delete(container *NetworkDriver) error

	ActionCreate(*NetworkDriver) (*NetworkDriver, error)

	ActionRemove(*NetworkDriver) (*NetworkDriver, error)
}

func newNetworkDriverClient(kuladoClient *KuladoClient) *NetworkDriverClient {
	return &NetworkDriverClient{
		kuladoClient: kuladoClient,
	}
}

func (c *NetworkDriverClient) Create(container *NetworkDriver) (*NetworkDriver, error) {
	resp := &NetworkDriver{}
	err := c.kuladoClient.doCreate(NETWORK_DRIVER_TYPE, container, resp)
	return resp, err
}

func (c *NetworkDriverClient) Update(existing *NetworkDriver, updates interface{}) (*NetworkDriver, error) {
	resp := &NetworkDriver{}
	err := c.kuladoClient.doUpdate(NETWORK_DRIVER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NetworkDriverClient) List(opts *ListOpts) (*NetworkDriverCollection, error) {
	resp := &NetworkDriverCollection{}
	err := c.kuladoClient.doList(NETWORK_DRIVER_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *NetworkDriverCollection) Next() (*NetworkDriverCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NetworkDriverCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NetworkDriverClient) ById(id string) (*NetworkDriver, error) {
	resp := &NetworkDriver{}
	err := c.kuladoClient.doById(NETWORK_DRIVER_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *NetworkDriverClient) Delete(container *NetworkDriver) error {
	return c.kuladoClient.doResourceDelete(NETWORK_DRIVER_TYPE, &container.Resource)
}

func (c *NetworkDriverClient) ActionCreate(resource *NetworkDriver) (*NetworkDriver, error) {

	resp := &NetworkDriver{}

	err := c.kuladoClient.doAction(NETWORK_DRIVER_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *NetworkDriverClient) ActionRemove(resource *NetworkDriver) (*NetworkDriver, error) {

	resp := &NetworkDriver{}

	err := c.kuladoClient.doAction(NETWORK_DRIVER_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
