package soundcloud

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

type Video struct {
	Url   string
	Title string
}

func (s *soundcloud) Search(query string) (out []Video, _ error) {
	client := http.DefaultClient
	u, err := url.Parse("https://api-v2.soundcloud.com/search/tracks")
	if err != nil {
		panic(err)
	}
	u.RawQuery = url.Values{
		"q":         []string{query},
		"client_id": []string{s.id},
		"limit":     []string{strconv.Itoa(10)},
		"offset":    []string{strconv.Itoa(0)},
	}.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla Fire Foxy Lady")
	resp, err := client.Do(req)
	if err != nil {
		return out, err
	}
	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return out, err
	}

	for _, e := range result.Collection {
		fmt.Println(e.User.Username, e.Title, e.MonetizationModel)
		out = append(out, Video{
			Title: fmt.Sprintf("%s - %s", e.User.Username, e.Title),
			Url:   e.PermalinkURL,
		})
	}

	return out, nil
}

func GetClientID() (string, error) {
	client := http.DefaultClient

	urlPattern := regexp.MustCompile(`<script[^>]+src="([^"]+)"`)
	idPattern := regexp.MustCompile(`client_id\s*:\s*"([0-9a-zA-Z]{32})"`)

	resp, err := client.Get("https://soundcloud.com")
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	found := urlPattern.FindAllSubmatch(data, -1)

	for i := len(found); i > 0; i-- {
		resp, err := http.Get(string(found[i-1][1]))
		if err == nil {
			data, err := io.ReadAll(resp.Body)
			if err == nil {
				found := idPattern.FindSubmatch(data)
				if found != nil {
					return string(found[1]), nil
				}
			}
		}
	}
	return "", nil
}
