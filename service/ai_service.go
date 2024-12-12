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

// AnalyzeData mengirimkan permintaan ke endpoint AI untuk menganalisis data tabel berdasarkan query.
// AnalyzeData mengirimkan permintaan ke endpoint AI untuk menganalisis data tabel berdasarkan query.
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

    // Log respons API untuk debugging
    // fmt.Printf("Debug: Response from API: %v\n", tapasResponse)

    // Validasi struktur respons dan ambil data
    cells, exists := tapasResponse["cells"].([]interface{})
    if !exists || len(cells) == 0 {
        return "", errors.New("no valid answer found in the response")
    }

    // Validasi tipe data elemen pertama
    answer, ok := cells[0].(string)
    if !ok {
        return "", errors.New("answer is not a valid string")
    }

    return answer, nil
}


// ChatWithAI mengirimkan permintaan ke endpoint AI untuk mendapatkan respon berbasis konteks dan query.

// ChatWithAI mengirimkan permintaan ke model Phi-3.5-mini-instruct di Hugging Face


// ChatWithAI mengirimkan permintaan ke model Phi-3.5-mini-instruct di Hugging Face
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
