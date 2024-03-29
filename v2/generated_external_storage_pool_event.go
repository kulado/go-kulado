package client

const (
	EXTERNAL_STORAGE_POOL_EVENT_TYPE = "externalStoragePoolEvent"
)

type ExternalStoragePoolEvent struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	EventType string `json:"eventType,omitempty" yaml:"event_type,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	HostUuids []string `json:"hostUuids,omitempty" yaml:"host_uuids,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	ReportedAccountId string `json:"reportedAccountId,omitempty" yaml:"reported_account_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	StoragePool StoragePool `json:"storagePool,omitempty" yaml:"storage_pool,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalStoragePoolEventCollection struct {
	Collection
	Data   []ExternalStoragePoolEvent `json:"data,omitempty"`
	client *ExternalStoragePoolEventClient
}

type ExternalStoragePoolEventClient struct {
	kuladoClient *KuladoClient
}

type ExternalStoragePoolEventOperations interface {
	List(opts *ListOpts) (*ExternalStoragePoolEventCollection, error)
	Create(opts *ExternalStoragePoolEvent) (*ExternalStoragePoolEvent, error)
	Update(existing *ExternalStoragePoolEvent, updates interface{}) (*ExternalStoragePoolEvent, error)
	ById(id string) (*ExternalStoragePoolEvent, error)
	Delete(container *ExternalStoragePoolEvent) error

	ActionCreate(*ExternalStoragePoolEvent) (*ExternalEvent, error)

	ActionRemove(*ExternalStoragePoolEvent) (*ExternalEvent, error)
}

func newExternalStoragePoolEventClient(kuladoClient *KuladoClient) *ExternalStoragePoolEventClient {
	return &ExternalStoragePoolEventClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalStoragePoolEventClient) Create(container *ExternalStoragePoolEvent) (*ExternalStoragePoolEvent, error) {
	resp := &ExternalStoragePoolEvent{}
	err := c.kuladoClient.doCreate(EXTERNAL_STORAGE_POOL_EVENT_TYPE, container, resp)
	return resp, err
}

func (c *ExternalStoragePoolEventClient) Update(existing *ExternalStoragePoolEvent, updates interface{}) (*ExternalStoragePoolEvent, error) {
	resp := &ExternalStoragePoolEvent{}
	err := c.kuladoClient.doUpdate(EXTERNAL_STORAGE_POOL_EVENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalStoragePoolEventClient) List(opts *ListOpts) (*ExternalStoragePoolEventCollection, error) {
	resp := &ExternalStoragePoolEventCollection{}
	err := c.kuladoClient.doList(EXTERNAL_STORAGE_POOL_EVENT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalStoragePoolEventCollection) Next() (*ExternalStoragePoolEventCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalStoragePoolEventCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalStoragePoolEventClient) ById(id string) (*ExternalStoragePoolEvent, error) {
	resp := &ExternalStoragePoolEvent{}
	err := c.kuladoClient.doById(EXTERNAL_STORAGE_POOL_EVENT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalStoragePoolEventClient) Delete(container *ExternalStoragePoolEvent) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_STORAGE_POOL_EVENT_TYPE, &container.Resource)
}

func (c *ExternalStoragePoolEventClient) ActionCreate(resource *ExternalStoragePoolEvent) (*ExternalEvent, error) {

	resp := &ExternalEvent{}

	err := c.kuladoClient.doAction(EXTERNAL_STORAGE_POOL_EVENT_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ExternalStoragePoolEventClient) ActionRemove(resource *ExternalStoragePoolEvent) (*ExternalEvent, error) {

	resp := &ExternalEvent{}

	err := c.kuladoClient.doAction(EXTERNAL_STORAGE_POOL_EVENT_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
