package client

const (
	EXTERNAL_HOST_EVENT_TYPE = "externalHostEvent"
)

type ExternalHostEvent struct {
	Resource `yaml:"-"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	DeleteHost bool `json:"deleteHost,omitempty" yaml:"delete_host,omitempty"`

	EventType string `json:"eventType,omitempty" yaml:"event_type,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	HostId string `json:"hostId,omitempty" yaml:"host_id,omitempty"`

	HostLabel string `json:"hostLabel,omitempty" yaml:"host_label,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	ReportedAccountId string `json:"reportedAccountId,omitempty" yaml:"reported_account_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalHostEventCollection struct {
	Collection
	Data   []ExternalHostEvent `json:"data,omitempty"`
	client *ExternalHostEventClient
}

type ExternalHostEventClient struct {
	kuladoClient *KuladoClient
}

type ExternalHostEventOperations interface {
	List(opts *ListOpts) (*ExternalHostEventCollection, error)
	Create(opts *ExternalHostEvent) (*ExternalHostEvent, error)
	Update(existing *ExternalHostEvent, updates interface{}) (*ExternalHostEvent, error)
	ById(id string) (*ExternalHostEvent, error)
	Delete(container *ExternalHostEvent) error
}

func newExternalHostEventClient(kuladoClient *KuladoClient) *ExternalHostEventClient {
	return &ExternalHostEventClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalHostEventClient) Create(container *ExternalHostEvent) (*ExternalHostEvent, error) {
	resp := &ExternalHostEvent{}
	err := c.kuladoClient.doCreate(EXTERNAL_HOST_EVENT_TYPE, container, resp)
	return resp, err
}

func (c *ExternalHostEventClient) Update(existing *ExternalHostEvent, updates interface{}) (*ExternalHostEvent, error) {
	resp := &ExternalHostEvent{}
	err := c.kuladoClient.doUpdate(EXTERNAL_HOST_EVENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalHostEventClient) List(opts *ListOpts) (*ExternalHostEventCollection, error) {
	resp := &ExternalHostEventCollection{}
	err := c.kuladoClient.doList(EXTERNAL_HOST_EVENT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalHostEventCollection) Next() (*ExternalHostEventCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalHostEventCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalHostEventClient) ById(id string) (*ExternalHostEvent, error) {
	resp := &ExternalHostEvent{}
	err := c.kuladoClient.doById(EXTERNAL_HOST_EVENT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalHostEventClient) Delete(container *ExternalHostEvent) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_HOST_EVENT_TYPE, &container.Resource)
}
