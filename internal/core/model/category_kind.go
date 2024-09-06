package model

import "database/sql/driver"

type CategoryKind string

const (
    CategoryKindC1 CategoryKind = "C1"
    CategoryKindC2 CategoryKind = "C2"
    CategoryKindG1 CategoryKind = "G1"
)

// CategoryKind methods
func (ck CategoryKind) Value() (driver.Value, error) {
    return string(ck), nil
}

func (ck *CategoryKind) Scan(value interface{}) error {
    *ck = CategoryKind(value.(string))
    return nil
}