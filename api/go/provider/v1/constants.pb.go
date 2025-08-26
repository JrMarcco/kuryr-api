package providerv1

const (
	FieldId               = "id"
	FieldProviderName     = "provider_name"
	FieldChannel          = "channel"
	FieldEndpoint         = "endpoint"
	FieldRegionId         = "region_id"
	FieldAppId            = "app_id"
	FieldApiKey           = "api_key"
	FieldApiSecret        = "api_secret"
	FieldWeight           = "weight"
	FieldQpsLimit         = "qps_limit"
	FieldDailyLimit       = "daily_limit"
	FieldAuditCallbackUrl = "audit_callback_url"
	FieldActiveStatus     = "active_status"
)

var UpdatableFields = map[string]struct{}{
	FieldProviderName:     {},
	FieldChannel:          {},
	FieldEndpoint:         {},
	FieldRegionId:         {},
	FieldApiKey:           {},
	FieldApiSecret:        {},
	FieldWeight:           {},
	FieldQpsLimit:         {},
	FieldDailyLimit:       {},
	FieldAuditCallbackUrl: {},
	FieldActiveStatus:     {},
}
