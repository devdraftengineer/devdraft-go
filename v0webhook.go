// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/devdraft-go/internal/apijson"
	"github.com/stainless-sdks/devdraft-go/internal/apiquery"
	"github.com/stainless-sdks/devdraft-go/internal/requestconfig"
	"github.com/stainless-sdks/devdraft-go/option"
	"github.com/stainless-sdks/devdraft-go/packages/param"
	"github.com/stainless-sdks/devdraft-go/packages/respjson"
)

// V0WebhookService contains methods and other services that help with interacting
// with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0WebhookService] method instead.
type V0WebhookService struct {
	Options []option.RequestOption
}

// NewV0WebhookService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0WebhookService(opts ...option.RequestOption) (r V0WebhookService) {
	r = V0WebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook endpoint for receiving event notifications. Requires
// webhook:create scope.
func (r *V0WebhookService) New(ctx context.Context, body V0WebhookNewParams, opts ...option.RequestOption) (res *WebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves details for a specific webhook. Requires webhook:read scope.
func (r *V0WebhookService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing webhook configuration. Requires webhook:update scope.
func (r *V0WebhookService) Update(ctx context.Context, id string, body V0WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieves a list of all webhooks for your application. Requires webhook:read
// scope.
func (r *V0WebhookService) List(ctx context.Context, query V0WebhookListParams, opts ...option.RequestOption) (res *[]WebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a webhook configuration. Requires webhook:delete scope.
func (r *V0WebhookService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WebhookResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WebhookResponse struct {
	// Unique identifier for the webhook
	ID string `json:"id,required"`
	// Timestamp when the webhook was created
	CreatedAt string `json:"created_at,required"`
	// Webhook delivery statistics
	DeliveryStats any `json:"delivery_stats,required"`
	// Whether webhook payloads are encrypted
	Encrypted bool `json:"encrypted,required"`
	// Whether the webhook is currently active
	IsActive bool `json:"isActive,required"`
	// Name of the webhook
	Name string `json:"name,required"`
	// Timestamp when the webhook was last updated
	UpdatedAt string `json:"updated_at,required"`
	// URL where webhook events are sent
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		DeliveryStats respjson.Field
		Encrypted     respjson.Field
		IsActive      respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0WebhookNewParams struct {
	// Whether webhook payloads should be encrypted
	Encrypted bool `json:"encrypted,required"`
	// Whether the webhook is active and will receive events
	IsActive bool `json:"isActive,required"`
	// Name of the webhook for identification purposes
	Name string `json:"name,required"`
	// URL where webhook events will be sent
	URL string `json:"url,required" format:"uri"`
	// Secret key used to sign webhook payloads for verification
	SigningSecret param.Opt[string] `json:"signing_secret,omitzero"`
	paramObj
}

func (r V0WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V0WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0WebhookUpdateParams struct {
	// Whether webhook payloads should be encrypted
	Encrypted param.Opt[bool] `json:"encrypted,omitzero"`
	// Whether the webhook is active and will receive events
	IsActive param.Opt[bool] `json:"isActive,omitzero"`
	// Name of the webhook for identification purposes
	Name param.Opt[string] `json:"name,omitzero"`
	// Secret key used to sign webhook payloads for verification
	SigningSecret param.Opt[string] `json:"signing_secret,omitzero"`
	// URL where webhook events will be sent
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	paramObj
}

func (r V0WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V0WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0WebhookListParams struct {
	// Number of records to skip (default: 0)
	Skip param.Opt[float64] `query:"skip,omitzero" json:"-"`
	// Number of records to return (default: 10)
	Take param.Opt[float64] `query:"take,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0WebhookListParams]'s query parameters as `url.Values`.
func (r V0WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
