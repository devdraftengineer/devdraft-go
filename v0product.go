// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package devdraft

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/devdraftengineer/devdraft-go/internal/apiform"
	"github.com/devdraftengineer/devdraft-go/internal/apiquery"
	"github.com/devdraftengineer/devdraft-go/internal/requestconfig"
	"github.com/devdraftengineer/devdraft-go/option"
	"github.com/devdraftengineer/devdraft-go/packages/param"
)

// V0ProductService contains methods and other services that help with interacting
// with the devdraft API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV0ProductService] method instead.
type V0ProductService struct {
	Options []option.RequestOption
}

// NewV0ProductService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV0ProductService(opts ...option.RequestOption) (r V0ProductService) {
	r = V0ProductService{}
	r.Options = opts
	return
}

// Creates a new product with optional image uploads.
//
// This endpoint allows you to create products with detailed information and
// multiple images.
//
// ## Use Cases
//
// - Add new products to your catalog
// - Create products with multiple images
// - Set up product pricing and descriptions
//
// ## Supported Image Formats
//
// - JPEG/JPG
// - PNG
// - WebP
// - Maximum 10 images per product
// - Maximum file size: 5MB per image
//
// ## Example Request (multipart/form-data)
//
// ```
// name: "Premium Widget"
// description: "High-quality widget for all your needs"
// price: "99.99"
// images: [file1.jpg, file2.jpg]  // Optional, up to 10 images
// ```
//
// ## Required Fields
//
// - `name`: Product name
// - `price`: Product price (decimal number)
//
// ## Optional Fields
//
// - `description`: Detailed product description
// - `images`: Product images (up to 10 files)
func (r *V0ProductService) New(ctx context.Context, body V0ProductNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieves detailed information about a specific product.
//
// This endpoint allows you to fetch complete product details including all images.
//
// ## Use Cases
//
// - View product details
// - Display product information
// - Check product availability
//
// ## Example Response
//
// ```json
//
//	{
//	  "id": "prod_123456",
//	  "name": "Premium Widget",
//	  "description": "High-quality widget for all your needs",
//	  "price": "99.99",
//	  "images": [
//	    "https://storage.example.com/images/file1.jpg",
//	    "https://storage.example.com/images/file2.jpg"
//	  ],
//	  "createdAt": "2024-03-20T10:00:00Z",
//	  "updatedAt": "2024-03-20T10:00:00Z"
//	}
//
// ```
func (r *V0ProductService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Updates an existing product's information and optionally adds new images.
//
// This endpoint allows you to modify product details and manage product images.
//
// ## Use Cases
//
// - Update product information
// - Change product pricing
// - Modify product images
// - Update product description
//
// ## Example Request (multipart/form-data)
//
// ```
// name: "Updated Premium Widget"
// description: "Updated description"
// price: "129.99"
// images: [file1.jpg, file2.jpg]  // Optional, up to 10 images
// ```
//
// ## Notes
//
// - Only include fields that need to be updated
// - New images will replace existing images
// - Maximum 10 images per product
// - Images are automatically optimized
func (r *V0ProductService) Update(ctx context.Context, id string, body V0ProductUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return
}

// Retrieves a list of products with pagination.
//
// This endpoint allows you to fetch products with optional pagination.
//
// ## Use Cases
//
// - List all products
// - Browse product catalog
// - Implement product search
//
// ## Query Parameters
//
// - `skip`: Number of records to skip (default: 0)
// - `take`: Number of records to take (default: 10)
//
// ## Example Response
//
// ```json
//
//	{
//	  "data": [
//	    {
//	      "id": "prod_123456",
//	      "name": "Premium Widget",
//	      "description": "High-quality widget for all your needs",
//	      "price": "99.99",
//	      "images": [
//	        "https://storage.example.com/images/file1.jpg",
//	        "https://storage.example.com/images/file2.jpg"
//	      ],
//	      "createdAt": "2024-03-20T10:00:00Z"
//	    }
//	  ],
//	  "total": 100,
//	  "skip": 0,
//	  "take": 10
//	}
//
// ```
func (r *V0ProductService) List(ctx context.Context, query V0ProductListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v0/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

// Deletes a product and its associated images.
//
// This endpoint allows you to remove a product and all its associated data.
//
// ## Use Cases
//
// - Remove discontinued products
// - Clean up product catalog
// - Delete test products
//
// ## Notes
//
// - This action cannot be undone
// - All product images will be deleted
// - Associated data will be removed
func (r *V0ProductService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/products/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Adds new images to an existing product.
//
// This endpoint allows you to upload additional images for a product that already
// exists.
//
// ## Use Cases
//
// - Add more product images
// - Update product gallery
// - Enhance product presentation
//
// ## Supported Image Formats
//
// - JPEG/JPG
// - PNG
// - WebP
// - Maximum 10 images per upload
// - Maximum file size: 5MB per image
//
// ## Example Request (multipart/form-data)
//
// ```
// images: [file1.jpg, file2.jpg]  // Up to 10 images
// ```
//
// ## Notes
//
// - Images are appended to existing product images
// - Total images per product cannot exceed 10
// - Images are automatically optimized and resized
func (r *V0ProductService) UploadImages(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v0/products/%s/images", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type V0ProductNewParams struct {
	// Detailed description of the product. Supports markdown formatting for rich text
	// display.
	Description string `json:"description,required"`
	// Product name as it will appear to customers. Should be clear and descriptive.
	Name string `json:"name,required"`
	// Product price in the specified currency. Must be greater than 0.
	Price float64 `json:"price,required"`
	// Product type
	ProductType param.Opt[string] `json:"productType,omitzero"`
	// Quantity available
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Product status
	Status param.Opt[string] `json:"status,omitzero"`
	// Stock count
	StockCount param.Opt[float64] `json:"stockCount,omitzero"`
	// Product type
	Type param.Opt[string] `json:"type,omitzero"`
	// Unit of measurement
	Unit param.Opt[string] `json:"unit,omitzero"`
	// Weight of the product
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// Currency code for the price. Defaults to USD if not specified.
	//
	// Any of "USD", "EUR", "GBP", "CAD", "AUD", "JPY".
	Currency V0ProductNewParamsCurrency `json:"currency,omitzero"`
	// Array of image URLs
	Images []string `json:"images,omitzero"`
	paramObj
}

func (r V0ProductNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Currency code for the price. Defaults to USD if not specified.
type V0ProductNewParamsCurrency string

const (
	V0ProductNewParamsCurrencyUsd V0ProductNewParamsCurrency = "USD"
	V0ProductNewParamsCurrencyEur V0ProductNewParamsCurrency = "EUR"
	V0ProductNewParamsCurrencyGbp V0ProductNewParamsCurrency = "GBP"
	V0ProductNewParamsCurrencyCad V0ProductNewParamsCurrency = "CAD"
	V0ProductNewParamsCurrencyAud V0ProductNewParamsCurrency = "AUD"
	V0ProductNewParamsCurrencyJpy V0ProductNewParamsCurrency = "JPY"
)

type V0ProductUpdateParams struct {
	// Detailed description of the product. Supports markdown formatting for rich text
	// display.
	Description param.Opt[string] `json:"description,omitzero"`
	// Product name as it will appear to customers. Should be clear and descriptive.
	Name param.Opt[string] `json:"name,omitzero"`
	// Product price in the specified currency. Must be greater than 0.
	Price param.Opt[float64] `json:"price,omitzero"`
	// Product type
	ProductType param.Opt[string] `json:"productType,omitzero"`
	// Quantity available
	Quantity param.Opt[float64] `json:"quantity,omitzero"`
	// Product status
	Status param.Opt[string] `json:"status,omitzero"`
	// Stock count
	StockCount param.Opt[float64] `json:"stockCount,omitzero"`
	// Product type
	Type param.Opt[string] `json:"type,omitzero"`
	// Unit of measurement
	Unit param.Opt[string] `json:"unit,omitzero"`
	// Weight of the product
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// Currency code for the price. Defaults to USD if not specified.
	//
	// Any of "USD", "EUR", "GBP", "CAD", "AUD", "JPY".
	Currency V0ProductUpdateParamsCurrency `json:"currency,omitzero"`
	// Array of image URLs
	Images []string `json:"images,omitzero"`
	paramObj
}

func (r V0ProductUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Currency code for the price. Defaults to USD if not specified.
type V0ProductUpdateParamsCurrency string

const (
	V0ProductUpdateParamsCurrencyUsd V0ProductUpdateParamsCurrency = "USD"
	V0ProductUpdateParamsCurrencyEur V0ProductUpdateParamsCurrency = "EUR"
	V0ProductUpdateParamsCurrencyGbp V0ProductUpdateParamsCurrency = "GBP"
	V0ProductUpdateParamsCurrencyCad V0ProductUpdateParamsCurrency = "CAD"
	V0ProductUpdateParamsCurrencyAud V0ProductUpdateParamsCurrency = "AUD"
	V0ProductUpdateParamsCurrencyJpy V0ProductUpdateParamsCurrency = "JPY"
)

type V0ProductListParams struct {
	// Number of records to skip
	Skip param.Opt[float64] `query:"skip,omitzero" json:"-"`
	// Number of records to take
	Take param.Opt[float64] `query:"take,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V0ProductListParams]'s query parameters as `url.Values`.
func (r V0ProductListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
