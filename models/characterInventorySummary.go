package models

type CharacterInventorySummary struct {
    Response struct {
        Data struct {
            Items []struct {
                ItemHash       int64 `json:"itemHash"`
                ItemID         string `json:"itemId"`
                Quantity       int `json:"quantity"`
                DamageType     int `json:"damageType"`
                DamageTypeHash int64 `json:"damageTypeHash"`
                IsGridComplete bool `json:"isGridComplete"`
                TransferStatus int `json:"transferStatus"`
                State          int `json:"state"`
                CharacterIndex int `json:"characterIndex"`
                BucketHash     int64 `json:"bucketHash"`
                PrimaryStat    struct {
                    StatHash     int64 `json:"statHash"`
                    Value        int `json:"value"`
                    MaximumValue int `json:"maximumValue"`
                } `json:"primaryStat,omitempty"`
            } `json:"items"`
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
