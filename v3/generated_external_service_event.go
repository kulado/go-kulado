package client

const (
	EXTERNAL_SERVICE_EVENT_TYPE = "externalServiceEvent"
)

type ExternalServiceEvent struct {
	Resource `yaml:"-"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Environment interface{} `json:"environment,omitempty" yaml:"environment,omitempty"`

	EventType string `json:"eventType,omitempty" yaml:"event_type,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	ReportedAccountId string `json:"reportedAccountId,omitempty" yaml:"reported_account_id,omitempty"`

	Service interface{} `json:"service,omitempty" yaml:"service,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalServiceEventCollection struct {
	Collection
	Data   []ExternalServiceEvent `json:"data,omitempty"`
	client *ExternalServiceEventClient
}

type ExternalServiceEventClient struct {
	kuladoClient *KuladoClient
}

type ExternalServiceEventOperations interface {
	List(opts *ListOpts) (*ExternalServiceEventCollection, error)
	Create(opts *ExternalServiceEvent) (*ExternalServiceEvent, error)
	Update(existing *ExternalServiceEvent, updates interface{}) (*ExternalServiceEvent, error)
	ById(id string) (*ExternalServiceEvent, error)
	Delete(container *ExternalServiceEvent) error
}

func newExternalServiceEventClient(kuladoClient *KuladoClient) *ExternalServiceEventClient {
	return &ExternalServiceEventClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalServiceEventClient) Create(container *ExternalServiceEvent) (*ExternalServiceEvent, error) {
	resp := &ExternalServiceEvent{}
	err := c.kuladoClient.doCreate(EXTERNAL_SERVICE_EVENT_TYPE, container, resp)
	return resp, err
}

func (c *ExternalServiceEventClient) Update(existing *ExternalServiceEvent, updates interface{}) (*ExternalServiceEvent, error) {
	resp := &ExternalServiceEvent{}
	err := c.kuladoClient.doUpdate(EXTERNAL_SERVICE_EVENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalServiceEventClient) List(opts *ListOpts) (*ExternalServiceEventCollection, error) {
	resp := &ExternalServiceEventCollection{}
	err := c.kuladoClient.doList(EXTERNAL_SERVICE_EVENT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalServiceEventCollection) Next() (*ExternalServiceEventCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalServiceEventCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalServiceEventClient) ById(id string) (*ExternalServiceEvent, error) {
	resp := &ExternalServiceEvent{}
	err := c.kuladoClient.doById(EXTERNAL_SERVICE_EVENT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalServiceEventClient) Delete(container *ExternalServiceEvent) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_SERVICE_EVENT_TYPE, &container.Resource)
}
