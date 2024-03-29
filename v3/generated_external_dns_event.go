package client

const (
	EXTERNAL_DNS_EVENT_TYPE = "externalDnsEvent"
)

type ExternalDnsEvent struct {
	Resource `yaml:"-"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	EventType string `json:"eventType,omitempty" yaml:"event_type,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	ReportedAccountId string `json:"reportedAccountId,omitempty" yaml:"reported_account_id,omitempty"`

	ServiceName string `json:"serviceName,omitempty" yaml:"service_name,omitempty"`

	StackName string `json:"stackName,omitempty" yaml:"stack_name,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ExternalDnsEventCollection struct {
	Collection
	Data   []ExternalDnsEvent `json:"data,omitempty"`
	client *ExternalDnsEventClient
}

type ExternalDnsEventClient struct {
	kuladoClient *KuladoClient
}

type ExternalDnsEventOperations interface {
	List(opts *ListOpts) (*ExternalDnsEventCollection, error)
	Create(opts *ExternalDnsEvent) (*ExternalDnsEvent, error)
	Update(existing *ExternalDnsEvent, updates interface{}) (*ExternalDnsEvent, error)
	ById(id string) (*ExternalDnsEvent, error)
	Delete(container *ExternalDnsEvent) error
}

func newExternalDnsEventClient(kuladoClient *KuladoClient) *ExternalDnsEventClient {
	return &ExternalDnsEventClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ExternalDnsEventClient) Create(container *ExternalDnsEvent) (*ExternalDnsEvent, error) {
	resp := &ExternalDnsEvent{}
	err := c.kuladoClient.doCreate(EXTERNAL_DNS_EVENT_TYPE, container, resp)
	return resp, err
}

func (c *ExternalDnsEventClient) Update(existing *ExternalDnsEvent, updates interface{}) (*ExternalDnsEvent, error) {
	resp := &ExternalDnsEvent{}
	err := c.kuladoClient.doUpdate(EXTERNAL_DNS_EVENT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ExternalDnsEventClient) List(opts *ListOpts) (*ExternalDnsEventCollection, error) {
	resp := &ExternalDnsEventCollection{}
	err := c.kuladoClient.doList(EXTERNAL_DNS_EVENT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ExternalDnsEventCollection) Next() (*ExternalDnsEventCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ExternalDnsEventCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ExternalDnsEventClient) ById(id string) (*ExternalDnsEvent, error) {
	resp := &ExternalDnsEvent{}
	err := c.kuladoClient.doById(EXTERNAL_DNS_EVENT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ExternalDnsEventClient) Delete(container *ExternalDnsEvent) error {
	return c.kuladoClient.doResourceDelete(EXTERNAL_DNS_EVENT_TYPE, &container.Resource)
}
