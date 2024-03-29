package client

const (
	IN_SERVICE_UPGRADE_STRATEGY_TYPE = "inServiceUpgradeStrategy"
)

type InServiceUpgradeStrategy struct {
	Resource `yaml:"-"`

	BatchSize int64 `json:"batchSize,omitempty" yaml:"batch_size,omitempty"`

	IntervalMillis int64 `json:"intervalMillis,omitempty" yaml:"interval_millis,omitempty"`

	LaunchConfig *LaunchConfig `json:"launchConfig,omitempty" yaml:"launch_config,omitempty"`

	PreviousLaunchConfig *LaunchConfig `json:"previousLaunchConfig,omitempty" yaml:"previous_launch_config,omitempty"`

	PreviousSecondaryLaunchConfigs []SecondaryLaunchConfig `json:"previousSecondaryLaunchConfigs,omitempty" yaml:"previous_secondary_launch_configs,omitempty"`

	SecondaryLaunchConfigs []SecondaryLaunchConfig `json:"secondaryLaunchConfigs,omitempty" yaml:"secondary_launch_configs,omitempty"`

	StartFirst bool `json:"startFirst,omitempty" yaml:"start_first,omitempty"`
}

type InServiceUpgradeStrategyCollection struct {
	Collection
	Data   []InServiceUpgradeStrategy `json:"data,omitempty"`
	client *InServiceUpgradeStrategyClient
}

type InServiceUpgradeStrategyClient struct {
	kuladoClient *KuladoClient
}

type InServiceUpgradeStrategyOperations interface {
	List(opts *ListOpts) (*InServiceUpgradeStrategyCollection, error)
	Create(opts *InServiceUpgradeStrategy) (*InServiceUpgradeStrategy, error)
	Update(existing *InServiceUpgradeStrategy, updates interface{}) (*InServiceUpgradeStrategy, error)
	ById(id string) (*InServiceUpgradeStrategy, error)
	Delete(container *InServiceUpgradeStrategy) error
}

func newInServiceUpgradeStrategyClient(kuladoClient *KuladoClient) *InServiceUpgradeStrategyClient {
	return &InServiceUpgradeStrategyClient{
		kuladoClient: kuladoClient,
	}
}

func (c *InServiceUpgradeStrategyClient) Create(container *InServiceUpgradeStrategy) (*InServiceUpgradeStrategy, error) {
	resp := &InServiceUpgradeStrategy{}
	err := c.kuladoClient.doCreate(IN_SERVICE_UPGRADE_STRATEGY_TYPE, container, resp)
	return resp, err
}

func (c *InServiceUpgradeStrategyClient) Update(existing *InServiceUpgradeStrategy, updates interface{}) (*InServiceUpgradeStrategy, error) {
	resp := &InServiceUpgradeStrategy{}
	err := c.kuladoClient.doUpdate(IN_SERVICE_UPGRADE_STRATEGY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *InServiceUpgradeStrategyClient) List(opts *ListOpts) (*InServiceUpgradeStrategyCollection, error) {
	resp := &InServiceUpgradeStrategyCollection{}
	err := c.kuladoClient.doList(IN_SERVICE_UPGRADE_STRATEGY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *InServiceUpgradeStrategyCollection) Next() (*InServiceUpgradeStrategyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &InServiceUpgradeStrategyCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *InServiceUpgradeStrategyClient) ById(id string) (*InServiceUpgradeStrategy, error) {
	resp := &InServiceUpgradeStrategy{}
	err := c.kuladoClient.doById(IN_SERVICE_UPGRADE_STRATEGY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *InServiceUpgradeStrategyClient) Delete(container *InServiceUpgradeStrategy) error {
	return c.kuladoClient.doResourceDelete(IN_SERVICE_UPGRADE_STRATEGY_TYPE, &container.Resource)
}
