package client

const (
	ROLLING_RESTART_STRATEGY_TYPE = "rollingRestartStrategy"
)

type RollingRestartStrategy struct {
	Resource

	BatchSize int64 `json:"batchSize,omitempty" yaml:"batch_size,omitempty"`

	IntervalMillis int64 `json:"intervalMillis,omitempty" yaml:"interval_millis,omitempty"`
}

type RollingRestartStrategyCollection struct {
	Collection
	Data   []RollingRestartStrategy `json:"data,omitempty"`
	client *RollingRestartStrategyClient
}

type RollingRestartStrategyClient struct {
	kuladoClient *KuladoClient
}

type RollingRestartStrategyOperations interface {
	List(opts *ListOpts) (*RollingRestartStrategyCollection, error)
	Create(opts *RollingRestartStrategy) (*RollingRestartStrategy, error)
	Update(existing *RollingRestartStrategy, updates interface{}) (*RollingRestartStrategy, error)
	ById(id string) (*RollingRestartStrategy, error)
	Delete(container *RollingRestartStrategy) error
}

func newRollingRestartStrategyClient(kuladoClient *KuladoClient) *RollingRestartStrategyClient {
	return &RollingRestartStrategyClient{
		kuladoClient: kuladoClient,
	}
}

func (c *RollingRestartStrategyClient) Create(container *RollingRestartStrategy) (*RollingRestartStrategy, error) {
	resp := &RollingRestartStrategy{}
	err := c.kuladoClient.doCreate(ROLLING_RESTART_STRATEGY_TYPE, container, resp)
	return resp, err
}

func (c *RollingRestartStrategyClient) Update(existing *RollingRestartStrategy, updates interface{}) (*RollingRestartStrategy, error) {
	resp := &RollingRestartStrategy{}
	err := c.kuladoClient.doUpdate(ROLLING_RESTART_STRATEGY_TYPE, &existing.Resource, updates, resp)
	return resp, err
}

func (c *RollingRestartStrategyClient) List(opts *ListOpts) (*RollingRestartStrategyCollection, error) {
	resp := &RollingRestartStrategyCollection{}
	err := c.kuladoClient.doList(ROLLING_RESTART_STRATEGY_TYPE, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *RollingRestartStrategyCollection) Next() (*RollingRestartStrategyCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &RollingRestartStrategyCollection{}
		err := cc.client.kuladoClient.doNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *RollingRestartStrategyClient) ById(id string) (*RollingRestartStrategy, error) {
	resp := &RollingRestartStrategy{}
	err := c.kuladoClient.doById(ROLLING_RESTART_STRATEGY_TYPE, id, resp)
	if apiError, ok := err.(*ApiError); ok {
		if apiError.StatusCode == 404 {
			return nil, nil
		}
	}
	return resp, err
}

func (c *RollingRestartStrategyClient) Delete(container *RollingRestartStrategy) error {
	return c.kuladoClient.doResourceDelete(ROLLING_RESTART_STRATEGY_TYPE, &container.Resource)
}
