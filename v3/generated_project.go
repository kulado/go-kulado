package client

const (
	PROJECT_TYPE = "project"
)

type Project struct {
	Resource `yaml:"-"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	ClusterOwner bool `json:"clusterOwner,omitempty" yaml:"cluster_owner,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Members []ProjectMember `json:"members,omitempty" yaml:"members,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	ProjectLinks []string `json:"projectLinks,omitempty" yaml:"project_links,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	ServicesPortRange *ServicesPortRange `json:"servicesPortRange,omitempty" yaml:"services_port_range,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

type ProjectCollection struct {
	Collection
	Data   []Project `json:"data,omitempty"`
	client *ProjectClient
}

type ProjectClient struct {
	kuladoClient *KuladoClient
}

type ProjectOperations interface {
	List(opts *ListOpts) (*ProjectCollection, error)
	Create(opts *Project) (*Project, error)
	Update(existing *Project, updates interface{}) (*Project, error)
	ById(id string) (*Project, error)
	Delete(container *Project) error

	ActionActivate(*Project) (*Account, error)

	ActionCreate(*Project) (*Account, error)

	ActionDeactivate(*Project) (*Account, error)

	ActionPurge(*Project) (*Account, error)

	ActionRemove(*Project) (*Account, error)

	ActionSetmembers(*Project, *SetProjectMembersInput) (*SetProjectMembersInput, error)

	ActionUpdate(*Project) (*Account, error)
}

func newProjectClient(kuladoClient *KuladoClient) *ProjectClient {
	return &ProjectClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ProjectClient) Create(container *Project) (*Project, error) {
	resp := &Project{}
	err := c.kuladoClient.doCreate(PROJECT_TYPE, container, resp)
	return resp, err
}

func (c *ProjectClient) Update(existing *Project, updates interface{}) (*Project, error) {
	resp := &Project{}
	err := c.kuladoClient.doUpdate(PROJECT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProjectClient) List(opts *ListOpts) (*ProjectCollection, error) {
	resp := &ProjectCollection{}
	err := c.kuladoClient.doList(PROJECT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProjectCollection) Next() (*ProjectCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProjectCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProjectClient) ById(id string) (*Project, error) {
	resp := &Project{}
	err := c.kuladoClient.doById(PROJECT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ProjectClient) Delete(container *Project) error {
	return c.kuladoClient.doResourceDelete(PROJECT_TYPE, &container.Resource)
}

func (c *ProjectClient) ActionActivate(resource *Project) (*Account, error) {

	resp := &Account{}

	err := c.kuladoClient.doAction(PROJECT_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ProjectClient) ActionCreate(resource *Project) (*Account, error) {

	resp := &Account{}

	err := c.kuladoClient.doAction(PROJECT_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ProjectClient) ActionDeactivate(resource *Project) (*Account, error) {

	resp := &Account{}

	err := c.kuladoClient.doAction(PROJECT_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ProjectClient) ActionPurge(resource *Project) (*Account, error) {

	resp := &Account{}

	err := c.kuladoClient.doAction(PROJECT_TYPE, "purge", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ProjectClient) ActionRemove(resource *Project) (*Account, error) {

	resp := &Account{}

	err := c.kuladoClient.doAction(PROJECT_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *ProjectClient) ActionSetmembers(resource *Project, input *SetProjectMembersInput) (*SetProjectMembersInput, error) {

	resp := &SetProjectMembersInput{}

	err := c.kuladoClient.doAction(PROJECT_TYPE, "setmembers", &resource.Resource, input, resp)

	return resp, err
}

func (c *ProjectClient) ActionUpdate(resource *Project) (*Account, error) {

	resp := &Account{}

	err := c.kuladoClient.doAction(PROJECT_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
