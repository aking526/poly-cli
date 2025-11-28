package polymarket

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// -- HTTP Client --
type Client struct {
	httpClient *http.Client
	baseUrl string
}

func NewClient(baseUrl string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: baseUrl,
	}
}

func NewGammaClient() *Client {
	return NewClient("https://gamma-api.polymarket.com")
}

// -- Data Models --
type Market struct {
	ID string `json:"id"`
	Question string `json:"question"`
	Slug string `json:"slug"`
	Volume string `json:"volume"`
	Outcome string `json:"outcome"`
	Price string `json:"price"`
}

// -- Query Functions --
func (c *Client) GetTrendingMarkets(limit int) ([]Market, error) {
	url := fmt.Sprintf("%s/markets?limit=%d&active=true&closed=false&order=volume&ascending=false", c.baseUrl, limit)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON into MarketAPIResponse slice
	var apiResponses []MarketAPIResponse
	if err := json.Unmarshal(body, &apiResponses); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Convert to Market slice
	markets := make([]Market, 0, len(apiResponses))
	for _, apiResp := range apiResponses {
		market := convertToMarket(apiResp)
		markets = append(markets, market)
	}

	return markets, nil
}

// convertToMarket converts a MarketAPIResponse to a simplified Market struct
func convertToMarket(apiResp MarketAPIResponse) Market {
	// Extract the first outcome and price if available
	outcome := ""
	price := ""
	
	// Parse outcomes JSON string
	var outcomes []string
	if apiResp.Outcomes != "" {
		json.Unmarshal([]byte(apiResp.Outcomes), &outcomes)
		if len(outcomes) > 0 {
			outcome = outcomes[0]
		}
	}
	
	// Parse outcome prices JSON string
	var outcomePrices []string
	if apiResp.OutcomePrices != "" {
		json.Unmarshal([]byte(apiResp.OutcomePrices), &outcomePrices)
		if len(outcomePrices) > 0 {
			price = outcomePrices[0]
		}
	}
	
	// Fallback to lastTradePrice if no outcome prices
	if price == "" && apiResp.LastTradePrice > 0 {
		price = strconv.FormatFloat(apiResp.LastTradePrice, 'f', -1, 64)
	}

	return Market{
		ID:       apiResp.ID,
		Question: apiResp.Question,
		Slug:     apiResp.Slug,
		Volume:   apiResp.Volume,
		Outcome:  outcome,
		Price:    price,
	}
}
