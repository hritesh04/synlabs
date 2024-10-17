package parser

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/hritesh04/synlabs/internal/domain"
)

type ResumeParser struct {
	ParserUrl string
	header    map[string]string
}

func NewResumeParser(apikey, url string) *ResumeParser {
	header := map[string]string{
		"Content-Type": "application/octet-stream",
		"apikey":       apikey,
	}
	return &ResumeParser{
		ParserUrl: url,
		header:    header,
	}
}

func (r *ResumeParser) Parse(file *multipart.FileHeader) (*domain.Profile, error) {
	result := make(map[string]interface{})

	content, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer content.Close()
	req, err := http.NewRequest(http.MethodPost, r.ParserUrl, content)
	if err != nil {
		return nil, err
	}
	for key, header := range r.header {
		req.Header.Add(key, header)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	profile := &domain.Profile{
		Name:              getString(result, "name"),
		Email:             getString(result, "email"),
		Phone:             getString(result, "phone"),
		Education:         getString(result, "education"),
		Skills:            getString(result, "skills"),
		Experience:        getString(result, "experience"),
		ResumeFileAddress: getString(result, "resume_file_address"),
	}
	return profile, nil
}

func getString(m map[string]interface{}, key string) string {
	if value, ok := m[key]; ok {
		if str, ok := value.(string); ok {
			fmt.Println(str)
			return str
		}
	}
	return ""
}
