package client

const (
	AZURE_CONFIG_TYPE = "azureConfig"
)

type AzureConfig struct {
	Resource

	AvailabilitySet string `json:"availabilitySet,omitempty" yaml:"availability_set,omitempty"`

	ClientId string `json:"clientId,omitempty" yaml:"client_id,omitempty"`

	ClientSecret string `json:"clientSecret,omitempty" yaml:"client_secret,omitempty"`

	CustomData string `json:"customData,omitempty" yaml:"custom_data,omitempty"`

	DockerPort string `json:"dockerPort,omitempty" yaml:"docker_port,omitempty"`

	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`

	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	NoPublicIp bool `json:"noPublicIp,omitempty" yaml:"no_public_ip,omitempty"`

	OpenPort []string `json:"openPort,omitempty" yaml:"open_port,omitempty"`

	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"private_ip_address,omitempty"`

	ResourceGroup string `json:"resourceGroup,omitempty" yaml:"resource_group,omitempty"`

	Size string `json:"size,omitempty" yaml:"size,omitempty"`

	SshUser string `json:"sshUser,omitempty" yaml:"ssh_user,omitempty"`

	StaticPublicIp bool `json:"staticPublicIp,omitempty" yaml:"static_public_ip,omitempty"`

	StorageType string `json:"storageType,omitempty" yaml:"storage_type,omitempty"`

	Subnet string `json:"subnet,omitempty" yaml:"subnet,omitempty"`

	SubnetPrefix string `json:"subnetPrefix,omitempty" yaml:"subnet_prefix,omitempty"`

	SubscriptionId string `json:"subscriptionId,omitempty" yaml:"subscription_id,omitempty"`

	UsePrivateIp bool `json:"usePrivateIp,omitempty" yaml:"use_private_ip,omitempty"`

	Vnet string `json:"vnet,omitempty" yaml:"vnet,omitempty"`
}

type AzureConfigCollection struct {
	Collection
	Data   []AzureConfig `json:"data,omitempty"`
	client *AzureConfigClient
}

type AzureConfigClient struct {
	kuladoClient *KuladoClient
}

type AzureConfigOperations interface {
	List(opts *ListOpts) (*AzureConfigCollection, error)
	Create(opts *AzureConfig) (*AzureConfig, error)
	Update(existing *AzureConfig, updates interface{}) (*AzureConfig, error)
	ById(id string) (*AzureConfig, error)
	Delete(container *AzureConfig) error
}

func newAzureConfigClient(kuladoClient *KuladoClient) *AzureConfigClient {
	return &AzureConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *AzureConfigClient) Create(container *AzureConfig) (*AzureConfig, error) {
	resp := &AzureConfig{}
	err := c.kuladoClient.doCreate(AZURE_CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *AzureConfigClient) Update(existing *AzureConfig, updates interface{}) (*AzureConfig, error) {
	resp := &AzureConfig{}
	err := c.kuladoClient.doUpdate(AZURE_CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *AzureConfigClient) List(opts *ListOpts) (*AzureConfigCollection, error) {
	resp := &AzureConfigCollection{}
	err := c.kuladoClient.doList(AZURE_CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *AzureConfigCollection) Next() (*AzureConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &AzureConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *AzureConfigClient) ById(id string) (*AzureConfig, error) {
	resp := &AzureConfig{}
	err := c.kuladoClient.doById(AZURE_CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *AzureConfigClient) Delete(container *AzureConfig) error {
	return c.kuladoClient.doResourceDelete(AZURE_CONFIG_TYPE, &container.Resource)
}
