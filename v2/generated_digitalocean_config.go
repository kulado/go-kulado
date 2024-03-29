package client

const (
	DIGITALOCEAN_CONFIG_TYPE = "digitaloceanConfig"
)

type DigitaloceanConfig struct {
	Resource

	AccessToken string `json:"accessToken,omitempty" yaml:"access_token,omitempty"`

	Backups bool `json:"backups,omitempty" yaml:"backups,omitempty"`

	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	Ipv6 bool `json:"ipv6,omitempty" yaml:"ipv6,omitempty"`

	PrivateNetworking bool `json:"privateNetworking,omitempty" yaml:"private_networking,omitempty"`

	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	Size string `json:"size,omitempty" yaml:"size,omitempty"`

	SshKeyFingerprint string `json:"sshKeyFingerprint,omitempty" yaml:"ssh_key_fingerprint,omitempty"`

	SshKeyPath string `json:"sshKeyPath,omitempty" yaml:"ssh_key_path,omitempty"`

	SshPort string `json:"sshPort,omitempty" yaml:"ssh_port,omitempty"`

	SshUser string `json:"sshUser,omitempty" yaml:"ssh_user,omitempty"`

	Userdata string `json:"userdata,omitempty" yaml:"userdata,omitempty"`
}

type DigitaloceanConfigCollection struct {
	Collection
	Data   []DigitaloceanConfig `json:"data,omitempty"`
	client *DigitaloceanConfigClient
}

type DigitaloceanConfigClient struct {
	kuladoClient *KuladoClient
}

type DigitaloceanConfigOperations interface {
	List(opts *ListOpts) (*DigitaloceanConfigCollection, error)
	Create(opts *DigitaloceanConfig) (*DigitaloceanConfig, error)
	Update(existing *DigitaloceanConfig, updates interface{}) (*DigitaloceanConfig, error)
	ById(id string) (*DigitaloceanConfig, error)
	Delete(container *DigitaloceanConfig) error
}

func newDigitaloceanConfigClient(kuladoClient *KuladoClient) *DigitaloceanConfigClient {
	return &DigitaloceanConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *DigitaloceanConfigClient) Create(container *DigitaloceanConfig) (*DigitaloceanConfig, error) {
	resp := &DigitaloceanConfig{}
	err := c.kuladoClient.doCreate(DIGITALOCEAN_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *DigitaloceanConfigClient) Update(existing *DigitaloceanConfig, updates interface{}) (*DigitaloceanConfig, error) {
	resp := &DigitaloceanConfig{}
	err := c.kuladoClient.doUpdate(DIGITALOCEAN_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *DigitaloceanConfigClient) List(opts *ListOpts) (*DigitaloceanConfigCollection, error) {
	resp := &DigitaloceanConfigCollection{}
	err := c.kuladoClient.doList(DIGITALOCEAN_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *DigitaloceanConfigCollection) Next() (*DigitaloceanConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &DigitaloceanConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *DigitaloceanConfigClient) ById(id string) (*DigitaloceanConfig, error) {
	resp := &DigitaloceanConfig{}
	err := c.kuladoClient.doById(DIGITALOCEAN_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *DigitaloceanConfigClient) Delete(container *DigitaloceanConfig) error {
	return c.kuladoClient.doResourceDelete(DIGITALOCEAN_CONFIG_TYPE, &container.Resource)
}
