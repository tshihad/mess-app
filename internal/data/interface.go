package data

import "net/http"

// Repo - Data repository
type Repo interface {
}

// Cache := TODO
type Cache interface {
}

// APIRespose defined Success and Fail responses
type APIRespose interface {
	Success(http.ResponseWriter, int, interface{})
	Fail(http.ResponseWriter, int, string, int)
}
