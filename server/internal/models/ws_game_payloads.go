package models

type StartGamePayload struct {
	QuizID string `json:"quiz_id"`
}

type QuestionPayload struct {
	ID        string   `json:"id"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	TimeLimit int      `json:"time_limit"`
}

type AnswerSubmissionPayload struct {
	QuestionID string `json:"question_id"`
	Option     int    `json:"option"`
}
