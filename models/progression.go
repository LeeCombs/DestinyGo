package models

type Progression struct {
    DailyProgress       int `json:"dailyProgress"`
    WeeklyProgress      int `json:"weeklyProgress"`
    CurrentProgress     int `json:"currentProgress"`
    Level               int `json:"level"`
    Step                int `json:"step"`
    ProgressToNextLevel int `json:"progressToNextLevel"`
    NextLevelAt         int `json:"nextLevelAt"`
    ProgressionHash     int64 `json:"progressionHash"`
}
