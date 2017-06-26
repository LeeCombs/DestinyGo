package models

type Stat struct {
    StatHash     int64 `json:"statHash"`
    Value        int `json:"value"`
    MaximumValue int `json:"maximumValue"`
}
