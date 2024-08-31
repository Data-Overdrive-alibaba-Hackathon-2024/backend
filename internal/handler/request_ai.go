package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/data-overdrive-alibaba-hackathon-2024/config"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"go.uber.org/zap"
	"io"
	"net/http"
)

var (
	modelUrl = config.ModelUrl()
	modelKey = config.ModelKey()
)

func (h *questionHandler) RequestAI(level int) (model.GenerateQuestionAIResponse, error) {
	var questionText model.GenerateQuestionAIResponse

	payload := model.GenerateQuestionAIRequest{
		Input: model.GenerateQuestionAIInput{
			Prompt: fmt.Sprintf("give me question with topic level %d", level),
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		h.logger.Error("failed to marshal request", zap.Error(err))
		return questionText, err
	}

	request, err := http.NewRequest("POST", modelUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		h.logger.Error("failed to create request", zap.Error(err))
		return questionText, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+modelKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		h.logger.Error("failed to send request", zap.Error(err))
		return questionText, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		h.logger.Error("failed to read response", zap.Error(err))
		return questionText, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		h.logger.Error("failed to unmarshal response", zap.Error(err))
		return questionText, err
	}

	output, ok := result["output"].(map[string]interface{})
	if !ok {
		h.logger.Error("invalid response format")
		return questionText, err
	}

	text, ok := output["text"].(string)
	if !ok {
		h.logger.Error("invalid response format")
		return questionText, err
	}

	if err := json.Unmarshal([]byte(text), &questionText); err != nil {
		h.logger.Error("failed to unmarshal response", zap.Error(err))
		return questionText, err
	}

	return questionText, nil
}
