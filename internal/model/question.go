package model

type GenerateQuestionRequest struct {
	UserId string `json:"user_id"`
	Level  int    `json:"level"`
}

type GenerateQuestionAIRequest struct {
	Input GenerateQuestionAIInput `json:"input"`
}

type GenerateQuestionAIInput struct {
	Prompt string `json:"prompt"`
}

type GenerateQuestionAIResponse struct {
	Question      string                           `json:"question"`
	Options       OptionGenerateQuestionAIResponse `json:"options"`
	CorrectAnswer string                           `json:"correct_answer"`
}

type OptionGenerateQuestionAIResponse struct {
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
	Option4 string `json:"option4"`
}

type InsertQuestionInput struct {
	UserId        string                           `json:"user_id"`
	Level         int                              `json:"level"`
	Question      string                           `json:"question"`
	Options       OptionGenerateQuestionAIResponse `json:"options"`
	CorrectAnswer string                           `json:"correct_answer"`
}
