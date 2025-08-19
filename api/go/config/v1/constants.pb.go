package configv1

const (
	FieldId        = "id"
	FieldBizId     = "biz_id"
	FieldRateLimit = "rate_limit"
)

// 渠道配置字段 ( ChannelConfig )
const (
	FieldChanneConfigItems       = "channel_config.items"
	FieldChanneConfigRetryPolicy = "channel_config.retry_policy"
)

// 配额配置字段 ( QuotaConfig )
const (
	FieldQuotaConfigMonthly = "quota_config.monthly"
	FieldQuotaConfigDaily   = "quota_config.daily"
)

const (
	FieldQuotaConfigMonthlySms   = "quota_config.monthly.sms"
	FieldQuotaConfigMonthlyEmail = "quota_config.monthly.email"
)

const (
	FieldQuotaConfigDailySms   = "quota_config.daily.sms"
	FieldQuotaConfigDailyEmail = "quota_config.daily.email"
)

// 回调配置字段 ( CallbackConfig )
const (
	FieldCallbackConfigServiceName = "callback_config.service_name"
	FieldCallbackConfigRetryPolicy = "callback_config.retry_policy"
)

var InsertableFields = map[string]struct{}{
	FieldBizId:                     {},
	FieldRateLimit:                 {},
	FieldChanneConfigItems:         {},
	FieldChanneConfigRetryPolicy:   {},
	FieldQuotaConfigMonthly:        {},
	FieldQuotaConfigDaily:          {},
	FieldQuotaConfigMonthlySms:     {},
	FieldQuotaConfigMonthlyEmail:   {},
	FieldQuotaConfigDailySms:       {},
	FieldQuotaConfigDailyEmail:     {},
	FieldCallbackConfigServiceName: {},
	FieldCallbackConfigRetryPolicy: {},
}

var UpdatableFields = map[string]struct{}{
	FieldRateLimit:                 {},
	FieldChanneConfigItems:         {},
	FieldChanneConfigRetryPolicy:   {},
	FieldQuotaConfigMonthly:        {},
	FieldQuotaConfigDaily:          {},
	FieldQuotaConfigMonthlySms:     {},
	FieldQuotaConfigMonthlyEmail:   {},
	FieldQuotaConfigDailySms:       {},
	FieldQuotaConfigDailyEmail:     {},
	FieldCallbackConfigServiceName: {},
	FieldCallbackConfigRetryPolicy: {},
}
