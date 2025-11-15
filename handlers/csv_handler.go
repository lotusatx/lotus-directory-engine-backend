package handlers

import (
	"fmt"
	"encoding/csv"
	"github.com/lotusatx/lotus-directory-engine-backend/models"
)

func readCsvFile(filePath string) ([]models.User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	var groups []models.Group
	for i, record := range records {
		if i == 0 {
			continue 		
		}
		group := models.Group{
			Name: record[0],
			Description: record[1],
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		groups = append(groups, group)
	}

	return groups, nil
}