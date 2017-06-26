package models

type Equipment struct {
    ItemHash int64 `json:"itemHash"`
    Dyes     []struct {
        ChannelHash int64 `json:"channelHash"`
        DyeHash     int64 `json:"dyeHash"`
    } `json:"dyes"`
}
