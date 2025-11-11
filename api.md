# V0

Methods:

- <code title="get /api/v0/wallets">client.V0.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0Service.GetWallets">GetWallets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Health

Response Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0HealthCheckResponse">V0HealthCheckResponse</a>
- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0HealthCheckPublicResponse">V0HealthCheckPublicResponse</a>

Methods:

- <code title="get /api/v0/health">client.V0.Health.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0HealthCheckResponse">V0HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/health/public">client.V0.Health.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0HealthService.CheckPublic">CheckPublic</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0HealthCheckPublicResponse">V0HealthCheckPublicResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## TestPayment

Response Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#PaymentResponse">PaymentResponse</a>
- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TestPaymentRefundResponse">V0TestPaymentRefundResponse</a>

Methods:

- <code title="get /api/v0/test-payment/{id}">client.V0.TestPayment.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TestPaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#PaymentResponse">PaymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v0/test-payment">client.V0.TestPayment.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TestPaymentService.Process">Process</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TestPaymentProcessParams">V0TestPaymentProcessParams</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#PaymentResponse">PaymentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v0/test-payment/{id}/refund">client.V0.TestPayment.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TestPaymentService.Refund">Refund</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TestPaymentRefundResponse">V0TestPaymentRefundResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Customers

Params Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#CustomerStatus">CustomerStatus</a>
- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#CustomerType">CustomerType</a>

Methods:

- <code title="post /api/v0/customers">client.V0.Customers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerNewParams">V0CustomerNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/customers/{id}">client.V0.Customers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="patch /api/v0/customers/{id}">client.V0.Customers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerUpdateParams">V0CustomerUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/customers">client.V0.Customers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerListParams">V0CustomerListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### LiquidationAddresses

Response Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#LiquidationAddressResponse">LiquidationAddressResponse</a>

Methods:

- <code title="post /api/v0/customers/{customerId}/liquidation_addresses">client.V0.Customers.LiquidationAddresses.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerLiquidationAddressService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerLiquidationAddressNewParams">V0CustomerLiquidationAddressNewParams</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#LiquidationAddressResponse">LiquidationAddressResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/customers/{customerId}/liquidation_addresses/{liquidationAddressId}">client.V0.Customers.LiquidationAddresses.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerLiquidationAddressService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, liquidationAddressID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerLiquidationAddressGetParams">V0CustomerLiquidationAddressGetParams</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#LiquidationAddressResponse">LiquidationAddressResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/customers/{customerId}/liquidation_addresses">client.V0.Customers.LiquidationAddresses.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0CustomerLiquidationAddressService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#LiquidationAddressResponse">LiquidationAddressResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PaymentLinks

Methods:

- <code title="post /api/v0/payment-links">client.V0.PaymentLinks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentLinkNewParams">V0PaymentLinkNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/payment-links/{id}">client.V0.PaymentLinks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentLinkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /api/v0/payment-links/{id}">client.V0.PaymentLinks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentLinkService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentLinkUpdateParams">V0PaymentLinkUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/payment-links">client.V0.PaymentLinks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentLinkListParams">V0PaymentLinkListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## PaymentIntents

Params Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#BridgePaymentRail">BridgePaymentRail</a>
- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#StableCoinCurrency">StableCoinCurrency</a>

Methods:

- <code title="post /api/v0/payment-intents/bank">client.V0.PaymentIntents.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentIntentService.NewBank">NewBank</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentIntentNewBankParams">V0PaymentIntentNewBankParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/v0/payment-intents/stablecoin">client.V0.PaymentIntents.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentIntentService.NewStable">NewStable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0PaymentIntentNewStableParams">V0PaymentIntentNewStableParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#WebhookResponse">WebhookResponse</a>

Methods:

- <code title="post /api/v0/webhooks">client.V0.Webhooks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0WebhookNewParams">V0WebhookNewParams</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#WebhookResponse">WebhookResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/webhooks/{id}">client.V0.Webhooks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#WebhookResponse">WebhookResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/v0/webhooks/{id}">client.V0.Webhooks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0WebhookUpdateParams">V0WebhookUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#WebhookResponse">WebhookResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/webhooks">client.V0.Webhooks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0WebhookListParams">V0WebhookListParams</a>) ([]<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#WebhookResponse">WebhookResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v0/webhooks/{id}">client.V0.Webhooks.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#WebhookResponse">WebhookResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transfers

Methods:

