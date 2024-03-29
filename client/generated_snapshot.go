package client

const (
	SNAPSHOT_TYPE = "snapshot"
)

type Snapshot struct {
	Resource

	AccountId string `json:"accountId,omitempty" yaml:"account_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	TransitioningProgress int64 `json:"transitioningProgress,omitempty" yaml:"transitioning_progress,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`

	VolumeId string `json:"volumeId,omitempty" yaml:"volume_id,omitempty"`
}

type SnapshotCollection struct {
	Collection
	Data   []Snapshot `json:"data,omitempty"`
	client *SnapshotClient
}

type SnapshotClient struct {
	kuladoClient *KuladoClient
}

type SnapshotOperations interface {
	List(opts *ListOpts) (*SnapshotCollection, error)
	Create(opts *Snapshot) (*Snapshot, error)
	Update(existing *Snapshot, updates interface{}) (*Snapshot, error)
	ById(id string) (*Snapshot, error)
	Delete(container *Snapshot) error

	ActionBackup(*Snapshot, *SnapshotBackupInput) (*Backup, error)

	ActionCreate(*Snapshot) (*Snapshot, error)

	ActionRemove(*Snapshot) (*Snapshot, error)
}

func newSnapshotClient(kuladoClient *KuladoClient) *SnapshotClient {
	return &SnapshotClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SnapshotClient) Create(container *Snapshot) (*Snapshot, error) {
	resp := &Snapshot{}
	err := c.kuladoClient.doCreate(SNAPSHOT_TYPE, container, resp)
	return resp, err
}

func (c *SnapshotClient) Update(existing *Snapshot, updates interface{}) (*Snapshot, error) {
	resp := &Snapshot{}
	err := c.kuladoClient.doUpdate(SNAPSHOT_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SnapshotClient) List(opts *ListOpts) (*SnapshotCollection, error) {
	resp := &SnapshotCollection{}
	err := c.kuladoClient.doList(SNAPSHOT_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SnapshotCollection) Next() (*SnapshotCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SnapshotCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SnapshotClient) ById(id string) (*Snapshot, error) {
	resp := &Snapshot{}
	err := c.kuladoClient.doById(SNAPSHOT_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SnapshotClient) Delete(container *Snapshot) error {
	return c.kuladoClient.doResourceDelete(SNAPSHOT_TYPE, &container.Resource)
}

func (c *SnapshotClient) ActionBackup(resource *Snapshot, input *SnapshotBackupInput) (*Backup, error) {

	resp := &Backup{}

	err := c.kuladoClient.doAction(SNAPSHOT_TYPE, "backup", &resource.Resource, input, resp)

	return resp, err
}

func (c *SnapshotClient) ActionCreate(resource *Snapshot) (*Snapshot, error) {

	resp := &Snapshot{}

	err := c.kuladoClient.doAction(SNAPSHOT_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *SnapshotClient) ActionRemove(resource *Snapshot) (*Snapshot, error) {

	resp := &Snapshot{}

	err := c.kuladoClient.doAction(SNAPSHOT_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}
