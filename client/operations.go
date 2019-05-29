package client

type operation interface {
	Process(state *albionState)
}

//go:generate stringer -type=OperationType
type OperationType uint16

const (
	opUnused                                  = 0
	opPing                                    = 1
	opJoin                                    = 2
	opCreateAccount                           = 3
	opLogin                                   = 4
	opSendCrashLog                            = 5
	opCreateCharacter                         = 6
	opDeleteCharacter                         = 7
	opSelectCharacter                         = 8
	opRedeemKeycode                           = 9
	opGetGameServerByCluster                  = 10
	opGetActiveSubscription                   = 11
	opGetShopPurchaseUrl                      = 12
	opGetBuyTrialDetails                      = 13
	opGetReferralSeasonDetails                = 14
	opGetReferralLink                         = 15
	opGetAvailableTrialKeys                   = 16
	opGetShopTilesForCategory                 = 17
	opMove                                    = 18
	opAttackStart                             = 19
	opCastStart                               = 20
	opCastCancel                              = 21
	opTerminateToggleSpell                    = 22
	opChannelingCancel                        = 23
	opAttackBuildingStart                     = 24
	opInventoryDestroyItem                    = 25
	opInventoryMoveItem                       = 26
	opInventorySplitStack                     = 27
	opGetClusterData                          = 28
	opChangeCluster                           = 29
	opConsoleCommand                          = 30
	opChatMessage                             = 31
	opReportClientError                       = 32
	opRegisterToObject                        = 33
	opUnRegisterFromObject                    = 34
	opCraftBuildingChangeSettings             = 35
	opCraftBuildingTakeMoney                  = 36
	opRepairBuildingChangeSettings            = 37
	opRepairBuildingTakeMoney                 = 38
	opActionBuildingChangeSettings            = 39
	opHarvestStart                            = 40
	opHarvestCancel                           = 41
	opTakeSilver                              = 42
	opActionOnBuildingStart                   = 43
	opActionOnBuildingCancel                  = 44
	opItemRerollQualityStart                  = 45
	opItemRerollQualityCancel                 = 46
	opInstallResourceStart                    = 47
	opInstallResourceCancel                   = 48
	opInstallSilver                           = 49
	opBuildingFillNutrition                   = 50
	opBuildingChangeRenovationState           = 51
	opBuildingBuySkin                         = 52
	opBuildingClaim                           = 53
	opBuildingGiveup                          = 54
	opBuildingNutritionSilverStorageDeposit   = 55
	opBuildingNutritionSilverStorageWithdraw  = 56
	opBuildingNutritionSilverRewardSet        = 57
	opConstructionSiteCreate                  = 58
	opPlaceableItemPlace                      = 59
	opPlaceableItemPlaceCancel                = 60
	opPlaceableObjectPickup                   = 61
	opFurnitureObjectUse                      = 62
	opFarmableHarvest                         = 63
	opFarmableFinishGrownItem                 = 64
	opFarmableDestroy                         = 65
	opFarmableGetProduct                      = 66
	opFarmableFill                            = 67
	opLaborerObjectPlace                      = 68
	opLaborerObjectPlaceCancel                = 69
	opTearDownConstructionSite                = 70
	opCastleGateUse                           = 71
	opAuctionCreateOffer                      = 72
	opAuctionCreateRequest                    = 73
	opAuctionGetOffers                        = 74
	opAuctionGetRequests                      = 75
	opAuctionBuyOffer                         = 76
	opAuctionAbortAuction                     = 77
	opAuctionModifyAuction                    = 78
	opAuctionAbortOffer                       = 79
	opAuctionAbortRequest                     = 80
	opAuctionSellRequest                      = 81
	opAuctionGetFinishedAuctions              = 82
	opAuctionFetchAuction                     = 83
	opAuctionGetMyOpenOffers                  = 84
	opAuctionGetMyOpenRequests                = 85
	opAuctionGetMyOpenAuctions                = 86
	opAuctionGetItemsAverage                  = 87
	opAuctionGetItemAverageStats              = 88
	opAuctionGetItemAverageValue              = 89
	opContainerOpen                           = 90
	opContainerClose                          = 91
	opContainerManageSubContainer             = 92
	opRespawn                                 = 93
	opSuicide                                 = 94
	opJoinGuild                               = 95
	opLeaveGuild                              = 96
	opCreateGuild                             = 97
	opInviteToGuild                           = 98
	opDeclineGuildInvitation                  = 99
	opKickFromGuild                           = 100
	opDuellingChallengePlayer                 = 101
	opDuellingAcceptChallenge                 = 102
	opDuellingDenyChallenge                   = 103
	opChangeClusterTax                        = 104
	opClaimTerritory                          = 105
	opGiveUpTerritory                         = 106
	opChangeTerritoryAccessRights             = 107
	opGetMonolithInfo                         = 108
	opGetClaimInfo                            = 109
	opGetAttackInfo                           = 110
	opGetTerritorySeasonPoints                = 111
	opGetAttackSchedule                       = 112
	opScheduleAttack                          = 113
	opGetMatches                              = 114
	opGetMatchDetails                         = 115
	opJoinMatch                               = 116
	opLeaveMatch                              = 117
	opChangeChatSettings                      = 118
	opLogoutStart                             = 119
	opLogoutCancel                            = 120
	opClaimOrbStart                           = 121
	opClaimOrbCancel                          = 122
	opMatchLootChestOpeningStart              = 123
	opMatchLootChestOpeningCancel             = 124
	opDepositToGuildAccount                   = 125
	opWithdrawalFromAccount                   = 126
	opChangeGuildPayUpkeepFlag                = 127
	opChangeGuildTax                          = 128
	opGetMyTerritories                        = 129
	opMorganaCommand                          = 130
	opGetServerInfo                           = 131
	opInviteMercenaryToMatch                  = 132
	opSubscribeToCluster                      = 133
	opAnswerMercenaryInvitation               = 134
	opGetCharacterEquipment                   = 135
	opGetCharacterSteamAchievements           = 136
	opGetCharacterStats                       = 137
	opGetKillHistoryDetails                   = 138
	opLearnMasteryLevel                       = 139
	opReSpecAchievement                       = 140
	opChangeAvatar                            = 141
	opGetRankings                             = 142
	opGetRank                                 = 143
	opGetGvgSeasonRankings                    = 144
	opGetGvgSeasonRank                        = 145
	opGetGvgSeasonHistoryRankings             = 146
	opGetGvgSeasonGuildMemberHistory          = 147
	opKickFromGvGMatch                        = 148
	opGetChestLogs                            = 149
	opGetAccessRightLogs                      = 150
	opGetGuildAccountLogs                     = 151
	opInviteToPlayerTrade                     = 152
	opPlayerTradeCancel                       = 153
	opPlayerTradeInvitationAccept             = 154
	opPlayerTradeAddItem                      = 155
	opPlayerTradeRemoveItem                   = 156
	opPlayerTradeAcceptTrade                  = 157
	opPlayerTradeSetSilverOrGold              = 158
	opSendMiniMapPing                         = 159
	opStuck                                   = 160
	opBuyRealEstate                           = 161
	opClaimRealEstate                         = 162
	opGiveUpRealEstate                        = 163
	opChangeRealEstateOutline                 = 164
	opGetMailInfos                            = 165
	opReadMail                                = 166
	opSendNewMail                             = 167
	opDeleteMail                              = 168
	opClaimAttachmentFromMail                 = 169
	opUpdateLfgInfo                           = 170
	opGetLfgInfos                             = 171
	opGetMyGuildLfgInfo                       = 172
	opGetLfgDescriptionText                   = 173
	opLfgApplyToGuild                         = 174
	opAnswerLfgGuildApplication               = 175
	opGetClusterInfo                          = 176
	opRegisterChatPeer                        = 177
	opSendChatMessage                         = 178
	opJoinChatChannel                         = 179
	opLeaveChatChannel                        = 180
	opSendWhisperMessage                      = 181
	opSay                                     = 182
	opPlayEmote                               = 183
	opStopEmote                               = 184
	opGetClusterMapInfo                       = 185
	opAccessRightsChangeSettings              = 186
	opMount                                   = 187
	opMountCancel                             = 188
	opBuyJourney                              = 189
	opSetSaleStatusForEstate                  = 190
	opResolveGuildOrPlayerName                = 191
	opGetRespawnInfos                         = 192
	opMakeHome                                = 193
	opLeaveHome                               = 194
	opResurrectionReply                       = 195
	opAllianceCreate                          = 196
	opAllianceDisband                         = 197
	opAllianceGetMemberInfos                  = 198
	opAllianceInvite                          = 199
	opAllianceAnswerInvitation                = 200
	opAllianceCancelInvitation                = 201
	opAllianceKickGuild                       = 202
	opAllianceLeave                           = 203
	opAllianceChangeGoldPaymentFlag           = 204
	opAllianceGetDetailInfo                   = 205
	opGetIslandInfos                          = 206
	opAbandonMyIsland                         = 207
	opBuyMyIsland                             = 208
	opBuyGuildIsland                          = 209
	opAbandonGuildIsland                      = 210
	opUpgradeMyIsland                         = 211
	opUpgradeGuildIsland                      = 212
	opMoveMyIsland                            = 213
	opMoveGuildIsland                         = 214
	opTerritoryFillNutrition                  = 215
	opTeleportBack                            = 216
	opPartyInvitePlayer                       = 217
	opPartyAnswerInvitation                   = 218
	opPartyLeave                              = 219
	opPartyKickPlayer                         = 220
	opPartyMakeLeader                         = 221
	opPartyChangeLootSetting                  = 222
	opPartyMarkObject                         = 223
	opPartySetRole                            = 224
	opGetGuildMOTD                            = 225
	opSetGuildMOTD                            = 226
	opExitEnterStart                          = 227
	opExitEnterCancel                         = 228
	opQuestGiverRequest                       = 229
	opGoldMarketGetBuyOffer                   = 230
	opGoldMarketGetBuyOfferFromSilver         = 231
	opGoldMarketGetSellOffer                  = 232
	opGoldMarketGetSellOfferFromSilver        = 233
	opGoldMarketBuyGold                       = 234
	opGoldMarketSellGold                      = 235
	opGoldMarketCreateSellOrder               = 236
	opGoldMarketCreateBuyOrder                = 237
	opGoldMarketGetInfos                      = 238
	opGoldMarketCancelOrder                   = 239
	opGoldMarketGetAverageInfo                = 240
	opSiegeCampClaimStart                     = 241
	opSiegeCampClaimCancel                    = 242
	opTreasureChestUsingStart                 = 243
	opTreasureChestUsingCancel                = 244
	opUseLootChest                            = 245
	opUseShrine                               = 246
	opLaborerStartJob                         = 247
	opLaborerTakeJobLoot                      = 248
	opLaborerDismiss                          = 249
	opLaborerMove                             = 250
	opLaborerBuyItem                          = 251
	opLaborerUpgrade                          = 252
	opBuyPremium                              = 253
	opBuyTrial                                = 254
	opRealEstateGetAuctionData                = 255
	opRealEstateBidOnAuction                  = 256
	opGetSiegeCampCooldown                    = 257
	opFriendInvite                            = 258
	opFriendAnswerInvitation                  = 259
	opFriendCancelnvitation                   = 260
	opFriendRemove                            = 261
	opInventoryStack                          = 262
	opInventorySort                           = 263
	opEquipmentItemChangeSpell                = 264
	opExpeditionRegister                      = 265
	opExpeditionRegisterCancel                = 266
	opJoinExpedition                          = 267
	opDeclineExpeditionInvitation             = 268
	opVoteStart                               = 269
	opVoteDoVote                              = 270
	opRatingDoRate                            = 271
	opEnteringExpeditionStart                 = 272
	opEnteringExpeditionCancel                = 273
	opActivateExpeditionCheckPoint            = 274
	opArenaRegister                           = 275
	opArenaRegisterCancel                     = 276
	opArenaLeave                              = 277
	opJoinArenaMatch                          = 278
	opDeclineArenaInvitation                  = 279
	opEnteringArenaStart                      = 280
	opEnteringArenaCancel                     = 281
	opArenaCustomMatch                        = 282
	opArenaCustomMatchCreate                  = 283
	opUpdateCharacterStatement                = 284
	opBoostFarmable                           = 285
	opGetStrikeHistory                        = 286
	opUseFunction                             = 287
	opUsePortalEntrance                       = 288
	opResetPortalBinding                      = 289
	opQueryPortalBinding                      = 290
	opClaimPaymentTransaction                 = 291
	opChangeUseFlag                           = 292
	opClientPerformanceStats                  = 293
	opExtendedHardwareStats                   = 294
	opClientLowMemoryWarning                  = 295
	opTerritoryClaimStart                     = 296
	opTerritoryClaimCancel                    = 297
	opRequestAppStoreProducts                 = 298
	opVerifyProductPurchase                   = 299
	opQueryGuildPlayerStats                   = 300
	opTrackAchievements                       = 301
	opSetAchievementsAutoLearn                = 302
	opDepositItemToGuildCurrency              = 303
	opWithdrawalItemFromGuildCurrency         = 304
	opAuctionSellSpecificItemRequest          = 305
	opFishingStart                            = 306
	opFishingCasting                          = 307
	opFishingCast                             = 308
	opFishingCatch                            = 309
	opFishingPull                             = 310
	opFishingGiveLine                         = 311
	opFishingFinish                           = 312
	opFishingCancel                           = 313
	opCreateGuildAccessTag                    = 314
	opDeleteGuildAccessTag                    = 315
	opRenameGuildAccessTag                    = 316
	opFlagGuildAccessTagGuildPermission       = 317
	opAssignGuildAccessTag                    = 318
	opRemoveGuildAccessTagFromPlayer          = 319
	opModifyGuildAccessTagEditors             = 320
	opRequestPublicAccessTags                 = 321
	opChangeAccessTagPublicFlag               = 322
	opUpdateGuildAccessTag                    = 323
	opSteamStartMicrotransaction              = 324
	opSteamFinishMicrotransaction             = 325
	opSteamIdHasActiveAccount                 = 326
	opCheckEmailAccountState                  = 327
	opLinkAccountToSteamId                    = 328
	opBuyGvgSeasonBooster                     = 329
	opChangeFlaggingPrepare                   = 330
	opOverCharge                              = 331
	opOverChargeEnd                           = 332
	opRequestTrusted                          = 333
	opChangeGuildLogo                         = 334
	opPartyFinderRegisterForUpdates           = 335
	opPartyFinderUnregisterForUpdates         = 336
	opPartyFinderEnlistNewPartySearch         = 337
	opPartyFinderDeletePartySearch            = 338
	opPartyFinderChangePartySearch            = 339
	opPartyFinderChangeRole                   = 340
	opPartyFinderApplyForGroup                = 341
	opPartyFinderAcceptOrDeclineApplyForGroup = 342
	opPartyFinderGetEquipmentSnapshot         = 343
	opPartyFinderRegisterApplicants           = 344
	opPartyFinderUnregisterApplicants         = 345
	opPartyFinderFulltextSearch               = 346
	opPartyFinderRequestEquipmentSnapshot     = 347
	opGetPersonalSeasonTrackerData            = 348
	opUseConsumableFromInventory              = 349
	opClaimPersonalSeasonReward               = 350
	opEasyAntiCheatMessageToServer            = 351
	opRetaliationAttackClaimTerritory         = 352
	opSetNextTutorialState                    = 353
	opAddPlayerToMuteList                     = 354
	opRemovePlayerFromMuteList                = 355
	opMakeTerritoryHome                       = 356
	opLeaveTerritoryHome                      = 357
	opProductShopUserEvent                    = 358
	opSetFavoriteIsland                       = 359
)
