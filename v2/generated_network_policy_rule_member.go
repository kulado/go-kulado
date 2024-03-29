package client

const (
	NETWORK_POLICY_RULE_MEMBER_TYPE = "networkPolicyRuleMember"
)

type NetworkPolicyRuleMember struct {
	Resource

	Selector string `json:"selector,omitempty" yaml:"selector,omitempty"`
}

type NetworkPolicyRuleMemberCollection struct {
	Collection
	Data   []NetworkPolicyRuleMember `json:"data,omitempty"`
	client *NetworkPolicyRuleMemberClient
}

type NetworkPolicyRuleMemberClient struct {
	kuladoClient *KuladoClient
}

type NetworkPolicyRuleMemberOperations interface {
	List(opts *ListOpts) (*NetworkPolicyRuleMemberCollection, error)
	Create(opts *NetworkPolicyRuleMember) (*NetworkPolicyRuleMember, error)
	Update(existing *NetworkPolicyRuleMember, updates interface{}) (*NetworkPolicyRuleMember, error)
	ById(id string) (*NetworkPolicyRuleMember, error)
	Delete(container *NetworkPolicyRuleMember) error
}

func newNetworkPolicyRuleMemberClient(kuladoClient *KuladoClient) *NetworkPolicyRuleMemberClient {
	return &NetworkPolicyRuleMemberClient{
		kuladoClient: kuladoClient,
	}
}

func (c *NetworkPolicyRuleMemberClient) Create(container *NetworkPolicyRuleMember) (*NetworkPolicyRuleMember, error) {
	resp := &NetworkPolicyRuleMember{}
	err := c.kuladoClient.doCreate(NETWORK_POLICY_RULE_MEMBER_TYPE, container, resp)
	return resp, err
}

func (c *NetworkPolicyRuleMemberClient) Update(existing *NetworkPolicyRuleMember, updates interface{}) (*NetworkPolicyRuleMember, error) {
	resp := &NetworkPolicyRuleMember{}
	err := c.kuladoClient.doUpdate(NETWORK_POLICY_RULE_MEMBER_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *NetworkPolicyRuleMemberClient) List(opts *ListOpts) (*NetworkPolicyRuleMemberCollection, error) {
	resp := &NetworkPolicyRuleMemberCollection{}
	err := c.kuladoClient.doList(NETWORK_POLICY_RULE_MEMBER_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *NetworkPolicyRuleMemberCollection) Next() (*NetworkPolicyRuleMemberCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &NetworkPolicyRuleMemberCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *NetworkPolicyRuleMemberClient) ById(id string) (*NetworkPolicyRuleMember, error) {
	resp := &NetworkPolicyRuleMember{}
	err := c.kuladoClient.doById(NETWORK_POLICY_RULE_MEMBER_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *NetworkPolicyRuleMemberClient) Delete(container *NetworkPolicyRuleMember) error {
	return c.kuladoClient.doResourceDelete(NETWORK_POLICY_RULE_MEMBER_TYPE, &container.Resource)
}
