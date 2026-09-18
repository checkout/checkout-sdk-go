// Package inventory provides types and a client for the Checkout.com Inventory API: stock
// levels, atomic multi-variant reservations (hold/commit/release), stock adjustments, and
// per-variant product knowledge (merchandising metadata for AI agents).
//
// Every operation in this package requires OAuth client-credentials authorization scoped to
// agentic:inventory. None of the ten Inventory endpoints accept ApiSecretKey/ApiPublicKey
// authorization.
package inventory

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

const (
	inventoryPath = "inventory"
	adjustments   = "adjustments"
	reservations  = "reservations"
	commitPath    = "commit"
	releasePath   = "release"
	productPath   = "product"
)

// InventoryLevelState is the stock state of a variant, derived from on_hand, reserved and
// safety_stock.
type InventoryLevelState string

const (
	// InStock means available stock is above zero and above any low-stock threshold applied
	// by the caller.
	InStock InventoryLevelState = "in_stock"
	// Limited means available stock is low.
	Limited InventoryLevelState = "limited"
	// OutOfStock means available stock is zero.
	OutOfStock InventoryLevelState = "out_of_stock"
)

// InventorySource indicates whether a variant's stock is managed directly through this API or
// synced from an external system.
type InventorySource string

const (
	// Managed means stock levels are set and adjusted directly through this API.
	Managed InventorySource = "managed"
	// Sync means stock levels are kept up to date from an external inventory system.
	Sync InventorySource = "sync"
)

// InventoryReservationState is the lifecycle state of a reservation.
type InventoryReservationState string

const (
	// Held means the reservation is active and its items count against reserved stock.
	Held InventoryReservationState = "held"
	// Committed means the reservation has been converted into a permanent stock deduction.
	Committed InventoryReservationState = "committed"
	// Released means the reservation was cancelled and its held stock returned.
	Released InventoryReservationState = "released"
	// Expired means the reservation was held past its expires_at without being committed or
	// released. A reservation in this state is reported as expired even if its stored state
	// still reads held.
	Expired InventoryReservationState = "expired"
)

// InventoryProductCondition is the condition of the merchandise described by a product
// knowledge record.
type InventoryProductCondition string

const (
	// NewCondition means the item is new.
	NewCondition InventoryProductCondition = "new"
	// UsedCondition means the item is used.
	UsedCondition InventoryProductCondition = "used"
	// RefurbishedCondition means the item has been refurbished.
	RefurbishedCondition InventoryProductCondition = "refurbished"
)

// InventoryHalLink is a HAL-style hypermedia link describing an action available on the
// resource it is attached to.
type InventoryHalLink struct {
	// HRef is the absolute URI of the linked resource.
	// [Optional]
	HRef *string `json:"href,omitempty"`
	// Actions lists the HTTP methods supported at HRef.
	// [Optional]
	Actions []string `json:"actions,omitempty"`
	// Types lists the media types supported at HRef.
	// [Optional]
	Types []string `json:"types,omitempty"`
}

// InventoryMoney is a monetary amount used by product knowledge pricing fields.
type InventoryMoney struct {
	// Amount is the amount in the currency's minor unit.
	// [Required]
	Amount int64 `json:"amount"`
	// Currency is the three-letter ISO 4217 currency code.
	// [Required]
	// min 3 characters, max 3 characters
	Currency string `json:"currency"`
}

// InventoryAdjustmentRequest is the request body of POST /inventory/adjustments.
type InventoryAdjustmentRequest struct {
	// VariantId is the identifier of the variant to adjust. The variant must already exist.
	// [Required]
	// max 128 characters
	VariantId string `json:"variant_id"`
	// Delta is the signed change to apply to on_hand. A negative delta that would drive
	// on_hand below zero is rejected with a 409. The value must be non-zero (enforced by the
	// API, not a formal schema constraint).
	// [Required]
	Delta int64 `json:"delta"`
	// Reason is a free-text reason recorded in the adjustment ledger. It must not contain
	// personal data.
	// [Required]
	// min 1 character, max 256 characters
	Reason string `json:"reason"`
}

// InventoryLevelsLinks is the _links object of an InventoryLevels response.
type InventoryLevelsLinks struct {
	// Self links back to this variant's stock levels.
	// [Optional]
	Self *InventoryHalLink `json:"self,omitempty"`
	// Set links to the operation that sets this variant's stock levels.
	// [Optional]
	Set *InventoryHalLink `json:"set,omitempty"`
}

