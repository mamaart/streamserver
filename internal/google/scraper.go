package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Video struct {
	Title   string `json:"title"`
	VideoID string `json:"videoId"`
}

func req(query string) *http.Request {
	u, err := url.Parse("https://youtube.com/results")
	if err != nil {
		log.Fatal(err)
	}

	u.RawQuery = url.Values{
		"search_query": []string{query},
	}.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", "Mozilla Fire Foxy Lady")
	return req
}

func Parse(data []byte) (out []Video, _ error) {
	s := string(data)
	start := strings.Index(s, "ytInitialData") + len("ytInitialData") + 3
	end := strings.Index(s[start:], "};") + start + 1

	if start < 3 || end < start {
		return out, fmt.Errorf("Failed to locate ytInitialData")
	}

	var model Data
	if err := json.Unmarshal([]byte(s[start:end]), &model); err != nil {
		return out, err
	}

	return model.GetVideos(), nil
}

func (d Data) GetVideos() (out []Video) {
	for _, e := range d.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents {
		for _, e := range e.ItemSectionRenderer.Contents {
			if e.VideoRenderer.VideoID != "" && len(e.VideoRenderer.Title.Runs) > 0 {
				out = append(out, Video{
					Title:   e.VideoRenderer.Title.Runs[0].Text,
					VideoID: e.VideoRenderer.VideoID,
				})
			}
		}
	}
	return out
}

func Search(query string) ([]Video, error) {
	resp, err := http.DefaultClient.Do(req(query))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("search result parsing failed: %w", err)
	}

	results, err := Parse(body)
	if err != nil {
		return nil, fmt.Errorf("search result parsing failed: %w", err)
	}

	if len(results) == 0 {
		return nil, errors.New("no results")
	}

	return results, nil
}
