package models

import "github.com/google/uuid"

type BranchLocation struct {
	ID        uuid.UUID `gorm:"primary_key; type:uuid; default:uuid_generate_v4()" json:"id"`
	BranchID  uuid.UUID `gorm:"foreignKey:id" json:"branch_id"`
	Latitude  float64   `gorm:"type:decimal(10,8); default:0" json:"latitude" validate:"numeric"`
	Longitude float64   `gorm:"type:decimal(11,8); default:0" json:"longitude" validate:"numeric"`
	Distance  float64   `json:"distance,omitempty"`
}
