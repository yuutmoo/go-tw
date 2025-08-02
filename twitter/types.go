package twitter

type ProfileResponse struct {
	Data struct {
		UserResultByScreenName struct {
			Result struct {
				Typename string `json:"__typename"`
				Legacy   struct {
				} `json:"legacy"`
				Privacy struct {
					Protected bool `json:"protected"`
				} `json:"privacy"`
				RelationshipPerspectives struct {
					Blocking   bool `json:"blocking"`
					BlockedBy  bool `json:"blocked_by"`
					Following  bool `json:"following"`
					FollowedBy bool `json:"followed_by"`
				} `json:"relationship_perspectives"`
				RestId string `json:"rest_id"`
				Core   struct {
					Name       string `json:"name"`
					ScreenName string `json:"screen_name"`
				} `json:"core"`
				IsVerifiedOrganization bool `json:"is_verified_organization"`
				Profilemodules         struct {
					V1 []interface{} `json:"v1"`
				} `json:"profilemodules"`
				Id string `json:"id"`
			} `json:"result"`
			Id string `json:"id"`
		} `json:"user_result_by_screen_name"`
	} `json:"data"`
}

type FollowResponse struct {
	Errors []struct {
		Message   string `json:"message"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Path       []string `json:"path"`
		Extensions struct {
			Name    string `json:"name"`
			Source  string `json:"source"`
			Code    int    `json:"code"`
			Kind    string `json:"kind"`
			Tracing struct {
				TraceId string `json:"trace_id"`
			} `json:"tracing"`
		} `json:"extensions"`
		Code    int    `json:"code"`
		Kind    string `json:"kind"`
		Name    string `json:"name"`
		Source  string `json:"source"`
		Tracing struct {
			TraceId string `json:"trace_id"`
		} `json:"tracing"`
	} `json:"errors"`
}

type LikeRequest struct {
	Variables struct {
		TweetId string `json:"tweet_id"`
	} `json:"variables"`
	QueryId string `json:"queryId"`
}

type LikeResponse struct {
	Data struct {
		FavoriteTweet string `json:"favorite_tweet"`
	} `json:"data"`
	Errors []struct {
		Message   string `json:"message"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Path       []string `json:"path"`
		Extensions struct {
			Name    string `json:"name"`
			Source  string `json:"source"`
			Code    int    `json:"code"`
			Kind    string `json:"kind"`
			Tracing struct {
				TraceId string `json:"trace_id"`
			} `json:"tracing"`
		} `json:"extensions"`
		Code    int    `json:"code"`
		Kind    string `json:"kind"`
		Name    string `json:"name"`
		Source  string `json:"source"`
		Tracing struct {
			TraceId string `json:"trace_id"`
		} `json:"tracing"`
	} `json:"errors"`
}

type CommentRequest struct {
	Variables struct {
		TweetText string `json:"tweet_text"`
		Reply     struct {
			InReplyToTweetId    string        `json:"in_reply_to_tweet_id"`
			ExcludeReplyUserIds []interface{} `json:"exclude_reply_user_ids"`
		} `json:"reply"`
		DarkRequest bool `json:"dark_request"`
		Media       struct {
			MediaEntities     []interface{} `json:"media_entities"`
			PossiblySensitive bool          `json:"possibly_sensitive"`
		} `json:"media"`
		SemanticAnnotationIds  []interface{} `json:"semantic_annotation_ids"`
		DisallowedReplyOptions interface{}   `json:"disallowed_reply_options"`
	} `json:"variables"`
	Features `json:"features"`
	QueryId  string `json:"queryId"`
}

