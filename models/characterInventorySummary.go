package models

type CharacterInventorySummary struct {
    Response struct {
        Data struct {
            Items        []Item `json:"items"`
            Currencies   []struct {
                ItemHash int64 `json:"itemHash"`
                Value    int `json:"value"`
            } `json:"currencies"`
        } `json:"data"`
    } `json:"Response"`
    ErrorCode       int `json:"ErrorCode"`
    ThrottleSeconds int `json:"ThrottleSeconds"`
    ErrorStatus     string `json:"ErrorStatus"`
    Message         string `json:"Message"`
    MessageData     struct {
    } `json:"MessageData"`
}
