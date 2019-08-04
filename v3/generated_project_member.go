package client

const (
	PROJECT_MEMBER_TYPE = "projectMember"
)

type ProjectMember struct {
	Resource `yaml:"-"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	ExternalIdType string `json:"externalIdType,omitempty" yaml:"external_id_type,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	ProjectId string `json:"projectId,omitempty" yaml:"project_id,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ProjectMemberCollection struct {
	Collection
	Data   []ProjectMember `json:"data,omitempty"`
	client *ProjectMemberClient
}

type ProjectMemberClient struct {
	kuladoClient *KuladoClient
}

type ProjectMemberOperations interface {
	List(opts *ListOpts) (*ProjectMemberCollection, error)
	Create(opts *ProjectMember) (*ProjectMember, error)
	Update(existing *ProjectMember, updates interface{}) (*ProjectMember, error)
	ById(id string) (*ProjectMember, error)
	Delete(container *ProjectMember) error
}

func newProjectMemberClient(kuladoClient *KuladoClient) *ProjectMemberClient {
	return &ProjectMemberClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ProjectMemberClient) Create(container *ProjectMember) (*ProjectMember, error) {
	resp := &ProjectMember{}
	err := c.kuladoClient.doCreate(PROJECT_MEMBER_TYPE, container, resp)
	return resp, err
}

func (c *ProjectMemberClient) Update(existing *ProjectMember, updates interface{}) (*ProjectMember, error) {
	resp := &ProjectMember{}
	err := c.kuladoClient.doUpdate(PROJECT_MEMBER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ProjectMemberClient) List(opts *ListOpts) (*ProjectMemberCollection, error) {
	resp := &ProjectMemberCollection{}
	err := c.kuladoClient.doList(PROJECT_MEMBER_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ProjectMemberCollection) Next() (*ProjectMemberCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ProjectMemberCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ProjectMemberClient) ById(id string) (*ProjectMember, error) {
	resp := &ProjectMember{}
	err := c.kuladoClient.doById(PROJECT_MEMBER_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ProjectMemberClient) Delete(container *ProjectMember) error {
	return c.kuladoClient.doResourceDelete(PROJECT_MEMBER_TYPE, &container.Resource)
}
