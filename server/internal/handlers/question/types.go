package question

type CreateQuestionRequest struct {
	Text      string   `json:"text" binding:"required"`
	Options   []string `json:"options" binding:"required,min=2"`
	Answer    int      `json:"answer" binding:"required"`
	TimeLimit int      `json:"time_limit" binding:"required,min=5,max=120"`
}
