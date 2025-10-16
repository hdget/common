package biz

const (
	ContextKeyAppId         = "hd-app-id"   // 应用ID
	ContextKeyClient        = "hd-client"   // 客户端
	ContextKeyRelease       = "hd-release"  // 版本号
	ContextKeyTsn           = "hd-tsn"      // tenant sn
	ContextKeyTid           = "hd-tid"      // tenant id
	ContextKeyUid           = "hd-uid"      // user id
	ContextKeyUsn           = "hd-usn"      // user sn
	ContextKeyRoleIds       = "hd-role-ids" // role ids
	ContextKeyCaller        = "dapr-caller-app-id"
	ContextKeyDbTransaction = "hd-db-tx"
)
