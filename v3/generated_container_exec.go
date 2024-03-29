package client

const (
	CONTAINER_EXEC_TYPE = "containerExec"
)

type ContainerExec struct {
	Resource `yaml:"-"`

	AttachStdin bool `json:"attachStdin,omitempty" yaml:"attach_stdin,omitempty"`

	AttachStdout bool `json:"attachStdout,omitempty" yaml:"attach_stdout,omitempty"`

	Command []string `json:"command,omitempty" yaml:"command,omitempty"`

	Tty bool `json:"tty,omitempty" yaml:"tty,omitempty"`
}

type ContainerExecCollection struct {
	Collection
	Data   []ContainerExec `json:"data,omitempty"`
	client *ContainerExecClient
}

type ContainerExecClient struct {
	kuladoClient *KuladoClient
}

type ContainerExecOperations interface {
	List(opts *ListOpts) (*ContainerExecCollection, error)
	Create(opts *ContainerExec) (*ContainerExec, error)
	Update(existing *ContainerExec, updates interface{}) (*ContainerExec, error)
	ById(id string) (*ContainerExec, error)
	Delete(container *ContainerExec) error
}

func newContainerExecClient(kuladoClient *KuladoClient) *ContainerExecClient {
	return &ContainerExecClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ContainerExecClient) Create(container *ContainerExec) (*ContainerExec, error) {
	resp := &ContainerExec{}
	err := c.kuladoClient.doCreate(CONTAINER_EXEC_TYPE, container, resp)
	return resp, err
}

func (c *ContainerExecClient) Update(existing *ContainerExec, updates interface{}) (*ContainerExec, error) {
	resp := &ContainerExec{}
	err := c.kuladoClient.doUpdate(CONTAINER_EXEC_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ContainerExecClient) List(opts *ListOpts) (*ContainerExecCollection, error) {
	resp := &ContainerExecCollection{}
	err := c.kuladoClient.doList(CONTAINER_EXEC_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ContainerExecCollection) Next() (*ContainerExecCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ContainerExecCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ContainerExecClient) ById(id string) (*ContainerExec, error) {
	resp := &ContainerExec{}
	err := c.kuladoClient.doById(CONTAINER_EXEC_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ContainerExecClient) Delete(container *ContainerExec) error {
	return c.kuladoClient.doResourceDelete(CONTAINER_EXEC_TYPE, &container.Resource)
}
