package client

const (
	CLUSTER_MEMBERSHIP_TYPE = "clusterMembership"
)

type ClusterMembership struct {
	Resource

	Clustered bool `json:"clustered,omitempty" yaml:"clustered,omitempty"`

	Config string `json:"config,omitempty" yaml:"config,omitempty"`

	Heartbeat int64 `json:"heartbeat,omitempty" yaml:"heartbeat,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ClusterMembershipCollection struct {
	Collection
	Data   []ClusterMembership `json:"data,omitempty"`
	client *ClusterMembershipClient
}

type ClusterMembershipClient struct {
	kuladoClient *KuladoClient
}

type ClusterMembershipOperations interface {
	List(opts *ListOpts) (*ClusterMembershipCollection, error)
	Create(opts *ClusterMembership) (*ClusterMembership, error)
	Update(existing *ClusterMembership, updates interface{}) (*ClusterMembership, error)
	ById(id string) (*ClusterMembership, error)
	Delete(container *ClusterMembership) error
}

func newClusterMembershipClient(kuladoClient *KuladoClient) *ClusterMembershipClient {
	return &ClusterMembershipClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ClusterMembershipClient) Create(container *ClusterMembership) (*ClusterMembership, error) {
	resp := &ClusterMembership{}
	err := c.kuladoClient.doCreate(CLUSTER_MEMBERSHIP_TYPE, container, resp)
	return resp, err
}

func (c *ClusterMembershipClient) Update(existing *ClusterMembership, updates interface{}) (*ClusterMembership, error) {
	resp := &ClusterMembership{}
	err := c.kuladoClient.doUpdate(CLUSTER_MEMBERSHIP_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterMembershipClient) List(opts *ListOpts) (*ClusterMembershipCollection, error) {
	resp := &ClusterMembershipCollection{}
	err := c.kuladoClient.doList(CLUSTER_MEMBERSHIP_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterMembershipCollection) Next() (*ClusterMembershipCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterMembershipCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterMembershipClient) ById(id string) (*ClusterMembership, error) {
	resp := &ClusterMembership{}
	err := c.kuladoClient.doById(CLUSTER_MEMBERSHIP_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ClusterMembershipClient) Delete(container *ClusterMembership) error {
	return c.kuladoClient.doResourceDelete(CLUSTER_MEMBERSHIP_TYPE, &container.Resource)
}
