package client

const (
	AMAZONEC2CONFIG_TYPE = "amazonec2Config"
)

type Amazonec2Config struct {
	Resource `yaml:"-"`

	AccessKey string `json:"accessKey,omitempty" yaml:"access_key,omitempty"`

	Ami string `json:"ami,omitempty" yaml:"ami,omitempty"`

	BlockDurationMinutes string `json:"blockDurationMinutes,omitempty" yaml:"block_duration_minutes,omitempty"`

	DeviceName string `json:"deviceName,omitempty" yaml:"device_name,omitempty"`

	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" yaml:"iam_instance_profile,omitempty"`

	InsecureTransport bool `json:"insecureTransport,omitempty" yaml:"insecure_transport,omitempty"`

	InstanceType string `json:"instanceType,omitempty" yaml:"instance_type,omitempty"`

	KeypairName string `json:"keypairName,omitempty" yaml:"keypair_name,omitempty"`

	Monitoring bool `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`

	OpenPort []string `json:"openPort,omitempty" yaml:"open_port,omitempty"`

	PrivateAddressOnly bool `json:"privateAddressOnly,omitempty" yaml:"private_address_only,omitempty"`

	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	RequestSpotInstance bool `json:"requestSpotInstance,omitempty" yaml:"request_spot_instance,omitempty"`

	Retries string `json:"retries,omitempty" yaml:"retries,omitempty"`

	RootSize string `json:"rootSize,omitempty" yaml:"root_size,omitempty"`

	SecretKey string `json:"secretKey,omitempty" yaml:"secret_key,omitempty"`

	SecurityGroup []string `json:"securityGroup,omitempty" yaml:"security_group,omitempty"`

	SessionToken string `json:"sessionToken,omitempty" yaml:"session_token,omitempty"`

	SpotPrice string `json:"spotPrice,omitempty" yaml:"spot_price,omitempty"`

	SshKeypath string `json:"sshKeypath,omitempty" yaml:"ssh_keypath,omitempty"`

	SshUser string `json:"sshUser,omitempty" yaml:"ssh_user,omitempty"`

	SubnetId string `json:"subnetId,omitempty" yaml:"subnet_id,omitempty"`

	Tags string `json:"tags,omitempty" yaml:"tags,omitempty"`

	UseEbsOptimizedInstance bool `json:"useEbsOptimizedInstance,omitempty" yaml:"use_ebs_optimized_instance,omitempty"`

	UsePrivateAddress bool `json:"usePrivateAddress,omitempty" yaml:"use_private_address,omitempty"`

	Userdata string `json:"userdata,omitempty" yaml:"userdata,omitempty"`

	VolumeType string `json:"volumeType,omitempty" yaml:"volume_type,omitempty"`

	VpcId string `json:"vpcId,omitempty" yaml:"vpc_id,omitempty"`

	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}

type Amazonec2ConfigCollection struct {
	Collection
	Data   []Amazonec2Config `json:"data,omitempty"`
	client *Amazonec2ConfigClient
}

type Amazonec2ConfigClient struct {
	kuladoClient *KuladoClient
}

type Amazonec2ConfigOperations interface {
	List(opts *ListOpts) (*Amazonec2ConfigCollection, error)
	Create(opts *Amazonec2Config) (*Amazonec2Config, error)
	Update(existing *Amazonec2Config, updates interface{}) (*Amazonec2Config, error)
	ById(id string) (*Amazonec2Config, error)
	Delete(container *Amazonec2Config) error
}

func newAmazonec2ConfigClient(kuladoClient *KuladoClient) *Amazonec2ConfigClient {
	return &Amazonec2ConfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *Amazonec2ConfigClient) Create(container *Amazonec2Config) (*Amazonec2Config, error) {
	resp := &Amazonec2Config{}
	err := c.kuladoClient.doCreate(AMAZONEC2CONFIG_TYPE, container, resp)
	return resp, err
}

func (c *Amazonec2ConfigClient) Update(existing *Amazonec2Config, updates interface{}) (*Amazonec2Config, error) {
	resp := &Amazonec2Config{}
	err := c.kuladoClient.doUpdate(AMAZONEC2CONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *Amazonec2ConfigClient) List(opts *ListOpts) (*Amazonec2ConfigCollection, error) {
	resp := &Amazonec2ConfigCollection{}
	err := c.kuladoClient.doList(AMAZONEC2CONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *Amazonec2ConfigCollection) Next() (*Amazonec2ConfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &Amazonec2ConfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *Amazonec2ConfigClient) ById(id string) (*Amazonec2Config, error) {
	resp := &Amazonec2Config{}
	err := c.kuladoClient.doById(AMAZONEC2CONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *Amazonec2ConfigClient) Delete(container *Amazonec2Config) error {
	return c.kuladoClient.doResourceDelete(AMAZONEC2CONFIG_TYPE, &container.Resource)
}
