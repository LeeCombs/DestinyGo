package constants

type DestinyClass int
const (
    ClassTitan DestinyClass = iota
    ClassHunter
    ClassWarlock
    ClassUnknown
)

type DefineType int
const (
    DefTypeNone DefineType = iota
    _ // Intentional blank
    DefTypeActivity
    DefTypeClass
    DefTypeGender
    DefTypeInventoryBucket
    DefTypeInventoryItem
    DefTypeProgression
    DefTypeRace
    DefTypeStat
    DefTypeTalentGrid
    DefTypeStatGroup
    DefTypeUnlockFlag
    DefTypeVendor
    DefTypeDestination
    DefTypePlace
    DefTypeDirectorBook
    DefTypeMaterialRequirement
    DefTypeSandboxPerk
    DefTypeArtDye
    DefTypeArtDyeChannel
    DefTypeActivityBundle
    DefTypeGearAsset
    DefTypeGrimoireCard
)
