package models

type AccountSummaryResponse struct {
	Response struct {
		Data struct {
			MembershipID string `json:"membershipId"`
			MembershipType int `json:"membershipType"`
			Characters []Character `json:"characters"`
			Inventory struct {
				Items []interface{} `json:"items"`
				Currencies []struct {
					ItemHash int64 `json:"itemHash"`
					Value int `json:"value"`
				} `json:"currencies"`
			} `json:"inventory"`
			GrimoireScore int `json:"grimoireScore"`
			Versions int `json:"versions"`
		} `json:"data"`
	} `json:"Response"`
	ErrorCode int `json:"ErrorCode"`
	ThrottleSeconds int `json:"ThrottleSeconds"`
	ErrorStatus string `json:"ErrorStatus"`
	Message string `json:"Message"`
	MessageData struct {
	} `json:"MessageData"`
}