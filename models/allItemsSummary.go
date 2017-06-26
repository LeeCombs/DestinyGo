package models

type AllItemsSummary struct {
    Response struct {
        Data struct {
            Items      []Item `json:"items"`
            Characters []Character `json:"characters"`
        } `json:"data"`
    } `json:"Response"`
    ErrorCode       int `json:"ErrorCode"`
    ThrottleSeconds int `json:"ThrottleSeconds"`
    ErrorStatus     string `json:"ErrorStatus"`
    Message         string `json:"Message"`
    MessageData     struct {
    } `json:"MessageData"`
}
