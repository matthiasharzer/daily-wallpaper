package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const bingBaseURL = "https://www.bing.com"
const bingAPITemplateURL = "https://www.bing.com/HPImageArchive.aspx?format=js&mkt=%s&idx=%d&n=%d"

type ResponseImage struct {
	URLBase   string `json:"urlbase"`
	Title     string `json:"title"`
	Copyright string `json:"copyright"`
}

func (r ResponseImage) FullSizeURL() string {
	return fmt.Sprintf("%s%s_UHD.jpg", bingBaseURL, r.URLBase)
}

type Response struct {
	Images []ResponseImage `json:"images"`
}

type API struct {
	market string
}

func NewAPI(market string) *API {
	return &API{
		market: market,
	}
}

func (a *API) buildURL(startIndex, count int) string {
	return fmt.Sprintf(bingAPITemplateURL, a.market, startIndex, count)
}

func (a *API) GetWallpapers(startIndex, count int) ([]ResponseImage, error) {
	if count <= 0 {
		return nil, fmt.Errorf("count must be greater than 0")
	}
	// limitation of the Bing API
	if startIndex < 0 || startIndex > 15 {
		return nil, fmt.Errorf("startIndex must be between 0 and 15")
	}

	url := a.buildURL(startIndex, count)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch wallpapers: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return response.Images, nil
}
