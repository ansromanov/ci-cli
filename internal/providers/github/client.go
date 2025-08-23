package github

type Client struct {
	Token string
}

func NewClient(token string) *Client {
	return &Client{Token: token}
}
