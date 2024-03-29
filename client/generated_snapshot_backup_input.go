package client

const (
	SNAPSHOT_BACKUP_INPUT_TYPE = "snapshotBackupInput"
)

type SnapshotBackupInput struct {
	Resource

	BackupTargetId string `json:"backupTargetId,omitempty" yaml:"backup_target_id,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`
}

type SnapshotBackupInputCollection struct {
	Collection
	Data   []SnapshotBackupInput `json:"data,omitempty"`
	client *SnapshotBackupInputClient
}

type SnapshotBackupInputClient struct {
	kuladoClient *KuladoClient
}

type SnapshotBackupInputOperations interface {
	List(opts *ListOpts) (*SnapshotBackupInputCollection, error)
	Create(opts *SnapshotBackupInput) (*SnapshotBackupInput, error)
	Update(existing *SnapshotBackupInput, updates interface{}) (*SnapshotBackupInput, error)
	ById(id string) (*SnapshotBackupInput, error)
	Delete(container *SnapshotBackupInput) error

	ActionCreate(*SnapshotBackupInput) (*Backup, error)

	ActionRemove(*SnapshotBackupInput) (*Backup, error)
}

func newSnapshotBackupInputClient(kuladoClient *KuladoClient) *SnapshotBackupInputClient {
	return &SnapshotBackupInputClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SnapshotBackupInputClient) Create(container *SnapshotBackupInput) (*SnapshotBackupInput, error) {
	resp := &SnapshotBackupInput{}
	err := c.kuladoClient.doCreate(SNAPSHOT_BACKUP_INPUT_TYPE, container, resp)
	return resp, err
}

func (c *SnapshotBackupInputClient) Update(existing *SnapshotBackupInput, updates interface{}) (*SnapshotBackupInput, error) {
	resp := &SnapshotBackupInput{}
	err := c.kuladoClient.doUpdate(SNAPSHOT_BACKUP_INPUT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SnapshotBackupInputClient) List(opts *ListOpts) (*SnapshotBackupInputCollection, error) {
	resp := &SnapshotBackupInputCollection{}
	err := c.kuladoClient.doList(SNAPSHOT_BACKUP_INPUT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SnapshotBackupInputCollection) Next() (*SnapshotBackupInputCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SnapshotBackupInputCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SnapshotBackupInputClient) ById(id string) (*SnapshotBackupInput, error) {
	resp := &SnapshotBackupInput{}
	err := c.kuladoClient.doById(SNAPSHOT_BACKUP_INPUT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SnapshotBackupInputClient) Delete(container *SnapshotBackupInput) error {
	return c.kuladoClient.doResourceDelete(SNAPSHOT_BACKUP_INPUT_TYPE, &container.Resource)
}

func (c *SnapshotBackupInputClient) ActionCreate(resource *SnapshotBackupInput) (*Backup, error) {

	resp := &Backup{}

	err := c.kuladoClient.doAction(SNAPSHOT_BACKUP_INPUT_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *SnapshotBackupInputClient) ActionRemove(resource *SnapshotBackupInput) (*Backup, error) {

	resp := &Backup{}

	err := c.kuladoClient.doAction(SNAPSHOT_BACKUP_INPUT_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
