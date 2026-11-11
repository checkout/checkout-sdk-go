package accounts

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
)

type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
	filesClient   client.HttpClient
}

func NewClient(
	configuration *configuration.Configuration,
	apiClient client.HttpClient,
	filesClient client.HttpClient,
) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
		filesClient:   filesClient,
	}
}

// DefaultSchemaVersion is 3.0, the latest Accounts API schema version. Pass it as the schemaVersion
// argument (or "") to negotiate 3.0 via the Accept header.
const DefaultSchemaVersion = "3.0"

func buildSchemaVersionHeaders(schemaVersion string) *Headers {
	if schemaVersion == "" {
		schemaVersion = DefaultSchemaVersion
	}
	return &Headers{Accept: "application/json;schema_version=" + schemaVersion}
}

func (c *Client) CreateEntity(request OnboardEntityRequest, schemaVersion string) (*OnboardEntityResponse, error) {
	return c.CreateEntityWithContext(context.Background(), request, schemaVersion)
}

func (c *Client) CreateEntityWithContext(
	ctx context.Context,
	request OnboardEntityRequest,
	schemaVersion string,
) (*OnboardEntityResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	payload := struct {
		OnboardEntityRequest
		Headers *Headers `json:"-"`
	}{request, buildSchemaVersionHeaders(schemaVersion)}

	var response OnboardEntityResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath),
		auth,
		payload,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetSubEntityMembers sends GET /accounts/entities/{entityId}/members: get sub-entity Members.
//
// entityId is the {entityId} path value: the ID of the sub-entity.
//
// It returns the OnboardSubEntityDetailsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) GetSubEntityMembers(entityId string) (*OnboardSubEntityDetailsResponse, error) {
	return c.GetSubEntityMembersWithContext(context.Background(), entityId)
}

// GetSubEntityMembersWithContext is GetSubEntityMembers with a caller-supplied context.
func (c *Client) GetSubEntityMembersWithContext(
	ctx context.Context,
	entityId string,
) (*OnboardSubEntityDetailsResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response OnboardSubEntityDetailsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId, membersPath), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ReinviteSubEntityMember sends PUT /accounts/entities/{entityId}/members/{userId}: reinvite a
// sub-entity member.
//
// entityId is the {entityId} path value: the ID of the sub-entity. userId is the {userId} path
// value: the ID of the invited user. request is sent as the JSON body.
//
// It returns the OnboardSubEntityResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// userId returns a CheckoutArgumentError and sends no request.
func (c *Client) ReinviteSubEntityMember(
	entityId string,
	userId string,
	request OnboardSubEntityRequest,
) (*OnboardSubEntityResponse, error) {
	return c.ReinviteSubEntityMemberWithContext(context.Background(), entityId, userId, request)
}

