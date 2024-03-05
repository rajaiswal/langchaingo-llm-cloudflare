package cloudflareclient

type Client struct {
	AccountID          string
	Token              string
	ModelName          string
	EmbeddingModelName string
	Url                string
}

func NewClient(accountID, url, token, modelName, embeddingModelName string) *Client {
	return &Client{
		AccountID:          accountID,
		Url:                url,
		Token:              token,
		ModelName:          modelName,
		EmbeddingModelName: embeddingModelName,
	}
}