type Features struct {
	PremiumContentApiReadEnabled                                   bool `json:"premium_content_api_read_enabled"`
	CommunitiesWebEnableTweetCommunityResultsFetch                 bool `json:"communities_web_enable_tweet_community_results_fetch"`
	C9STweetAnatomyModeratorBadgeEnabled                           bool `json:"c9s_tweet_anatomy_moderator_badge_enabled"`
	ResponsiveWebGrokAnalyzeButtonFetchTrendsEnabled               bool `json:"responsive_web_grok_analyze_button_fetch_trends_enabled"`
	ResponsiveWebGrokAnalyzePostFollowupsEnabled                   bool `json:"responsive_web_grok_analyze_post_followups_enabled"`
	ResponsiveWebJetfuelFrame                                      bool `json:"responsive_web_jetfuel_frame"`
	ResponsiveWebGrokShareAttachmentEnabled                        bool `json:"responsive_web_grok_share_attachment_enabled"`
	ResponsiveWebEditTweetApiEnabled                               bool `json:"responsive_web_edit_tweet_api_enabled"`
	GraphqlIsTranslatableRwebTweetIsTranslatableEnabled            bool `json:"graphql_is_translatable_rweb_tweet_is_translatable_enabled"`
	ViewCountsEverywhereApiEnabled                                 bool `json:"view_counts_everywhere_api_enabled"`
	LongformNotetweetsConsumptionEnabled                           bool `json:"longform_notetweets_consumption_enabled"`
	ResponsiveWebTwitterArticleTweetConsumptionEnabled             bool `json:"responsive_web_twitter_article_tweet_consumption_enabled"`
	TweetAwardsWebTippingEnabled                                   bool `json:"tweet_awards_web_tipping_enabled"`
	ResponsiveWebGrokShowGrokTranslatedPost                        bool `json:"responsive_web_grok_show_grok_translated_post"`
	ResponsiveWebGrokAnalysisButtonFromBackend                     bool `json:"responsive_web_grok_analysis_button_from_backend"`
	CreatorSubscriptionsQuoteTweetPreviewEnabled                   bool `json:"creator_subscriptions_quote_tweet_preview_enabled"`
	LongformNotetweetsRichTextReadEnabled                          bool `json:"longform_notetweets_rich_text_read_enabled"`
	LongformNotetweetsInlineMediaEnabled                           bool `json:"longform_notetweets_inline_media_enabled"`
	PaymentsEnabled                                                bool `json:"payments_enabled"`
	ProfileLabelImprovementsPcfLabelInPostEnabled                  bool `json:"profile_label_improvements_pcf_label_in_post_enabled"`
	RwebTipjarConsumptionEnabled                                   bool `json:"rweb_tipjar_consumption_enabled"`
	VerifiedPhoneLabelEnabled                                      bool `json:"verified_phone_label_enabled"`
	ArticlesPreviewEnabled                                         bool `json:"articles_preview_enabled"`
	ResponsiveWebGrokCommunityNoteAutoTranslationIsEnabled         bool `json:"responsive_web_grok_community_note_auto_translation_is_enabled"`
	ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled      bool `json:"responsive_web_graphql_skip_user_profile_image_extensions_enabled"`
	FreedomOfSpeechNotReachFetchEnabled                            bool `json:"freedom_of_speech_not_reach_fetch_enabled"`
	StandardizedNudgesMisinfo                                      bool `json:"standardized_nudges_misinfo"`
	TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled bool `json:"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled"`
	ResponsiveWebGrokImageAnnotationEnabled                        bool `json:"responsive_web_grok_image_annotation_enabled"`
	ResponsiveWebGraphqlTimelineNavigationEnabled                  bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
	ResponsiveWebEnhanceCardsEnabled                               bool `json:"responsive_web_enhance_cards_enabled"`
	ResponsiveWebGraphqlExcludeDirectiveEnabled                    bool `json:"responsive_web_graphql_exclude_directive_enabled"`
	ResponsiveWebMediaDownloadVideoEnabled                         bool `json:"responsive_web_media_download_video_enabled"`
	TweetypieUnmentionOptimizationEnabled                         bool `json:"tweetypie_unmention_optimization_enabled"`
}

