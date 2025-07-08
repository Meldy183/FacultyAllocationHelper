package userprofile

type AddProfileRequest struct {
	NameEnglish string `json:"name_english"`
	Email       string `json:"email"`
	Alias       string `json:"alias"`
	Institute   string `json:"institute"`
	Position    string `json:"position"`
}
