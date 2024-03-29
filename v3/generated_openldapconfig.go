package client

const (
	OPENLDAPCONFIG_TYPE = "openldapconfig"
)

type Openldapconfig struct {
	Resource `yaml:"-"`

	AccessMode string `json:"accessMode,omitempty" yaml:"access_mode,omitempty"`

	AllowedIdentities []Identity `json:"allowedIdentities,omitempty" yaml:"allowed_identities,omitempty"`

	ConnectionTimeout int64 `json:"connectionTimeout,omitempty" yaml:"connection_timeout,omitempty"`

	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	GroupDNField string `json:"groupDNField,omitempty" yaml:"group_dnfield,omitempty"`

	GroupMemberMappingAttribute string `json:"groupMemberMappingAttribute,omitempty" yaml:"group_member_mapping_attribute,omitempty"`

	GroupMemberUserAttribute string `json:"groupMemberUserAttribute,omitempty" yaml:"group_member_user_attribute,omitempty"`

	GroupNameField string `json:"groupNameField,omitempty" yaml:"group_name_field,omitempty"`

	GroupObjectClass string `json:"groupObjectClass,omitempty" yaml:"group_object_class,omitempty"`

	GroupSearchDomain string `json:"groupSearchDomain,omitempty" yaml:"group_search_domain,omitempty"`

	GroupSearchField string `json:"groupSearchField,omitempty" yaml:"group_search_field,omitempty"`

	LoginDomain string `json:"loginDomain,omitempty" yaml:"login_domain,omitempty"`

	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	Port int64 `json:"port,omitempty" yaml:"port,omitempty"`

	Server string `json:"server,omitempty" yaml:"server,omitempty"`

	ServiceAccountPassword string `json:"serviceAccountPassword,omitempty" yaml:"service_account_password,omitempty"`

	ServiceAccountUsername string `json:"serviceAccountUsername,omitempty" yaml:"service_account_username,omitempty"`

	Tls bool `json:"tls,omitempty" yaml:"tls,omitempty"`

	UserDisabledBitMask int64 `json:"userDisabledBitMask,omitempty" yaml:"user_disabled_bit_mask,omitempty"`

	UserEnabledAttribute string `json:"userEnabledAttribute,omitempty" yaml:"user_enabled_attribute,omitempty"`

	UserLoginField string `json:"userLoginField,omitempty" yaml:"user_login_field,omitempty"`

	UserMemberAttribute string `json:"userMemberAttribute,omitempty" yaml:"user_member_attribute,omitempty"`

	UserNameField string `json:"userNameField,omitempty" yaml:"user_name_field,omitempty"`

	UserObjectClass string `json:"userObjectClass,omitempty" yaml:"user_object_class,omitempty"`

	UserSearchField string `json:"userSearchField,omitempty" yaml:"user_search_field,omitempty"`
}

type OpenldapconfigCollection struct {
	Collection
	Data   []Openldapconfig `json:"data,omitempty"`
	client *OpenldapconfigClient
}

type OpenldapconfigClient struct {
	kuladoClient *KuladoClient
}

type OpenldapconfigOperations interface {
	List(opts *ListOpts) (*OpenldapconfigCollection, error)
	Create(opts *Openldapconfig) (*Openldapconfig, error)
	Update(existing *Openldapconfig, updates interface{}) (*Openldapconfig, error)
	ById(id string) (*Openldapconfig, error)
	Delete(container *Openldapconfig) error
}

func newOpenldapconfigClient(kuladoClient *KuladoClient) *OpenldapconfigClient {
	return &OpenldapconfigClient{
		kuladoClient: kuladoClient,
	}
}

func (c *OpenldapconfigClient) Create(container *Openldapconfig) (*Openldapconfig, error) {
	resp := &Openldapconfig{}
	err := c.kuladoClient.doCreate(OPENLDAPCONFIG_TYPE, container, resp)
	return resp, err
}

func (c *OpenldapconfigClient) Update(existing *Openldapconfig, updates interface{}) (*Openldapconfig, error) {
	resp := &Openldapconfig{}
	err := c.kuladoClient.doUpdate(OPENLDAPCONFIG_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *OpenldapconfigClient) List(opts *ListOpts) (*OpenldapconfigCollection, error) {
	resp := &OpenldapconfigCollection{}
	err := c.kuladoClient.doList(OPENLDAPCONFIG_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *OpenldapconfigCollection) Next() (*OpenldapconfigCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &OpenldapconfigCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *OpenldapconfigClient) ById(id string) (*Openldapconfig, error) {
	resp := &Openldapconfig{}
	err := c.kuladoClient.doById(OPENLDAPCONFIG_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *OpenldapconfigClient) Delete(container *Openldapconfig) error {
	return c.kuladoClient.doResourceDelete(OPENLDAPCONFIG_TYPE, &container.Resource)
}
