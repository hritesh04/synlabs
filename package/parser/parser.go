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
	var profile domain.Profile

	content, err := file.Open()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, r.ParserUrl, content)
	if err != nil {
		return nil, err
	}
	for key, header := range r.header {
		req.Header.Add(key, header)
	}
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, err
}
