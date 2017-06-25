package models

type CharacterProgression struct {
    Response struct {
        Data struct {
            Progressions []struct {
                DailyProgress       int `json:"dailyProgress"`
                WeeklyProgress      int `json:"weeklyProgress"`
                CurrentProgress     int `json:"currentProgress"`
                Level               int `json:"level"`
                Step                int `json:"step"`
                ProgressToNextLevel int `json:"progressToNextLevel"`
                NextLevelAt         int `json:"nextLevelAt"`
                ProgressionHash     int64 `json:"progressionHash"`
            } `json:"progressions"`
            LevelProgression struct {
                DailyProgress       int `json:"dailyProgress"`
                WeeklyProgress      int `json:"weeklyProgress"`
                CurrentProgress     int `json:"currentProgress"`
                Level               int `json:"level"`
                Step                int `json:"step"`
                ProgressToNextLevel int `json:"progressToNextLevel"`
                NextLevelAt         int `json:"nextLevelAt"`
                ProgressionHash     int64 `json:"progressionHash"`
            } `json:"levelProgression"`
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