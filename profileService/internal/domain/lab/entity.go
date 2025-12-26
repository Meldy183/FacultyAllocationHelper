package lab

import "github.com/google/uuid"

type Lab struct {
	LabID       uuid.UUID
	Name        string
	InstituteID uuid.UUID
}
