// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[evUnused-0]
	_ = x[evLeave-1]
	_ = x[evJoinFinished-2]
	_ = x[evMove-3]
	_ = x[evTeleport-4]
	_ = x[evChangeEquipment-5]
	_ = x[evHealthUpdate-6]
	_ = x[evEnergyUpdate-7]
	_ = x[evDamageShieldUpdate-8]
	_ = x[evCraftingFocusUpdate-9]
	_ = x[evActiveSpellEffectsUpdate-10]
	_ = x[evResetCooldowns-11]
	_ = x[evAttack-12]
	_ = x[evCastStart-13]
	_ = x[evCastCancel-14]
	_ = x[evCastTimeUpdate-15]
	_ = x[evCastFinished-16]
	_ = x[evCastSpell-17]
	_ = x[evCastHit-18]
	_ = x[evCastHits-19]
	_ = x[evChannelingEnded-20]
	_ = x[evAttackBuilding-21]
	_ = x[evInventoryPutItem-22]
	_ = x[evInventoryDeleteItem-23]
	_ = x[evNewCharacter-24]
	_ = x[evNewEquipmentItem-25]
	_ = x[evNewSimpleItem-26]
	_ = x[evNewFurnitureItem-27]
	_ = x[evNewJournalItem-28]
	_ = x[evNewLaborerItem-29]
	_ = x[evNewSimpleHarvestableObject-30]
	_ = x[evNewSimpleHarvestableObjectList-31]
	_ = x[evNewHarvestableObject-32]
	_ = x[evNewSilverObject-33]
	_ = x[evNewBuilding-34]
	_ = x[evHarvestableChangeState-35]
	_ = x[evMobChangeState-36]
	_ = x[evFactionBuildingInfo-37]
	_ = x[evCraftBuildingInfo-38]
	_ = x[evRepairBuildingInfo-39]
	_ = x[evMeldBuildingInfo-40]
	_ = x[evConstructionSiteInfo-41]
	_ = x[evPlayerBuildingInfo-42]
	_ = x[evFarmBuildingInfo-43]
	_ = x[evTutorialBuildingInfo-44]
	_ = x[evLaborerObjectInfo-45]
	_ = x[evLaborerObjectJobInfo-46]
	_ = x[evMarketPlaceBuildingInfo-47]
	_ = x[evHarvestStart-48]
	_ = x[evHarvestCancel-49]
	_ = x[evHarvestFinished-50]
	_ = x[evTakeSilver-51]
	_ = x[evActionOnBuildingStart-52]
	_ = x[evActionOnBuildingCancel-53]
	_ = x[evActionOnBuildingFinished-54]
	_ = x[evItemRerollQualityStart-55]
	_ = x[evItemRerollQualityCancel-56]
	_ = x[evItemRerollQualityFinished-57]
	_ = x[evInstallResourceStart-58]
	_ = x[evInstallResourceCancel-59]
	_ = x[evInstallResourceFinished-60]
	_ = x[evCraftItemFinished-61]
	_ = x[evLogoutCancel-62]
	_ = x[evChatMessage-63]
	_ = x[evChatSay-64]
	_ = x[evChatWhisper-65]
	_ = x[evChatMuted-66]
	_ = x[evPlayEmote-67]
	_ = x[evStopEmote-68]
	_ = x[evSystemMessage-69]
	_ = x[evUtilityTextMessage-70]
	_ = x[evUpdateMoney-71]
	_ = x[evUpdateFame-72]
	_ = x[evUpdateLearningPoints-73]
	_ = x[evUpdateReSpecPoints-74]
	_ = x[evUpdateCurrency-75]
	_ = x[evUpdateFactionStanding-76]
	_ = x[evRespawn-77]
	_ = x[evServerDebugLog-78]
	_ = x[evCharacterEquipmentChanged-79]
	_ = x[evRegenerationHealthChanged-80]
	_ = x[evRegenerationEnergyChanged-81]
	_ = x[evRegenerationMountHealthChanged-82]
	_ = x[evRegenerationCraftingChanged-83]
	_ = x[evRegenerationHealthEnergyComboChanged-84]
	_ = x[evRegenerationPlayerComboChanged-85]
	_ = x[evDurabilityChanged-86]
	_ = x[evNewLoot-87]
	_ = x[evAttachItemContainer-88]
	_ = x[evDetachItemContainer-89]
	_ = x[evGuildUpdate-90]
	_ = x[evGuildPlayerUpdated-91]
	_ = x[evInvitedToGuild-92]
	_ = x[evGuildMemberWorldUpdate-93]
	_ = x[evUpdateMatchDetails-94]
	_ = x[evObjectEvent-95]
	_ = x[evNewMonolithObject-96]
	_ = x[evNewSiegeCampObject-97]
	_ = x[evNewOrbObject-98]
	_ = x[evNewCastleObject-99]
	_ = x[evNewSpellEffectArea-100]
	_ = x[evNewChainSpell-101]
	_ = x[evUpdateChainSpell-102]
	_ = x[evNewTreasureChest-103]
	_ = x[evStartMatch-104]
	_ = x[evStartTerritoryMatchInfos-105]
	_ = x[evStartArenaMatchInfos-106]
	_ = x[evEndTerritoryMatch-107]
	_ = x[evEndArenaMatch-108]
	_ = x[evMatchUpdate-109]
	_ = x[evActiveMatchUpdate-110]
	_ = x[evNewMob-111]
	_ = x[evDebugAggroInfo-112]
	_ = x[evDebugVariablesInfo-113]
	_ = x[evDebugReputationInfo-114]
	_ = x[evDebugDiminishingReturnInfo-115]
	_ = x[evDebugSmartClusterQueueInfo-116]
	_ = x[evClaimOrbStart-117]
	_ = x[evClaimOrbFinished-118]
	_ = x[evClaimOrbCancel-119]
	_ = x[evOrbUpdate-120]
	_ = x[evOrbClaimed-121]
	_ = x[evNewWarCampObject-122]
	_ = x[evNewMatchLootChestObject-123]
	_ = x[evNewArenaExit-124]
	_ = x[evGuildMemberTerritoryUpdate-125]
	_ = x[evInvitedMercenaryToMatch-126]
	_ = x[evClusterInfoUpdate-127]
	_ = x[evForcedMovement-128]
	_ = x[evForcedMovementCancel-129]
	_ = x[evCharacterStats-130]
	_ = x[evCharacterStatsKillHistory-131]
	_ = x[evCharacterStatsDeathHistory-132]
	_ = x[evGuildStats-133]
	_ = x[evKillHistoryDetails-134]
	_ = x[evFullAchievementInfo-135]
	_ = x[evFinishedAchievement-136]
	_ = x[evAchievementProgressInfo-137]
	_ = x[evFullAchievementProgressInfo-138]
	_ = x[evFullTrackedAchievementInfo-139]
	_ = x[evFullAutoLearnAchievementInfo-140]
	_ = x[evQuestGiverQuestOffered-141]
	_ = x[evQuestGiverDebugInfo-142]
	_ = x[evConsoleEvent-143]
	_ = x[evTimeSync-144]
	_ = x[evChangeAvatar-145]
	_ = x[evChangeMountSkin-146]
	_ = x[evGameEvent-147]
	_ = x[evKilledPlayer-148]
	_ = x[evDied-149]
	_ = x[evKnockedDown-150]
	_ = x[evMatchPlayerJoinedEvent-151]
	_ = x[evMatchPlayerStatsEvent-152]
	_ = x[evMatchPlayerStatsCompleteEvent-153]
	_ = x[evMatchTimeLineEventEvent-154]
	_ = x[evMatchPlayerMainGearStatsEvent-155]
	_ = x[evMatchPlayerChangedAvatarEvent-156]
	_ = x[evInvitationPlayerTrade-157]
	_ = x[evPlayerTradeStart-158]
	_ = x[evPlayerTradeCancel-159]
	_ = x[evPlayerTradeUpdate-160]
	_ = x[evPlayerTradeFinished-161]
	_ = x[evPlayerTradeAcceptChange-162]
	_ = x[evMiniMapPing-163]
	_ = x[evMarketPlaceNotification-164]
	_ = x[evDuellingChallengePlayer-165]
	_ = x[evNewDuellingPost-166]
	_ = x[evDuelStarted-167]
	_ = x[evDuelEnded-168]
	_ = x[evDuelDenied-169]
	_ = x[evDuelLeftArea-170]
	_ = x[evDuelReEnteredArea-171]
	_ = x[evNewRealEstate-172]
	_ = x[evMiniMapOwnedBuildingsPositions-173]
	_ = x[evRealEstateListUpdate-174]
	_ = x[evGuildLogoUpdate-175]
	_ = x[evGuildLogoChanged-176]
	_ = x[evPlaceableObjectPlace-177]
	_ = x[evPlaceableObjectPlaceCancel-178]
	_ = x[evFurnitureObjectBuffProviderInfo-179]
	_ = x[evFurnitureObjectCheatProviderInfo-180]
	_ = x[evFarmableObjectInfo-181]
	_ = x[evNewUnreadMails-182]
	_ = x[evGuildLogoObjectUpdate-183]
	_ = x[evStartLogout-184]
	_ = x[evNewChatChannels-185]
	_ = x[evJoinedChatChannel-186]
	_ = x[evLeftChatChannel-187]
	_ = x[evRemovedChatChannel-188]
	_ = x[evAccessStatus-189]
	_ = x[evMounted-190]
	_ = x[evMountStart-191]
	_ = x[evMountCancel-192]
	_ = x[evNewTravelpoint-193]
	_ = x[evNewIslandAccessPoint-194]
	_ = x[evNewExit-195]
	_ = x[evUpdateHome-196]
	_ = x[evUpdateChatSettings-197]
	_ = x[evResurrectionOffer-198]
	_ = x[evResurrectionReply-199]
	_ = x[evLootEquipmentChanged-200]
	_ = x[evUpdateUnlockedGuildLogos-201]
	_ = x[evUpdateUnlockedAvatars-202]
	_ = x[evUpdateUnlockedAvatarRings-203]
	_ = x[evUpdateUnlockedBuildings-204]
	_ = x[evNewIslandManagement-205]
	_ = x[evNewTeleportStone-206]
	_ = x[evCloak-207]
	_ = x[evPartyInvitation-208]
	_ = x[evPartyJoined-209]
	_ = x[evPartyDisbanded-210]
	_ = x[evPartyPlayerJoined-211]
	_ = x[evPartyChangedOrder-212]
	_ = x[evPartyPlayerLeft-213]
	_ = x[evPartyLeaderChanged-214]
	_ = x[evPartyLootSettingChangedPlayer-215]
	_ = x[evPartySilverGained-216]
	_ = x[evPartyPlayerUpdated-217]
	_ = x[evPartyInvitationPlayerBusy-218]
	_ = x[evPartyMarkedObjectsUpdated-219]
	_ = x[evPartyOnClusterPartyJoined-220]
	_ = x[evPartySetRoleFlag-221]
	_ = x[evSpellCooldownUpdate-222]
	_ = x[evNewHellgate-223]
	_ = x[evNewHellgateExit-224]
	_ = x[evNewExpeditionExit-225]
	_ = x[evNewExpeditionNarrator-226]
	_ = x[evExitEnterStart-227]
	_ = x[evExitEnterCancel-228]
	_ = x[evExitEnterFinished-229]
	_ = x[evHellClusterTimeUpdate-230]
	_ = x[evNewQuestGiverObject-231]
	_ = x[evFullQuestInfo-232]
	_ = x[evQuestProgressInfo-233]
	_ = x[evQuestGiverInfoForPlayer-234]
	_ = x[evFullExpeditionInfo-235]
	_ = x[evExpeditionQuestProgressInfo-236]
	_ = x[evInvitedToExpedition-237]
	_ = x[evExpeditionRegistrationInfo-238]
	_ = x[evEnteringExpeditionStart-239]
	_ = x[evEnteringExpeditionCancel-240]
	_ = x[evRewardGranted-241]
	_ = x[evArenaRegistrationInfo-242]
	_ = x[evEnteringArenaStart-243]
	_ = x[evEnteringArenaCancel-244]
	_ = x[evEnteringArenaLockStart-245]
	_ = x[evEnteringArenaLockCancel-246]
	_ = x[evInvitedToArenaMatch-247]
	_ = x[evPlayerCounts-248]
	_ = x[evInCombatStateUpdate-249]
	_ = x[evOtherGrabbedLoot-250]
	_ = x[evSiegeCampClaimStart-251]
	_ = x[evSiegeCampClaimCancel-252]
	_ = x[evSiegeCampClaimFinished-253]
	_ = x[evSiegeCampScheduleResult-254]
	_ = x[evTreasureChestUsingStart-255]
	_ = x[evTreasureChestUsingFinished-256]
	_ = x[evTreasureChestUsingCancel-257]
	_ = x[evTreasureChestUsingOpeningComplete-258]
	_ = x[evTreasureChestForceCloseInventory-259]
	_ = x[evPremiumChanged-260]
	_ = x[evPremiumExtended-261]
	_ = x[evPremiumLifeTimeRewardGained-262]
	_ = x[evLaborerGotUpgraded-263]
	_ = x[evJournalGotFull-264]
	_ = x[evJournalFillError-265]
	_ = x[evFriendRequest-266]
	_ = x[evFriendRequestInfos-267]
	_ = x[evFriendInfos-268]
	_ = x[evFriendRequestAnswered-269]
	_ = x[evFriendOnlineStatus-270]
	_ = x[evFriendRequestCanceled-271]
	_ = x[evFriendRemoved-272]
	_ = x[evFriendUpdated-273]
	_ = x[evPartyLootItems-274]
	_ = x[evPartyLootItemsRemoved-275]
	_ = x[evReputationUpdate-276]
	_ = x[evDefenseUnitAttackBegin-277]
	_ = x[evDefenseUnitAttackEnd-278]
	_ = x[evDefenseUnitAttackDamage-279]
	_ = x[evUnrestrictedPvpZoneUpdate-280]
	_ = x[evReputationImplicationUpdate-281]
	_ = x[evNewMountObject-282]
	_ = x[evMountHealthUpdate-283]
	_ = x[evMountCooldownUpdate-284]
	_ = x[evNewExpeditionAgent-285]
	_ = x[evNewExpeditionCheckPoint-286]
	_ = x[evExpeditionStartEvent-287]
	_ = x[evVoteEvent-288]
	_ = x[evRatingEvent-289]
	_ = x[evNewArenaAgent-290]
	_ = x[evBoostFarmable-291]
	_ = x[evUseFunction-292]
	_ = x[evNewPortalEntrance-293]
	_ = x[evNewPortalExit-294]
	_ = x[evNewRandomDungeonExit-295]
	_ = x[evWaitingQueueUpdate-296]
	_ = x[evPlayerMovementRateUpdate-297]
	_ = x[evObserveStart-298]
	_ = x[evMinimapZergs-299]
	_ = x[evPaymentTransactions-300]
	_ = x[evPerformanceStatsUpdate-301]
	_ = x[evOverloadModeUpdate-302]
	_ = x[evDebugDrawEvent-303]
	_ = x[evRecordCameraMove-304]
	_ = x[evRecordStart-305]
	_ = x[evTerritoryClaimStart-306]
	_ = x[evTerritoryClaimCancel-307]
	_ = x[evTerritoryClaimFinished-308]
	_ = x[evTerritoryScheduleResult-309]
	_ = x[evUpdateAccountState-310]
	_ = x[evStartDeterministicRoam-311]
	_ = x[evGuildFullAccessTagsUpdated-312]
	_ = x[evGuildAccessTagUpdated-313]
	_ = x[evGvgSeasonUpdate-314]
	_ = x[evGvgSeasonCheatCommand-315]
	_ = x[evSeasonPointsByKillingBooster-316]
	_ = x[evFishingStart-317]
	_ = x[evFishingCast-318]
	_ = x[evFishingCatch-319]
	_ = x[evFishingFinished-320]
	_ = x[evFishingCancel-321]
	_ = x[evNewFloatObject-322]
	_ = x[evNewFishingZoneObject-323]
	_ = x[evFishingMiniGame-324]
	_ = x[evSteamAchievementCompleted-325]
	_ = x[evUpdatePuppet-326]
	_ = x[evChangeFlaggingFinished-327]
	_ = x[evNewOutpostObject-328]
	_ = x[evOutpostUpdate-329]
	_ = x[evOutpostClaimed-330]
	_ = x[evOutpostReward-331]
	_ = x[evOverChargeEnd-332]
	_ = x[evOverChargeStatus-333]
	_ = x[evPartyFinderFullUpdate-334]
	_ = x[evPartyFinderUpdate-335]
	_ = x[evPartyFinderApplicantsUpdate-336]
	_ = x[evPartyFinderEquipmentSnapshot-337]
	_ = x[evPartyFinderJoinRequestDeclined-338]
	_ = x[evNewUnlockedPersonalSeasonRewards-339]
	_ = x[evPersonalSeasonPointsGained-340]
	_ = x[evEasyAntiCheatMessageToClient-341]
	_ = x[evMatchLootChestOpeningStart-342]
	_ = x[evMatchLootChestOpeningFinished-343]
	_ = x[evMatchLootChestOpeningCancel-344]
	_ = x[evNotifyCrystalMatchReward-345]
	_ = x[evCrystalRealmFeedback-346]
	_ = x[evNewLocationMarker-347]
	_ = x[evNewTutorialBlocker-348]
	_ = x[evNewTileSwitch-349]
	_ = x[evNewInformationProvider-350]
	_ = x[evNewDynamicGuildLogo-351]
	_ = x[evTutorialUpdate-352]
	_ = x[evTriggerHintBox-353]
	_ = x[evRandomDungeonPositionInfo-354]
	_ = x[evNewLootChest-355]
	_ = x[evUpdateLootChest-356]
	_ = x[evLootChestOpened-357]
	_ = x[evNewShrine-358]
	_ = x[evUpdateShrine-359]
	_ = x[evMutePlayerUpdate-360]
	_ = x[evShopTileUpdate-361]
	_ = x[evShopUpdate-362]
	_ = x[evEasyAntiCheatKick-363]
	_ = x[evUnlockVanityUnlock-364]
	_ = x[evCustomizationChanged-365]
	_ = x[evBaseVaultInfo-366]
	_ = x[evGuildVaultInfo-367]
	_ = x[evBankVaultInfo-368]
	_ = x[evRecoveryVaultPlayerInfo-369]
	_ = x[evRecoveryVaultGuildInfo-370]
	_ = x[evUpdateWardrobe-371]
	_ = x[evCastlePhaseChanged-372]
	_ = x[evGuildAccountLogEvent-373]
	_ = x[evNewHideoutObject-374]
	_ = x[evNewHideoutManagement-375]
	_ = x[evNewHideoutExit-376]
	_ = x[evInitHideoutAttackStart-377]
	_ = x[evInitHideoutAttackCancel-378]
	_ = x[evInitHideoutAttackFinished-379]
	_ = x[evHideoutManagementUpdate-380]
	_ = x[evIpChanged-381]
	_ = x[evSmartClusterQueueUpdateInfo-382]
	_ = x[evSmartClusterQueueActiveInfo-383]
	_ = x[evSmartClusterQueueKickWarning-384]
	_ = x[evSmartClusterQueueInvite-385]
	_ = x[evReceivedGvgSeasonPoints-386]
	_ = x[evTerritoryBonusLevelUpdate-387]
	_ = x[evOpenWorldAttackScheduleStart-388]
	_ = x[evOpenWorldAttackScheduleFinished-389]
	_ = x[evOpenWorldAttackScheduleCancel-390]
	_ = x[evOpenWorldAttackConquerStart-391]
	_ = x[evOpenWorldAttackConquerFinished-392]
	_ = x[evOpenWorldAttackConquerCancel-393]
	_ = x[evOpenWorldAttackConquerStatus-394]
	_ = x[evOpenWorldAttackStart-395]
	_ = x[evOpenWorldAttackEnd-396]
	_ = x[evNewRandomResourceBlocker-397]
}

