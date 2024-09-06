package model

import "database/sql/driver"

type PlanCourseType string

// Constants for PlanCourseType
const (
    PlanCourseTypeNormal PlanCourseType = "Normal"
    PlanCourseTypeFE     PlanCourseType = "FE"
    PlanCourseTypeMinor  PlanCourseType = "Minor"
)

// PlanCourseType methods
func (pct PlanCourseType) Value() (driver.Value, error) {
    return string(pct), nil
}

func (pct *PlanCourseType) Scan(value interface{}) error {
    *pct = PlanCourseType(value.(string))
    return nil
}