package client

const (
	EXTERNAL_HANDLER_PROCESS_TYPE = "externalHandlerProcess"
)

type ExternalHandlerProcess struct {
	Resource

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalHandlerProcessCollection struct {
	Collection
	Data   []ExternalHandlerProcess `json:"data,omitempty"`
	client *ExternalHandlerProcessClient
}

type ExternalHandlerProcessClient struct {
	kuladoClient *KuladoClient
}

type ExternalHandlerProcessOperations interface {
	List(opts *ListOpts) (*ExternalHandlerProcessCollection, error)
	Create(opts *ExternalHandlerProcess) (*ExternalHandlerProcess, error)
	Update(existing *ExternalHandlerProcess, updates interface{}) (*ExternalHandlerProcess, error)
	ById(id string) (*ExternalHandlerProcess, error)
	Delete(container *ExternalHandlerProcess) error

	ActionActivate(*ExternalHandlerProcess) (*ExternalHandlerProcess, error)

	ActionCreate(*ExternalHandlerProcess) (*ExternalHandlerProcess, error)

	ActionDeactivate(*ExternalHandlerProcess) (*ExternalHandlerProcess, error)

	ActionPurge(*ExternalHandlerProcess) (*ExternalHandlerProcess, error)

	ActionRemove(*ExternalHandlerProcess) (*ExternalHandlerProcess, error)

	ActionUpdate(*ExternalHandlerProcess) (*ExternalHandlerProcess, error)
}

func newExternalHandlerProcessClient(kuladoClient *KuladoClient) *ExternalHandlerProcessClient {
	return &ExternalHandlerProcessClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalHandlerProcessClient) Create(container *ExternalHandlerProcess) (*ExternalHandlerProcess, error) {
	resp := &ExternalHandlerProcess{}
	err := c.kuladoClient.doCreate(EXTERNAL_HANDLER_PROCESS_TYPE, container, resp)
	return resp, err
}

func (c *ExternalHandlerProcessClient) Update(existing *ExternalHandlerProcess, updates interface{}) (*ExternalHandlerProcess, error) {
	resp := &ExternalHandlerProcess{}
	err := c.kuladoClient.doUpdate(EXTERNAL_HANDLER_PROCESS_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalHandlerProcessClient) List(opts *ListOpts) (*ExternalHandlerProcessCollection, error) {
	resp := &ExternalHandlerProcessCollection{}
	err := c.kuladoClient.doList(EXTERNAL_HANDLER_PROCESS_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalHandlerProcessCollection) Next() (*ExternalHandlerProcessCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalHandlerProcessCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalHandlerProcessClient) ById(id string) (*ExternalHandlerProcess, error) {
	resp := &ExternalHandlerProcess{}
	err := c.kuladoClient.doById(EXTERNAL_HANDLER_PROCESS_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalHandlerProcessClient) Delete(container *ExternalHandlerProcess) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_HANDLER_PROCESS_TYPE, &container.Resource)
}

func (c *ExternalHandlerProcessClient) ActionActivate(resource *ExternalHandlerProcess) (*ExternalHandlerProcess, error) {

	resp := &ExternalHandlerProcess{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_PROCESS_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerProcessClient) ActionCreate(resource *ExternalHandlerProcess) (*ExternalHandlerProcess, error) {

	resp := &ExternalHandlerProcess{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_PROCESS_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerProcessClient) ActionDeactivate(resource *ExternalHandlerProcess) (*ExternalHandlerProcess, error) {

	resp := &ExternalHandlerProcess{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_PROCESS_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerProcessClient) ActionPurge(resource *ExternalHandlerProcess) (*ExternalHandlerProcess, error) {

	resp := &ExternalHandlerProcess{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_PROCESS_TYPE, "purge", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerProcessClient) ActionRemove(resource *ExternalHandlerProcess) (*ExternalHandlerProcess, error) {

	resp := &ExternalHandlerProcess{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_PROCESS_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalHandlerProcessClient) ActionUpdate(resource *ExternalHandlerProcess) (*ExternalHandlerProcess, error) {

	resp := &ExternalHandlerProcess{}

	err := c.kuladoClient.doAction(EXTERNAL_HANDLER_PROCESS_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