type CommentResponse struct {
	Data struct {
		CreateTweet struct {
			TweetResults struct {
				Result struct {
					RestId string `json:"rest_id"`
					Core   struct {
						UserResults struct {
							Result struct {
								Typename                   string `json:"__typename"`
								Id                         string `json:"id"`
								RestId                     string `json:"rest_id"`
								AffiliatesHighlightedLabel struct {
								} `json:"affiliates_highlighted_label"`
								Avatar struct {
									ImageUrl string `json:"image_url"`
								} `json:"avatar"`
								Core struct {
									CreatedAt  string `json:"created_at"`
									Name       string `json:"name"`
									ScreenName string `json:"screen_name"`
								} `json:"core"`
								DmPermissions struct {
									CanDm bool `json:"can_dm"`
								} `json:"dm_permissions"`
								HasGraduatedAccess bool `json:"has_graduated_access"`
								IsBlueVerified     bool `json:"is_blue_verified"`
								Legacy             struct {
									DefaultProfile      bool   `json:"default_profile"`
									DefaultProfileImage bool   `json:"default_profile_image"`
									Description         string `json:"description"`
									Entities            struct {
										Description struct {
											Urls []interface{} `json:"urls"`
										} `json:"description"`
									} `json:"entities"`
									FastFollowersCount      int           `json:"fast_followers_count"`
									FavouritesCount         int           `json:"favourites_count"`
									FollowersCount          int           `json:"followers_count"`
									FriendsCount            int           `json:"friends_count"`
									HasCustomTimelines      bool          `json:"has_custom_timelines"`
									IsTranslator            bool          `json:"is_translator"`
									ListedCount             int           `json:"listed_count"`
									MediaCount              int           `json:"media_count"`
									NeedsPhoneVerification  bool          `json:"needs_phone_verification"`
									NormalFollowersCount    int           `json:"normal_followers_count"`
									PinnedTweetIdsStr       []interface{} `json:"pinned_tweet_ids_str"`
									PossiblySensitive       bool          `json:"possibly_sensitive"`
									ProfileBannerUrl        string        `json:"profile_banner_url"`
									ProfileInterstitialType string        `json:"profile_interstitial_type"`
									StatusesCount           int           `json:"statuses_count"`
									TranslatorType          string        `json:"translator_type"`
									WantRetweets            bool          `json:"want_retweets"`
									WithheldInCountries     []interface{} `json:"withheld_in_countries"`
								} `json:"legacy"`
								Location struct {
									Location string `json:"location"`
								} `json:"location"`
								MediaPermissions struct {
									CanMediaTag bool `json:"can_media_tag"`
								} `json:"media_permissions"`
								ParodyCommentaryFanLabel string `json:"parody_commentary_fan_label"`
								ProfileImageShape        string `json:"profile_image_shape"`
								Privacy                  struct {
									Protected bool `json:"protected"`
								} `json:"privacy"`
								RelationshipPerspectives struct {
									Following bool `json:"following"`
								} `json:"relationship_perspectives"`
								TipjarSettings struct {
								} `json:"tipjar_settings"`
								Verification struct {
									Verified bool `json:"verified"`
								} `json:"verification"`
							} `json:"result"`
						} `json:"user_results"`
					} `json:"core"`
					UnmentionData struct {
					} `json:"unmention_data"`
					EditControl struct {
						EditTweetIds       []string `json:"edit_tweet_ids"`
						EditableUntilMsecs string   `json:"editable_until_msecs"`
						IsEditEligible     bool     `json:"is_edit_eligible"`
						EditsRemaining     string   `json:"edits_remaining"`
					} `json:"edit_control"`
					IsTranslatable bool `json:"is_translatable"`
					Views          struct {
						State string `json:"state"`
					} `json:"views"`
					Source string `json:"source"`
					Legacy struct {
						BookmarkCount     int    `json:"bookmark_count"`
						Bookmarked        bool   `json:"bookmarked"`
						CreatedAt         string `json:"created_at"`
						ConversationIdStr string `json:"conversation_id_str"`
						DisplayTextRange  []int  `json:"display_text_range"`
						Entities          struct {
							Hashtags     []interface{} `json:"hashtags"`
							Symbols      []interface{} `json:"symbols"`
							Timestamps   []interface{} `json:"timestamps"`
							Urls         []interface{} `json:"urls"`
							UserMentions []struct {
								IdStr      string `json:"id_str"`
								Name       string `json:"name"`
								ScreenName string `json:"screen_name"`
								Indices    []int  `json:"indices"`
							} `json:"user_mentions"`
						} `json:"entities"`
						FavoriteCount        int    `json:"favorite_count"`
						Favorited            bool   `json:"favorited"`
						FullText             string `json:"full_text"`
						InReplyToScreenName  string `json:"in_reply_to_screen_name"`
						InReplyToStatusIdStr string `json:"in_reply_to_status_id_str"`
						InReplyToUserIdStr   string `json:"in_reply_to_user_id_str"`
						IsQuoteStatus        bool   `json:"is_quote_status"`
						Lang                 string `json:"lang"`
						QuoteCount           int    `json:"quote_count"`
						ReplyCount           int    `json:"reply_count"`
						RetweetCount         int    `json:"retweet_count"`
						Retweeted            bool   `json:"retweeted"`
						UserIdStr            string `json:"user_id_str"`
						IdStr                string `json:"id_str"`
					} `json:"legacy"`
					UnmentionInfo struct {
					} `json:"unmention_info"`
				} `json:"result"`
			} `json:"tweet_results"`
		} `json:"create_tweet"`
	} `json:"data"`
	Errors []struct {
		Message   string `json:"message"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Path       []string `json:"path"`
		Extensions struct {
			Name    string `json:"name"`
			Source  string `json:"source"`
			Code    int    `json:"code"`
			Kind    string `json:"kind"`
			Tracing struct {
				TraceId string `json:"trace_id"`
			} `json:"tracing"`
		} `json:"extensions"`
		Code    int    `json:"code"`
		Kind    string `json:"kind"`
		Name    string `json:"name"`
		Source  string `json:"source"`
		Tracing struct {
			TraceId string `json:"trace_id"`
		} `json:"tracing"`
	} `json:"errors"`
}

type RetweetRequest struct {
	Variables struct {
		TweetId     string `json:"tweet_id"`
		DarkRequest bool   `json:"dark_request"`
	} `json:"variables"`
	QueryId string `json:"queryId"`
}

type RetweetResponse struct {
	Data struct {
		CreateRetweet struct {
			RetweetResults struct {
				Result struct {
					RestId string `json:"rest_id"`
					Legacy struct {
						FullText string `json:"full_text"`
					} `json:"legacy"`
				} `json:"result"`
			} `json:"retweet_results"`
		} `json:"create_retweet"`
	} `json:"data"`
	Errors []struct {
		Message   string `json:"message"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Path       []string `json:"path"`
		Extensions struct {
			Name    string `json:"name"`
			Source  string `json:"source"`
			Code    int    `json:"code"`
			Kind    string `json:"kind"`
			Tracing struct {
				TraceId string `json:"trace_id"`
			} `json:"tracing"`
		} `json:"extensions"`
		Code    int    `json:"code"`
		Kind    string `json:"kind"`
		Name    string `json:"name"`
		Source  string `json:"source"`
		Tracing struct {
			TraceId string `json:"trace_id"`
		} `json:"tracing"`
	} `json:"errors"`
}

type ResponseError struct {
	Errors []struct {
		Message   string `json:"message"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Path       []string `json:"path"`
		Extensions struct {
			Name    string `json:"name"`
			Source  string `json:"source"`
			Code    int    `json:"code"`
			Kind    string `json:"kind"`
			Tracing struct {
				TraceId string `json:"trace_id"`
			} `json:"tracing"`
		} `json:"extensions"`
		Code    int    `json:"code"`
		Kind    string `json:"kind"`
		Name    string `json:"name"`
		Source  string `json:"source"`
		Tracing struct {
			TraceId string `json:"trace_id"`
		} `json:"tracing"`
	} `json:"errors"`
}