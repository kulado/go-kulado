package client

const (
	EXTERNAL_HANDLER_TYPE = "externalHandler"
)

type ExternalHandler struct {
	Resource

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Priority int64 `json:"priority,omitempty" yaml:"priority,omitempty"`

	ProcessConfigs []ExternalHandlerProcessConfig `json:"processConfigs,omitempty" yaml:"process_configs,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	Retries int64 `json:"retries,omitempty" yaml:"retries,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	TimeoutMillis int64 `json:"timeoutMillis,omitempty" yaml:"timeout_millis,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalHandlerCollection struct {
	Collection
	Data   []ExternalHandler `json:"data,omitempty"`
	client *ExternalHandlerClient
}

type ExternalHandlerClient struct {
	kuladoClient *KuladoClient
}

type ExternalHandlerOperations interface {
	List(opts *ListOpts) (*ExternalHandlerCollection, error)
	Create(opts *ExternalHandler) (*ExternalHandler, error)
	Update(existing *ExternalHandler, updates interface{}) (*ExternalHandler, error)
	ById(id string) (*ExternalHandler, error)
	Delete(container *ExternalHandler) error

	ActionActivate(*ExternalHandler) (*ExternalHandler, error)

	ActionCreate(*ExternalHandler) (*ExternalHandler, error)

	ActionDeactivate(*ExternalHandler) (*ExternalHandler, error)

	ActionPurge(*ExternalHandler) (*ExternalHandler, error)

	ActionRemove(*ExternalHandler) (*ExternalHandler, error)

	ActionUpdate(*ExternalHandler) (*ExternalHandler, error)
}

func newExternalHandlerClient(kuladoClient *KuladoClient) *ExternalHandlerClient {
	return &ExternalHandlerClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalHandlerClient) Create(container *ExternalHandler) (*ExternalHandler, error) {
	resp := &ExternalHandler{}
	err := c.kuladoClient.doCreate(EXTERNAL_HANDLER_TYPE, container, resp)
	return resp, err
}

func (c *ExternalHandlerClient) Update(existing *ExternalHandler, updates interface{}) (*ExternalHandler, error) {
	resp := &ExternalHandler{}
	err := c.kuladoClient.doUpdate(EXTERNAL_HANDLER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalHandlerClient) List(opts *ListOpts) (*ExternalHandlerCollection, error) {
	resp := &ExternalHandlerCollection{}
	err := c.kuladoClient.doList(EXTERNAL_HANDLER_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalHandlerCollection) Next() (*ExternalHandlerCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalHandlerCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalHandlerClient) ById(id string) (*ExternalHandler, error) {
	resp := &ExternalHandler{}
	err := c.kuladoClient.doById(EXTERNAL_HANDLER_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalHandlerClient) Delete(container *ExternalHandler) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_HANDLER_TYPE, &container.Resource)
}

func (c *ExternalHandlerClient) ActionActivate(resource *ExternalHandler) (*ExternalHandler, error) {

	resp := &ExternalHandler{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerClient) ActionCreate(resource *ExternalHandler) (*ExternalHandler, error) {

	resp := &ExternalHandler{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerClient) ActionDeactivate(resource *ExternalHandler) (*ExternalHandler, error) {

	resp := &ExternalHandler{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerClient) ActionPurge(resource *ExternalHandler) (*ExternalHandler, error) {

	resp := &ExternalHandler{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_TYPE, "purge", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerClient) ActionRemove(resource *ExternalHandler) (*ExternalHandler, error) {

	resp := &ExternalHandler{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerClient) ActionUpdate(resource *ExternalHandler) (*ExternalHandler, error) {

	resp := &ExternalHandler{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
