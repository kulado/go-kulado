package catalog

const (
	TEMPLATE_VERSION_TYPE = "templateVersion"
)

type TemplateVersion struct {
	Resource `yaml:"-"`

	Actions map[string]string `json:"actions,omitempty" yaml:"actions,omitempty"`

	Bindings map[string]string `json:"bindings,omitempty" yaml:"bindings,omitempty"`

	Files map[string]string `json:"files,omitempty" yaml:"files,omitempty"`

	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	Links map[string]string `json:"links,omitempty" yaml:"links,omitempty"`

	MaximumKuladoVersion string `json:"maximumKuladoVersion,omitempty" yaml:"maximum_kulado_version,omitempty"`

	MinimumKuladoVersion string `json:"minimumKuladoVersion,omitempty" yaml:"minimum_kulado_version,omitempty"`

	Questions []Question `json:"questions,omitempty" yaml:"questions,omitempty"`

	TemplateId string `json:"templateId,omitempty" yaml:"template_id,omitempty"`

	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	UpgradeFrom string `json:"upgradeFrom,omitempty" yaml:"upgrade_from,omitempty"`

	UpgradeVersionLinks map[string]string `json:"upgradeVersionLinks,omitempty" yaml:"upgrade_version_links,omitempty"`

	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}

type TemplateVersionCollection struct {
	Collection
	Data   []TemplateVersion `json:"data,omitempty"`
	client *TemplateVersionClient
}

type TemplateVersionClient struct {
	kuladoClient *KuladoClient
}

type TemplateVersionOperations interface {
	List(opts *ListOpts) (*TemplateVersionCollection, error)
	Create(opts *TemplateVersion) (*TemplateVersion, error)
	Update(existing *TemplateVersion, updates interface{}) (*TemplateVersion, error)
	ById(id string) (*TemplateVersion, error)
	Delete(container *TemplateVersion) error
}

func newTemplateVersionClient(kuladoClient *KuladoClient) *TemplateVersionClient {
	return &TemplateVersionClient{
		kuladoClient: kuladoClient,
	}
}

func (c *TemplateVersionClient) Create(container *TemplateVersion) (*TemplateVersion, error) {
	resp := &TemplateVersion{}
	err := c.kuladoClient.doCreate(TEMPLATE_VERSION_TYPE, container, resp)
	return resp, err
}

func (c *TemplateVersionClient) Update(existing *TemplateVersion, updates interface{}) (*TemplateVersion, error) {
	resp := &TemplateVersion{}
	err := c.kuladoClient.doUpdate(TEMPLATE_VERSION_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *TemplateVersionClient) List(opts *ListOpts) (*TemplateVersionCollection, error) {
	resp := &TemplateVersionCollection{}
	err := c.kuladoClient.doList(TEMPLATE_VERSION_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *TemplateVersionCollection) Next() (*TemplateVersionCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &TemplateVersionCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *TemplateVersionClient) ById(id string) (*TemplateVersion, error) {
	resp := &TemplateVersion{}
	err := c.kuladoClient.doById(TEMPLATE_VERSION_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *TemplateVersionClient) Delete(container *TemplateVersion) error {
	return c.kuladoClient.doResourceDelete(TEMPLATE_VERSION_TYPE, &container.Resource)
}
