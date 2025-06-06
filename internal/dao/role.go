// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"tg_bot_backend/internal/dao/internal"
)

// internalRoleDao is an internal type for wrapping the internal DAO implementation.
type internalRoleDao = *internal.RoleDao

// roleDao is the data access object for the table role.
// You can define custom methods on it to extend its functionality as needed.
type roleDao struct {
	internalRoleDao
}

var (
	// Role is a globally accessible object for table role operations.
	Role = roleDao{
		internal.NewRoleDao(),
	}
)

// Add your custom methods and functionality below.
