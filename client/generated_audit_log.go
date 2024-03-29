package client

const (
	AUDIT_LOG_TYPE = "auditLog"
)

type AuditLog struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	AuthType string `json:"authType,omitempty" yaml:"auth_type,omitempty"`

	AuthenticatedAsAccountId string `json:"authenticatedAsAccountId,omitempty" yaml:"authenticated_as_account_id,omitempty"`

	AuthenticatedAsIdentityId string `json:"authenticatedAsIdentityId,omitempty" yaml:"authenticated_as_identity_id,omitempty"`

	ClientIp string `json:"clientIp,omitempty" yaml:"client_ip,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	EventType string `json:"eventType,omitempty" yaml:"event_type,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	RequestObject string `json:"requestObject,omitempty" yaml:"request_object,omitempty"`

	ResourceId int64 `json:"resourceId,omitempty" yaml:"resource_id,omitempty"`

	ResourceType string `json:"resourceType,omitempty" yaml:"resource_type,omitempty"`

	ResponseCode string `json:"responseCode,omitempty" yaml:"response_code,omitempty"`

	ResponseObject string `json:"responseObject,omitempty" yaml:"response_object,omitempty"`
}

type AuditLogCollection struct {
	Collection
	Data   []AuditLog `json:"data,omitempty"`
	client *AuditLogClient
}

type AuditLogClient struct {
	kuladoClient *KuladoClient
}

type AuditLogOperations interface {
	List(opts *ListOpts) (*AuditLogCollection, error)
	Create(opts *AuditLog) (*AuditLog, error)
	Update(existing *AuditLog, updates interface{}) (*AuditLog, error)
	ById(id string) (*AuditLog, error)
	Delete(container *AuditLog) error
}

func newAuditLogClient(kuladoClient *KuladoClient) *AuditLogClient {
	return &AuditLogClient{
		kuladoClient: kuladoClient,
	}
}

func (c *AuditLogClient) Create(container *AuditLog) (*AuditLog, error) {
	resp := &AuditLog{}
	err := c.kuladoClient.doCreate(AUDIT_LOG_TYPE, container, resp)
	return resp, err
}

func (c *AuditLogClient) Update(existing *AuditLog, updates interface{}) (*AuditLog, error) {
	resp := &AuditLog{}
	err := c.kuladoClient.doUpdate(AUDIT_LOG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AuditLogClient) List(opts *ListOpts) (*AuditLogCollection, error) {
	resp := &AuditLogCollection{}
	err := c.kuladoClient.doList(AUDIT_LOG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AuditLogCollection) Next() (*AuditLogCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AuditLogCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AuditLogClient) ById(id string) (*AuditLog, error) {
	resp := &AuditLog{}
	err := c.kuladoClient.doById(AUDIT_LOG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *AuditLogClient) Delete(container *AuditLog) error {
	return c.kuladoClient.doResourceDelete(AUDIT_LOG_TYPE, &container.Resource)
}