// ReinviteSubEntityMemberWithContext is ReinviteSubEntityMember with a caller-supplied context.
func (c *Client) ReinviteSubEntityMemberWithContext(
	ctx context.Context,
	entityId string,
	userId string,
	request OnboardSubEntityRequest,
) (*OnboardSubEntityResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	userId, err = common.EncodePathSegment(userId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response OnboardSubEntityResponse
	err = c.apiClient.PutWithContext(ctx, common.BuildPath(
		accountsPath,
		entitiesPath,
		entityId,
		membersPath,
		userId,
	), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetEntity sends GET /accounts/entities/{id}: get entity details.
//
// entityId is the {id} path value: the ID of the entity. schemaVersion is sent in the Accept header
// as the schema_version; an empty value uses DefaultSchemaVersion.
//
// It returns the OnboardEntityDetails, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) GetEntity(entityId string, schemaVersion string) (*OnboardEntityDetails, error) {
	return c.GetEntityWithContext(context.Background(), entityId, schemaVersion)
}

// GetEntityWithContext is GetEntity with a caller-supplied context.
func (c *Client) GetEntityWithContext(ctx context.Context, entityId string, schemaVersion string) (*OnboardEntityDetails, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	payload := struct {
		Headers *Headers `json:"-"`
	}{buildSchemaVersionHeaders(schemaVersion)}

	var response OnboardEntityDetails
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId), auth, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateEntity sends PUT /accounts/entities/{id}: update entity details.
//
// entityId is the {id} path value: the ID of the entity. request is sent as the JSON body.
// schemaVersion is sent in the Accept header as the schema_version; an empty value uses
// DefaultSchemaVersion.
//
// It returns the OnboardEntityResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) UpdateEntity(entityId string, request OnboardEntityRequest, schemaVersion string) (*OnboardEntityResponse, error) {
	return c.UpdateEntityWithContext(context.Background(), entityId, request, schemaVersion)
}

// UpdateEntityWithContext is UpdateEntity with a caller-supplied context.
func (c *Client) UpdateEntityWithContext(
	ctx context.Context,
	entityId string,
	request OnboardEntityRequest,
	schemaVersion string,
) (*OnboardEntityResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	payload := struct {
		OnboardEntityRequest
		Headers *Headers `json:"-"`
	}{request, buildSchemaVersionHeaders(schemaVersion)}

	var response OnboardEntityResponse
	err = c.apiClient.PutWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId), auth, payload, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreatePaymentInstruments sends POST /accounts/entities/{entityId}/instruments.
//
// entityId is the {entityId} path value. request is sent as the JSON body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
//
// Deprecated: Use CreatePaymentInstrument IdResponse for CreatePaymentInstrument instead.
func (c *Client) CreatePaymentInstruments(
	entityId string,
	request PaymentInstrument,
) (*common.MetadataResponse, error) {
	return c.CreatePaymentInstrumentsWithContext(context.Background(), entityId, request)
}

// CreatePaymentInstrumentsWithContext is CreatePaymentInstruments with a caller-supplied context.
//
// Deprecated: Use CreatePaymentInstrumentWithContext IdResponse for CreatePaymentInstrumentsWithContext instead.
func (c *Client) CreatePaymentInstrumentsWithContext(
	ctx context.Context,
	entityId string,
	request PaymentInstrument,
) (*common.MetadataResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, instrumentsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreatePaymentInstrument sends POST /accounts/entities/{id}/payment-instruments: add a payment
// instrument.
//
// entityId is the {id} path value: the sub-entity's ID. request is sent as the JSON body.
//
// It returns the common.IdResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) CreatePaymentInstrument(
	entityId string,
	request PaymentInstrumentRequest,
) (*common.IdResponse, error) {
	return c.CreatePaymentInstrumentWithContext(context.Background(), entityId, request)
}

// CreatePaymentInstrumentWithContext is CreatePaymentInstrument with a caller-supplied context.
func (c *Client) CreatePaymentInstrumentWithContext(
	ctx context.Context,
	entityId string,
	request PaymentInstrumentRequest,
) (*common.IdResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, paymentInstrumentsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RetrievePaymentInstrumentDetails sends GET
// /accounts/entities/{entityId}/payment-instruments/{id}: get payment instrument details.
//
// entityId is the {entityId} path value: the sub-entity's ID. paymentInstrumentId is the {id} path
// value: the payment instrument's ID.
//
// It returns the PaymentInstrumentDetailsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// paymentInstrumentId returns a CheckoutArgumentError and sends no request.
func (c *Client) RetrievePaymentInstrumentDetails(
	entityId string,
	paymentInstrumentId string,
) (*PaymentInstrumentDetailsResponse, error) {
	return c.RetrievePaymentInstrumentDetailsWithContext(context.Background(), entityId, paymentInstrumentId)
}

// RetrievePaymentInstrumentDetailsWithContext is RetrievePaymentInstrumentDetails with a caller-supplied context.
func (c *Client) RetrievePaymentInstrumentDetailsWithContext(
	ctx context.Context,
	entityId string,
	paymentInstrumentId string,
) (*PaymentInstrumentDetailsResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	paymentInstrumentId, err = common.EncodePathSegment(paymentInstrumentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentInstrumentDetailsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, paymentInstrumentsPath, paymentInstrumentId),
		auth, nil,

		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdatePaymentInstrumentDetails sends PATCH
// /accounts/entities/{entityId}/payment-instruments/{id}: update payment instrument details.
//
// entityId is the {entityId} path value: the sub-entity's ID. instrumentId is the {id} path value:
// the payment instrument's ID. request is sent as the JSON body.
//
// It returns the common.IdResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// instrumentId returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdatePaymentInstrumentDetails(
	entityId, instrumentId string,
	request UpdatePaymentInstrumentRequest,
) (*common.IdResponse, error) {
	return c.UpdatePaymentInstrumentDetailsWithContext(context.Background(), entityId, instrumentId, request)
}

// UpdatePaymentInstrumentDetailsWithContext is UpdatePaymentInstrumentDetails with a caller-supplied context.
func (c *Client) UpdatePaymentInstrumentDetailsWithContext(
	ctx context.Context,
	entityId, instrumentId string,
	request UpdatePaymentInstrumentRequest,
) (*common.IdResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	instrumentId, err = common.EncodePathSegment(instrumentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PatchWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, paymentInstrumentsPath, instrumentId),
		auth,
		request,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// QueryPaymentInstruments sends GET /accounts/entities/{id}/payment-instruments: query payment
// instruments.
//
// entityId is the {id} path value: the sub-entity's ID. query is sent as the query string.
//
// It returns the PaymentInstrumentQueryResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) QueryPaymentInstruments(
	entityId string,
	query PaymentInstrumentsQuery,
) (*PaymentInstrumentQueryResponse, error) {
	return c.QueryPaymentInstrumentsWithContext(context.Background(), entityId, query)
}

// QueryPaymentInstrumentsWithContext is QueryPaymentInstruments with a caller-supplied context.
func (c *Client) QueryPaymentInstrumentsWithContext(
	ctx context.Context,
	entityId string,
	query PaymentInstrumentsQuery,
) (*PaymentInstrumentQueryResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(
		common.BuildPath(accountsPath, entitiesPath, entityId, paymentInstrumentsPath),
		query,
	)
	if err != nil {
		return nil, err
	}

	var response PaymentInstrumentQueryResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RetrievePayoutSchedule sends GET /accounts/entities/{id}/payout-schedules: retrieve a
// sub-entity's payout schedule.
//
// entityId is the {id} path value: the ID of the sub-entity.
//
// It returns the PayoutSchedule, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) RetrievePayoutSchedule(entityId string) (*PayoutSchedule, error) {
	return c.RetrievePayoutScheduleWithContext(context.Background(), entityId)
}

// RetrievePayoutScheduleWithContext is RetrievePayoutSchedule with a caller-supplied context.
func (c *Client) RetrievePayoutScheduleWithContext(
	ctx context.Context,
	entityId string,
) (*PayoutSchedule, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PayoutSchedule
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, payoutSchedulesPath),
		auth, nil,

		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdatePayoutSchedule sends PUT /accounts/entities/{id}/payout-schedules: update a sub-entity's
// payout schedule.
//
// entityId is the {id} path value: the ID of the sub-entity. currency and updateSchedule are sent
// as the JSON body, an object whose only key is currency and whose value is updateSchedule.
//
// It returns the common.IdResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) UpdatePayoutSchedule(
	entityId string,
	currency common.Currency,
	updateSchedule CurrencySchedule,
) (*common.IdResponse, error) {
	return c.UpdatePayoutScheduleWithContext(context.Background(), entityId, currency, updateSchedule)
}

// UpdatePayoutScheduleWithContext is UpdatePayoutSchedule with a caller-supplied context.
func (c *Client) UpdatePayoutScheduleWithContext(
	ctx context.Context,
	entityId string,
	currency common.Currency,
	updateSchedule CurrencySchedule,
) (*common.IdResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	request := map[common.Currency]CurrencySchedule{
		currency: updateSchedule,
	}

	var response common.IdResponse
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, payoutSchedulesPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateReserveRule sends POST /accounts/entities/{id}/reserve-rules: add a reserve rule.
//
// entityId is the {id} path value: the sub-entity's ID. request is sent as the JSON body.
//
// It returns the common.IdResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) CreateReserveRule(entityId string, request ReserveRuleRequest) (*common.IdResponse, error) {
	return c.CreateReserveRuleWithContext(context.Background(), entityId, request)
}

// CreateReserveRuleWithContext is CreateReserveRule with a caller-supplied context.
func (c *Client) CreateReserveRuleWithContext(
	ctx context.Context,
	entityId string,
	request ReserveRuleRequest,
) (*common.IdResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, reserveRulesPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetReserveRules sends GET /accounts/entities/{id}/reserve-rules: query reserve rules.
//
// entityId is the {id} path value: the sub-entity's ID.
//
// It returns the ReserveRulesResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) GetReserveRules(entityId string) (*ReserveRulesResponse, error) {
	return c.GetReserveRulesWithContext(context.Background(), entityId)
}

// GetReserveRulesWithContext is GetReserveRules with a caller-supplied context.
func (c *Client) GetReserveRulesWithContext(ctx context.Context, entityId string) (*ReserveRulesResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ReserveRulesResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, reserveRulesPath),
		auth, nil,

		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetReserveRuleDetails sends GET /accounts/entities/{entityId}/reserve-rules/{id}: get reserve
// rule details.
//
// entityId is the {entityId} path value: the sub-entity's ID. reserveRuleId is the {id} path value:
// the reserve rule ID.
//
// It returns the ReserveRuleResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// reserveRuleId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetReserveRuleDetails(entityId, reserveRuleId string) (*ReserveRuleResponse, error) {
	return c.GetReserveRuleDetailsWithContext(context.Background(), entityId, reserveRuleId)
}

