package client

const (
	DEPLOYMENT_SYNC_RESPONSE_TYPE = "deploymentSyncResponse"
)

type DeploymentSyncResponse struct {
	Resource `yaml:"-"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	InstanceStatus []InstanceStatus `json:"instanceStatus,omitempty" yaml:"instance_status,omitempty"`

	NodeName string `json:"nodeName,omitempty" yaml:"node_name,omitempty"`
}

type DeploymentSyncResponseCollection struct {
	Collection
	Data   []DeploymentSyncResponse `json:"data,omitempty"`
	client *DeploymentSyncResponseClient
}

type DeploymentSyncResponseClient struct {
	kuladoClient *KuladoClient
}

type DeploymentSyncResponseOperations interface {
	List(opts *ListOpts) (*DeploymentSyncResponseCollection, error)
	Create(opts *DeploymentSyncResponse) (*DeploymentSyncResponse, error)
	Update(existing *DeploymentSyncResponse, updates interface{}) (*DeploymentSyncResponse, error)
	ById(id string) (*DeploymentSyncResponse, error)
	Delete(container *DeploymentSyncResponse) error
}

func newDeploymentSyncResponseClient(kuladoClient *KuladoClient) *DeploymentSyncResponseClient {
	return &DeploymentSyncResponseClient{
		kuladoClient: kuladoClient,
	}
}

func (c *DeploymentSyncResponseClient) Create(container *DeploymentSyncResponse) (*DeploymentSyncResponse, error) {
	resp := &DeploymentSyncResponse{}
	err := c.kuladoClient.doCreate(DEPLOYMENT_SYNC_RESPONSE_TYPE, container, resp)
	return resp, err
}

func (c *DeploymentSyncResponseClient) Update(existing *DeploymentSyncResponse, updates interface{}) (*DeploymentSyncResponse, error) {
	resp := &DeploymentSyncResponse{}
	err := c.kuladoClient.doUpdate(DEPLOYMENT_SYNC_RESPONSE_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DeploymentSyncResponseClient) List(opts *ListOpts) (*DeploymentSyncResponseCollection, error) {
	resp := &DeploymentSyncResponseCollection{}
	err := c.kuladoClient.doList(DEPLOYMENT_SYNC_RESPONSE_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DeploymentSyncResponseCollection) Next() (*DeploymentSyncResponseCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DeploymentSyncResponseCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DeploymentSyncResponseClient) ById(id string) (*DeploymentSyncResponse, error) {
	resp := &DeploymentSyncResponse{}
	err := c.kuladoClient.doById(DEPLOYMENT_SYNC_RESPONSE_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *DeploymentSyncResponseClient) Delete(container *DeploymentSyncResponse) error {
	return c.kuladoClient.doResourceDelete(DEPLOYMENT_SYNC_RESPONSE_TYPE, &container.Resource)
}
