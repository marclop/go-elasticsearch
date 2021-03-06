// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package indices

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// GetTemplateOption is a non-required GetTemplate option that gets applied to an HTTP request.
type GetTemplateOption func(r *transport.Request)

// WithGetTemplateName - the comma separated names of the index templates.
func WithGetTemplateName(name []string) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplateFlatSettings - return settings in flat format (default: false).
func WithGetTemplateFlatSettings(flatSettings bool) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplateLocal - return local information, do not retrieve the state from master node (default: false).
func WithGetTemplateLocal(local bool) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplateMasterTimeout - explicit operation timeout for connection to master node.
func WithGetTemplateMasterTimeout(masterTimeout time.Duration) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplateErrorTrace - include the stack trace of returned errors.
func WithGetTemplateErrorTrace(errorTrace bool) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplateFilterPath - a comma-separated list of filters used to reduce the respone.
func WithGetTemplateFilterPath(filterPath []string) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplateHuman - return human readable values for statistics.
func WithGetTemplateHuman(human bool) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplateIgnore - ignores the specified HTTP status codes.
func WithGetTemplateIgnore(ignore []int) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplatePretty - pretty format the returned JSON response.
func WithGetTemplatePretty(pretty bool) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithGetTemplateSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithGetTemplateSourceParam(sourceParam string) GetTemplateOption {
	return func(r *transport.Request) {
	}
}

// GetTemplate - index templates allow you to define templates that will automatically be applied when new indices are created. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/indices-templates.html for more info.
//
// options: optional parameters.
func (i *Indices) GetTemplate(options ...GetTemplateOption) (*GetTemplateResponse, error) {
	req := i.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := i.transport.Do(req)
	return &GetTemplateResponse{resp}, err
}

// GetTemplateResponse is the response for GetTemplate.
type GetTemplateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *GetTemplateResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
