package killgrave

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Request struct {
	Method   string `json:"method"`
	Endpoint string `json:"endpoint"`
}

type Response struct {
	Status   int               `json:"status"`
	Headers  map[string]string `json:"headers"`
	BodyFile string            `json:"bodyFile"`
}

type Imposter struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

// generateImposterFile creates the auto.imp.json file in the directory
func generateImposterFile(endpoint string) error {
	imposters := []Imposter{
		{
			Request: Request{
				Method:   "GET",
				Endpoint: endpoint,
			},
			Response: Response{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				BodyFile: "-",
			},
		},
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	outputPath := filepath.Join(workingDir, "imposters", "auto.imp.json")

	data, err := json.MarshalIndent(imposters, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(outputPath, data, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Println("File created at:", outputPath)
	return nil
}
