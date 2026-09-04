package nas

import (
	"github.com/checkout/checkout-sdk-go/v3/accounts"
	"github.com/checkout/checkout-sdk-go/v3/agenticcommerce"
	"github.com/checkout/checkout-sdk-go/v3/apm/bacs"
	"github.com/checkout/checkout-sdk-go/v3/apm/ideal"
	"github.com/checkout/checkout-sdk-go/v3/apm/klarna"
	"github.com/checkout/checkout-sdk-go/v3/apm/sepa"
	"github.com/checkout/checkout-sdk-go/v3/balances"
	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/compliancerequests"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/customers"
	"github.com/checkout/checkout-sdk-go/v3/disputes"
	"github.com/checkout/checkout-sdk-go/v3/financial"
	"github.com/checkout/checkout-sdk-go/v3/forex"
	"github.com/checkout/checkout-sdk-go/v3/forward"
	"github.com/checkout/checkout-sdk-go/v3/identities/addressdocumentverification"
	"github.com/checkout/checkout-sdk-go/v3/identities/amlscreening"
	"github.com/checkout/checkout-sdk-go/v3/identities/applicants"
	"github.com/checkout/checkout-sdk-go/v3/identities/faceauthentication"
	"github.com/checkout/checkout-sdk-go/v3/identities/iddocumentverification"
	"github.com/checkout/checkout-sdk-go/v3/identities/identityverification"
	instruments "github.com/checkout/checkout-sdk-go/v3/instruments/nas"
	"github.com/checkout/checkout-sdk-go/v3/issuing"
	"github.com/checkout/checkout-sdk-go/v3/issuing/cardholdertokens"
	"github.com/checkout/checkout-sdk-go/v3/metadata"
	"github.com/checkout/checkout-sdk-go/v3/networktokens"
	"github.com/checkout/checkout-sdk-go/v3/onboardingsimulator"
	"github.com/checkout/checkout-sdk-go/v3/paymentmethods"
	"github.com/checkout/checkout-sdk-go/v3/payments/applepay"
	"github.com/checkout/checkout-sdk-go/v3/payments/contexts"
	"github.com/checkout/checkout-sdk-go/v3/payments/googlepay"
	"github.com/checkout/checkout-sdk-go/v3/payments/hosted"
	"github.com/checkout/checkout-sdk-go/v3/payments/links"
	payments "github.com/checkout/checkout-sdk-go/v3/payments/nas"
	payment_sessions "github.com/checkout/checkout-sdk-go/v3/payments/sessions"
	"github.com/checkout/checkout-sdk-go/v3/payments/setups"
	"github.com/checkout/checkout-sdk-go/v3/reports"
	"github.com/checkout/checkout-sdk-go/v3/sessions"
	"github.com/checkout/checkout-sdk-go/v3/standaloneaccountupdater"
	"github.com/checkout/checkout-sdk-go/v3/tokens"
	"github.com/checkout/checkout-sdk-go/v3/transfers"
	"github.com/checkout/checkout-sdk-go/v3/workflows"
)

type Api struct {
	Accounts                    *accounts.Client
	Balances                    *balances.Client
	Customers                   *customers.Client
	Disputes                    *disputes.Client
	Financial                   *financial.Client
	Forex                       *forex.Client
	Hosted                      *hosted.Client
	Instruments                 *instruments.Client
	Links                       *links.Client
	Metadata                    *metadata.Client
	Payments                    *payments.Client
	Sessions                    *sessions.Client
	Tokens                      *tokens.Client
	Transfers                   *transfers.Client
	WorkFlows                   *workflows.Client
	Reports                     *reports.Client
	Issuing                     *issuing.Client
	CardholderTokens            *cardholdertokens.Client
	Contexts                    *contexts.Client
	PaymentSessions             *payment_sessions.Client
	PaymentSetups               *setups.Client
	Forward                     *forward.Client
	ApplePay                    *applepay.Client
	GooglePay                   *googlepay.Client
	NetworkTokens               *networktokens.Client
	StandaloneAccountUpdater    *standaloneaccountupdater.Client
	AgenticCommerce             *agenticcommerce.Client
	ComplianceRequests          *compliancerequests.Client
	PaymentMethods              *paymentmethods.Client
	AmlScreening                *amlscreening.Client
	Applicants                  *applicants.Client
	FaceAuthentication          *faceauthentication.Client
	IdDocumentVerification      *iddocumentverification.Client
	AddressDocumentVerification *addressdocumentverification.Client
	IdentityVerification        *identityverification.Client

	Ideal  *ideal.Client
	Klarna *klarna.Client
	Sepa   *sepa.Client
	Bacs   *bacs.Client

	OnboardingSimulator *onboardingsimulator.Client
}

