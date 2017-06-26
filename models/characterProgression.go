package models

type CharacterProgression struct {
    Response struct {
        Data struct {
            Progressions           []Progression `json:"progressions"`
            LevelProgression       Progression `json:"levelProgression"`
            BaseCharacterLevel     int `json:"baseCharacterLevel"`
            IsPrestigeLevel        bool `json:"isPrestigeLevel"`
            FactionProgressionHash int64 `json:"factionProgressionHash"`
            PercentToNextLevel     float64 `json:"percentToNextLevel"`
        } `json:"data"`
    } `json:"Response"`
    ErrorCode       int `json:"ErrorCode"`
    ThrottleSeconds int `json:"ThrottleSeconds"`
    ErrorStatus     string `json:"ErrorStatus"`
    Message         string `json:"Message"`
    MessageData     struct {
    } `json:"MessageData"`
}
