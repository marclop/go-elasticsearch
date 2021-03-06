// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// GetOption is a non-required Get option that gets applied to an HTTP request.
type GetOption func(r *transport.Request)

// WithGetFeature - a comma-separated list of features.
func WithGetFeature(feature []string) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetAllowNoIndices - ignore if a wildcard expression resolves to no concrete indices (default: false).
func WithGetAllowNoIndices(allowNoIndices bool) GetOption {
	return func(r *transport.Request) {
	}
}

// GetExpandWildcards - whether wildcard expressions should get expanded to open or closed indices (default: open).
type GetExpandWildcards int

const (
	// GetExpandWildcardsOpen can be used to set GetExpandWildcards to "open"
	GetExpandWildcardsOpen = iota
	// GetExpandWildcardsClosed can be used to set GetExpandWildcards to "closed"
	GetExpandWildcardsClosed = iota
	// GetExpandWildcardsNone can be used to set GetExpandWildcards to "none"
	GetExpandWildcardsNone = iota
	// GetExpandWildcardsAll can be used to set GetExpandWildcards to "all"
	GetExpandWildcardsAll = iota
)

// WithGetExpandWildcards - whether wildcard expressions should get expanded to open or closed indices (default: open).
func WithGetExpandWildcards(expandWildcards GetExpandWildcards) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetFlatSettings - return settings in flat format (default: false).
func WithGetFlatSettings(flatSettings bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetIgnoreUnavailable - ignore unavailable indexes (default: false).
func WithGetIgnoreUnavailable(ignoreUnavailable bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetIncludeDefaults - whether to return all default setting for each of the indices.
func WithGetIncludeDefaults(includeDefaults bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetLocal - return local information, do not retrieve the state from master node (default: false).
func WithGetLocal(local bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetErrorTrace - include the stack trace of returned errors.
func WithGetErrorTrace(errorTrace bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetFilterPath - a comma-separated list of filters used to reduce the respone.
func WithGetFilterPath(filterPath []string) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetHuman - return human readable values for statistics.
func WithGetHuman(human bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetIgnore - ignores the specified HTTP status codes.
func WithGetIgnore(ignore []int) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetPretty - pretty format the returned JSON response.
func WithGetPretty(pretty bool) GetOption {
	return func(r *transport.Request) {
	}
}

// WithGetSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithGetSourceParam(sourceParam string) GetOption {
	return func(r *transport.Request) {
	}
}

// Get - the get index API allows to retrieve information about one or more indexes. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-get-index.html for more info.
//
// index: a comma-separated list of index names.
//
// options: optional parameters.
func (i *Indices) Get(index []string, options ...GetOption) (*GetResponse, error) {
	req := i.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := i.transport.Do(req)
	return &GetResponse{resp}, err
}

// GetResponse is the response for Get.
type GetResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *GetResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
