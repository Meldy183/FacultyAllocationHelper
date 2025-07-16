package logpage

type GetLogpagesResponce struct {
	Logpages *[]Logpage `json:"logpages"`
}
type Logpage struct {
	LogID     int64  `json:"log_id"`
	UserID    string `json:"user_id"`
	Action    string `json:"action"`
	SubjectID int64  `json:"subject_id"`
	CreatedAt string `json:"created_at"`
}
