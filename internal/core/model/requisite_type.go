package model

import "database/sql/driver"

type RequisiteType string

const (
    RequisiteTypeAny RequisiteType = "Any"
    RequisiteTypeAll RequisiteType = "All"
    RequisiteTypeCo  RequisiteType = "Co"
)

// RequisiteType methods
func (rt RequisiteType) Value() (driver.Value, error) {
    return string(rt), nil
}

func (rt *RequisiteType) Scan(value interface{}) error {
    *rt = RequisiteType(value.(string))
    return nil
}