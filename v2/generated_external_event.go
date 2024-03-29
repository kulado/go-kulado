package client

const (
	EXTERNAL_EVENT_TYPE = "externalEvent"
)

type ExternalEvent struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	EventType string `json:"eventType,omitempty" yaml:"event_type,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	ReportedAccountId string `json:"reportedAccountId,omitempty" yaml:"reported_account_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalEventCollection struct {
	Collection
	Data   []ExternalEvent `json:"data,omitempty"`
	client *ExternalEventClient
}

type ExternalEventClient struct {
	kuladoClient *KuladoClient
}

type ExternalEventOperations interface {
	List(opts *ListOpts) (*ExternalEventCollection, error)
	Create(opts *ExternalEvent) (*ExternalEvent, error)
	Update(existing *ExternalEvent, updates interface{}) (*ExternalEvent, error)
	ById(id string) (*ExternalEvent, error)
	Delete(container *ExternalEvent) error

	ActionCreate(*ExternalEvent) (*ExternalEvent, error)

	ActionRemove(*ExternalEvent) (*ExternalEvent, error)
}

func newExternalEventClient(kuladoClient *KuladoClient) *ExternalEventClient {
	return &ExternalEventClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalEventClient) Create(container *ExternalEvent) (*ExternalEvent, error) {
	resp := &ExternalEvent{}
	err := c.kuladoClient.doCreate(EXTERNAL_EVENT_TYPE, container, resp)
	return resp, err
}

func (c *ExternalEventClient) Update(existing *ExternalEvent, updates interface{}) (*ExternalEvent, error) {
	resp := &ExternalEvent{}
	err := c.kuladoClient.doUpdate(EXTERNAL_EVENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalEventClient) List(opts *ListOpts) (*ExternalEventCollection, error) {
	resp := &ExternalEventCollection{}
	err := c.kuladoClient.doList(EXTERNAL_EVENT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalEventCollection) Next() (*ExternalEventCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalEventCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalEventClient) ById(id string) (*ExternalEvent, error) {
	resp := &ExternalEvent{}
	err := c.kuladoClient.doById(EXTERNAL_EVENT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalEventClient) Delete(container *ExternalEvent) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_EVENT_TYPE, &container.Resource)
}

func (c *ExternalEventClient) ActionCreate(resource *ExternalEvent) (*ExternalEvent, error) {

	resp := &ExternalEvent{}

	err := c.kuladoClient.doAction(EXTERNAL_EVENT_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalEventClient) ActionRemove(resource *ExternalEvent) (*ExternalEvent, error) {

	resp := &ExternalEvent{}

	err := c.kuladoClient.doAction(EXTERNAL_EVENT_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
