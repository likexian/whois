package whois

import "errors"

var (
	// ErrDomainEmpty domain is an empty string
	ErrDomainEmpty = errors.New("domain is empty")

	// ErrWhoisServerNotFound no whois server found
	ErrWhoisServerNotFound = errors.New("no whois server found for domain")
)
