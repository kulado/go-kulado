package client

const (
	EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE = "externalHandlerExternalHandlerProcessMap"
)

type ExternalHandlerExternalHandlerProcessMap struct {
	Resource

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	EventName string `json:"eventName,omitempty" yaml:"event_name,omitempty"`

	ExternalHandlerId string `json:"externalHandlerId,omitempty" yaml:"external_handler_id,omitempty"`

	ExternalHandlerProcessId string `json:"externalHandlerProcessId,omitempty" yaml:"external_handler_process_id,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	OnError string `json:"onError,omitempty" yaml:"on_error,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalHandlerExternalHandlerProcessMapCollection struct {
	Collection
	Data   []ExternalHandlerExternalHandlerProcessMap `json:"data,omitempty"`
	client *ExternalHandlerExternalHandlerProcessMapClient
}

type ExternalHandlerExternalHandlerProcessMapClient struct {
	kuladoClient *KuladoClient
}

type ExternalHandlerExternalHandlerProcessMapOperations interface {
	List(opts *ListOpts) (*ExternalHandlerExternalHandlerProcessMapCollection, error)
	Create(opts *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error)
	Update(existing *ExternalHandlerExternalHandlerProcessMap, updates interface{}) (*ExternalHandlerExternalHandlerProcessMap, error)
	ById(id string) (*ExternalHandlerExternalHandlerProcessMap, error)
	Delete(container *ExternalHandlerExternalHandlerProcessMap) error

	ActionActivate(*ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error)

	ActionCreate(*ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error)

	ActionDeactivate(*ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error)

	ActionPurge(*ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error)

	ActionRemove(*ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error)

	ActionUpdate(*ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error)
}

func newExternalHandlerExternalHandlerProcessMapClient(kuladoClient *KuladoClient) *ExternalHandlerExternalHandlerProcessMapClient {
	return &ExternalHandlerExternalHandlerProcessMapClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) Create(container *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error) {
	resp := &ExternalHandlerExternalHandlerProcessMap{}
	err := c.kuladoClient.doCreate(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, container, resp)
	return resp, err
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) Update(existing *ExternalHandlerExternalHandlerProcessMap, updates interface{}) (*ExternalHandlerExternalHandlerProcessMap, error) {
	resp := &ExternalHandlerExternalHandlerProcessMap{}
	err := c.kuladoClient.doUpdate(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) List(opts *ListOpts) (*ExternalHandlerExternalHandlerProcessMapCollection, error) {
	resp := &ExternalHandlerExternalHandlerProcessMapCollection{}
	err := c.kuladoClient.doList(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalHandlerExternalHandlerProcessMapCollection) Next() (*ExternalHandlerExternalHandlerProcessMapCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalHandlerExternalHandlerProcessMapCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) ById(id string) (*ExternalHandlerExternalHandlerProcessMap, error) {
	resp := &ExternalHandlerExternalHandlerProcessMap{}
	err := c.kuladoClient.doById(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) Delete(container *ExternalHandlerExternalHandlerProcessMap) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, &container.Resource)
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) ActionActivate(resource *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error) {

	resp := &ExternalHandlerExternalHandlerProcessMap{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) ActionCreate(resource *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error) {

	resp := &ExternalHandlerExternalHandlerProcessMap{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) ActionDeactivate(resource *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error) {

	resp := &ExternalHandlerExternalHandlerProcessMap{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) ActionPurge(resource *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error) {

	resp := &ExternalHandlerExternalHandlerProcessMap{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, "purge", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) ActionRemove(resource *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error) {

	resp := &ExternalHandlerExternalHandlerProcessMap{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerExternalHandlerProcessMapClient) ActionUpdate(resource *ExternalHandlerExternalHandlerProcessMap) (*ExternalHandlerExternalHandlerProcessMap, error) {

	resp := &ExternalHandlerExternalHandlerProcessMap{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_EXTERNAL_HANDLER_PROCESS_MAP_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
