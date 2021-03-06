// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// MtermvectorsOption is a non-required Mtermvectors option that gets applied to an HTTP request.
type MtermvectorsOption func(r *transport.Request)

// WithMtermvectorsIndex - the index in which the document resides.
func WithMtermvectorsIndex(index string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsType - the type of the document.
func WithMtermvectorsType(documentType string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsFieldStatistics - specifies if document count, sum of document frequencies and sum of total term frequencies should be returned. Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsFieldStatistics(fieldStatistics bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsFields - a comma-separated list of fields to return. Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsFields(fields []string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsIds - a comma-separated list of documents ids. You must define ids as parameter or set "ids" or "docs" in the request body.
func WithMtermvectorsIds(ids []string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsOffsets - specifies if term offsets should be returned. Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsOffsets(offsets bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsParent - parent id of documents. Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsParent(parent string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsPayloads - specifies if term payloads should be returned. Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsPayloads(payloads bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsPositions - specifies if term positions should be returned. Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsPositions(positions bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsPreference - specify the node or shard the operation should be performed on (default: random) .Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsPreference(preference string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsRealtime - specifies if requests are real-time as opposed to near-real-time (default: true).
func WithMtermvectorsRealtime(realtime bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsRouting - specific routing value. Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsRouting(routing string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsTermStatistics - specifies if total term frequency and document frequency should be returned. Applies to all returned documents unless otherwise specified in body "params" or "docs".
func WithMtermvectorsTermStatistics(termStatistics bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsVersion - explicit version number for concurrency control.
func WithMtermvectorsVersion(version int) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// MtermvectorsVersionType - specific version type.
type MtermvectorsVersionType int

const (
	// MtermvectorsVersionTypeInternal can be used to set MtermvectorsVersionType to "internal"
	MtermvectorsVersionTypeInternal = iota
	// MtermvectorsVersionTypeExternal can be used to set MtermvectorsVersionType to "external"
	MtermvectorsVersionTypeExternal = iota
	// MtermvectorsVersionTypeExternalGte can be used to set MtermvectorsVersionType to "external_gte"
	MtermvectorsVersionTypeExternalGte = iota
	// MtermvectorsVersionTypeForce can be used to set MtermvectorsVersionType to "force"
	MtermvectorsVersionTypeForce = iota
)

// WithMtermvectorsVersionType - specific version type.
func WithMtermvectorsVersionType(versionType MtermvectorsVersionType) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsBody - define ids, documents, parameters or a list of parameters per document here. You must at least provide a list of document ids. See documentation.
func WithMtermvectorsBody(body map[string]interface{}) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsErrorTrace - include the stack trace of returned errors.
func WithMtermvectorsErrorTrace(errorTrace bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsFilterPath - a comma-separated list of filters used to reduce the respone.
func WithMtermvectorsFilterPath(filterPath []string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsHuman - return human readable values for statistics.
func WithMtermvectorsHuman(human bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsIgnore - ignores the specified HTTP status codes.
func WithMtermvectorsIgnore(ignore []int) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsPretty - pretty format the returned JSON response.
func WithMtermvectorsPretty(pretty bool) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// WithMtermvectorsSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithMtermvectorsSourceParam(sourceParam string) MtermvectorsOption {
	return func(r *transport.Request) {
	}
}

// Mtermvectors - multi termvectors API allows to get multiple termvectors at once. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/docs-multi-termvectors.html for more info.
//
// options: optional parameters.
func (a *API) Mtermvectors(options ...MtermvectorsOption) (*MtermvectorsResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &MtermvectorsResponse{resp}, err
}

// MtermvectorsResponse is the response for Mtermvectors.
type MtermvectorsResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *MtermvectorsResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
