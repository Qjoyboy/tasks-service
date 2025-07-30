package task

type Task struct {
	ID     string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
	UserID string `json:"user_id"`
}

type TaskRequest struct {
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
}
