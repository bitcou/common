package graphql_errors

import "errors"

var (
	// ErrorNoAuth Request does not provide an API KEY on the headers.
	ErrorNoAuth = errors.New("authorized. please request an account and API KEY at bitcou.com")
	// ErrorProductUnavailable product is not available for the client.
	ErrorProductUnavailable = errors.New("product unavailable. please review your account details")
	// ErrorAdminOnly endpoint is reserved for admin only
	ErrorAdminOnly = errors.New("access denied")
)
