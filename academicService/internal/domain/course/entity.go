package course

import "github.com/shopspring/decimal"

type Course struct {
	CourseID            int64
	Code                string
	Name                string
	OfficialName        string
	HardnessCoefficient decimal.Decimal
	InstituteId         int64
	LecHours            int64
	LabHours            int64
}
