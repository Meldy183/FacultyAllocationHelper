package userprofile

type AddProfileRequest struct {
	NameEnglish      string `json:"name_english"`
	Email            string `json:"email"`
	Alias            string `json:"alias"`
	InstituteID      int    `json:"institute_id"`
	Position         string `json:"position"`
	IsRepresentative bool   `json:"isRepr"`
}