// GetReserveRuleDetailsWithContext is GetReserveRuleDetails with a caller-supplied context.
func (c *Client) GetReserveRuleDetailsWithContext(
	ctx context.Context,
	entityId, reserveRuleId string,
) (*ReserveRuleResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	reserveRuleId, err = common.EncodePathSegment(reserveRuleId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ReserveRuleResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, reserveRulesPath, reserveRuleId),
		auth, nil,

		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateReserveRule sends PUT /accounts/entities/{entityId}/reserve-rules/{id}: update reserve
// rule.
//
// entityId is the {entityId} path value: the sub-entity's ID. reserveRuleId is the {id} path value:
// the reserve rule ID. etag is sent as the If-Match header. request is sent as the JSON body.
//
// It returns the common.IdResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// reserveRuleId returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateReserveRule(
	entityId, reserveRuleId, etag string,
	request ReserveRuleRequest,
) (*common.IdResponse, error) {
	return c.UpdateReserveRuleWithContext(context.Background(), entityId, reserveRuleId, etag, request)
}

// UpdateReserveRuleWithContext is UpdateReserveRule with a caller-supplied context.
func (c *Client) UpdateReserveRuleWithContext(
	ctx context.Context,
	entityId, reserveRuleId, etag string,
	request ReserveRuleRequest,
) (*common.IdResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	reserveRuleId, err = common.EncodePathSegment(reserveRuleId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	payload := struct {
		ReserveRuleRequest
		Headers *Headers `json:"-"`
	}{request, &Headers{IfMatch: etag}}

	var response common.IdResponse
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(accountsPath, entitiesPath, entityId, reserveRulesPath, reserveRuleId),
		auth,
		payload,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SubmitFile uploads a file to the Files API (POST /files on the Files host), as a multipart
// request. The returned ID is what document Front and Back fields take.
func (c *Client) SubmitFile(file File) (*common.IdResponse, error) {
	return c.SubmitFileWithContext(context.Background(), file)
}

// SubmitFileWithContext is SubmitFile with a context for cancellation and deadlines.
func (c *Client) SubmitFileWithContext(
	ctx context.Context,
	file File,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	req, err := common.BuildFileUploadRequest(&file)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.filesClient.UploadWithContext(ctx, common.BuildPath(filesPath), auth, req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UploadFile requests an upload link for a sub-entity file (POST /entities/{entityId}/files on the
// Files host). It sends only request.Purpose, as the JSON body {"purpose": ...}; request.File is not
// read or sent. The response carries the file ID, the maximum size, the accepted MIME types and, in
// Links["upload"], the pre-signed URL: send the file content to that URL with an HTTP PUT whose body is
// the raw file bytes and whose Content-Type is one of DocumentTypesForPurpose.
//
// entityId is the {entityId} path value: the ID of the sub-entity.
//
// It returns the UploadFileResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) UploadFile(entityId string, request File) (*UploadFileResponse, error) {
	return c.UploadFileWithContext(context.Background(), entityId, request)
}

// UploadFileWithContext is UploadFile with a context for cancellation and deadlines.
func (c *Client) UploadFileWithContext(
	ctx context.Context,
	entityId string,
	request File,
) (*UploadFileResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	payload := struct {
		Purpose common.Purpose `json:"purpose"`
	}{request.Purpose}

	var response UploadFileResponse
	err = c.filesClient.PostWithContext(
		ctx,
		common.BuildPath(entitiesPath, entityId, filesPath),
		auth,
		payload,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RetrieveFile retrieves the details of a sub-entity's file
// (GET /entities/{entityId}/files/{fileId} on the Files host).
//
// entityId is the {entityId} path value: the ID of the sub-entity. fileId is the {fileId} path
// value: the ID of the file.
//
// It returns the FileDetailsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// fileId returns a CheckoutArgumentError and sends no request.
func (c *Client) RetrieveFile(entityId, fileId string) (*FileDetailsResponse, error) {
	return c.RetrieveFileWithContext(context.Background(), entityId, fileId)
}

// RetrieveFileWithContext is RetrieveFile with a context for cancellation and deadlines.
func (c *Client) RetrieveFileWithContext(ctx context.Context, entityId, fileId string) (*FileDetailsResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	fileId, err = common.EncodePathSegment(fileId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response FileDetailsResponse
	err = c.filesClient.GetWithContext(ctx, common.BuildPath(entitiesPath, entityId, filesPath, fileId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetEntityRequirements sends GET /accounts/entities/{id}/requirements: get a list of requirements.
//
// entityId is the {id} path value: the sub-entity's ID. schemaVersion is sent in the Accept header
// as the schema_version; an empty value uses DefaultSchemaVersion.
//
// It returns the EntityRequirementListResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) GetEntityRequirements(entityId string, schemaVersion string) (*EntityRequirementListResponse, error) {
	return c.GetEntityRequirementsWithContext(context.Background(), entityId, schemaVersion)
}

// GetEntityRequirementsWithContext is GetEntityRequirements with a caller-supplied context.
func (c *Client) GetEntityRequirementsWithContext(ctx context.Context, entityId string, schemaVersion string) (*EntityRequirementListResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	payload := struct {
		Headers *Headers `json:"-"`
	}{buildSchemaVersionHeaders(schemaVersion)}

	var response EntityRequirementListResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId, requirementsPath), auth, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetEntityRequirementDetails sends GET /accounts/entities/{id}/requirements/{requirementId}: get
// requirement details.
//
// entityId is the {id} path value: the sub-entity's ID. requirementId is the {requirementId} path
// value: the requirement's ID, as returned by GET /accounts/entities/{id}/requirements.
//
// It returns the EntityRequirementDetailsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// requirementId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetEntityRequirementDetails(entityId, requirementId string) (*EntityRequirementDetailsResponse, error) {
	return c.GetEntityRequirementDetailsWithContext(context.Background(), entityId, requirementId)
}

// GetEntityRequirementDetailsWithContext is GetEntityRequirementDetails with a caller-supplied context.
func (c *Client) GetEntityRequirementDetailsWithContext(ctx context.Context, entityId, requirementId string) (*EntityRequirementDetailsResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	requirementId, err = common.EncodePathSegment(requirementId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response EntityRequirementDetailsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId, requirementsPath, requirementId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ResolveEntityRequirement sends PUT /accounts/entities/{id}/requirements/{requirementId}: resolve
// a requirement.
//
// entityId is the {id} path value: the sub-entity's ID. requirementId is the {requirementId} path
// value: the requirement's ID, as returned by GET /accounts/entities/{id}/requirements. request is
// sent as the JSON body.
//
// It returns the EntityRequirementUpdateResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// requirementId returns a CheckoutArgumentError and sends no request.
func (c *Client) ResolveEntityRequirement(entityId, requirementId string, request EntityRequirementUpdateRequest) (*EntityRequirementUpdateResponse, error) {
	return c.ResolveEntityRequirementWithContext(context.Background(), entityId, requirementId, request)
}

// ResolveEntityRequirementWithContext is ResolveEntityRequirement with a caller-supplied context.
func (c *Client) ResolveEntityRequirementWithContext(ctx context.Context, entityId, requirementId string, request EntityRequirementUpdateRequest) (*EntityRequirementUpdateResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	requirementId, err = common.EncodePathSegment(requirementId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response EntityRequirementUpdateResponse
	err = c.apiClient.PutWithContext(ctx, common.BuildPath(accountsPath, entitiesPath, entityId, requirementsPath, requirementId), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
