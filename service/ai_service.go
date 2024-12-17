package service

import (
	"a21hc3NpZ25tZW50/model"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type AIService struct {
	Client HTTPClient
}

func (s *AIService) AnalyzeData(table map[string][]string, query, token string) (string, error) {
    if len(table) == 0 {
        return "", errors.New("table is empty")
    }

    requestBody := model.AIRequest{
        Inputs: model.Inputs{
            Table: table,
            Query: query,
        },
    }
    body, err := json.Marshal(requestBody)
    if err != nil {
        return "", errors.New("failed to marshal request body: " + err.Error())
    }

    req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/google/tapas-large-finetuned-wtq", bytes.NewBuffer(body))
    if err != nil {
        return "", errors.New("failed to create HTTP request: " + err.Error())
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    resp, err := s.Client.Do(req)
    if err != nil {
        return "", errors.New("failed to send HTTP request: " + err.Error())
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", errors.New("received non-200 response: " + resp.Status)
    }

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", errors.New("failed to read response body: " + err.Error())
    }

    var tapasResponse map[string]interface{}
    err = json.Unmarshal(responseBody, &tapasResponse)
    if err != nil {
        return "", errors.New("failed to unmarshal response: " + err.Error())
    }

    answer, exists := tapasResponse["cells"].([]interface{})
    if !exists {
        return "", errors.New("no answer found in the response")
    }
    // Assuming the result is a string within the cells array
    if len(answer) > 0 {
        return answer[0].(string), nil
    }
    
    return "", errors.New("no valid answer in the response")
}

func (s *AIService) ChatWithAI(context, query, token string) (model.ChatResponse, error) {
    input := context + "\n" + query
    requestBody := map[string]string{
        "inputs": input,
    }
    body, err := json.Marshal(requestBody)
    if err != nil {
        return model.ChatResponse{}, errors.New("failed to marshal request body: " + err.Error())
    }

    req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/microsoft/Phi-3.5-mini-instruct", bytes.NewBuffer(body))
    if err != nil {
        return model.ChatResponse{}, errors.New("failed to create HTTP request: " + err.Error())
    }
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    resp, err := s.Client.Do(req)
    if err != nil {
        return model.ChatResponse{}, errors.New("failed to send HTTP request: " + err.Error())
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return model.ChatResponse{}, errors.New("received non-200 response: " + resp.Status)
    }

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return model.ChatResponse{}, errors.New("failed to read response body: " + err.Error())
    }

    var chatResponse []map[string]string
    err = json.Unmarshal(responseBody, &chatResponse)
    if err != nil {
        return model.ChatResponse{}, errors.New("failed to unmarshal response: " + err.Error())
    }

    if len(chatResponse) > 0 {
        generatedText, exists := chatResponse[0]["generated_text"]
        if exists {
            return model.ChatResponse{
                GeneratedText: generatedText,
            }, nil
        }
    }

    return model.ChatResponse{}, errors.New("no response from AI")
}
