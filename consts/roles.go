package consts

// RoleModules
var RoleModules = struct {
	Forms string
}{
	Forms: "|forms",
}

// RolePermissions
var RolePermissions = struct {
	Create,
	Edit,
	View,
	Delete string
}{
	Create: "create",
	Edit:   "edit",
	View:   "view",
	Delete: "delete",
}
