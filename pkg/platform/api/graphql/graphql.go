package graphql

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ActiveState/cli/pkg/platform/api"
	"github.com/ActiveState/cli/pkg/platform/authentication"

	"github.com/machinebox/graphql"
)

type Request interface {
	Query() string
	Vars() map[string]interface{}
}

type Header map[string][]string

type graphqlRequest = graphql.Request

type graphqlClient = graphql.Client

type BearerTokenProvider interface {
	BearerToken() string
}

type GQLClient struct {
	*graphqlClient
	common        Header
	tokenProvider BearerTokenProvider
	timeout       time.Duration
}

// Get is a legacy method, to be removed once we have commands that don't rely on globals
func Get() *GQLClient {
	url := api.GetServiceURL(api.ServiceGraphQL)
	return New(url.String(), map[string][]string{}, authentication.Get(), 0)
}

func New(url string, common Header, bearerToken BearerTokenProvider, timeout time.Duration) *GQLClient {
	if timeout == 0 {
		timeout = time.Second * 60
	}

	gqlClient := graphql.NewClient(url)
	gqlClient.Log = func(s string) {
		fmt.Fprintln(os.Stderr, s)
	}

	return &GQLClient{
		graphqlClient: gqlClient,
		common:        common,
		tokenProvider: bearerToken,
		timeout:       timeout,
	}
}

func (c *GQLClient) Run(request Request, response interface{}) error {
	graphRequest := graphql.NewRequest(request.Query())
	for key, value := range request.Vars() {
		graphRequest.Var(key, value)
	}

	ctx := context.Background()
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(ctx, c.timeout)
	defer cancel()

	bearerToken := c.tokenProvider.BearerToken()
	if bearerToken != "" {
		graphRequest.Header.Set("Authorization", "Bearer "+bearerToken)
	}

	return c.graphqlClient.Run(ctx, graphRequest, response)
}
