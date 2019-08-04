package client

const (
	CLUSTER_IDENTITY_TYPE = "clusterIdentity"
)

type ClusterIdentity struct {
	Resource `yaml:"-"`

	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`

	UserId string `json:"userId,omitempty" yaml:"user_id,omitempty"`

	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}

type ClusterIdentityCollection struct {
	Collection
	Data   []ClusterIdentity `json:"data,omitempty"`
	client *ClusterIdentityClient
}

type ClusterIdentityClient struct {
	kuladoClient *KuladoClient
}

type ClusterIdentityOperations interface {
	List(opts *ListOpts) (*ClusterIdentityCollection, error)
	Create(opts *ClusterIdentity) (*ClusterIdentity, error)
	Update(existing *ClusterIdentity, updates interface{}) (*ClusterIdentity, error)
	ById(id string) (*ClusterIdentity, error)
	Delete(container *ClusterIdentity) error
}

func newClusterIdentityClient(kuladoClient *KuladoClient) *ClusterIdentityClient {
	return &ClusterIdentityClient{
		kuladoClient: kuladoClient,
	}
}

func (c *ClusterIdentityClient) Create(container *ClusterIdentity) (*ClusterIdentity, error) {
	resp := &ClusterIdentity{}
	err := c.kuladoClient.doCreate(CLUSTER_IDENTITY_TYPE, container, resp)
	return resp, err
}

func (c *ClusterIdentityClient) Update(existing *ClusterIdentity, updates interface{}) (*ClusterIdentity, error) {
	resp := &ClusterIdentity{}
	err := c.kuladoClient.doUpdate(CLUSTER_IDENTITY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterIdentityClient) List(opts *ListOpts) (*ClusterIdentityCollection, error) {
	resp := &ClusterIdentityCollection{}
	err := c.kuladoClient.doList(CLUSTER_IDENTITY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterIdentityCollection) Next() (*ClusterIdentityCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterIdentityCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterIdentityClient) ById(id string) (*ClusterIdentity, error) {
	resp := &ClusterIdentity{}
	err := c.kuladoClient.doById(CLUSTER_IDENTITY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *ClusterIdentityClient) Delete(container *ClusterIdentity) error {
	return c.kuladoClient.doResourceDelete(CLUSTER_IDENTITY_TYPE, &container.Resource)
}
