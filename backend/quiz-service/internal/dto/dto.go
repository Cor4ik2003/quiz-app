package dto

type Answer struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type Question struct {
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

type CreateQuizRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type AnswerInput struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type QuestionInput struct {
	Text    string        `json:"text"`
	Answers []AnswerInput `json:"answers"`
}

type QuizInput struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Code        string          `json:"code"`
	Questions   []QuestionInput `json:"questions"`
}