// InventoryLevels is the response of GET/PUT /inventory/{variant_id} and POST
// /inventory/adjustments. adjustInventory and createInventoryReservation-style idempotent
// replay: this schema is shared by both the 201 (created/adjusted) and 200 (idempotent replay)
// responses of adjustInventory.
type InventoryLevels struct {
	HttpMetadata common.HttpMetadata
	// VariantId is the identifier of the variant these levels belong to.
	// [Optional]
	VariantId string `json:"variant_id,omitempty"`
	// OnHand is the physical stock currently held.
	// [Optional]
	OnHand int64 `json:"on_hand,omitempty"`
	// Reserved is the sum of quantities held by active reservations.
	// [Optional]
	Reserved int64 `json:"reserved,omitempty"`
	// SafetyStock is the buffer withheld from sale.
	// [Optional]
	SafetyStock int64 `json:"safety_stock,omitempty"`
	// Available is max(0, on_hand - reserved - safety_stock).
	// [Optional]
	Available int64 `json:"available,omitempty"`
	// State is the derived stock state of the variant.
	// [Optional]
	// Enum: "in_stock" "limited" "out_of_stock"
	State InventoryLevelState `json:"state,omitempty"`
	// Source indicates whether stock is managed directly or synced from an external system.
	// [Optional]
	// Enum: "managed" "sync"
	Source InventorySource `json:"source,omitempty"`
	// CreatedOn is when the stock record was created.
	// [Optional]
	// Format: date-time (RFC 3339)
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// ModifiedOn is when the stock record was last modified.
	// [Optional]
	// Format: date-time (RFC 3339)
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
	// Product is the variant's product knowledge, embedded only when the request was made
	// with ?expand=product and product knowledge exists for the variant.
	// [Optional]
	Product *InventoryProductKnowledge `json:"product,omitempty"`
	// Links exposes the actions available on this resource.
	// [Optional]
	Links *InventoryLevelsLinks `json:"_links,omitempty"`
}

// InventoryLevelsQuery is the query string of GET /inventory/{variant_id}.
type InventoryLevelsQuery struct {
	// Expand, when set to "product", embeds the variant's product knowledge in the response.
	// [Optional]
	Expand string `url:"expand,omitempty"`
}

// InventoryReservationItem is a single variant/quantity pair inside a reservation.
type InventoryReservationItem struct {
	// VariantId is the identifier of the variant to reserve. The variant must already exist.
	// [Required]
	// max 128 characters
	VariantId string `json:"variant_id"`
	// Quantity is the number of units to reserve.
	// [Required]
	// min 1
	Quantity int64 `json:"quantity"`
}

// InventoryReservationRequest is the request body of POST /inventory/reservations.
type InventoryReservationRequest struct {
	// OwnerType identifies the kind of owner the reservation is held for.
	// [Required]
	// max 64 characters
	OwnerType string `json:"owner_type"`
	// OwnerReference identifies the specific owner the reservation is held for.
	// [Required]
	// max 256 characters
	OwnerReference string `json:"owner_reference"`
	// Items is the list of variant/quantity pairs to reserve. Variant IDs must be unique
	// within the request.
	// [Required]
	// min 1 item, max 45 items
	Items []InventoryReservationItem `json:"items"`
	// TtlSeconds is how long the reservation is held before it expires.
	// [Optional]
	// min 60, max 3600, default 900
	TtlSeconds int64 `json:"ttl_seconds,omitempty"`
}

// InventoryReservationLinks is the _links object of an InventoryReservation response. Held
// reservations expose self, commit and release; terminal-state reservations expose self only.
type InventoryReservationLinks struct {
	// Self links back to this reservation.
	// [Optional]
	Self *InventoryHalLink `json:"self,omitempty"`
	// Commit links to the operation that commits this reservation.
	// [Optional]
	Commit *InventoryHalLink `json:"commit,omitempty"`
	// Release links to the operation that releases this reservation.
	// [Optional]
	Release *InventoryHalLink `json:"release,omitempty"`
}

// InventoryReservation is the response of createInventoryReservation, getInventoryReservation,
// commitInventoryReservation and releaseInventoryReservation.
type InventoryReservation struct {
	HttpMetadata common.HttpMetadata
	// Id is the reservation identifier, formatted rsv_{base32-encoded GUID}.
	// [Optional]
	Id string `json:"id,omitempty"`
	// State is the lifecycle state of the reservation. A reservation held past its
	// expires_at reports as expired even before it is otherwise updated.
	// [Optional]
	// Enum: "held" "committed" "released" "expired"
	State InventoryReservationState `json:"state,omitempty"`
	// OwnerType echoes the owner_type supplied on creation.
	// [Optional]
	OwnerType string `json:"owner_type,omitempty"`
	// OwnerReference echoes the owner_reference supplied on creation.
	// [Optional]
	OwnerReference string `json:"owner_reference,omitempty"`
	// Items is the list of variant/quantity pairs held by this reservation.
	// [Optional]
	Items []InventoryReservationItem `json:"items,omitempty"`
	// ExpiresAt is when the reservation expires if not committed or released first.
	// [Optional]
	// Format: date-time (RFC 3339)
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// CreatedOn is when the reservation was created.
	// [Optional]
	// Format: date-time (RFC 3339)
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// Links exposes the actions available on this resource.
	// [Optional]
	Links *InventoryReservationLinks `json:"_links,omitempty"`
}

