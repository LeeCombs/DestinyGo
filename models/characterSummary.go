package models

type CharacterSummary struct {
    Response struct {
        Data Character `json:"data"`
    } `json:"Response"`
    ErrorCode int `json:"ErrorCode"`
    ThrottleSeconds int `json:"ThrottleSeconds"`
    ErrorStatus string `json:"ErrorStatus"`
    Message string `json:"Message"`
    MessageData struct {
    } `json:"MessageData"`
}
