package types

type Decl int

const (
	smiDeclUnknown Decl = iota
	smiDeclModuleIdentity
	smiDeclObjectType
	smiDeclObjectIdentity
	smiDeclTrapType
	smiDeclNotificationType
	smiDeclGroup
	smiDeclCompliance
	smiDeclCapabilities
	smiDeclObjectGroup
	smiDeclNotificationGroup
)
