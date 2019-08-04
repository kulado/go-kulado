package client

const (
	METADATA_SYNC_REQUEST_TYPE = "metadataSyncRequest"
)

type MetadataSyncRequest struct {
	Resource `yaml:"-"`

	Full bool `json:"full,omitempty" yaml:"full,omitempty"`

	Generation string `json:"generation,omitempty" yaml:"generation,omitempty"`

	Removes map[string]interface{} `json:"removes,omitempty" yaml:"removes,omitempty"`

	Updates map[string]interface{} `json:"updates,omitempty" yaml:"updates,omitempty"`
}

type MetadataSyncRequestCollection struct {
	Collection
	Data   []MetadataSyncRequest `json:"data,omitempty"`
	client *MetadataSyncRequestClient
}

type MetadataSyncRequestClient struct {
	kuladoClient *KuladoClient
}

type MetadataSyncRequestOperations interface {
	List(opts *ListOpts) (*MetadataSyncRequestCollection, error)
	Create(opts *MetadataSyncRequest) (*MetadataSyncRequest, error)
	Update(existing *MetadataSyncRequest, updates interface{}) (*MetadataSyncRequest, error)
	ById(id string) (*MetadataSyncRequest, error)
	Delete(container *MetadataSyncRequest) error
}

func newMetadataSyncRequestClient(kuladoClient *KuladoClient) *MetadataSyncRequestClient {
	return &MetadataSyncRequestClient{
		kuladoClient: kuladoClient,
	}
}

func (c *MetadataSyncRequestClient) Create(container *MetadataSyncRequest) (*MetadataSyncRequest, error) {
	resp := &MetadataSyncRequest{}
	err := c.kuladoClient.doCreate(METADATA_SYNC_REQUEST_TYPE, container, resp)
	return resp, err
}

func (c *MetadataSyncRequestClient) Update(existing *MetadataSyncRequest, updates interface{}) (*MetadataSyncRequest, error) {
	resp := &MetadataSyncRequest{}
	err := c.kuladoClient.doUpdate(METADATA_SYNC_REQUEST_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *MetadataSyncRequestClient) List(opts *ListOpts) (*MetadataSyncRequestCollection, error) {
	resp := &MetadataSyncRequestCollection{}
	err := c.kuladoClient.doList(METADATA_SYNC_REQUEST_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *MetadataSyncRequestCollection) Next() (*MetadataSyncRequestCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &MetadataSyncRequestCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *MetadataSyncRequestClient) ById(id string) (*MetadataSyncRequest, error) {
	resp := &MetadataSyncRequest{}
	err := c.kuladoClient.doById(METADATA_SYNC_REQUEST_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *MetadataSyncRequestClient) Delete(container *MetadataSyncRequest) error {
	return c.kuladoClient.doResourceDelete(METADATA_SYNC_REQUEST_TYPE, &container.Resource)
}
