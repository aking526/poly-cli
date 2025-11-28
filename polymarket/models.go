package polymarket

// MarketAPIResponse represents the structure of a full market response from the Polymarket API.
type MarketAPIResponse struct {
	ID                         string         `json:"id"`
	Question                   string         `json:"question"`
	ConditionID                string         `json:"conditionId"`
	Slug                       string         `json:"slug"`
	EndDate                    string         `json:"endDate"`
	StartDate                  string         `json:"startDate"`
	Image                      string         `json:"image"`
	Icon                       string         `json:"icon"`
	Description                string         `json:"description"`
	Outcomes                   string         `json:"outcomes"` // array JSON in string format
	OutcomePrices              string         `json:"outcomePrices"` // array JSON in string format
	Volume                     string         `json:"volume"`
	Active                     bool           `json:"active"`
	Closed                     bool           `json:"closed"`
	MarketMakerAddress         string         `json:"marketMakerAddress"`
	CreatedAt                  string         `json:"createdAt"`
	UpdatedAt                  string         `json:"updatedAt"`
	ClosedTime                 string         `json:"closedTime"`
	New                        bool           `json:"new"`
	Featured                   bool           `json:"featured"`
	SubmittedBy                string         `json:"submitted_by"`
	Archived                   bool           `json:"archived"`
	ResolvedBy                 string         `json:"resolvedBy"`
	Restricted                 bool           `json:"restricted"`
	GroupItemTitle             string         `json:"groupItemTitle"`
	GroupItemThreshold         string         `json:"groupItemThreshold"`
	QuestionID                 string         `json:"questionID"`
	UmaEndDate                 string         `json:"umaEndDate"`
	EnableOrderBook            bool           `json:"enableOrderBook"`
	OrderPriceMinTickSize      float64        `json:"orderPriceMinTickSize"`
	OrderMinSize               float64        `json:"orderMinSize"`
	UmaResolutionStatus        string         `json:"umaResolutionStatus"`
	VolumeNum                  float64        `json:"volumeNum"`
	EndDateIso                 string         `json:"endDateIso"`
	StartDateIso               string         `json:"startDateIso"`
	HasReviewedDates           bool           `json:"hasReviewedDates"`
	Volume1wk                  float64        `json:"volume1wk"`
	Volume1mo                  float64        `json:"volume1mo"`
	Volume1yr                  float64        `json:"volume1yr"`
	ClobTokenIds               string         `json:"clobTokenIds"` // array JSON in string format
	UmaBond                    string         `json:"umaBond"`
	UmaReward                  string         `json:"umaReward"`
	Volume1wkClob              float64        `json:"volume1wkClob"`
	Volume1moClob              float64        `json:"volume1moClob"`
	Volume1yrClob              float64        `json:"volume1yrClob"`
	VolumeClob                 float64        `json:"volumeClob"`
	CustomLiveness             int            `json:"customLiveness"`
	AcceptingOrders            bool           `json:"acceptingOrders"`
	NegRisk                    bool           `json:"negRisk"`
	NegRiskRequestID           string         `json:"negRiskRequestID"`
	Events                     []MarketAPIResponseEvent        `json:"events"`
	Ready                      bool           `json:"ready"`
	Funded                     bool           `json:"funded"`
	AcceptingOrdersTimestamp   string         `json:"acceptingOrdersTimestamp"`
	Cyom                       bool           `json:"cyom"`
	PagerDutyNotificationEnabled bool         `json:"pagerDutyNotificationEnabled"`
	Approved                   bool           `json:"approved"`
	RewardsMinSize             float64        `json:"rewardsMinSize"`
	RewardsMaxSpread           float64        `json:"rewardsMaxSpread"`
	Spread                     float64        `json:"spread"`
	AutomaticallyResolved      bool           `json:"automaticallyResolved"`
	LastTradePrice             float64        `json:"lastTradePrice"`
	BestAsk                    float64        `json:"bestAsk"`
	AutomaticallyActive        bool           `json:"automaticallyActive"`
	ClearBookOnStart           bool           `json:"clearBookOnStart"`
	ManualActivation           bool           `json:"manualActivation"`
	NegRiskOther               bool           `json:"negRiskOther"`
	UmaResolutionStatuses      string         `json:"umaResolutionStatuses"`
	PendingDeployment          bool           `json:"pendingDeployment"`
	Deploying                  bool           `json:"deploying"`
	DeployingTimestamp         string         `json:"deployingTimestamp"`
	RfqEnabled                 bool           `json:"rfqEnabled"`
	HoldingRewardsEnabled      bool           `json:"holdingRewardsEnabled"`
	FeesEnabled                bool           `json:"feesEnabled"`
}

// Event and Series structs for nested data
type MarketAPIResponseEvent struct {
	ID                string   `json:"id"`
	Ticker            string   `json:"ticker"`
	Slug              string   `json:"slug"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	StartDate         string   `json:"startDate"`
	CreationDate      string   `json:"creationDate"`
	EndDate           string   `json:"endDate"`
	Image             string   `json:"image"`
	Icon              string   `json:"icon"`
	Active            bool     `json:"active"`
	Closed            bool     `json:"closed"`
	Archived          bool     `json:"archived"`
	New               bool     `json:"new"`
	Featured          bool     `json:"featured"`
	Restricted        bool     `json:"restricted"`
	Volume            float64  `json:"volume"`
	OpenInterest      float64  `json:"openInterest"`
	CreatedAt         string   `json:"createdAt"`
	UpdatedAt         string   `json:"updatedAt"`
	Volume1wk         float64  `json:"volume1wk"`
	Volume1mo         float64  `json:"volume1mo"`
	Volume1yr         float64  `json:"volume1yr"`
	EnableOrderBook   bool     `json:"enableOrderBook"`
	NegRisk           bool     `json:"negRisk"`
	CommentCount      int      `json:"commentCount"`
	Series            []MarketAPIResponseSeries `json:"series"`
	Cyom              bool     `json:"cyom"`
	ClosedTime        string   `json:"closedTime"`
	ShowAllOutcomes   bool     `json:"showAllOutcomes"`
	ShowMarketImages  bool     `json:"showMarketImages"`
	AutomaticallyResolved bool `json:"automaticallyResolved"`
	EnableNegRisk     bool     `json:"enableNegRisk"`
	AutomaticallyActive bool   `json:"automaticallyActive"`
	SeriesSlug        string   `json:"seriesSlug"`
	NegRiskAugmented  bool     `json:"negRiskAugmented"`
	PendingDeployment bool     `json:"pendingDeployment"`
	Deploying         bool     `json:"deploying"`
}

type MarketAPIResponseSeries struct {
	ID          string `json:"id"`
	Ticker      string `json:"ticker"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	SeriesType  string `json:"seriesType"`
	Recurrence  string `json:"recurrence"`
	Image       string `json:"image"`
	Icon        string `json:"icon"`
	Active      bool   `json:"active"`
	Closed      bool   `json:"closed"`
	Archived    bool   `json:"archived"`
	Featured    bool   `json:"featured"`
	Restricted  bool   `json:"restricted"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	CommentCount int   `json:"commentCount"`
}

