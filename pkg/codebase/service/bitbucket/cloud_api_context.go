package bitbucket

import (
	"context"

	bbapi "github.com/khulnasoft/codebase/go-bitbucket"
)

// BuildCloudAPIContext builds context.Context used to call Bitbucket Cloud Code Insights API
func BuildCloudAPIContext(ctx context.Context, user, password, token string) context.Context {
	if user != "" && password != "" {
		ctx = withBasicAuth(ctx, user, password)
	}

	if token != "" {
		ctx = withAccessToken(ctx, token)
	}

	return ctx
}

// WithBasicAuth adds basic auth credentials to context
func withBasicAuth(ctx context.Context, username, password string) context.Context {
	return context.WithValue(ctx, bbapi.ContextBasicAuth,
		bbapi.BasicAuth{
			UserName: username,
			Password: password,
		})
}

// WithAccessToken adds basic auth credentials to context
func withAccessToken(ctx context.Context, accessToken string) context.Context {
	return context.WithValue(ctx, bbapi.ContextAccessToken, accessToken)
}
