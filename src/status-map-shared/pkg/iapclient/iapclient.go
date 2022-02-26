package iapclient

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/api/idtoken"
)

type client struct {
	audience string
}

func New(audience string) *client {
	return &client{
		audience: audience,
	}
}

func (c *client) Apply(ctx context.Context, request *http.Request) ([]byte, error) {
	client, err := idtoken.NewClient(ctx, c.audience)
	if err != nil {
		return nil, fmt.Errorf("idtoken.NewClient: %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %v", err)
	}
	defer response.Body.Close()

	out, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return out, nil
}
