package client

const (
	LABEL_TYPE = "label"
)

type Label struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

type LabelCollection struct {
	Collection
	Data   []Label `json:"data,omitempty"`
	client *LabelClient
}

type LabelClient struct {
	kuladoClient *KuladoClient
}

type LabelOperations interface {
	List(opts *ListOpts) (*LabelCollection, error)
	Create(opts *Label) (*Label, error)
	Update(existing *Label, updates interface{}) (*Label, error)
	ById(id string) (*Label, error)
	Delete(container *Label) error

	ActionCreate(*Label) (*Label, error)

	ActionRemove(*Label) (*Label, error)
}

func newLabelClient(kuladoClient *KuladoClient) *LabelClient {
	return &LabelClient{
		kuladoClient: kuladoClient,
	}
}

func (c *LabelClient) Create(container *Label) (*Label, error) {
	resp := &Label{}
	err := c.kuladoClient.doCreate(LABEL_TYPE, container, resp)
	return resp, err
}

func (c *LabelClient) Update(existing *Label, updates interface{}) (*Label, error) {
	resp := &Label{}
	err := c.kuladoClient.doUpdate(LABEL_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *LabelClient) List(opts *ListOpts) (*LabelCollection, error) {
	resp := &LabelCollection{}
	err := c.kuladoClient.doList(LABEL_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *LabelCollection) Next() (*LabelCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &LabelCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *LabelClient) ById(id string) (*Label, error) {
	resp := &Label{}
	err := c.kuladoClient.doById(LABEL_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *LabelClient) Delete(container *Label) error {
	return c.kuladoClient.doResourceDelete(LABEL_TYPE, &container.Resource)
}

func (c *LabelClient) ActionCreate(resource *Label) (*Label, error) {

	resp := &Label{}

	err := c.kuladoClient.doAction(LABEL_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *LabelClient) ActionRemove(resource *Label) (*Label, error) {

	resp := &Label{}

	err := c.kuladoClient.doAction(LABEL_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