- <code title="post /api/v0/transfers/direct-bank">client.V0.Transfers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferService.NewDirectBank">NewDirectBank</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferNewDirectBankParams">V0TransferNewDirectBankParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/v0/transfers/direct-wallet">client.V0.Transfers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferService.NewDirectWallet">NewDirectWallet</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferNewDirectWalletParams">V0TransferNewDirectWalletParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/v0/transfers/external-bank-transfer">client.V0.Transfers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferService.NewExternalBankTransfer">NewExternalBankTransfer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferNewExternalBankTransferParams">V0TransferNewExternalBankTransferParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/v0/transfers/external-stablecoin-transfer">client.V0.Transfers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferService.NewExternalStablecoinTransfer">NewExternalStablecoinTransfer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferNewExternalStablecoinTransferParams">V0TransferNewExternalStablecoinTransferParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/v0/transfers/stablecoin-conversion">client.V0.Transfers.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferService.NewStablecoinConversion">NewStablecoinConversion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TransferNewStablecoinConversionParams">V0TransferNewStablecoinConversionParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Balance

Response Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#AggregatedBalance">AggregatedBalance</a>
- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0BalanceGetAllStablecoinBalancesResponse">V0BalanceGetAllStablecoinBalancesResponse</a>

Methods:

- <code title="get /api/v0/balance">client.V0.Balance.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0BalanceService.GetAllStablecoinBalances">GetAllStablecoinBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0BalanceGetAllStablecoinBalancesResponse">V0BalanceGetAllStablecoinBalancesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/balance/eurc">client.V0.Balance.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0BalanceService.GetEurc">GetEurc</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#AggregatedBalance">AggregatedBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/balance/usdc">client.V0.Balance.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0BalanceService.GetUsdc">GetUsdc</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#AggregatedBalance">AggregatedBalance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ExchangeRate

Response Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#ExchangeRateResponse">ExchangeRateResponse</a>

Methods:

- <code title="get /api/v0/exchange-rate/eur-to-usd">client.V0.ExchangeRate.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ExchangeRateService.GetEurToUsd">GetEurToUsd</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#ExchangeRateResponse">ExchangeRateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/exchange-rate">client.V0.ExchangeRate.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ExchangeRateService.GetExchangeRate">GetExchangeRate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ExchangeRateGetExchangeRateParams">V0ExchangeRateGetExchangeRateParams</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#ExchangeRateResponse">ExchangeRateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/exchange-rate/usd-to-eur">client.V0.ExchangeRate.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ExchangeRateService.GetUsdToEur">GetUsdToEur</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#ExchangeRateResponse">ExchangeRateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Products

Methods:

- <code title="post /api/v0/products">client.V0.Products.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductNewParams">V0ProductNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/products/{id}">client.V0.Products.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /api/v0/products/{id}">client.V0.Products.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductUpdateParams">V0ProductUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/products">client.V0.Products.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductListParams">V0ProductListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /api/v0/products/{id}">client.V0.Products.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /api/v0/products/{id}/images">client.V0.Products.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0ProductService.UploadImages">UploadImages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Invoices

Params Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#CreateInvoiceParam">CreateInvoiceParam</a>

Methods:

- <code title="post /api/v0/invoices">client.V0.Invoices.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0InvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0InvoiceNewParams">V0InvoiceNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/invoices/{id}">client.V0.Invoices.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0InvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /api/v0/invoices/{id}">client.V0.Invoices.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0InvoiceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0InvoiceUpdateParams">V0InvoiceUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/invoices">client.V0.Invoices.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0InvoiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0InvoiceListParams">V0InvoiceListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Taxes

Response Types:

- <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxNewResponse">V0TaxNewResponse</a>

Methods:

- <code title="post /api/v0/taxes">client.V0.Taxes.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxNewParams">V0TaxNewParams</a>) (<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxNewResponse">V0TaxNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v0/taxes/{id}">client.V0.Taxes.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /api/v0/taxes/{id}">client.V0.Taxes.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxUpdateParams">V0TaxUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /api/v0/taxes">client.V0.Taxes.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go">devdraft</a>.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxListParams">V0TaxListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /api/v0/taxes/{id}">client.V0.Taxes.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /api/v0/taxes">client.V0.Taxes.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxService.DeleteAll">DeleteAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /api/v0/taxes">client.V0.Taxes.<a href="https://pkg.go.dev/github.com/devdraftengineer/devdraft-go#V0TaxService.UpdateAll">UpdateAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
