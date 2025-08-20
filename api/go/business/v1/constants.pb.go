package businessv1

const (
	FieldId           = "id"
	FieldBizKey       = "biz_key"
	FieldBizName      = "biz_name"
	FieldBizType      = "biz_type"
	FieldBizSecret    = "biz_secret"
	FieldContact      = "contact"
	FieldContactEmail = "contact_email"
	FieldCreatorId    = "creator_id"
	FieldCreatedAt    = "created_at"
	FieldUpdatedAt    = "updated_at"
)

var UpdatableFields = map[string]struct{}{
	FieldBizName:      {},
	FieldContact:      {},
	FieldContactEmail: {},
}
