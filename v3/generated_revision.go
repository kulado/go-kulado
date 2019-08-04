package client

const (
	REVISION_TYPE = "revision"
)

type Revision struct {
	Resource `yaml:"-"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Config *Service `json:"config,omitempty" yaml:"config,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	ServiceId string `json:"serviceId,omitempty" yaml:"service_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type RevisionCollection struct {
	Collection
	Data   []Revision `json:"data,omitempty"`
	client *RevisionClient
}

type RevisionClient struct {
	kuladoClient *KuladoClient
}

type RevisionOperations interface {
	List(opts *ListOpts) (*RevisionCollection, error)
	Create(opts *Revision) (*Revision, error)
	Update(existing *Revision, updates interface{}) (*Revision, error)
	ById(id string) (*Revision, error)
	Delete(container *Revision) error
}

func newRevisionClient(kuladoClient *KuladoClient) *RevisionClient {
	return &RevisionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *RevisionClient) Create(container *Revision) (*Revision, error) {
	resp := &Revision{}
	err := c.kuladoClient.doCreate(REVISION_TYPE, container, resp)
	return resp, err
}

func (c *RevisionClient) Update(existing *Revision, updates interface{}) (*Revision, error) {
	resp := &Revision{}
	err := c.kuladoClient.doUpdate(REVISION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RevisionClient) List(opts *ListOpts) (*RevisionCollection, error) {
	resp := &RevisionCollection{}
	err := c.kuladoClient.doList(REVISION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RevisionCollection) Next() (*RevisionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RevisionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RevisionClient) ById(id string) (*Revision, error) {
	resp := &Revision{}
	err := c.kuladoClient.doById(REVISION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RevisionClient) Delete(container *Revision) error {
	return c.kuladoClient.doResourceDelete(REVISION_TYPE, &container.Resource)
}
