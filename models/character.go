package models

import (
    "time"
)

type Character struct {
    CharacterBase struct {
        MembershipID             string `json:"membershipId"`
        MembershipType           int `json:"membershipType"`
        CharacterID              string `json:"characterId"`
        DateLastPlayed           time.Time `json:"dateLastPlayed"`
        MinutesPlayedThisSession string `json:"minutesPlayedThisSession"`
        MinutesPlayedTotal       string `json:"minutesPlayedTotal"`
        PowerLevel               int `json:"powerLevel"`
        RaceHash                 int64 `json:"raceHash"`
        GenderHash               int64 `json:"genderHash"`
        ClassHash                int64 `json:"classHash"`
        CurrentActivityHash      int64 `json:"currentActivityHash"`
        LastCompletedStoryHash   int64 `json:"lastCompletedStoryHash"`
        Stats struct {
            STATDEFENSE    Stat `json:"STAT_DEFENSE"`
            STATINTELLECT  Stat `json:"STAT_INTELLECT"`
            STATDISCIPLINE Stat `json:"STAT_DISCIPLINE"`
            STATSTRENGTH   Stat `json:"STAT_STRENGTH"`
            STATLIGHT      Stat `json:"STAT_LIGHT"`
            STATARMOR      Stat `json:"STAT_ARMOR"`
            STATAGILITY    Stat `json:"STAT_AGILITY"`
            STATRECOVERY   Stat `json:"STAT_RECOVERY"`
            STATOPTICS     Stat `json:"STAT_OPTICS"`
        } `json:"stats"`
        Customization struct {
            Personality  int64 `json:"personality"`
            Face         int64 `json:"face"`
            SkinColor    int64 `json:"skinColor"`
            LipColor     int64 `json:"lipColor"`
            EyeColor     int64 `json:"eyeColor"`
            HairColor    int64 `json:"hairColor"`
            FeatureColor int64 `json:"featureColor"`
            DecalColor   int64 `json:"decalColor"`
            WearHelmet   bool `json:"wearHelmet"`
            HairIndex    int `json:"hairIndex"`
            FeatureIndex int `json:"featureIndex"`
            DecalIndex   int `json:"decalIndex"`
        } `json:"customization"`
        GrimoireScore int `json:"grimoireScore"`
        PeerView struct {
            Equipment []Equipment`json:"equipment"`
        } `json:"peerView"`
        GenderType         int `json:"genderType"`
        ClassType          int `json:"classType"`
        BuildStatGroupHash int64 `json:"buildStatGroupHash"`
    } `json:"characterBase"`
    LevelProgression   Progression `json:"levelProgression"`
    EmblemPath         string `json:"emblemPath"`
    BackgroundPath     string `json:"backgroundPath"`
    EmblemHash         int64 `json:"emblemHash"`
    CharacterLevel     int `json:"characterLevel"`
    BaseCharacterLevel int `json:"baseCharacterLevel"`
    IsPrestigeLevel    bool `json:"isPrestigeLevel"`
    PercentToNextLevel float64 `json:"percentToNextLevel"`
}