// InventorySetLevelsRequest is the request body of PUT /inventory/{variant_id}.
type InventorySetLevelsRequest struct {
	// OnHand is the physical stock to set.
	// [Required]
	// min 0
	OnHand int64 `json:"on_hand"`
	// SafetyStock is the buffer to withhold from sale. Defaults to 0 when the variant is
	// created by this call; left unchanged on an existing variant if omitted. A pointer
	// distinguishes "omitted" from an explicit 0.
	// [Optional]
	// min 0
	SafetyStock *int64 `json:"safety_stock,omitempty"`
	// Reason is a free-text reason recorded in the ledger. It must not contain personal data.
	// [Optional]
	// max 256 characters
	Reason string `json:"reason,omitempty"`
}

// InventoryProductKnowledge is the response of getInventoryProduct and setInventoryProduct, and
// is embeddable in InventoryLevels.product when ?expand=product is requested.
type InventoryProductKnowledge struct {
	HttpMetadata common.HttpMetadata
	// VariantId is the identifier of the variant this product knowledge describes.
	// [Required]
	VariantId string `json:"variant_id,omitempty"`
	// Title is the merchandising title of the product.
	// [Required]
	Title string `json:"title,omitempty"`
	// Description is the merchandising description of the product.
	// [Required]
	Description string `json:"description,omitempty"`
	// ProductUrl is the canonical URL of the product page.
	// [Required]
	ProductUrl string `json:"product_url,omitempty"`
	// ImageUrl is the URL of the primary product image.
	// [Required]
	ImageUrl string `json:"image_url,omitempty"`
	// AdditionalImageUrls lists URLs of additional product images.
	// [Optional]
	AdditionalImageUrls []string `json:"additional_image_urls,omitempty"`
	// VideoUrl is the URL of a product video.
	// [Optional]
	VideoUrl string `json:"video_url,omitempty"`
	// Model3dUrl is the URL of a 3D model of the product.
	// [Optional]
	Model3dUrl string `json:"model_3d_url,omitempty"`
	// Sku is the merchant's stock-keeping unit for the product.
	// [Optional]
	Sku string `json:"sku,omitempty"`
	// Gtin is the product's Global Trade Item Number.
	// [Optional]
	Gtin string `json:"gtin,omitempty"`
	// Mpn is the product's Manufacturer Part Number.
	// [Optional]
	Mpn string `json:"mpn,omitempty"`
	// Brand is the product's brand.
	// [Optional]
	Brand string `json:"brand,omitempty"`
	// Category is the product's merchandising category.
	// [Optional]
	Category string `json:"category,omitempty"`
	// Price is the product's list price.
	// [Optional]
	Price *InventoryMoney `json:"price,omitempty"`
	// SalePrice is the product's discounted price, when on sale. Must share Price's currency
	// and must not exceed Price.
	// [Optional]
	SalePrice *InventoryMoney `json:"sale_price,omitempty"`
	// SalePriceStartsAt is when SalePrice becomes active. Pairs with SalePrice.
	// [Optional]
	// Format: date-time (RFC 3339)
	SalePriceStartsAt *time.Time `json:"sale_price_starts_at,omitempty"`
	// SalePriceEndsAt is when SalePrice stops being active. Pairs with SalePrice.
	// [Optional]
	// Format: date-time (RFC 3339)
	SalePriceEndsAt *time.Time `json:"sale_price_ends_at,omitempty"`
	// GroupId, when set, groups this variant with other variants of the same product; Color
	// and Size are both required when GroupId is set (enforced by the API, not a formal
	// schema constraint).
	// [Optional]
	GroupId string `json:"group_id,omitempty"`
	// GroupTitle is the merchandising title shared by all variants in GroupId.
	// [Optional]
	GroupTitle string `json:"group_title,omitempty"`
	// Color is the variant's color. Required when GroupId is set (enforced by the API, not a
	// formal schema constraint).
	// [Optional]
	Color string `json:"color,omitempty"`
	// Size is the variant's size. Required when GroupId is set (enforced by the API, not a
	// formal schema constraint).
	// [Optional]
	Size string `json:"size,omitempty"`
	// SizeSystem is the sizing system Size is expressed in.
	// [Optional]
	SizeSystem string `json:"size_system,omitempty"`
	// Gender is the target gender for the product.
	// [Optional]
	Gender string `json:"gender,omitempty"`
	// Condition is the condition of the merchandise. Defaults to "new".
	// [Required]
	// Enum: "new" "used" "refurbished"
	Condition InventoryProductCondition `json:"condition,omitempty"`
	// Material is the product's material.
	// [Optional]
	Material string `json:"material,omitempty"`
	// AgeGroup is the target age group for the product.
	// [Optional]
	AgeGroup string `json:"age_group,omitempty"`
	// Length is the product's length.
	// [Optional]
	Length float64 `json:"length,omitempty"`
	// Width is the product's width.
	// [Optional]
	Width float64 `json:"width,omitempty"`
	// Height is the product's height.
	// [Optional]
	Height float64 `json:"height,omitempty"`
	// DimensionUnit is the unit Length, Width and Height are expressed in.
	// [Optional]
	DimensionUnit string `json:"dimension_unit,omitempty"`
	// Weight is the product's weight.
	// [Optional]
	Weight float64 `json:"weight,omitempty"`
	// WeightUnit is the unit Weight is expressed in.
	// [Optional]
	WeightUnit string `json:"weight_unit,omitempty"`
	// ExpirationDate is when the product expires, if applicable.
	// [Optional]
	// Format: date-time (RFC 3339)
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// HarmonizedSystemCode is the product's Harmonized System customs code.
	// [Optional]
	HarmonizedSystemCode string `json:"harmonized_system_code,omitempty"`
	// CountryOfOrigin is the product's two-letter ISO 3166-1 alpha-2 country of origin.
	// [Optional]
	CountryOfOrigin string `json:"country_of_origin,omitempty"`
	// SellerName is the name of the seller of record.
	// [Optional]
	SellerName string `json:"seller_name,omitempty"`
	// SellerUrl is the URL of the seller of record.
	// [Optional]
	SellerUrl string `json:"seller_url,omitempty"`
	// SellerPrivacyPolicy is the URL of the seller's privacy policy.
	// [Optional]
	SellerPrivacyPolicy string `json:"seller_privacy_policy,omitempty"`
	// SellerTos is the URL of the seller's terms of service.
	// [Optional]
	SellerTos string `json:"seller_tos,omitempty"`
	// CreatedOn is when this product knowledge record was created.
	// [Required]
	// Format: date-time (RFC 3339)
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// ModifiedOn is when this product knowledge record was last modified.
	// [Required]
	// Format: date-time (RFC 3339)
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
	// Links exposes the actions available on this resource.
	// [Required]
	Links *InventoryProductKnowledgeLinks `json:"_links,omitempty"`
}

