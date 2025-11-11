// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/devdraft-go/internal/apijson"
	"github.com/stainless-sdks/devdraft-go/internal/apiquery"
	"github.com/stainless-sdks/devdraft-go/internal/requestconfig"
	"github.com/stainless-sdks/devdraft-go/option"
	"github.com/stainless-sdks/devdraft-go/packages/param"
	"github.com/stainless-sdks/devdraft-go/packages/respjson"
)

// V0TaxService contains methods and other services that help with interacting with
// the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0TaxService] method instead.
type V0TaxService struct {
	Options []option.RequestOption
}

// NewV0TaxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV0TaxService(opts ...option.RequestOption) (r V0TaxService) {
	r = V0TaxService{}
	r.Options = opts
	return
}

// Creates a new tax rate that can be applied to products, invoices, and payment
// links.
//
// ## Use Cases
//
// - Set up sales tax for different regions
// - Create VAT rates for international customers
// - Configure state and local tax rates
// - Manage tax compliance requirements
//
// ## Example: Create Basic Sales Tax
//
// ```json
//
//	{
//	  "name": "Sales Tax",
//	  "description": "Standard sales tax for retail transactions",
//	  "percentage": 8.5,
//	  "active": true
//	}
//
// ```
//
// ## Required Fields
//
// - `name`: Tax name for identification (1-100 characters)
// - `percentage`: Tax rate percentage (0-100)
//
// ## Optional Fields
//
// - `description`: Explanation of what this tax covers (max 255 characters)
// - `active`: Whether this tax is currently active (default: true)
// - `appIds`: Array of app IDs where this tax should be available
func (r *V0TaxService) New(ctx context.Context, body V0TaxNewParams, opts ...option.RequestOption) (res *V0TaxNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v0/taxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information about a specific tax.
//
// ## Use Cases
//
// - View tax details
// - Verify tax configuration
// - Check tax status before applying to products
//
// ## Path Parameters
//
// - `id`: Tax UUID (required) - Must be a valid UUID format
//
// ## Example Request
//
// `GET /api/v0/taxes/123e4567-e89b-12d3-a456-426614174000`
//
// ## Example Response
//
// ```json
//
//	{
//	  "id": "tax_123456",
//	  "name": "Sales Tax",
//	  "description": "Standard sales tax for retail transactions",
//	  "percentage": 8.5,
//	  "active": true,
//	  "created_at": "2024-03-20T10:00:00Z",
//	  "updated_at": "2024-03-20T10:00:00Z"
//	}
//
// ```
func (r *V0TaxService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/taxes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Updates an existing tax's information.
//
// ## Use Cases
//
// - Modify tax percentage rates
// - Update tax descriptions
// - Activate/deactivate taxes
// - Change tax names
//
// ## Path Parameters
//
// - `id`: Tax UUID (required) - Must be a valid UUID format
//
// ## Example Request
//
// `PUT /api/v0/taxes/123e4567-e89b-12d3-a456-426614174000`
//
// ## Example Request Body
//
// ```json
//
//	{
//	  "name": "Updated Sales Tax",
//	  "description": "Updated sales tax rate",
//	  "percentage": 9.0,
//	  "active": true
//	}
//
// ```
//
// ## Notes
//
// - Only include fields that need to be updated
// - All fields are optional in updates
// - Percentage changes affect future calculations only
func (r *V0TaxService) Update(ctx context.Context, id string, body V0TaxUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/taxes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Retrieves a list of taxes with optional filtering and pagination.
//
// ## Use Cases
//
// - List all available tax rates
// - Search taxes by name
// - Filter active/inactive taxes
// - Export tax configuration
//
// ## Query Parameters
//
// - `skip`: Number of records to skip (default: 0, min: 0)
// - `take`: Number of records to return (default: 10, min: 1, max: 100)
// - `name`: Filter taxes by name (partial match, case-insensitive)
// - `active`: Filter by active status (true/false)
//
// ## Example Request
//
// `GET /api/v0/taxes?skip=0&take=20&active=true&name=Sales`
//
// ## Example Response
//
// ```json
// [
//
//	{
//	  "id": "tax_123456",
//	  "name": "Sales Tax",
//	  "description": "Standard sales tax for retail transactions",
//	  "percentage": 8.5,
//	  "active": true,
//	  "created_at": "2024-03-20T10:00:00Z",
//	  "updated_at": "2024-03-20T10:00:00Z"
//	}
//
// ]
// ```
func (r *V0TaxService) List(ctx context.Context, query V0TaxListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/taxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Deletes an existing tax.
//
// ## Use Cases
//
// - Remove obsolete tax rates
// - Clean up unused taxes
// - Comply with regulatory changes
//
// ## Path Parameters
//
// - `id`: Tax UUID (required) - Must be a valid UUID format
//
// ## Example Request
//
// `DELETE /api/v0/taxes/123e4567-e89b-12d3-a456-426614174000`
//
// ## Warning
//
// This action cannot be undone. Consider deactivating the tax instead of deleting
// it if it has been used in transactions.
func (r *V0TaxService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/taxes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// This endpoint requires a tax ID in the URL path. Use DELETE /api/v0/taxes/{id}
// instead.
func (r *V0TaxService) DeleteAll(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/taxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// This endpoint requires a tax ID in the URL path. Use PUT /api/v0/taxes/{id}
// instead.
func (r *V0TaxService) UpdateAll(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/taxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

type V0TaxNewResponse struct {
	ID          string    `json:"id"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	Percentage  float64   `json:"percentage"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Active      respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Percentage  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V0TaxNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V0TaxNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0TaxNewParams struct {
	// Tax name. Used to identify and reference this tax rate.
	Name string `json:"name,required"`
	// Tax percentage rate. Must be between 0 and 100.
	Percentage float64 `json:"percentage,required"`
	// Whether this tax is currently active and can be applied.
	Active param.Opt[bool] `json:"active,omitzero"`
	// Optional description explaining what this tax covers.
	Description param.Opt[string] `json:"description,omitzero"`
	// Array of app IDs where this tax should be available. If not provided, tax will
	// be available for the current app.
	AppIDs []string `json:"appIds,omitzero" format:"uuid"`
	paramObj
}

func (r V0TaxNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V0TaxNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0TaxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0TaxUpdateParams struct {
	// Whether this tax is currently active and can be applied
	Active param.Opt[bool] `json:"active,omitzero"`
	// Detailed description of what this tax covers
	Description param.Opt[string] `json:"description,omitzero"`
	// Tax name for identification and display purposes
	Name param.Opt[string] `json:"name,omitzero"`
	// Tax rate as a percentage (0-100)
	Percentage param.Opt[float64] `json:"percentage,omitzero"`
	// Array of app IDs where this tax should be available
	AppIDs []string `json:"appIds,omitzero" format:"uuid"`
	paramObj
}

func (r V0TaxUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V0TaxUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0TaxUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0TaxListParams struct {
	// Filter by active status
	Active param.Opt[bool] `query:"active,omitzero" json:"-"`
	// Filter taxes by name (partial match, case-insensitive)
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of records to skip for pagination
	Skip param.Opt[float64] `query:"skip,omitzero" json:"-"`
	// Number of records to return (max 100)
	Take param.Opt[float64] `query:"take,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0TaxListParams]'s query parameters as `url.Values`.
func (r V0TaxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
