package configv1

const (
	FieldId             = "id"
	FieldBizId          = "biz_id"
	FieldChannelConfig  = "channel_config"
	FieldQuotaConfig    = "quota_config"
	FieldCallbackConfig = "callback_config"
	FieldRateLimit      = "rate_limit"
	FieldCreatedAt      = "created_at"
	FieldUpdatedAt      = "updated_at"
)

var UpdatableFields = map[string]struct{}{
	FieldChannelConfig:  {},
	FieldQuotaConfig:    {},
	FieldCallbackConfig: {},
	FieldRateLimit:      {},
}
