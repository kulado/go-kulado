package catalog

type KuladoClient struct {
	KuladoBaseClient

	ApiVersion      ApiVersionOperations
	Catalog         CatalogOperations
	Template        TemplateOperations
	Question        QuestionOperations
	TemplateVersion TemplateVersionOperations
	Error           ErrorOperations
}

func constructClient(kuladoBaseClient *KuladoBaseClientImpl) *KuladoClient {
	client := &KuladoClient{
		KuladoBaseClient: kuladoBaseClient,
	}

	client.ApiVersion = newApiVersionClient(client)
	client.Catalog = newCatalogClient(client)
	client.Template = newTemplateClient(client)
	client.Question = newQuestionClient(client)
	client.TemplateVersion = newTemplateVersionClient(client)
	client.Error = newErrorClient(client)

	return client
}

func NewKuladoClient(opts *ClientOpts) (*KuladoClient, error) {
	kuladoBaseClient := &KuladoBaseClientImpl{
		Types: map[string]Schema{},
	}
	client := constructClient(kuladoBaseClient)

	err := setupKuladoBaseClient(kuladoBaseClient, opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}
