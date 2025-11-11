// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/devdraftengineer/devdraft-go/internal/apijson"
	"github.com/devdraftengineer/devdraft-go/internal/apiquery"
	"github.com/devdraftengineer/devdraft-go/internal/requestconfig"
	"github.com/devdraftengineer/devdraft-go/option"
	"github.com/devdraftengineer/devdraft-go/packages/param"
)

// V0CustomerService contains methods and other services that help with interacting
// with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0CustomerService] method instead.
type V0CustomerService struct {
	Options              []option.RequestOption
	LiquidationAddresses V0CustomerLiquidationAddressService
}

// NewV0CustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0CustomerService(opts ...option.RequestOption) (r V0CustomerService) {
	r = V0CustomerService{}
	r.Options = opts
	r.LiquidationAddresses = NewV0CustomerLiquidationAddressService(opts...)
	return
}

// Creates a new customer in the system with their personal and contact
// information.
//
// This endpoint allows you to register new customers and store their details for
// future transactions.
//
// ## Use Cases
//
// - Register new customers for payment processing
// - Store customer information for recurring payments
// - Maintain customer profiles for transaction history
//
// ## Example: Create New Customer
//
// ```json
//
//	{
//	  "first_name": "John",
//	  "last_name": "Doe",
//	  "email": "john.doe@example.com",
//	  "phone_number": "+1-555-123-4567",
//	  "customer_type": "Startup",
//	  "status": "ACTIVE"
//	}
//
// ```
//
// ## Required Fields
//
// - `first_name`: Customer's first name (1-100 characters)
// - `last_name`: Customer's last name (1-100 characters)
// - `phone_number`: Customer's phone number (max 20 characters)
//
// ## Optional Fields
//
//   - `email`: Valid email address (max 255 characters)
//   - `customer_type`: Type of customer account (Individual, Startup, Small
//     Business, Medium Business, Enterprise, Non-Profit, Government)
//   - `status`: Customer status (ACTIVE, BLACKLISTED, DEACTIVATED)
func (r *V0CustomerService) New(ctx context.Context, body V0CustomerNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieves detailed information about a specific customer.
//
// This endpoint allows you to fetch complete customer details including their
// contact information and status.
//
// ## Use Cases
//
// - View customer details
// - Verify customer information
// - Check customer status before processing payments
//
// ## Path Parameters
//
// - `id`: Customer UUID (required) - Must be a valid UUID format
//
// ## Example Request
//
// `GET /api/v0/customers/123e4567-e89b-12d3-a456-426614174000`
//
// ## Example Response
//
// ```json
//
//	{
//	  "id": "cust_123456",
//	  "first_name": "John",
//	  "last_name": "Doe",
//	  "email": "john.doe@example.com",
//	  "phone_number": "+1-555-123-4567",
//	  "customer_type": "Enterprise",
//	  "status": "ACTIVE",
//	  "last_spent": 1250.5,
//	  "last_purchase_date": "2024-03-15T14:30:00Z",
//	  "created_at": "2024-03-20T10:00:00Z",
//	  "updated_at": "2024-03-20T10:00:00Z"
//	}
//
// ```
func (r *V0CustomerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Updates an existing customer's information.
//
// This endpoint allows you to modify customer details while preserving their core
// information.
//
// ## Use Cases
//
// - Update customer contact information
// - Change customer type
// - Modify customer status
//
// ## Path Parameters
//
// - `id`: Customer UUID (required) - Must be a valid UUID format
//
// ## Example Request
//
// `PATCH /api/v0/customers/123e4567-e89b-12d3-a456-426614174000`
//
// ## Example Request Body
//
// ```json
//
//	{
//	  "first_name": "John",
//	  "last_name": "Smith",
//	  "phone_number": "+1-987-654-3210",
//	  "customer_type": "Small Business"
//	}
//
// ```
//
// ## Notes
//
// - Only include fields that need to be updated
// - All fields are optional in updates
// - Status changes may require additional verification
func (r *V0CustomerService) Update(ctx context.Context, id string, body V0CustomerUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/customers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieves a list of customers with optional filtering and pagination.
//
// This endpoint allows you to search and filter customers based on various
// criteria.
//
// ## Use Cases
//
// - List all customers with pagination
// - Search customers by name or email
// - Filter customers by status
// - Export customer data
//
// ## Query Parameters
//
// - `skip`: Number of records to skip (default: 0, min: 0)
// - `take`: Number of records to take (default: 10, min: 1, max: 100)
// - `status`: Filter by customer status (ACTIVE, BLACKLISTED, DEACTIVATED)
// - `name`: Filter by customer name (partial match, case-insensitive)
// - `email`: Filter by customer email (exact match, case-insensitive)
//
// ## Example Request
//
// `GET /api/v0/customers?skip=0&take=20&status=ACTIVE&name=John`
//
// ## Example Response
//
// ```json
//
//	{
//	  "data": [
//	    {
//	      "id": "cust_123456",
//	      "first_name": "John",
//	      "last_name": "Doe",
//	      "email": "john.doe@example.com",
//	      "phone_number": "+1-555-123-4567",
//	      "customer_type": "Startup",
//	      "status": "ACTIVE",
//	      "created_at": "2024-03-20T10:00:00Z",
//	      "updated_at": "2024-03-20T10:00:00Z"
//	    }
//	  ],
//	  "total": 100,
//	  "skip": 0,
//	  "take": 10
//	}
//
// ```
func (r *V0CustomerService) List(ctx context.Context, query V0CustomerListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Current status of the customer account. Controls access to services and
// features.
type CustomerStatus string

const (
	CustomerStatusActive      CustomerStatus = "ACTIVE"
	CustomerStatusBlacklisted CustomerStatus = "BLACKLISTED"
	CustomerStatusDeactivated CustomerStatus = "DEACTIVATED"
	CustomerStatusDeleted     CustomerStatus = "DELETED"
)

// Type of customer account. Determines available features and compliance
// requirements.
type CustomerType string

const (
	CustomerTypeIndividual     CustomerType = "Individual"
	CustomerTypeStartup        CustomerType = "Startup"
	CustomerTypeSmallBusiness  CustomerType = "Small Business"
	CustomerTypeMediumBusiness CustomerType = "Medium Business"
	CustomerTypeEnterprise     CustomerType = "Enterprise"
	CustomerTypeNonProfit      CustomerType = "Non-Profit"
	CustomerTypeGovernment     CustomerType = "Government"
)

type V0CustomerNewParams struct {
	// Customer's first name. Used for personalization and legal documentation.
	FirstName string `json:"first_name,required"`
	// Customer's last name. Used for personalization and legal documentation.
	LastName string `json:"last_name,required"`
	// Customer's phone number. Used for SMS notifications and verification. Include
	// country code for international numbers.
	PhoneNumber string `json:"phone_number,required"`
	// Customer's email address. Used for notifications, receipts, and account
	// management. Must be a valid email format.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Type of customer account. Determines available features and compliance
	// requirements.
	//
	// Any of "Individual", "Startup", "Small Business", "Medium Business",
	// "Enterprise", "Non-Profit", "Government".
	CustomerType CustomerType `json:"customer_type,omitzero"`
	// Current status of the customer account. Controls access to services and
	// features.
	//
	// Any of "ACTIVE", "BLACKLISTED", "DEACTIVATED", "DELETED".
	Status CustomerStatus `json:"status,omitzero"`
	paramObj
}

func (r V0CustomerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V0CustomerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0CustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0CustomerUpdateParams struct {
	// Customer's email address. Used for notifications, receipts, and account
	// management. Must be a valid email format.
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Customer's first name. Used for personalization and legal documentation.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Customer's last name. Used for personalization and legal documentation.
	LastName param.Opt[string] `json:"last_name,omitzero"`
	// Customer's phone number. Used for SMS notifications and verification. Include
	// country code for international numbers.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	// Type of customer account. Determines available features and compliance
	// requirements.
	//
	// Any of "Individual", "Startup", "Small Business", "Medium Business",
	// "Enterprise", "Non-Profit", "Government".
	CustomerType CustomerType `json:"customer_type,omitzero"`
	// Current status of the customer account. Controls access to services and
	// features.
	//
	// Any of "ACTIVE", "BLACKLISTED", "DEACTIVATED", "DELETED".
	Status CustomerStatus `json:"status,omitzero"`
	paramObj
}

func (r V0CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V0CustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V0CustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V0CustomerListParams struct {
	// Filter customers by email (exact match, case-insensitive)
	Email param.Opt[string] `query:"email,omitzero" json:"-"`
	// Filter customers by name (partial match, case-insensitive)
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Number of records to skip for pagination
	Skip param.Opt[float64] `query:"skip,omitzero" json:"-"`
	// Number of records to return (max 100)
	Take param.Opt[float64] `query:"take,omitzero" json:"-"`
	// Filter customers by status
	//
	// Any of "ACTIVE", "BLACKLISTED", "DEACTIVATED", "DELETED".
	Status CustomerStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0CustomerListParams]'s query parameters as `url.Values`.
func (r V0CustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
