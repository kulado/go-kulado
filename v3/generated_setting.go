package client

const (
	SETTING_TYPE = "setting"
)

type Setting struct {
	Resource `yaml:"-"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

type SettingCollection struct {
	Collection
	Data   []Setting `json:"data,omitempty"`
	client *SettingClient
}

type SettingClient struct {
	kuladoClient *KuladoClient
}

type SettingOperations interface {
	List(opts *ListOpts) (*SettingCollection, error)
	Create(opts *Setting) (*Setting, error)
	Update(existing *Setting, updates interface{}) (*Setting, error)
	ById(id string) (*Setting, error)
	Delete(container *Setting) error
}

func newSettingClient(kuladoClient *KuladoClient) *SettingClient {
	return &SettingClient{
		kuladoClient: kuladoClient,
	}
}

func (c *SettingClient) Create(container *Setting) (*Setting, error) {
	resp := &Setting{}
	err := c.kuladoClient.doCreate(SETTING_TYPE, container, resp)
	return resp, err
}

func (c *SettingClient) Update(existing *Setting, updates interface{}) (*Setting, error) {
	resp := &Setting{}
	err := c.kuladoClient.doUpdate(SETTING_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *SettingClient) List(opts *ListOpts) (*SettingCollection, error) {
	resp := &SettingCollection{}
	err := c.kuladoClient.doList(SETTING_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *SettingCollection) Next() (*SettingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &SettingCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *SettingClient) ById(id string) (*Setting, error) {
	resp := &Setting{}
	err := c.kuladoClient.doById(SETTING_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *SettingClient) Delete(container *Setting) error {
	return c.kuladoClient.doResourceDelete(SETTING_TYPE, &container.Resource)
}
