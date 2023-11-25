package dto

import "github.com/google/uuid"

type UserPermissionPayloadData struct {
	ModuleID  uuid.UUID `json:"module_id"`
	Operation string    `json:"operation"`
}

type CreateUserModulePermission struct {
	Data []UserPermissionPayloadData `json:"data"`
}
