package models 

type Item struct {
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
}
