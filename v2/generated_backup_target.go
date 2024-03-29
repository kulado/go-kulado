package client

const (
	BACKUP_TARGET_TYPE = "backupTarget"
)

type BackupTarget struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	NfsConfig *NfsConfig `json:"nfsConfig,omitempty" yaml:"nfs_config,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type BackupTargetCollection struct {
	Collection
	Data   []BackupTarget `json:"data,omitempty"`
	client *BackupTargetClient
}

type BackupTargetClient struct {
	kuladoClient *KuladoClient
}

type BackupTargetOperations interface {
	List(opts *ListOpts) (*BackupTargetCollection, error)
	Create(opts *BackupTarget) (*BackupTarget, error)
	Update(existing *BackupTarget, updates interface{}) (*BackupTarget, error)
	ById(id string) (*BackupTarget, error)
	Delete(container *BackupTarget) error

	ActionCreate(*BackupTarget) (*BackupTarget, error)

	ActionRemove(*BackupTarget) (*BackupTarget, error)
}

func newBackupTargetClient(kuladoClient *KuladoClient) *BackupTargetClient {
	return &BackupTargetClient{
		kuladoClient: kuladoClient,
	}
}

func (c *BackupTargetClient) Create(container *BackupTarget) (*BackupTarget, error) {
	resp := &BackupTarget{}
	err := c.kuladoClient.doCreate(BACKUP_TARGET_TYPE, container, resp)
	return resp, err
}

func (c *BackupTargetClient) Update(existing *BackupTarget, updates interface{}) (*BackupTarget, error) {
	resp := &BackupTarget{}
	err := c.kuladoClient.doUpdate(BACKUP_TARGET_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *BackupTargetClient) List(opts *ListOpts) (*BackupTargetCollection, error) {
	resp := &BackupTargetCollection{}
	err := c.kuladoClient.doList(BACKUP_TARGET_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *BackupTargetCollection) Next() (*BackupTargetCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &BackupTargetCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *BackupTargetClient) ById(id string) (*BackupTarget, error) {
	resp := &BackupTarget{}
	err := c.kuladoClient.doById(BACKUP_TARGET_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *BackupTargetClient) Delete(container *BackupTarget) error {
	return c.kuladoClient.doResourceDelete(BACKUP_TARGET_TYPE, &container.Resource)
}

func (c *BackupTargetClient) ActionCreate(resource *BackupTarget) (*BackupTarget, error) {

	resp := &BackupTarget{}

	err := c.kuladoClient.doAction(BACKUP_TARGET_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *BackupTargetClient) ActionRemove(resource *BackupTarget) (*BackupTarget, error) {

	resp := &BackupTarget{}

	err := c.kuladoClient.doAction(BACKUP_TARGET_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
