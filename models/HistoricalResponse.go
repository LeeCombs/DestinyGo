package models

type HistoricalResponse struct {
	ErrorCode   float64 `json:"ErrorCode"`
	ErrorStatus string  `json:"ErrorStatus"`
	Message     string  `json:"Message"`
	MessageData struct {
	} `json:"MessageData"`
	Response struct {
		AllArena struct {
			AllTime struct {
				AbilityKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"abilityKills"`
				ActivitiesCleared struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"activitiesCleared"`
				ActivitiesEntered struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"activitiesEntered"`
				AllParticipantsCount struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"allParticipantsCount"`
				AllParticipantsScore struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"allParticipantsScore"`
				AllParticipantsTimePlayed struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"allParticipantsTimePlayed"`
				Assists struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"assists"`
				AverageDeathDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageDeathDistance"`
				AverageKillDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageKillDistance"`
				AverageLifespan struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageLifespan"`
				AverageScorePerKill struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageScorePerKill"`
				AverageScorePerLife struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageScorePerLife"`
				BestSingleGameKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"bestSingleGameKills"`
				BestSingleGameScore struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"bestSingleGameScore"`
				CourtOfOryxAttempts struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"courtOfOryxAttempts"`
				CourtOfOryxCompletions struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"courtOfOryxCompletions"`
				CourtOfOryxWinsTier1 struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"courtOfOryxWinsTier1"`
				CourtOfOryxWinsTier2 struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"courtOfOryxWinsTier2"`
				CourtOfOryxWinsTier3 struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"courtOfOryxWinsTier3"`
				Deaths struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"deaths"`
				FastestCompletion struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"fastestCompletion"`
				HighestCharacterLevel struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"highestCharacterLevel"`
				HighestLightLevel struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"highestLightLevel"`
				Kills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"kills"`
				KillsDeathsAssists struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"killsDeathsAssists"`
				KillsDeathsRatio struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"killsDeathsRatio"`
				LongestKillDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"longestKillDistance"`
				LongestKillSpree struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"longestKillSpree"`
				LongestSingleLife struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"longestSingleLife"`
				MostPrecisionKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"mostPrecisionKills"`
				ObjectivesCompleted struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"objectivesCompleted"`
				OrbsDropped struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"orbsDropped"`
				OrbsGathered struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"orbsGathered"`
				PrecisionKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"precisionKills"`
				PublicEventsCompleted struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"publicEventsCompleted"`
				PublicEventsJoined struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"publicEventsJoined"`
				RemainingTimeAfterQuitSeconds struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"remainingTimeAfterQuitSeconds"`
				ResurrectionsPerformed struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"resurrectionsPerformed"`
				ResurrectionsReceived struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"resurrectionsReceived"`
				Score struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"score"`
				SecondsPlayed struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"secondsPlayed"`
				Suicides struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"suicides"`
				TeamScore struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"teamScore"`
				TotalActivityDurationSeconds struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"totalActivityDurationSeconds"`
				TotalDeathDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"totalDeathDistance"`
				TotalKillDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"totalKillDistance"`
				WeaponBestType struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"weaponBestType"`
				WeaponKillsAutoRifle struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsAutoRifle"`
				WeaponKillsFusionRifle struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsFusionRifle"`
				WeaponKillsGrenade struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsGrenade"`
				WeaponKillsHandCannon struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsHandCannon"`
				WeaponKillsMachinegun struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsMachinegun"`
				WeaponKillsMelee struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsMelee"`
				WeaponKillsPulseRifle struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsPulseRifle"`
				WeaponKillsRelic struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsRelic"`
				WeaponKillsRocketLauncher struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsRocketLauncher"`
				WeaponKillsScoutRifle struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsScoutRifle"`
				WeaponKillsShotgun struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsShotgun"`
				WeaponKillsSideArm struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSideArm"`
				WeaponKillsSniper struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSniper"`
				WeaponKillsSubmachinegun struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSubmachinegun"`
				WeaponKillsSuper struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSuper"`
				WeaponKillsSword struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSword"`
			} `json:"allTime"`
		} `json:"allArena"`
		AllPvP struct {
			AllTime struct {
				AbilityKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"abilityKills"`
				ActivitiesEntered struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"activitiesEntered"`
				ActivitiesWon struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"activitiesWon"`
				AllParticipantsCount struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"allParticipantsCount"`
				AllParticipantsScore struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"allParticipantsScore"`
				AllParticipantsTimePlayed struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"allParticipantsTimePlayed"`
				Assists struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"assists"`
				AverageDeathDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageDeathDistance"`
				AverageKillDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageKillDistance"`
				AverageLifespan struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageLifespan"`
				AverageScorePerKill struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageScorePerKill"`
				AverageScorePerLife struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"averageScorePerLife"`
				BestSingleGameKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"bestSingleGameKills"`
				BestSingleGameScore struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"bestSingleGameScore"`
				CloseCalls struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"closeCalls"`
				CombatRating struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"combatRating"`
				Deaths struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"deaths"`
				DefensiveKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"defensiveKills"`
				DominationKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"dominationKills"`
				HighestCharacterLevel struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"highestCharacterLevel"`
				HighestLightLevel struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"highestLightLevel"`
				Kills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"kills"`
				KillsDeathsAssists struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"killsDeathsAssists"`
				KillsDeathsRatio struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"killsDeathsRatio"`
				LongestKillDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"longestKillDistance"`
				LongestKillSpree struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"longestKillSpree"`
				LongestSingleLife struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"longestSingleLife"`
				MostPrecisionKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"mostPrecisionKills"`
				ObjectivesCompleted struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"objectivesCompleted"`
				OffensiveKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"offensiveKills"`
				OrbsDropped struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"orbsDropped"`
				OrbsGathered struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"orbsGathered"`
				PrecisionKills struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"precisionKills"`
				RelicsCaptured struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"relicsCaptured"`
				RemainingTimeAfterQuitSeconds struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"remainingTimeAfterQuitSeconds"`
				ResurrectionsPerformed struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"resurrectionsPerformed"`
				ResurrectionsReceived struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"resurrectionsReceived"`
				Score struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"score"`
				SecondsPlayed struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"secondsPlayed"`
				Suicides struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"suicides"`
				TeamScore struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"teamScore"`
				TotalActivityDurationSeconds struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"totalActivityDurationSeconds"`
				TotalDeathDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"totalDeathDistance"`
				TotalKillDistance struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"totalKillDistance"`
				WeaponBestType struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"weaponBestType"`
				WeaponKillsAutoRifle struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsAutoRifle"`
				WeaponKillsFusionRifle struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsFusionRifle"`
				WeaponKillsGrenade struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsGrenade"`
				WeaponKillsHandCannon struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsHandCannon"`
				WeaponKillsMachinegun struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsMachinegun"`
				WeaponKillsMelee struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsMelee"`
				WeaponKillsPulseRifle struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsPulseRifle"`
				WeaponKillsRelic struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsRelic"`
				WeaponKillsRocketLauncher struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsRocketLauncher"`
				WeaponKillsScoutRifle struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsScoutRifle"`
				WeaponKillsShotgun struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsShotgun"`
				WeaponKillsSideArm struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSideArm"`
				WeaponKillsSniper struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSniper"`
				WeaponKillsSubmachinegun struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSubmachinegun"`
				WeaponKillsSuper struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSuper"`
				WeaponKillsSword struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"weaponKillsSword"`
				WinLossRatio struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					StatID string `json:"statId"`
				} `json:"winLossRatio"`
				ZonesCaptured struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"zonesCaptured"`
				ZonesNeutralized struct {
					Basic struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"basic"`
					Pga struct {
						DisplayValue string  `json:"displayValue"`
						Value        float64 `json:"value"`
					} `json:"pga"`
					StatID string `json:"statId"`
				} `json:"zonesNeutralized"`
			} `json:"allTime"`
		} `json:"allPvP"`
	} `json:"Response"`
}
