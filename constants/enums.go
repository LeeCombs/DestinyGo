package constants



type DefinitionType int
const (
    DefinitionTypeNone DefinitionType = iota
    _ // Intentional blank
    DefinitionTypeActivity
    DefinitionTypeClass
    DefinitionTypeGender
    DefinitionTypeInventoryBucket
    DefinitionTypeInventoryItem
    DefinitionTypeProgression
    DefinitionTypeRace
    DefinitionTypeStat
    DefinitionTypeTalentGrid
    DefinitionTypeStatGroup
    DefinitionTypeUnlockFlag
    DefinitionTypeVendor
    DefinitionTypeDestination
    DefinitionTypePlace
    DefinitionTypeDirectorBook
    DefinitionTypeMaterialRequirement
    DefinitionTypeSandboxPerk
    DefinitionTypeArtDye
    DefinitionTypeArtDyeChannel
    DefinitionTypeActivityBundle
    DefinitionTypeGearAsset
    DefinitionTypeGrimoireCard
)

type ItemType int
const (
    ItemTypeNone ItemType = iota
    ItemTypeCurrency
    ItemTypeArmor
    ItemTypeWeapon
    ItemTypeBounty
    ItemTypeCompletedBounty
    ItemTypeBountyReward
    ItemTypeMessage
    ItemTypeEngram
    ItemTypeConsumable
    ItemTypeExchangeMaterial
    ItemTypeMissionReward
    ItemTypeQuestStep
    ItemTypeQuestStepComplete
    ItemTypeEmblem
    ItemTypeQuest
)

type ItemSubType int
const (
    ItemSubTypeNone ItemSubType = iota
    ItemSubTypeCrucible
    ItemSubTypeVanguard
    ItemSubTypeIronBanner
    ItemSubTypeQueen
    ItemSubTypeExotic
    ItemSubTypeAutoRifle
    ItemSubTypeShotgun
    ItemSubTypeMachinegun
    ItemSubTypeHandCannon
    ItemSubTypeRocketLauncher
    ItemSubTypeFusionRifle
    ItemSubTypeSniperRifle
    ItemSubTypePulseRifle
    ItemSubTypeScoutRifle
    ItemSubTypeCamera
    ItemSubTypeCrm
    ItemSubTypeSidearm
    ItemSubTypeSword
    ItemSubTypeMask
)

type DamageType int
const (
    DamageTypeNone DamageType = iota
    DamageTypeKinetic
    DamageTypeArc
    DamageTypeThermal
    DamageTypeVoid
    DamageTypeRaid
)

type CardRarity int
const (
    CardRarityNone CardRarity = iota
    CardRarityCommon
    CardRaritySuperior
    CardRarityExotic
)

type Race int
const (
    RaceHuman Race = iota
    RaceAwoken
    RaceExo
    RaceUnknown
)

type Gender int
const (
    GenderMale Gender = iota
    GenderFemale
    GenderUnknown
)

type Class int
const (
    ClassTitan Class = iota
    ClassHunter
    ClassWarlock
    ClassUnknown
)

type StatCategory int
const (
    StatCategoryNone StatCategory = iota
    StatCategoryKills
    StatCategoryAssists
    StatCategoryDeaths
    StatCategoryCriticals
    StatCategoryKDa
    StatCategoryKD
    StatCategoryScore
    StatCategoryEntered
    StatCategoryTimePlayed
    StatCategoryMedalWins
    StatCategoryMedalGame
    StatCategoryMedalSpecialKills
    StatCategoryMedalSprees
    StatCategoryMedalMultiKills
    StatCategoryMedalAbilities
)