func CheckoutApi(configuration *configuration.Configuration) *Api {
	apiClient := buildBaseClient(configuration)

	api := Api{}
	api.Accounts = accounts.NewClient(configuration, apiClient, buildFilesClient(configuration))
	api.Balances = balances.NewClient(configuration, buildBalancesClient(configuration))
	api.Customers = customers.NewClient(configuration, apiClient)
	api.Disputes = disputes.NewClient(configuration, apiClient)
	api.Instruments = instruments.NewClient(configuration, apiClient)
	api.Financial = financial.NewClient(configuration, apiClient)
	api.Forex = forex.NewClient(configuration, apiClient)
	api.Hosted = hosted.NewClient(configuration, apiClient)
	api.Links = links.NewClient(configuration, apiClient)
	api.Metadata = metadata.NewClient(configuration, apiClient)
	api.Payments = payments.NewClient(configuration, apiClient)
	api.Sessions = sessions.NewClient(configuration, apiClient)
	api.Tokens = tokens.NewClient(configuration, apiClient)
	api.Transfers = transfers.NewClient(configuration, buildTransfersClient(configuration))
	api.WorkFlows = workflows.NewClient(configuration, apiClient)
	api.Reports = reports.NewClient(configuration, apiClient)
	api.Issuing = issuing.NewClient(configuration, apiClient)
	api.CardholderTokens = cardholdertokens.NewClient(configuration, apiClient)
	api.Contexts = contexts.NewClient(configuration, apiClient)
	api.PaymentSessions = payment_sessions.NewClient(configuration, apiClient)
	api.PaymentSetups = setups.NewClient(configuration, apiClient)
	api.Forward = forward.NewClient(configuration, buildForwardClient(configuration))
	api.ApplePay = applepay.NewClient(configuration, apiClient)
	api.GooglePay = googlepay.NewClient(configuration, apiClient)
	api.NetworkTokens = networktokens.NewClient(configuration, apiClient)
	api.StandaloneAccountUpdater = standaloneaccountupdater.NewClient(configuration, apiClient)
	api.AgenticCommerce = agenticcommerce.NewClient(configuration, apiClient)
	api.ComplianceRequests = compliancerequests.NewClient(configuration, apiClient)
	api.PaymentMethods = paymentmethods.NewClient(configuration, apiClient)
	identityClient := buildIdentityClient(configuration)
	api.AmlScreening = amlscreening.NewClient(configuration, identityClient)
	api.Applicants = applicants.NewClient(configuration, identityClient)
	api.FaceAuthentication = faceauthentication.NewClient(configuration, identityClient)
	api.IdDocumentVerification = iddocumentverification.NewClient(configuration, identityClient)
	api.AddressDocumentVerification = addressdocumentverification.NewClient(configuration, identityClient)
	api.IdentityVerification = identityverification.NewClient(configuration, identityClient)

	api.Ideal = ideal.NewClient(configuration, apiClient)
	api.Klarna = klarna.NewClient(configuration, apiClient)
	api.Sepa = sepa.NewClient(configuration, apiClient)
	api.Bacs = bacs.NewClient(configuration, apiClient)
	api.OnboardingSimulator = onboardingsimulator.NewClient(configuration, apiClient)
	return &api
}

func buildBaseClient(configuration *configuration.Configuration) client.HttpClient {
	if configuration.EnvironmentSubdomain != nil {
		return client.NewApiClient(configuration, configuration.EnvironmentSubdomain.ApiUrl)
	}
	return client.NewApiClient(configuration, configuration.Environment.BaseUri())
}

func buildFilesClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.FilesUri())
}

func buildBalancesClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.BalancesUri())
}

func buildTransfersClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.TransfersUri())
}

func buildForwardClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.ForwardUri())
}

func buildIdentityClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.IdentityUri())
}
