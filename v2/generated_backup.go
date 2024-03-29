package client

const (
	BACKUP_TYPE = "backup"
)

type Backup struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	BackupTargetId string `json:"backupTargetId,omitempty" yaml:"backup_target_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshot_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	VolumeId string `json:"volumeId,omitempty" yaml:"volume_id,omitempty"`
}

type BackupCollection struct {
	Collection
	Data   []Backup `json:"data,omitempty"`
	client *BackupClient
}

type BackupClient struct {
	kuladoClient *KuladoClient
}

type BackupOperations interface {
	List(opts *ListOpts) (*BackupCollection, error)
	Create(opts *Backup) (*Backup, error)
	Update(existing *Backup, updates interface{}) (*Backup, error)
	ById(id string) (*Backup, error)
	Delete(container *Backup) error

	ActionCreate(*Backup) (*Backup, error)

	ActionRemove(*Backup) (*Backup, error)
}

func newBackupClient(kuladoClient *KuladoClient) *BackupClient {
	return &BackupClient{
		kuladoClient: kuladoClient,
	}
}

func (c *BackupClient) Create(container *Backup) (*Backup, error) {
	resp := &Backup{}
	err := c.kuladoClient.doCreate(BACKUP_TYPE, container, resp)
	return resp, err
}

func (c *BackupClient) Update(existing *Backup, updates interface{}) (*Backup, error) {
	resp := &Backup{}
	err := c.kuladoClient.doUpdate(BACKUP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *BackupClient) List(opts *ListOpts) (*BackupCollection, error) {
	resp := &BackupCollection{}
	err := c.kuladoClient.doList(BACKUP_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *BackupCollection) Next() (*BackupCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &BackupCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *BackupClient) ById(id string) (*Backup, error) {
	resp := &Backup{}
	err := c.kuladoClient.doById(BACKUP_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *BackupClient) Delete(container *Backup) error {
	return c.kuladoClient.doResourceDelete(BACKUP_TYPE, &container.Resource)
}

func (c *BackupClient) ActionCreate(resource *Backup) (*Backup, error) {

	resp := &Backup{}

	err := c.kuladoClient.doAction(BACKUP_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *BackupClient) ActionRemove(resource *Backup) (*Backup, error) {

	resp := &Backup{}

	err := c.kuladoClient.doAction(BACKUP_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
