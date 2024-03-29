package client

const (
	PULL_TASK_TYPE = "pullTask"
)

type PullTask struct {
	Resource `yaml:"-"`

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	HostIds []string `json:"hostIds,omitempty" yaml:"host_ids,omitempty"`

	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Status map[string]string `json:"status,omitempty" yaml:"status,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type PullTaskCollection struct {
	Collection
	Data   []PullTask `json:"data,omitempty"`
	client *PullTaskClient
}

type PullTaskClient struct {
	kuladoClient *KuladoClient
}

type PullTaskOperations interface {
	List(opts *ListOpts) (*PullTaskCollection, error)
	Create(opts *PullTask) (*PullTask, error)
	Update(existing *PullTask, updates interface{}) (*PullTask, error)
	ById(id string) (*PullTask, error)
	Delete(container *PullTask) error

	ActionCreate(*PullTask) (*GenericObject, error)

	ActionRemove(*PullTask) (*GenericObject, error)
}

func newPullTaskClient(kuladoClient *KuladoClient) *PullTaskClient {
	return &PullTaskClient{
		kuladoClient: kuladoClient,
	}
}

func (c *PullTaskClient) Create(container *PullTask) (*PullTask, error) {
	resp := &PullTask{}
	err := c.kuladoClient.doCreate(PULL_TASK_TYPE, container, resp)
	return resp, err
}

func (c *PullTaskClient) Update(existing *PullTask, updates interface{}) (*PullTask, error) {
	resp := &PullTask{}
	err := c.kuladoClient.doUpdate(PULL_TASK_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PullTaskClient) List(opts *ListOpts) (*PullTaskCollection, error) {
	resp := &PullTaskCollection{}
	err := c.kuladoClient.doList(PULL_TASK_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PullTaskCollection) Next() (*PullTaskCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PullTaskCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PullTaskClient) ById(id string) (*PullTask, error) {
	resp := &PullTask{}
	err := c.kuladoClient.doById(PULL_TASK_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *PullTaskClient) Delete(container *PullTask) error {
	return c.kuladoClient.doResourceDelete(PULL_TASK_TYPE, &container.Resource)
}

func (c *PullTaskClient) ActionCreate(resource *PullTask) (*GenericObject, error) {

	resp := &GenericObject{}

	err := c.kuladoClient.doAction(PULL_TASK_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *PullTaskClient) ActionRemove(resource *PullTask) (*GenericObject, error) {

	resp := &GenericObject{}

	err := c.kuladoClient.doAction(PULL_TASK_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