const _EventType_name = "evUnusedevLeaveevJoinFinishedevMoveevTeleportevChangeEquipmentevHealthUpdateevEnergyUpdateevDamageShieldUpdateevCraftingFocusUpdateevActiveSpellEffectsUpdateevResetCooldownsevAttackevCastStartevCastCancelevCastTimeUpdateevCastFinishedevCastSpellevCastHitevCastHitsevChannelingEndedevAttackBuildingevInventoryPutItemevInventoryDeleteItemevNewCharacterevNewEquipmentItemevNewSimpleItemevNewFurnitureItemevNewJournalItemevNewLaborerItemevNewSimpleHarvestableObjectevNewSimpleHarvestableObjectListevNewHarvestableObjectevNewSilverObjectevNewBuildingevHarvestableChangeStateevMobChangeStateevFactionBuildingInfoevCraftBuildingInfoevRepairBuildingInfoevMeldBuildingInfoevConstructionSiteInfoevPlayerBuildingInfoevFarmBuildingInfoevTutorialBuildingInfoevLaborerObjectInfoevLaborerObjectJobInfoevMarketPlaceBuildingInfoevHarvestStartevHarvestCancelevHarvestFinishedevTakeSilverevActionOnBuildingStartevActionOnBuildingCancelevActionOnBuildingFinishedevItemRerollQualityStartevItemRerollQualityCancelevItemRerollQualityFinishedevInstallResourceStartevInstallResourceCancelevInstallResourceFinishedevCraftItemFinishedevLogoutCancelevChatMessageevChatSayevChatWhisperevChatMutedevPlayEmoteevStopEmoteevSystemMessageevUtilityTextMessageevUpdateMoneyevUpdateFameevUpdateLearningPointsevUpdateReSpecPointsevUpdateCurrencyevUpdateFactionStandingevRespawnevServerDebugLogevCharacterEquipmentChangedevRegenerationHealthChangedevRegenerationEnergyChangedevRegenerationMountHealthChangedevRegenerationCraftingChangedevRegenerationHealthEnergyComboChangedevRegenerationPlayerComboChangedevDurabilityChangedevNewLootevAttachItemContainerevDetachItemContainerevGuildUpdateevGuildPlayerUpdatedevInvitedToGuildevGuildMemberWorldUpdateevUpdateMatchDetailsevObjectEventevNewMonolithObjectevNewSiegeCampObjectevNewOrbObjectevNewCastleObjectevNewSpellEffectAreaevNewChainSpellevUpdateChainSpellevNewTreasureChestevStartMatchevStartTerritoryMatchInfosevStartArenaMatchInfosevEndTerritoryMatchevEndArenaMatchevMatchUpdateevActiveMatchUpdateevNewMobevDebugAggroInfoevDebugVariablesInfoevDebugReputationInfoevDebugDiminishingReturnInfoevDebugSmartClusterQueueInfoevClaimOrbStartevClaimOrbFinishedevClaimOrbCancelevOrbUpdateevOrbClaimedevNewWarCampObjectevNewMatchLootChestObjectevNewArenaExitevGuildMemberTerritoryUpdateevInvitedMercenaryToMatchevClusterInfoUpdateevForcedMovementevForcedMovementCancelevCharacterStatsevCharacterStatsKillHistoryevCharacterStatsDeathHistoryevGuildStatsevKillHistoryDetailsevFullAchievementInfoevFinishedAchievementevAchievementProgressInfoevFullAchievementProgressInfoevFullTrackedAchievementInfoevFullAutoLearnAchievementInfoevQuestGiverQuestOfferedevQuestGiverDebugInfoevConsoleEventevTimeSyncevChangeAvatarevChangeMountSkinevGameEventevKilledPlayerevDiedevKnockedDownevMatchPlayerJoinedEventevMatchPlayerStatsEventevMatchPlayerStatsCompleteEventevMatchTimeLineEventEventevMatchPlayerMainGearStatsEventevMatchPlayerChangedAvatarEventevInvitationPlayerTradeevPlayerTradeStartevPlayerTradeCancelevPlayerTradeUpdateevPlayerTradeFinishedevPlayerTradeAcceptChangeevMiniMapPingevMarketPlaceNotificationevDuellingChallengePlayerevNewDuellingPostevDuelStartedevDuelEndedevDuelDeniedevDuelLeftAreaevDuelReEnteredAreaevNewRealEstateevMiniMapOwnedBuildingsPositionsevRealEstateListUpdateevGuildLogoUpdateevGuildLogoChangedevPlaceableObjectPlaceevPlaceableObjectPlaceCancelevFurnitureObjectBuffProviderInfoevFurnitureObjectCheatProviderInfoevFarmableObjectInfoevNewUnreadMailsevGuildLogoObjectUpdateevStartLogoutevNewChatChannelsevJoinedChatChannelevLeftChatChannelevRemovedChatChannelevAccessStatusevMountedevMountStartevMountCancelevNewTravelpointevNewIslandAccessPointevNewExitevUpdateHomeevUpdateChatSettingsevResurrectionOfferevResurrectionReplyevLootEquipmentChangedevUpdateUnlockedGuildLogosevUpdateUnlockedAvatarsevUpdateUnlockedAvatarRingsevUpdateUnlockedBuildingsevNewIslandManagementevNewTeleportStoneevCloakevPartyInvitationevPartyJoinedevPartyDisbandedevPartyPlayerJoinedevPartyChangedOrderevPartyPlayerLeftevPartyLeaderChangedevPartyLootSettingChangedPlayerevPartySilverGainedevPartyPlayerUpdatedevPartyInvitationPlayerBusyevPartyMarkedObjectsUpdatedevPartyOnClusterPartyJoinedevPartySetRoleFlagevSpellCooldownUpdateevNewHellgateevNewHellgateExitevNewExpeditionExitevNewExpeditionNarratorevExitEnterStartevExitEnterCancelevExitEnterFinishedevHellClusterTimeUpdateevNewQuestGiverObjectevFullQuestInfoevQuestProgressInfoevQuestGiverInfoForPlayerevFullExpeditionInfoevExpeditionQuestProgressInfoevInvitedToExpeditionevExpeditionRegistrationInfoevEnteringExpeditionStartevEnteringExpeditionCancelevRewardGrantedevArenaRegistrationInfoevEnteringArenaStartevEnteringArenaCancelevEnteringArenaLockStartevEnteringArenaLockCancelevInvitedToArenaMatchevPlayerCountsevInCombatStateUpdateevOtherGrabbedLootevSiegeCampClaimStartevSiegeCampClaimCancelevSiegeCampClaimFinishedevSiegeCampScheduleResultevTreasureChestUsingStartevTreasureChestUsingFinishedevTreasureChestUsingCancelevTreasureChestUsingOpeningCompleteevTreasureChestForceCloseInventoryevPremiumChangedevPremiumExtendedevPremiumLifeTimeRewardGainedevLaborerGotUpgradedevJournalGotFullevJournalFillErrorevFriendRequestevFriendRequestInfosevFriendInfosevFriendRequestAnsweredevFriendOnlineStatusevFriendRequestCanceledevFriendRemovedevFriendUpdatedevPartyLootItemsevPartyLootItemsRemovedevReputationUpdateevDefenseUnitAttackBeginevDefenseUnitAttackEndevDefenseUnitAttackDamageevUnrestrictedPvpZoneUpdateevReputationImplicationUpdateevNewMountObjectevMountHealthUpdateevMountCooldownUpdateevNewExpeditionAgentevNewExpeditionCheckPointevExpeditionStartEventevVoteEventevRatingEventevNewArenaAgentevBoostFarmableevUseFunctionevNewPortalEntranceevNewPortalExitevNewRandomDungeonExitevWaitingQueueUpdateevPlayerMovementRateUpdateevObserveStartevMinimapZergsevPaymentTransactionsevPerformanceStatsUpdateevOverloadModeUpdateevDebugDrawEventevRecordCameraMoveevRecordStartevTerritoryClaimStartevTerritoryClaimCancelevTerritoryClaimFinishedevTerritoryScheduleResultevUpdateAccountStateevStartDeterministicRoamevGuildFullAccessTagsUpdatedevGuildAccessTagUpdatedevGvgSeasonUpdateevGvgSeasonCheatCommandevSeasonPointsByKillingBoosterevFishingStartevFishingCastevFishingCatchevFishingFinishedevFishingCancelevNewFloatObjectevNewFishingZoneObjectevFishingMiniGameevSteamAchievementCompletedevUpdatePuppetevChangeFlaggingFinishedevNewOutpostObjectevOutpostUpdateevOutpostClaimedevOutpostRewardevOverChargeEndevOverChargeStatusevPartyFinderFullUpdateevPartyFinderUpdateevPartyFinderApplicantsUpdateevPartyFinderEquipmentSnapshotevPartyFinderJoinRequestDeclinedevNewUnlockedPersonalSeasonRewardsevPersonalSeasonPointsGainedevEasyAntiCheatMessageToClientevMatchLootChestOpeningStartevMatchLootChestOpeningFinishedevMatchLootChestOpeningCancelevNotifyCrystalMatchRewardevCrystalRealmFeedbackevNewLocationMarkerevNewTutorialBlockerevNewTileSwitchevNewInformationProviderevNewDynamicGuildLogoevTutorialUpdateevTriggerHintBoxevRandomDungeonPositionInfoevNewLootChestevUpdateLootChestevLootChestOpenedevNewShrineevUpdateShrineevMutePlayerUpdateevShopTileUpdateevShopUpdateevEasyAntiCheatKickevUnlockVanityUnlockevCustomizationChangedevBaseVaultInfoevGuildVaultInfoevBankVaultInfoevRecoveryVaultPlayerInfoevRecoveryVaultGuildInfoevUpdateWardrobeevCastlePhaseChangedevGuildAccountLogEventevNewHideoutObjectevNewHideoutManagementevNewHideoutExitevInitHideoutAttackStartevInitHideoutAttackCancelevInitHideoutAttackFinishedevHideoutManagementUpdateevIpChangedevSmartClusterQueueUpdateInfoevSmartClusterQueueActiveInfoevSmartClusterQueueKickWarningevSmartClusterQueueInviteevReceivedGvgSeasonPointsevTerritoryBonusLevelUpdateevOpenWorldAttackScheduleStartevOpenWorldAttackScheduleFinishedevOpenWorldAttackScheduleCancelevOpenWorldAttackConquerStartevOpenWorldAttackConquerFinishedevOpenWorldAttackConquerCancelevOpenWorldAttackConquerStatusevOpenWorldAttackStartevOpenWorldAttackEndevNewRandomResourceBlocker"