// InventoryProductKnowledgeLinks is the _links object of an InventoryProductKnowledge response.
type InventoryProductKnowledgeLinks struct {
	// Self links back to this variant's product knowledge.
	// [Optional]
	Self *InventoryHalLink `json:"self,omitempty"`
	// Set links to the operation that sets this variant's product knowledge.
	// [Optional]
	Set *InventoryHalLink `json:"set,omitempty"`
	// Delete links to the operation that deletes this variant's product knowledge.
	// [Optional]
	Delete *InventoryHalLink `json:"delete,omitempty"`
}

// InventorySetProductRequest is the request body of PUT /inventory/{variant_id}/product. It
// carries the same optional merchandising fields as InventoryProductKnowledge (minus the
// server-assigned variant_id/created_on/modified_on/_links) plus request-only length limits on
// a few fields.
type InventorySetProductRequest struct {
	// Title is the merchandising title of the product.
	// [Required]
	// max 512 characters
	Title string `json:"title"`
	// Description is the merchandising description of the product.
	// [Required]
	// max 4000 characters
	Description string `json:"description"`
	// ProductUrl is the canonical URL of the product page.
	// [Required]
	// max 2048 characters
	ProductUrl string `json:"product_url"`
	// ImageUrl is the URL of the primary product image.
	// [Required]
	// max 2048 characters
	ImageUrl string `json:"image_url"`
	// AdditionalImageUrls lists URLs of additional product images.
	// [Optional]
	AdditionalImageUrls []string `json:"additional_image_urls,omitempty"`
	// VideoUrl is the URL of a product video.
	// [Optional]
	VideoUrl string `json:"video_url,omitempty"`
	// Model3dUrl is the URL of a 3D model of the product.
	// [Optional]
	Model3dUrl string `json:"model_3d_url,omitempty"`
	// Sku is the merchant's stock-keeping unit for the product.
	// [Optional]
	// max 128 characters
	Sku string `json:"sku,omitempty"`
	// Gtin is the product's Global Trade Item Number.
	// [Optional]
	Gtin string `json:"gtin,omitempty"`
	// Mpn is the product's Manufacturer Part Number.
	// [Optional]
	Mpn string `json:"mpn,omitempty"`
	// Brand is the product's brand.
	// [Optional]
	Brand string `json:"brand,omitempty"`
	// Category is the product's merchandising category.
	// [Optional]
	Category string `json:"category,omitempty"`
	// Price is the product's list price.
	// [Optional]
	Price *InventoryMoney `json:"price,omitempty"`
	// SalePrice is the product's discounted price, when on sale. Must share Price's currency
	// and must not exceed Price.
	// [Optional]
	SalePrice *InventoryMoney `json:"sale_price,omitempty"`
	// SalePriceStartsAt is when SalePrice becomes active. Pairs with SalePrice.
	// [Optional]
	// Format: date-time (RFC 3339)
	SalePriceStartsAt *time.Time `json:"sale_price_starts_at,omitempty"`
	// SalePriceEndsAt is when SalePrice stops being active. Pairs with SalePrice.
	// [Optional]
	// Format: date-time (RFC 3339)
	SalePriceEndsAt *time.Time `json:"sale_price_ends_at,omitempty"`
	// GroupId, when set, groups this variant with other variants of the same product; Color
	// and Size are both required when GroupId is set (enforced by the API, not a formal
	// schema constraint).
	// [Optional]
	GroupId string `json:"group_id,omitempty"`
	// GroupTitle is the merchandising title shared by all variants in GroupId.
	// [Optional]
	GroupTitle string `json:"group_title,omitempty"`
	// Color is the variant's color. Required when GroupId is set (enforced by the API, not a
	// formal schema constraint).
	// [Optional]
	Color string `json:"color,omitempty"`
	// Size is the variant's size. Required when GroupId is set (enforced by the API, not a
	// formal schema constraint).
	// [Optional]
	Size string `json:"size,omitempty"`
	// SizeSystem is the sizing system Size is expressed in.
	// [Optional]
	SizeSystem string `json:"size_system,omitempty"`
	// Gender is the target gender for the product.
	// [Optional]
	Gender string `json:"gender,omitempty"`
	// Condition is the condition of the merchandise. Must be an exact lowercase match.
	// Defaults to "new".
	// [Optional]
	// Enum: "new" "used" "refurbished"
	Condition InventoryProductCondition `json:"condition,omitempty"`
	// Material is the product's material.
	// [Optional]
	Material string `json:"material,omitempty"`
	// AgeGroup is the target age group for the product.
	// [Optional]
	AgeGroup string `json:"age_group,omitempty"`
	// Length is the product's length.
	// [Optional]
	Length float64 `json:"length,omitempty"`
	// Width is the product's width.
	// [Optional]
	Width float64 `json:"width,omitempty"`
	// Height is the product's height.
	// [Optional]
	Height float64 `json:"height,omitempty"`
	// DimensionUnit is the unit Length, Width and Height are expressed in.
	// [Optional]
	DimensionUnit string `json:"dimension_unit,omitempty"`
	// Weight is the product's weight.
	// [Optional]
	Weight float64 `json:"weight,omitempty"`
	// WeightUnit is the unit Weight is expressed in.
	// [Optional]
	WeightUnit string `json:"weight_unit,omitempty"`
	// ExpirationDate is when the product expires, if applicable.
	// [Optional]
	// Format: date-time (RFC 3339)
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// HarmonizedSystemCode is the product's Harmonized System customs code.
	// [Optional]
	HarmonizedSystemCode string `json:"harmonized_system_code,omitempty"`
	// CountryOfOrigin is the product's two-letter ISO 3166-1 alpha-2 country of origin.
	// [Optional]
	CountryOfOrigin string `json:"country_of_origin,omitempty"`
	// SellerName is the name of the seller of record.
	// [Optional]
	SellerName string `json:"seller_name,omitempty"`
	// SellerUrl is the URL of the seller of record.
	// [Optional]
	SellerUrl string `json:"seller_url,omitempty"`
	// SellerPrivacyPolicy is the URL of the seller's privacy policy.
	// [Optional]
	SellerPrivacyPolicy string `json:"seller_privacy_policy,omitempty"`
	// SellerTos is the URL of the seller's terms of service.
	// [Optional]
	SellerTos string `json:"seller_tos,omitempty"`
}
