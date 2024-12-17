package service

import (
	"encoding/csv"
	"errors"
	"strings"

	repository "a21hc3NpZ25tZW50/repository/fileRepository"
)

type FileService struct {
	Repo *repository.FileRepository
}

// ProcessFile memproses konten file CSV dan mengembalikannya dalam bentuk map
func (s *FileService) ProcessFile(fileContent string) (map[string][]string, error) {
    reader := csv.NewReader(strings.NewReader(fileContent))
    records, err := reader.ReadAll()
    if err != nil {
        return nil, errors.New("failed to read CSV content: " + err.Error())
    }

    if len(records) < 2 {
        return nil, errors.New("CSV file must have a header and at least one row of data")
    }

    result := make(map[string][]string)
    headers := records[0]

    for _, row := range records[1:] {
		for i, value := range row {
			// Ensure the column index exists in the header
			if i < len(headers) {
				header := headers[i]
				result[header] = append(result[header], value)
			}
		}
	}
	return result, nil
}