var _EventType_index = [...]uint16{0, 8, 15, 29, 35, 45, 62, 76, 90, 110, 131, 157, 173, 181, 192, 204, 220, 234, 245, 254, 264, 281, 297, 315, 336, 350, 368, 383, 401, 417, 433, 461, 493, 515, 532, 545, 569, 585, 606, 625, 645, 663, 685, 705, 723, 745, 764, 786, 811, 825, 840, 857, 869, 892, 916, 942, 966, 991, 1018, 1040, 1063, 1088, 1107, 1121, 1134, 1143, 1156, 1167, 1178, 1189, 1204, 1224, 1237, 1249, 1271, 1291, 1307, 1330, 1339, 1355, 1382, 1409, 1436, 1468, 1497, 1535, 1567, 1586, 1595, 1616, 1637, 1650, 1670, 1686, 1710, 1730, 1743, 1762, 1782, 1796, 1813, 1833, 1848, 1866, 1884, 1896, 1922, 1944, 1963, 1978, 1991, 2010, 2018, 2034, 2054, 2075, 2103, 2131, 2146, 2164, 2180, 2191, 2203, 2221, 2246, 2260, 2288, 2313, 2332, 2348, 2370, 2386, 2413, 2441, 2453, 2473, 2494, 2515, 2540, 2569, 2597, 2627, 2651, 2672, 2686, 2696, 2710, 2727, 2738, 2752, 2758, 2771, 2795, 2818, 2849, 2874, 2905, 2936, 2959, 2977, 2996, 3015, 3036, 3061, 3074, 3099, 3124, 3141, 3154, 3165, 3177, 3191, 3210, 3225, 3257, 3279, 3296, 3314, 3336, 3364, 3397, 3431, 3451, 3467, 3490, 3503, 3520, 3539, 3556, 3576, 3590, 3599, 3611, 3624, 3640, 3662, 3671, 3683, 3703, 3722, 3741, 3763, 3789, 3812, 3839, 3864, 3885, 3903, 3910, 3927, 3940, 3956, 3975, 3994, 4011, 4031, 4062, 4081, 4101, 4128, 4155, 4182, 4200, 4221, 4234, 4251, 4270, 4293, 4309, 4326, 4345, 4368, 4389, 4404, 4423, 4448, 4468, 4497, 4518, 4546, 4571, 4597, 4612, 4635, 4655, 4676, 4700, 4725, 4746, 4760, 4781, 4799, 4820, 4842, 4866, 4891, 4916, 4944, 4970, 5005, 5039, 5055, 5072, 5101, 5121, 5137, 5155, 5170, 5190, 5203, 5226, 5246, 5269, 5284, 5299, 5315, 5338, 5356, 5380, 5402, 5427, 5454, 5483, 5499, 5518, 5539, 5559, 5584, 5606, 5617, 5630, 5645, 5660, 5673, 5692, 5707, 5729, 5749, 5775, 5789, 5803, 5824, 5848, 5868, 5884, 5902, 5915, 5936, 5958, 5982, 6007, 6027, 6051, 6079, 6102, 6119, 6142, 6172, 6186, 6199, 6213, 6230, 6245, 6261, 6283, 6300, 6327, 6341, 6365, 6383, 6398, 6414, 6429, 6444, 6462, 6485, 6504, 6533, 6563, 6595, 6629, 6657, 6687, 6715, 6746, 6775, 6801, 6823, 6842, 6862, 6877, 6901, 6922, 6938, 6954, 6981, 6995, 7012, 7029, 7040, 7054, 7072, 7088, 7100, 7119, 7139, 7161, 7176, 7192, 7207, 7232, 7256, 7272, 7292, 7314, 7332, 7354, 7370, 7394, 7419, 7446, 7471, 7482, 7511, 7540, 7570, 7595, 7620, 7647, 7677, 7710, 7741, 7770, 7802, 7832, 7862, 7884, 7904, 7930}

func (i EventType) String() string {
	if i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}
