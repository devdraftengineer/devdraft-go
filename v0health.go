// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/devdraft-go/internal/apijson"
	"github.com/stainless-sdks/devdraft-go/internal/requestconfig"
	"github.com/stainless-sdks/devdraft-go/option"
	"github.com/stainless-sdks/devdraft-go/packages/respjson"
)

// V0HealthService contains methods and other services that help with interacting
// with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0HealthService] method instead.
type V0HealthService struct {
	Options []option.RequestOption
}

// NewV0HealthService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0HealthService(opts ...option.RequestOption) (r V0HealthService) {
	r = V0HealthService{}
	r.Options = opts
	return
}

// Authenticated health check endpoint
func (r *V0HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *V0HealthCheckResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Public health check endpoint
func (r *V0HealthService) CheckPublic(ctx context.Context, opts ...option.RequestOption) (res *V0HealthCheckPublicResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/health/public"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type V0HealthCheckResponse struct {
	// Indicates whether the request was properly authenticated. Always true for this
	// endpoint since authentication is required.
	Authenticated bool `json:"authenticated,required"`
	// Database connectivity status. Shows "connected" when database is accessible,
	// "error" when connection fails.
	//
	// Any of "connected", "error".
	Database V0HealthCheckResponseDatabase `json:"database,required"`
	// Human-readable message describing the current health status and any issues.
	Message string `json:"message,required"`
	// Overall health status of the service. Returns "ok" when healthy, "error" when
	// issues detected.
	//
	// Any of "ok", "error".
	Status V0HealthCheckResponseStatus `json:"status,required"`
	// ISO 8601 timestamp when the health check was performed.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Current version of the API service. Useful for debugging and compatibility
	// checks.
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Authenticated respjson.Field
		Database      respjson.Field
		Message       respjson.Field
		Status        respjson.Field
		Timestamp     respjson.Field
		Version       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V0HealthCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *V0HealthCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Database connectivity status. Shows "connected" when database is accessible,
// "error" when connection fails.
type V0HealthCheckResponseDatabase string

const (
	V0HealthCheckResponseDatabaseConnected V0HealthCheckResponseDatabase = "connected"
	V0HealthCheckResponseDatabaseError     V0HealthCheckResponseDatabase = "error"
)

// Overall health status of the service. Returns "ok" when healthy, "error" when
// issues detected.
type V0HealthCheckResponseStatus string

const (
	V0HealthCheckResponseStatusOk    V0HealthCheckResponseStatus = "ok"
	V0HealthCheckResponseStatusError V0HealthCheckResponseStatus = "error"
)

type V0HealthCheckPublicResponse struct {
	// Basic health status of the service. Returns "ok" when the service is responding.
	//
	// Any of "ok", "error".
	Status V0HealthCheckPublicResponseStatus `json:"status,required"`
	// ISO 8601 timestamp when the health check was performed.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Current version of the API service.
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Timestamp   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V0HealthCheckPublicResponse) RawJSON() string { return r.JSON.raw }
func (r *V0HealthCheckPublicResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic health status of the service. Returns "ok" when the service is responding.
type V0HealthCheckPublicResponseStatus string

const (
	V0HealthCheckPublicResponseStatusOk    V0HealthCheckPublicResponseStatus = "ok"
	V0HealthCheckPublicResponseStatusError V0HealthCheckPublicResponseStatus = "error"
)
