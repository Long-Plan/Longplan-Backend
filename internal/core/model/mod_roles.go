package model

import "database/sql/driver"

type ModRoles string

// Constants for ModRoles
const (
    ModRolesLecturer  ModRoles = "lecturer"
    ModRolesStaff     ModRoles = "staff"
    ModRolesRegistrar ModRoles = "registrar"
    ModRolesAdmin     ModRoles = "admin"
)

// ModRoles methods
func (mr ModRoles) Value() (driver.Value, error) {
    return string(mr), nil
}

func (mr *ModRoles) Scan(value interface{}) error {
    *mr = ModRoles(value.(string))
    return nil
}