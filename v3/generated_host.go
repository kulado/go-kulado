package client

const (
	HOST_TYPE = "host"
)

type Host struct {
	Resource `yaml:"-"`

	AgentId string `json:"agentId,omitempty" yaml:"agent_id,omitempty"`

	AgentIpAddress string `json:"agentIpAddress,omitempty" yaml:"agent_ip_address,omitempty"`

	AgentState string `json:"agentState,omitempty" yaml:"agent_state,omitempty"`

	Amazonec2Config *Amazonec2Config `json:"amazonec2Config,omitempty" yaml:"amazonec2config,omitempty"`

	AuthCertificateAuthority string `json:"authCertificateAuthority,omitempty" yaml:"auth_certificate_authority,omitempty"`

	AuthKey string `json:"authKey,omitempty" yaml:"auth_key,omitempty"`

	AzureConfig *AzureConfig `json:"azureConfig,omitempty" yaml:"azure_config,omitempty"`

	ClusterId string `json:"clusterId,omitempty" yaml:"cluster_id,omitempty"`

	Created string `json:"created,omitempty" yaml:"created,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`

	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	DigitaloceanConfig *DigitaloceanConfig `json:"digitaloceanConfig,omitempty" yaml:"digitalocean_config,omitempty"`

	DockerVersion string `json:"dockerVersion,omitempty" yaml:"docker_version,omitempty"`

	Driver string `json:"driver,omitempty" yaml:"driver,omitempty"`

	EngineEnv map[string]string `json:"engineEnv,omitempty" yaml:"engine_env,omitempty"`

	EngineInsecureRegistry []string `json:"engineInsecureRegistry,omitempty" yaml:"engine_insecure_registry,omitempty"`

	EngineInstallUrl string `json:"engineInstallUrl,omitempty" yaml:"engine_install_url,omitempty"`

	EngineLabel map[string]string `json:"engineLabel,omitempty" yaml:"engine_label,omitempty"`

	EngineOpt map[string]string `json:"engineOpt,omitempty" yaml:"engine_opt,omitempty"`

	EngineRegistryMirror []string `json:"engineRegistryMirror,omitempty" yaml:"engine_registry_mirror,omitempty"`

	EngineStorageDriver string `json:"engineStorageDriver,omitempty" yaml:"engine_storage_driver,omitempty"`

	ExternalId string `json:"externalId,omitempty" yaml:"external_id,omitempty"`

	ExtractedConfig string `json:"extractedConfig,omitempty" yaml:"extracted_config,omitempty"`

	HostTemplateId string `json:"hostTemplateId,omitempty" yaml:"host_template_id,omitempty"`

	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	Imported bool `json:"imported,omitempty" yaml:"imported,omitempty"`

	Info interface{} `json:"info,omitempty" yaml:"info,omitempty"`

	InstanceIds []string `json:"instanceIds,omitempty" yaml:"instance_ids,omitempty"`

	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`

	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	LocalStorageMb int64 `json:"localStorageMb,omitempty" yaml:"local_storage_mb,omitempty"`

	Memory int64 `json:"memory,omitempty" yaml:"memory,omitempty"`

	MilliCpu int64 `json:"milliCpu,omitempty" yaml:"milli_cpu,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	NodeName string `json:"nodeName,omitempty" yaml:"node_name,omitempty"`

	PacketConfig *PacketConfig `json:"packetConfig,omitempty" yaml:"packet_config,omitempty"`

	PublicEndpoints []PublicEndpoint `json:"publicEndpoints,omitempty" yaml:"public_endpoints,omitempty"`

	RemoveTime string `json:"removeTime,omitempty" yaml:"remove_time,omitempty"`

	Removed string `json:"removed,omitempty" yaml:"removed,omitempty"`

	StackId string `json:"stackId,omitempty" yaml:"stack_id,omitempty"`

	State string `json:"state,omitempty" yaml:"state,omitempty"`

	Transitioning string `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`

	TransitioningMessage string `json:"transitioningMessage,omitempty" yaml:"transitioning_message,omitempty"`

	Uuid string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type HostCollection struct {
	Collection
	Data   []Host `json:"data,omitempty"`
	client *HostClient
}

type HostClient struct {
	kuladoClient *KuladoClient
}

type HostOperations interface {
	List(opts *ListOpts) (*HostCollection, error)
	Create(opts *Host) (*Host, error)
	Update(existing *Host, updates interface{}) (*Host, error)
	ById(id string) (*Host, error)
	Delete(container *Host) error

	ActionActivate(*Host) (*Host, error)

	ActionCreate(*Host) (*Host, error)

	ActionDeactivate(*Host) (*Host, error)

	ActionDockersocket(*Host) (*HostAccess, error)

	ActionError(*Host) (*Host, error)

	ActionEvacuate(*Host) (*Host, error)

	ActionProvision(*Host) (*Host, error)

	ActionRemove(*Host) (*Host, error)

	ActionUpdate(*Host) (*Host, error)
}

func newHostClient(kuladoClient *KuladoClient) *HostClient {
	return &HostClient{
		kuladoClient: kuladoClient,
	}
}

func (c *HostClient) Create(container *Host) (*Host, error) {
	resp := &Host{}
	err := c.kuladoClient.doCreate(HOST_TYPE, container, resp)
	return resp, err
}

func (c *HostClient) Update(existing *Host, updates interface{}) (*Host, error) {
	resp := &Host{}
	err := c.kuladoClient.doUpdate(HOST_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HostClient) List(opts *ListOpts) (*HostCollection, error) {
	resp := &HostCollection{}
	err := c.kuladoClient.doList(HOST_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HostCollection) Next() (*HostCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HostCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HostClient) ById(id string) (*Host, error) {
	resp := &Host{}
	err := c.kuladoClient.doById(HOST_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *HostClient) Delete(container *Host) error {
	return c.kuladoClient.doResourceDelete(HOST_TYPE, &container.Resource)
}

func (c *HostClient) ActionActivate(resource *Host) (*Host, error) {

	resp := &Host{}

	err := c.kuladoClient.doAction(HOST_TYPE, "activate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *HostClient) ActionCreate(resource *Host) (*Host, error) {

	resp := &Host{}

	err := c.kuladoClient.doAction(HOST_TYPE, "create", &resource.Resource, nil, resp)

	return resp, err
}

func (c *HostClient) ActionDeactivate(resource *Host) (*Host, error) {

	resp := &Host{}

	err := c.kuladoClient.doAction(HOST_TYPE, "deactivate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *HostClient) ActionDockersocket(resource *Host) (*HostAccess, error) {

	resp := &HostAccess{}

	err := c.kuladoClient.doAction(HOST_TYPE, "dockersocket", &resource.Resource, nil, resp)

	return resp, err
}

func (c *HostClient) ActionError(resource *Host) (*Host, error) {

	resp := &Host{}

	err := c.kuladoClient.doAction(HOST_TYPE, "error", &resource.Resource, nil, resp)

	return resp, err
}

func (c *HostClient) ActionEvacuate(resource *Host) (*Host, error) {

	resp := &Host{}

	err := c.kuladoClient.doAction(HOST_TYPE, "evacuate", &resource.Resource, nil, resp)

	return resp, err
}

func (c *HostClient) ActionProvision(resource *Host) (*Host, error) {

	resp := &Host{}

	err := c.kuladoClient.doAction(HOST_TYPE, "provision", &resource.Resource, nil, resp)

	return resp, err
}

func (c *HostClient) ActionRemove(resource *Host) (*Host, error) {

	resp := &Host{}

	err := c.kuladoClient.doAction(HOST_TYPE, "remove", &resource.Resource, nil, resp)

	return resp, err
}

func (c *HostClient) ActionUpdate(resource *Host) (*Host, error) {

	resp := &Host{}

	err := c.kuladoClient.doAction(HOST_TYPE, "update", &resource.Resource, nil, resp)

	return resp, err
}
