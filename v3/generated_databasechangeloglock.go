package client

const (
	DATABASECHANGELOGLOCK_TYPE = "databasechangeloglock"
)

type Databasechangeloglock struct {
	Resource `yaml:"-"`

	Locked bool `json:"locked,omitempty" yaml:"locked,omitempty"`

	Lockedby string `json:"lockedby,omitempty" yaml:"lockedby,omitempty"`

	Lockgranted string `json:"lockgranted,omitempty" yaml:"lockgranted,omitempty"`
}

type DatabasechangeloglockCollection struct {
	Collection
	Data   []Databasechangeloglock `json:"data,omitempty"`
	client *DatabasechangeloglockClient
}

type DatabasechangeloglockClient struct {
	kuladoClient *KuladoClient
}

type DatabasechangeloglockOperations interface {
	List(opts *ListOpts) (*DatabasechangeloglockCollection, error)
	Create(opts *Databasechangeloglock) (*Databasechangeloglock, error)
	Update(existing *Databasechangeloglock, updates interface{}) (*Databasechangeloglock, error)
	ById(id string) (*Databasechangeloglock, error)
	Delete(container *Databasechangeloglock) error
}

func newDatabasechangeloglockClient(kuladoClient *KuladoClient) *DatabasechangeloglockClient {
	return &DatabasechangeloglockClient{
		kuladoClient: kuladoClient,
	}
}

func (c *DatabasechangeloglockClient) Create(container *Databasechangeloglock) (*Databasechangeloglock, error) {
	resp := &Databasechangeloglock{}
	err := c.kuladoClient.doCreate(DATABASECHANGELOGLOCK_TYPE, container, resp)
	return resp, err
}

func (c *DatabasechangeloglockClient) Update(existing *Databasechangeloglock, updates interface{}) (*Databasechangeloglock, error) {
	resp := &Databasechangeloglock{}
	err := c.kuladoClient.doUpdate(DATABASECHANGELOGLOCK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DatabasechangeloglockClient) List(opts *ListOpts) (*DatabasechangeloglockCollection, error) {
	resp := &DatabasechangeloglockCollection{}
	err := c.kuladoClient.doList(DATABASECHANGELOGLOCK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DatabasechangeloglockCollection) Next() (*DatabasechangeloglockCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DatabasechangeloglockCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DatabasechangeloglockClient) ById(id string) (*Databasechangeloglock, error) {
	resp := &Databasechangeloglock{}
	err := c.kuladoClient.doById(DATABASECHANGELOGLOCK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *DatabasechangeloglockClient) Delete(container *Databasechangeloglock) error {
	return c.kuladoClient.doResourceDelete(DATABASECHANGELOGLOCK_TYPE, &container.Resource)
}
