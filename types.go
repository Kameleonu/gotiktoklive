package gotiktoklive

import "time"

type Event interface {
	CreatedTimestamp() int64
	IsHistory() bool
}

type RoomEvent struct {
	Timestamp int64
	MessageID int64
	Type      string
	Message   string
	isHistory bool
}

func (r RoomEvent) CreatedTimestamp() int64 {
	return r.Timestamp
}

func (r RoomEvent) IsHistory() bool {
	return r.isHistory
}

type ChatEvent struct {
	MessageID    int64
	Timestamp    int64
	Comment      string
	User         *User
	UserIdentity *UserIdentity
	isHistory    bool
}

func (c ChatEvent) IsHistory() bool {
	return c.isHistory
}

func (c ChatEvent) CreatedTimestamp() int64 {
	return c.Timestamp
}

type userEventType string

const (
	USER_JOIN   userEventType = "user joined the stream"
	USER_SHARE  userEventType = "user shared the stream"
	USER_FOLLOW userEventType = "user followed the host"
)

type UserEvent struct {
	Timestamp int64
	MessageID int64
	Event     userEventType
	User      *User
	isHistory bool
}

func (u UserEvent) CreatedTimestamp() int64 {
	return u.Timestamp
}

func (u UserEvent) IsHistory() bool {
	return u.isHistory
}

type ViewersEvent struct {
	Timestamp int64
	MessageID int64
	Viewers   int
	isHistory bool
}

func (v ViewersEvent) TimeComparableID() int64 {
	return v.MessageID
}

func (v ViewersEvent) IsHistory() bool {
	return v.isHistory
}

func (v ViewersEvent) CreatedTimestamp() int64 {
	return v.Timestamp
}

type GiftEvent struct {
	MessageID    int64
	Timestamp    int64
	ID           int64
	Name         string
	Describe     string
	Diamonds     int
	RepeatCount  int
	RepeatEnd    bool
	Type         int
	ToUserID     int64
	User         *User
	UserIdentity *UserIdentity
	isHistory    bool
	GroupID      int64
	IsComboGift  bool
}

func (g GiftEvent) CreatedTimestamp() int64 {
	return g.Timestamp
}

func (g GiftEvent) IsHistory() bool {
	return g.isHistory
}

type LikeEvent struct {
	MessageID   int64
	Timestamp   int64
	Likes       int
	TotalLikes  int
	User        *User
	DisplayType string
	Label       string
	isHistory   bool
}

func (l LikeEvent) IsHistory() bool {
	return l.isHistory
}

func (l LikeEvent) CreatedTimestamp() int64 {
	return l.Timestamp
}

type QuestionEvent struct {
	MessageID int64
	Timestamp int64
	Quesion   string
	User      *User
	isHistory bool
}

func (q QuestionEvent) CreatedTimestamp() int64 {
	return q.Timestamp
}

func (q QuestionEvent) IsHistory() bool {
	return q.isHistory
}

type ControlEvent struct {
	MessageID   int64
	Timestamp   int64
	Action      int
	Description string
	isHistory   bool
}

func (c ControlEvent) IsHistory() bool {
	return c.isHistory
}

func (c ControlEvent) TimeComparableID() int64 {
	return c.MessageID
}

func (c ControlEvent) CreatedTimestamp() int64 {
	return c.Timestamp
}

type MicBattleEvent struct {
	MessageID int64
	Timestamp int64
	Users     []*User
	isHistory bool
}

func (m MicBattleEvent) IsHistory() bool {
	return m.isHistory
}

func (m MicBattleEvent) CreatedTimestamp() int64 {
	return m.Timestamp
}

type BattlesEvent struct {
	MessageID int64
	Timestamp int64
	Status    int
	Battles   []*Battle
	isHistory bool
}

func (b BattlesEvent) IsHistory() bool {
	return b.isHistory
}

func (b BattlesEvent) CreatedTimestamp() int64 {
	return b.Timestamp
}

type RoomBannerEvent struct {
	MessageID int64
	Timestamp int64
	Data      interface{}
	isHistory bool
}

func (r RoomBannerEvent) IsHistory() bool {
	return r.isHistory
}

func (r RoomBannerEvent) CreatedTimestamp() int64 {
	return r.Timestamp
}

type IntroEvent struct {
	MessageID int64
	Timestamp int64
	ID        int
	Title     string
	User      *User
	isHistory bool
}

func (i IntroEvent) IsHistory() bool {
	return i.isHistory
}

func (i IntroEvent) TimeComparableID() int64 {
	return i.MessageID
}

func (i IntroEvent) CreatedTimestamp() int64 {
	return i.Timestamp
}

type Battle struct {
	Host   int64
	Groups []*BattleGroup
}

type BattleGroup struct {
	Points int
	Users  []*User
}

type User struct {
	ID              int64
	Username        string
	Nickname        string
	AvatarThumb     *ProfilePicture
	ExtraAttributes *ExtraAttributes
	Badge           *BadgeAttributes
}

type UserIdentity struct {
	IsGiftGiver       bool
	IsSubscriber      bool
	IsMutualFollowing bool
	IsFollower        bool
	IsModerator       bool
	IsAnchor          bool
}

type ProfilePicture struct {
	Urls []string
}

type ExtraAttributes struct {
	FollowRole int
}

type BadgeAttributes struct {
	Badges []*UserBadge
}

type UserBadge struct {
	Type string
	Name string
}

type roomInfoRsp struct {
	RoomInfo *RoomInfo `json:"data"`
	Extra    struct {
		Now int64 `json:"now"`
	} `json:"extra"`
	StatusCode float64 `json:"status_code"`
}

type RoomInfo struct {
	AnchorABMap              interface{}   `json:"AnchorABMap"`
	AdminUserIds             []interface{} `json:"admin_user_ids"`
	AnchorScheduledTimeText  string        `json:"anchor_scheduled_time_text"`
	AnchorShareText          string        `json:"anchor_share_text"`
	AnchorTabType            float64       `json:"anchor_tab_type"`
	AnsweringQuestionContent string        `json:"answering_question_content"`
	AppID                    float64       `json:"app_id"`
	AutoCover                float64       `json:"auto_cover"`
	BookEndTime              float64       `json:"book_end_time"`
	BookTime                 float64       `json:"book_time"`
	BusinessLive             float64       `json:"business_live"`
	ChallengeInfo            string        `json:"challenge_info"`
	ClientVersion            float64       `json:"client_version"`
	CommentNameMode          float64       `json:"comment_name_mode"`
	CommerceInfo             struct {
		CommercePermission       float64 `json:"commerce_permission"`
		OecLiveEnterRoomInitData string  `json:"oec_live_enter_room_init_data"`
	} `json:"commerce_info"`
	CommonLabelList string `json:"common_label_list"`
	ContentTag      string `json:"content_tag"`
	Cover           struct {
		AvgColor   string   `json:"avg_color"`
		Height     float64  `json:"height"`
		ImageType  float64  `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      float64  `json:"width"`
	} `json:"cover"`
	CreateTime           int64         `json:"create_time"`
	DecoList             []interface{} `json:"deco_list"`
	DisablePreloadStream bool          `json:"disable_preload_stream"`
	FansclubMsgStyle     float64       `json:"fansclub_msg_style"`
	FeedRoomLabel        struct {
		AvgColor   string   `json:"avg_color"`
		Height     float64  `json:"height"`
		ImageType  float64  `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      float64  `json:"width"`
	} `json:"feed_room_label"`
	FeedRoomLabels      []interface{} `json:"feed_room_labels"`
	FilterMsgRules      []interface{} `json:"filter_msg_rules"`
	FinishReason        float64       `json:"finish_reason"`
	FinishTime          float64       `json:"finish_time"`
	FinishURL           string        `json:"finish_url"`
	FinishURLV2         string        `json:"finish_url_v2"`
	FollowMsgStyle      float64       `json:"follow_msg_style"`
	ForumExtraData      string        `json:"forum_extra_data"`
	GameTag             []interface{} `json:"game_tag"`
	GiftMsgStyle        float64       `json:"gift_msg_style"`
	GiftPollVoteEnabled bool          `json:"gift_poll_vote_enabled"`
	GroupSource         float64       `json:"group_source"`
	HasCommerceGoods    bool          `json:"has_commerce_goods"`
	Hashtag             struct {
		ID    float64 `json:"id"`
		Image struct {
			AvgColor   string   `json:"avg_color"`
			Height     float64  `json:"height"`
			ImageType  float64  `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      float64  `json:"width"`
		} `json:"image"`
		Namespace float64 `json:"namespace"`
		Title     string  `json:"title"`
	} `json:"hashtag"`
	HaveWishlist               bool    `json:"have_wishlist"`
	HotSentenceInfo            string  `json:"hot_sentence_info"`
	ID                         int64   `json:"id"`
	IDStr                      string  `json:"id_str"`
	InteractionQuestionVersion float64 `json:"interaction_question_version"`
	Introduction               string  `json:"introduction"`
	IsGatedRoom                bool    `json:"is_gated_room"`
	IsReplay                   bool    `json:"is_replay"`
	IsShowUserCardSwitch       bool    `json:"is_show_user_card_switch"`
	LastPingTime               float64 `json:"last_ping_time"`
	Layout                     float64 `json:"layout"`
	LikeCount                  float64 `json:"like_count"`
	LinkMic                    struct {
		AudienceIDList []interface{} `json:"audience_id_list"`
		BattleScores   []interface{} `json:"battle_scores"`
		BattleSettings struct {
			BattleID    float64 `json:"battle_id"`
			ChannelID   float64 `json:"channel_id"`
			Duration    float64 `json:"duration"`
			Finished    float64 `json:"finished"`
			MatchType   float64 `json:"match_type"`
			StartTime   float64 `json:"start_time"`
			StartTimeMs float64 `json:"start_time_ms"`
			Theme       string  `json:"theme"`
		} `json:"battle_settings"`
		ChannelID      float64       `json:"channel_id"`
		FollowedCount  float64       `json:"followed_count"`
		LinkedUserList []interface{} `json:"linked_user_list"`
		MultiLiveEnum  float64       `json:"multi_live_enum"`
		RivalAnchorID  float64       `json:"rival_anchor_id"`
		ShowUserList   []interface{} `json:"show_user_list"`
	} `json:"link_mic"`
	LinkerMap struct {
	} `json:"linker_map"`
	LinkmicLayout      float64       `json:"linkmic_layout"`
	LiveDistribution   []interface{} `json:"live_distribution"`
	LiveID             float64       `json:"live_id"`
	LiveReason         string        `json:"live_reason"`
	LiveRoomMode       float64       `json:"live_room_mode"`
	LiveTypeAudio      bool          `json:"live_type_audio"`
	LiveTypeLinkmic    bool          `json:"live_type_linkmic"`
	LiveTypeNormal     bool          `json:"live_type_normal"`
	LiveTypeSandbox    bool          `json:"live_type_sandbox"`
	LiveTypeScreenshot bool          `json:"live_type_screenshot"`
	LiveTypeSocialLive bool          `json:"live_type_social_live"`
	LiveTypeThirdParty bool          `json:"live_type_third_party"`
	LivingRoomAttrs    struct {
		AdminFlag   float64 `json:"admin_flag"`
		Rank        float64 `json:"rank"`
		RoomID      int64   `json:"room_id"`
		RoomIDStr   string  `json:"room_id_str"`
		SilenceFlag float64 `json:"silence_flag"`
	} `json:"living_room_attrs"`
	LotteryFinishTime    float64   `json:"lottery_finish_time"`
	MosaicStatus         float64   `json:"mosaic_status"`
	OsType               float64   `json:"os_type"`
	Owner                *UserData `json:"owner"`
	OwnerDeviceID        float64   `json:"owner_device_id"`
	OwnerDeviceIDStr     string    `json:"owner_device_id_str"`
	OwnerUserID          float64   `json:"owner_user_id"`
	OwnerUserIDStr       string    `json:"owner_user_id_str"`
	PreEnterTime         float64   `json:"pre_enter_time"`
	PreviewFlowTag       float64   `json:"preview_flow_tag"`
	RanklistAudienceType float64   `json:"ranklist_audience_type"`
	RelationTag          string    `json:"relation_tag"`
	Replay               bool      `json:"replay"`
	RoomAuditStatus      float64   `json:"room_audit_status"`
	RoomAuth             struct {
		Banner              float64 `json:"Banner"`
		BroadcastMessage    float64 `json:"BroadcastMessage"`
		Chat                bool    `json:"Chat"`
		ChatL2              bool    `json:"ChatL2"`
		ChatSubOnly         bool    `json:"ChatSubOnly"`
		Danmaku             bool    `json:"Danmaku"`
		Digg                bool    `json:"Digg"`
		DonationSticker     float64 `json:"DonationSticker"`
		Gift                bool    `json:"Gift"`
		GiftAnchorMt        float64 `json:"GiftAnchorMt"`
		GiftPoll            float64 `json:"GiftPoll"`
		GoldenEnvelope      float64 `json:"GoldenEnvelope"`
		InteractionQuestion bool    `json:"InteractionQuestion"`
		Landscape           float64 `json:"Landscape"`
		LandscapeChat       float64 `json:"LandscapeChat"`
		LuckMoney           bool    `json:"LuckMoney"`
		Poll                float64 `json:"Poll"`
		Promote             bool    `json:"Promote"`
		Props               bool    `json:"Props"`
		PublicScreen        float64 `json:"PublicScreen"`
		QuickChat           float64 `json:"QuickChat"`
		Rank                float64 `json:"Rank"`
		RoomContributor     bool    `json:"RoomContributor"`
		Share               bool    `json:"Share"`
		ShareEffect         float64 `json:"ShareEffect"`
		UserCard            bool    `json:"UserCard"`
		UserCount           float64 `json:"UserCount"`
		Viewers             bool    `json:"Viewers"`
		TransactionHistory  float64 `json:"transaction_history"`
	} `json:"room_auth"`
	RoomCreateAbParam string        `json:"room_create_ab_param"`
	RoomLayout        float64       `json:"room_layout"`
	RoomStickerList   []interface{} `json:"room_sticker_list"`
	RoomTabs          []interface{} `json:"room_tabs"`
	RoomTag           float64       `json:"room_tag"`
	ScrollConfig      string        `json:"scroll_config"`
	SearchID          float64       `json:"search_id"`
	ShareMsgStyle     float64       `json:"share_msg_style"`
	ShareURL          string        `json:"share_url"`
	ShortTitle        string        `json:"short_title"`
	ShortTouchItems   []interface{} `json:"short_touch_items"`
	SocialInteraction struct {
		MultiLive struct {
			UserSettings struct {
				MultiLiveApplyPermission float64 `json:"multi_live_apply_permission"`
			} `json:"user_settings"`
		} `json:"multi_live"`
	} `json:"social_interaction"`
	StartTime float64 `json:"start_time"`
	Stats     struct {
		DiggCount            float64 `json:"digg_count"`
		EnterCount           float64 `json:"enter_count"`
		FanTicket            float64 `json:"fan_ticket"`
		FollowCount          float64 `json:"follow_count"`
		GiftUvCount          float64 `json:"gift_uv_count"`
		ID                   int64   `json:"id"`
		IDStr                string  `json:"id_str"`
		LikeCount            float64 `json:"like_count"`
		ReplayFanTicket      float64 `json:"replay_fan_ticket"`
		ReplayViewers        float64 `json:"replay_viewers"`
		ShareCount           float64 `json:"share_count"`
		TotalUser            float64 `json:"total_user"`
		TotalUserDesp        string  `json:"total_user_desp"`
		UserCountComposition struct {
			MyFollow    float64 `json:"my_follow"`
			Other       float64 `json:"other"`
			VideoDetail float64 `json:"video_detail"`
		} `json:"user_count_composition"`
		Watermelon float64 `json:"watermelon"`
	} `json:"stats"`
	Status      float64       `json:"status"`
	StickerList []interface{} `json:"sticker_list"`
	StreamID    int64         `json:"stream_id"`
	StreamIDStr string        `json:"stream_id_str"`
	StreamURL   struct {
		CandidateResolution []string      `json:"candidate_resolution"`
		CompletePushUrls    []interface{} `json:"complete_push_urls"`
		DefaultResolution   string        `json:"default_resolution"`
		Extra               struct {
			AnchorInteractProfile   float64 `json:"anchor_interact_profile"`
			AudienceInteractProfile float64 `json:"audience_interact_profile"`
			BframeEnable            bool    `json:"bframe_enable"`
			BitrateAdaptStrategy    float64 `json:"bitrate_adapt_strategy"`
			Bytevc1Enable           bool    `json:"bytevc1_enable"`
			DefaultBitrate          float64 `json:"default_bitrate"`
			Fps                     float64 `json:"fps"`
			GopSec                  float64 `json:"gop_sec"`
			HardwareEncode          bool    `json:"hardware_encode"`
			Height                  float64 `json:"height"`
			MaxBitrate              float64 `json:"max_bitrate"`
			MinBitrate              float64 `json:"min_bitrate"`
			Roi                     bool    `json:"roi"`
			SwRoi                   bool    `json:"sw_roi"`
			VideoProfile            float64 `json:"video_profile"`
			Width                   float64 `json:"width"`
		} `json:"extra"`
		FlvPullURL struct {
			FullHd1 string `json:"FULL_HD1"`
			Hd1     string `json:"HD1"`
			Sd1     string `json:"SD1"`
			Sd2     string `json:"SD2"`
		} `json:"flv_pull_url"`
		FlvPullURLParams struct {
			FullHd1 string `json:"FULL_HD1"`
			Hd1     string `json:"HD1"`
			Sd1     string `json:"SD1"`
			Sd2     string `json:"SD2"`
		} `json:"flv_pull_url_params"`
		HlsPullURL    string `json:"hls_pull_url"`
		HlsPullURLMap struct {
		} `json:"hls_pull_url_map"`
		HlsPullURLParams string `json:"hls_pull_url_params"`
		ID               int64  `json:"id"`
		IDStr            string `json:"id_str"`
		LiveCoreSdkData  struct {
			PullData struct {
				Options struct {
					DefaultQuality struct {
						Level      float64 `json:"level"`
						Name       string  `json:"name"`
						Resolution string  `json:"resolution"`
						SdkKey     string  `json:"sdk_key"`
						VCodec     string  `json:"v_codec"`
					} `json:"default_quality"`
					Qualities []struct {
						Level      float64 `json:"level"`
						Name       string  `json:"name"`
						Resolution string  `json:"resolution"`
						SdkKey     string  `json:"sdk_key"`
						VCodec     string  `json:"v_codec"`
					} `json:"qualities"`
				} `json:"options"`
				StreamData string `json:"stream_data"`
			} `json:"pull_data"`
		} `json:"live_core_sdk_data"`
		Provider       float64       `json:"provider"`
		PushUrls       []interface{} `json:"push_urls"`
		ResolutionName struct {
			Auto    string `json:"AUTO"`
			FullHd1 string `json:"FULL_HD1"`
			Hd1     string `json:"HD1"`
			Origion string `json:"ORIGION"`
			Sd1     string `json:"SD1"`
			Sd2     string `json:"SD2"`
		} `json:"resolution_name"`
		RtmpPullURL       string  `json:"rtmp_pull_url"`
		RtmpPullURLParams string  `json:"rtmp_pull_url_params"`
		RtmpPushURL       string  `json:"rtmp_push_url"`
		RtmpPushURLParams string  `json:"rtmp_push_url_params"`
		StreamControlType float64 `json:"stream_control_type"`
	} `json:"stream_url"`
	StreamURLFilteredInfo struct {
		IsGatedRoom bool `json:"is_gated_room"`
		IsPaidEvent bool `json:"is_paid_event"`
	} `json:"stream_url_filtered_info"`
	Title             string    `json:"title"`
	TopFans           []*TopFan `json:"top_fans"`
	UseFilter         bool      `json:"use_filter"`
	UserCount         int       `json:"user_count"` // Viewers
	UserShareText     string    `json:"user_share_text"`
	VideoFeedTag      string    `json:"video_feed_tag"`
	WebcastCommentTcs float64   `json:"webcast_comment_tcs"`
	WebcastSdkVersion float64   `json:"webcast_sdk_version"`
	WithDrawSomething bool      `json:"with_draw_something"`
	WithKtv           bool      `json:"with_ktv"`
	WithLinkmic       bool      `json:"with_linkmic"`
}

type giftInfoRsp struct {
	GiftInfo *GiftInfo `json:"data"`
	Extra    struct {
		LogID string `json:"log_id"`
		Now   int64  `json:"now"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}

type GiftInfo struct {
	DoodleTemplates []interface{} `json:"doodle_templates"`
	Gifts           []struct {
		ActionType            int           `json:"action_type"`
		AppID                 int           `json:"app_id"`
		BusinessText          string        `json:"business_text"`
		ColorInfos            []interface{} `json:"color_infos"`
		Combo                 bool          `json:"combo"`
		Describe              string        `json:"describe"`
		DiamondCount          int           `json:"diamond_count"`
		Duration              int           `json:"duration"`
		EventName             string        `json:"event_name"`
		ForCustom             bool          `json:"for_custom"`
		ForLinkmic            bool          `json:"for_linkmic"`
		GiftRankRecommendInfo string        `json:"gift_rank_recommend_info"`
		GiftScene             int           `json:"gift_scene"`
		GoldEffect            string        `json:"gold_effect"`
		GraySchemeURL         string        `json:"gray_scheme_url"`
		GuideURL              string        `json:"guide_url"`
		Icon                  struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"icon"`
		ID    int `json:"id"`
		Image struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"image"`
		IsBroadcastGift    bool `json:"is_broadcast_gift"`
		IsDisplayedOnPanel bool `json:"is_displayed_on_panel"`
		IsEffectBefview    bool `json:"is_effect_befview"`
		IsGray             bool `json:"is_gray"`
		IsRandomGift       bool `json:"is_random_gift"`
		ItemType           int  `json:"item_type"`
		LockInfo           struct {
			Lock     bool `json:"lock"`
			LockType int  `json:"lock_type"`
		} `json:"lock_info"`
		Manual          string `json:"manual"`
		Name            string `json:"name"`
		Notify          bool   `json:"notify"`
		PrimaryEffectID int    `json:"primary_effect_id"`
		Region          string `json:"region"`
		SchemeURL       string `json:"scheme_url"`
		SpecialEffects  struct {
		} `json:"special_effects"`
		TriggerWords  []interface{} `json:"trigger_words"`
		Type          int           `json:"type"`
		GiftLabelIcon struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"gift_label_icon,omitempty"`
		PreviewImage struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"preview_image,omitempty"`
		TrackerParams struct {
			GiftProperty string `json:"gift_property"`
		} `json:"tracker_params,omitempty"`
		GiftPanelBanner struct {
			BgColorValues []interface{} `json:"bg_color_values"`
			DisplayText   struct {
				DefaultFormat struct {
					Bold               bool   `json:"bold"`
					Color              string `json:"color"`
					FontSize           int    `json:"font_size"`
					Italic             bool   `json:"italic"`
					ItalicAngle        int    `json:"italic_angle"`
					UseHeighLightColor bool   `json:"use_heigh_light_color"`
					UseRemoteClor      bool   `json:"use_remote_clor"`
					Weight             int    `json:"weight"`
				} `json:"default_format"`
				DefaultPattern string        `json:"default_pattern"`
				Key            string        `json:"key"`
				Pieces         []interface{} `json:"pieces"`
			} `json:"display_text"`
			LeftIcon struct {
				AvgColor   string   `json:"avg_color"`
				Height     int      `json:"height"`
				ImageType  int      `json:"image_type"`
				IsAnimated bool     `json:"is_animated"`
				OpenWebURL string   `json:"open_web_url"`
				URI        string   `json:"uri"`
				URLList    []string `json:"url_list"`
				Width      int      `json:"width"`
			} `json:"left_icon"`
			SchemaURL string `json:"schema_url"`
		} `json:"gift_panel_banner,omitempty"`
	} `json:"gifts"`
	GiftsInfo struct {
		ColorGiftIconAnimation struct {
			AvgColor   string   `json:"avg_color"`
			Height     int      `json:"height"`
			ImageType  int      `json:"image_type"`
			IsAnimated bool     `json:"is_animated"`
			OpenWebURL string   `json:"open_web_url"`
			URI        string   `json:"uri"`
			URLList    []string `json:"url_list"`
			Width      int      `json:"width"`
		} `json:"color_gift_icon_animation"`
		DefaultLocColorGiftID            int  `json:"default_loc_color_gift_id"`
		EnableFirstRechargeDynamicEffect bool `json:"enable_first_recharge_dynamic_effect"`
		FirstRechargeGiftInfo            struct {
			ExpireAt             int `json:"expire_at"`
			GiftID               int `json:"gift_id"`
			OriginalDiamondCount int `json:"original_diamond_count"`
		} `json:"first_recharge_gift_info"`
		GiftComboInfos []interface{} `json:"gift_combo_infos"`
		GiftGroupInfos []struct {
			GroupCount int    `json:"group_count"`
			GroupText  string `json:"group_text"`
		} `json:"gift_group_infos"`
		GiftIconInfo struct {
			EffectURI string `json:"effect_uri"`
			Icon      struct {
				AvgColor   string        `json:"avg_color"`
				Height     int           `json:"height"`
				ImageType  int           `json:"image_type"`
				IsAnimated bool          `json:"is_animated"`
				OpenWebURL string        `json:"open_web_url"`
				URI        string        `json:"uri"`
				URLList    []interface{} `json:"url_list"`
				Width      int           `json:"width"`
			} `json:"icon"`
			IconID       int    `json:"icon_id"`
			IconURI      string `json:"icon_uri"`
			Name         string `json:"name"`
			ValidEndAt   int    `json:"valid_end_at"`
			ValidStartAt int    `json:"valid_start_at"`
			WithEffect   bool   `json:"with_effect"`
		} `json:"gift_icon_info"`
		GiftPollInfo struct {
			GiftPollOptions []struct {
				GiftID         int `json:"gift_id"`
				PollResultIcon struct {
					AvgColor   string   `json:"avg_color"`
					Height     int      `json:"height"`
					ImageType  int      `json:"image_type"`
					IsAnimated bool     `json:"is_animated"`
					OpenWebURL string   `json:"open_web_url"`
					URI        string   `json:"uri"`
					URLList    []string `json:"url_list"`
					Width      int      `json:"width"`
				} `json:"poll_result_icon"`
			} `json:"gift_poll_options"`
		} `json:"gift_poll_info"`
		GiftWords                 string `json:"gift_words"`
		HideRechargeEntry         bool   `json:"hide_recharge_entry"`
		NewGiftID                 int    `json:"new_gift_id"`
		RecentlySentColorGiftID   int    `json:"recently_sent_color_gift_id"`
		RecommendedRandomGiftID   int    `json:"recommended_random_gift_id"`
		ShowFirstRechargeEntrance bool   `json:"show_first_recharge_entrance"`
		SpeedyGiftID              int    `json:"speedy_gift_id"`
	} `json:"gifts_info"`
	Pages []interface{} `json:"pages"`
}

type TopFan struct {
	FanTicket float64   `json:"fan_ticket"`
	User      *UserData `json:"user"`
}

type FeedItem struct {
	LiveStreams []*LiveStream `json:"data"`
	Extra       struct {
		Banner struct {
			Banners     []interface{} `json:"banners"`
			BannersType int           `json:"banners_type"`
			SwitchType  int           `json:"switch_type"`
			Title       string        `json:"title"`
			Total       int           `json:"total"`
		} `json:"banner"`
		Cost        int    `json:"cost"`
		HasMore     bool   `json:"has_more"`
		HashtagText string `json:"hashtag_text"`
		IsBackup    int    `json:"is_backup"`
		LogPb       struct {
			ImprID    string `json:"impr_id"`
			SessionID int    `json:"session_id"`
		} `json:"log_pb"`
		MaxTime     int64  `json:"max_time"`
		MinTime     int    `json:"min_time"`
		Now         int64  `json:"now"`
		Style       int    `json:"style"`
		Total       int    `json:"total"`
		UnreadExtra string `json:"unread_extra"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}

type LiveStream struct {
	t *TikTok

	Room      *RoomInfo `json:"data"`
	DebugInfo string    `json:"debug_info"`
	FlareInfo struct {
		IsFlare bool   `json:"is_flare"`
		TaskID  string `json:"task_id"`
	} `json:"flare_info"`
	IsPseudoLiving  bool   `json:"is_pseudo_living"`
	IsRecommendCard bool   `json:"is_recommend_card"`
	LiveReason      string `json:"live_reason"`
	Rid             string `json:"rid"`
	Type            int    `json:"type"`
}

type UserData struct {
	AllowFindByContacts                 bool `json:"allow_find_by_contacts"`
	AllowOthersDownloadVideo            bool `json:"allow_others_download_video"`
	AllowOthersDownloadWhenSharingVideo bool `json:"allow_others_download_when_sharing_video"`
	AllowShareShowProfile               bool `json:"allow_share_show_profile"`
	AllowShowInGossip                   bool `json:"allow_show_in_gossip"`
	AllowShowMyAction                   bool `json:"allow_show_my_action"`
	AllowStrangeComment                 bool `json:"allow_strange_comment"`
	AllowUnfollowerComment              bool `json:"allow_unfollower_comment"`
	AllowUseLinkmic                     bool `json:"allow_use_linkmic"`
	AvatarLarge                         struct {
		AvgColor   string   `json:"avg_color"`
		Height     int      `json:"height"`
		ImageType  int      `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      int      `json:"width"`
	} `json:"avatar_large"`
	AvatarMedium struct {
		AvgColor   string   `json:"avg_color"`
		Height     int      `json:"height"`
		ImageType  int      `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      int      `json:"width"`
	} `json:"avatar_medium"`
	AvatarThumb struct {
		AvgColor   string   `json:"avg_color"`
		Height     int      `json:"height"`
		ImageType  int      `json:"image_type"`
		IsAnimated bool     `json:"is_animated"`
		OpenWebURL string   `json:"open_web_url"`
		URI        string   `json:"uri"`
		URLList    []string `json:"url_list"`
		Width      int      `json:"width"`
	} `json:"avatar_thumb"`
	BadgeImageList           []interface{} `json:"badge_image_list"`
	BadgeList                []interface{} `json:"badge_list"`
	BgImgURL                 string        `json:"bg_img_url"`
	BioDescription           string        `json:"bio_description"`
	BlockStatus              int           `json:"block_status"`
	BorderList               []interface{} `json:"border_list"`
	CommentRestrict          int           `json:"comment_restrict"`
	CommerceWebcastConfigIds []interface{} `json:"commerce_webcast_config_ids"`
	Constellation            string        `json:"constellation"`
	CreateTime               int           `json:"create_time"`
	DisableIchat             int           `json:"disable_ichat"`
	Username                 string        `json:"display_id"`
	EnableIchatImg           int           `json:"enable_ichat_img"`
	Exp                      int           `json:"exp"`
	FanTicketCount           int           `json:"fan_ticket_count"`
	FoldStrangerChat         bool          `json:"fold_stranger_chat"`
	FollowInfo               struct {
		FollowStatus   int `json:"follow_status"`
		FollowerCount  int `json:"follower_count"`
		FollowingCount int `json:"following_count"`
		PushStatus     int `json:"push_status"`
	} `json:"follow_info"`
	FollowStatus        int           `json:"follow_status"`
	IchatRestrictType   int           `json:"ichat_restrict_type"`
	ID                  int64         `json:"id"`
	IDStr               string        `json:"id_str"`
	IsFollower          bool          `json:"is_follower"`
	IsFollowing         bool          `json:"is_following"`
	LinkMicStats        int           `json:"link_mic_stats"`
	MediaBadgeImageList []interface{} `json:"media_badge_image_list"`
	ModifyTime          int           `json:"modify_time"`
	NeedProfileGuide    bool          `json:"need_profile_guide"`
	NewRealTimeIcons    []interface{} `json:"new_real_time_icons"`
	Nickname            string        `json:"nickname"`
	OwnRoom             struct {
		RoomIds    []int64  `json:"room_ids"`
		RoomIdsStr []string `json:"room_ids_str"`
	} `json:"own_room"`
	PayGrade struct {
		GradeBanner        string        `json:"grade_banner"`
		GradeDescribe      string        `json:"grade_describe"`
		GradeIconList      []interface{} `json:"grade_icon_list"`
		Level              int           `json:"level"`
		Name               string        `json:"name"`
		NextName           string        `json:"next_name"`
		NextPrivileges     string        `json:"next_privileges"`
		Score              int           `json:"score"`
		ScreenChatType     int           `json:"screen_chat_type"`
		UpgradeNeedConsume int           `json:"upgrade_need_consume"`
	} `json:"pay_grade"`
	PayScore           int           `json:"pay_score"`
	PayScores          int           `json:"pay_scores"`
	PushCommentStatus  bool          `json:"push_comment_status"`
	PushDigg           bool          `json:"push_digg"`
	PushFollow         bool          `json:"push_follow"`
	PushFriendAction   bool          `json:"push_friend_action"`
	PushIchat          bool          `json:"push_ichat"`
	PushStatus         bool          `json:"push_status"`
	PushVideoPost      bool          `json:"push_video_post"`
	PushVideoRecommend bool          `json:"push_video_recommend"`
	RealTimeIcons      []interface{} `json:"real_time_icons"`
	SecUID             string        `json:"sec_uid"`
	Secret             int           `json:"secret"`
	ShareQrcodeURI     string        `json:"share_qrcode_uri"`
	SpecialID          string        `json:"special_id"`
	Status             int           `json:"status"`
	TicketCount        int           `json:"ticket_count"`
	TopFans            []interface{} `json:"top_fans"`
	TopVipNo           int           `json:"top_vip_no"`
	UserAttr           struct {
		IsAdmin      bool `json:"is_admin"`
		IsMuted      bool `json:"is_muted"`
		IsSuperAdmin bool `json:"is_super_admin"`
		MuteDuration int  `json:"mute_duration"`
	} `json:"user_attr"`
	UserRole                    int    `json:"user_role"`
	Verified                    bool   `json:"verified"`
	VerifiedContent             string `json:"verified_content"`
	VerifiedReason              string `json:"verified_reason"`
	WithCarManagementPermission bool   `json:"with_car_management_permission"`
	WithCommercePermission      bool   `json:"with_commerce_permission"`
	WithFusionShopEntry         bool   `json:"with_fusion_shop_entry"`
}

type rankListRsp struct {
	RankList RankList `json:"data"`
	Extra    struct {
		Now int64 `json:"now"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}

type RankList struct {
	AnchorShowContribution bool        `json:"anchor_show_contribution"`
	Anonymous              int         `json:"anonymous"`
	Currency               string      `json:"currency"`
	Ranks                  []*RankUser `json:"ranks"`
	RuleURL                string      `json:"rule_url"`
	Total                  int         `json:"total"`
}

type RankUser struct {
	GapDescription       string    `json:"gap_description"`
	Rank                 int       `json:"rank"`  // Absolute Rank
	Score                int       `json:"score"` // Coins Count
	User                 *UserData `json:"user"`
	UserRestrictionLevel int       `json:"user_restriction_level"`
}

type PriceList struct {
	PriceList []*PriceItem `json:"data"`
	Extra     struct {
		ApplePayHintURL                    string        `json:"apple_pay_hint_url"`
		BadgeIcon                          string        `json:"badge_icon"`
		Channel                            string        `json:"channel"`
		ChannelID                          int           `json:"channel_id"`
		CurrencyList                       []string      `json:"currency_list"`
		CustomizedIds                      []int         `json:"customized_ids"`
		DefaultCurrency                    string        `json:"default_currency"`
		DefaultPacketID                    int           `json:"default_packet_id"`
		ExtraDiamondList                   []interface{} `json:"extra_diamond_list"`
		FirstChargePacketID                int           `json:"first_charge_packet_id"`
		IsDefault                          bool          `json:"is_default"`
		IsRecommend                        bool          `json:"is_recommend"`
		LargePayURL                        string        `json:"large_pay_url"`
		MaxCustomizedDiamondCnt            int           `json:"max_customized_diamond_cnt"`
		MerchantID                         string        `json:"merchant_id"`
		MinCustomizedDiamondCnt            int           `json:"min_customized_diamond_cnt"`
		NeedAuth                           int           `json:"need_auth"`
		Now                                int64         `json:"now"`
		PloyTraceID                        int           `json:"ploy_trace_id"`
		RecentlyPurchasedPacketID          int           `json:"recently_purchased_packet_id"`
		RecommendedPacketID                int           `json:"recommended_packet_id"`
		ShouldDisplayCustomizedWebRecharge bool          `json:"should_display_customized_web_recharge"`
		ShowHint                           int           `json:"show_hint"`
		SignInfos                          []interface{} `json:"sign_infos"`
		TotalSigned                        int           `json:"total_signed"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
}

type PriceItem struct {
	CouponID      string `json:"coupon_id"`
	CurrencyPrice []struct {
		Currency      string `json:"currency"`
		KeepDot       int    `json:"keep_dot"`
		OriginalPrice int    `json:"original_price"`
		Price         int    `json:"price"` // Local Currency Price Cents
		PriceDot      int    `json:"price_dot"`
		PriceShowForm string `json:"price_show_form"`
	} `json:"currency_price"`
	Describe      string `json:"describe"`
	DiamondCount  int    `json:"diamond_count"` // Coin Count
	DiscountPrice int    `json:"discount_price"`
	ExchangePrice int    `json:"exchange_price"` // Local Currency Price Cents
	GivingCount   int    `json:"giving_count"`
	IapID         string `json:"iap_id"`
	ID            int    `json:"id"`
	Price         int    `json:"price"` // USD Cents
}

type UserInfo struct {
	ID           string `json:"id"`
	ShortID      string `json:"shortId"`
	UniqueID     string `json:"uniqueId"`
	Nickname     string `json:"nickname"`
	AvatarLarger string `json:"avatarLarger"`
	AvatarMedium string `json:"avatarMedium"`
	AvatarThumb  string `json:"avatarThumb"`
	Biography    string `json:"signature"`
	CreateTime   int    `json:"createTime"`
	Verified     bool   `json:"verified"`
	SecUID       string `json:"secUid"`
	Ftc          bool   `json:"ftc"`
	Relation     int    `json:"relation"`
	OpenFavorite bool   `json:"openFavorite"`
	BioLink      struct {
		Link string `json:"link"`
		Risk int    `json:"risk"`
	} `json:"bioLink"`
	CommentSetting int    `json:"commentSetting"`
	DuetSetting    int    `json:"duetSetting"`
	StitchSetting  int    `json:"stitchSetting"`
	PrivateAccount bool   `json:"privateAccount"`
	Secret         bool   `json:"secret"`
	IsADVirtual    bool   `json:"isADVirtual"`
	RoomID         string `json:"roomId"`

	Stats UserStats
}

type UserStats struct {
	FollowerCount  int  `json:"followerCount"`
	FollowingCount int  `json:"followingCount"`
	Heart          int  `json:"heart"`
	HeartCount     int  `json:"heartCount"`
	VideoCount     int  `json:"videoCount"`
	DiggCount      int  `json:"diggCount"`
	NeedFix        bool `json:"needFix"`
}

type SignedURL struct {
	SignedURL      string `json:"signedUrl"`
	MsToken        string `json:"msToken"`
	Signature      string `json:"_signature"`
	XBogus         string `json:"X-Bogus"`
	UserAgent      string `json:"User-Agent"`
	BrowserVersion string `json:"browserVersion"`
	BrowserName    string `json:"browserName"`
	Error          string `json:"error"`
}

// DisconnectEvent sent went disconnected from live. When this event occurs no other events will be emitted and the live
// instance should be closed with `Closed`. A new track user/room should be invoked to reconnect if desired. This event
// should always be emitted.
type DisconnectEvent struct {
	created time.Time
}

func (d DisconnectEvent) IsHistory() bool {
	return false
}

func (d DisconnectEvent) CreatedTimestamp() int64 {
	return d.created.Unix()
}

type LimitInfo struct {
	Max       int       `json:"max"`
	Remaining int       `json:"remaining"`
	ResetAt   time.Time `json:"reset_at"`
}

// SigningLimits are the rates and result from the configured signer.
type SigningLimits struct {
	Code    int
	Message string
	Day     LimitInfo
	Hour    LimitInfo
	Minute  LimitInfo
}

type liveRoomContainer struct {
	LiveRoomUserInfo *LiveRoomUserInfo `json:"liveRoomUserInfo,omitempty"`
}

type LiveRoomUserInfo struct {
	LiveRoomUser *LiveRoomUser `json:"user,omitempty"`
	Stats        Stats         `json:"stats"`
	LiveRoom     *LiveRoom     `json:"liveRoom,omitempty"`
}

type LiveRoomUser struct {
	AvatarLarger string `json:"avatarLarger"`
	AvatarMedium string `json:"avatarMedium"`
	AvatarThumb  string `json:"avatarThumb"`
	ID           string `json:"id"`
	Nickname     string `json:"nickname"`
	SecUID       string `json:"secUid"`
	Secret       bool   `json:"secret"`
	UniqueID     string `json:"uniqueId"`
	Verified     bool   `json:"verified"`
	RoomID       string `json:"roomId"`
	Signature    string `json:"signature"`
	Status       int    `json:"status"`
	FollowStatus int    `json:"followStatus"`
}

type Stats struct {
	FollowingCount int `json:"followingCount"`
	FollowerCount  int `json:"followerCount"`
}

type LiveRoom struct {
	CoverURL          string        `json:"coverUrl"`
	SquareCoverImg    string        `json:"squareCoverImg"`
	Title             string        `json:"title"`
	StartTime         int64         `json:"startTime"`
	Status            int           `json:"status"`
	PaidEvent         PaidEvent     `json:"paidEvent"`
	LiveSubOnly       int           `json:"liveSubOnly"`
	LiveRoomMode      int           `json:"liveRoomMode"`
	HashTagID         int           `json:"hashTagId"`
	GameTagID         int           `json:"gameTagId"`
	LiveRoomStats     LiveRoomStats `json:"liveRoomStats"`
	StreamData        StreamData    `json:"streamData"`
	StreamID          string        `json:"streamId"`
	MultiStreamScene  int           `json:"multiStreamScene"`
	MultiStreamSource int           `json:"multiStreamSource"`
	HevcStreamData    StreamData    `json:"hevcStreamData"`
}

type PaidEvent struct {
	EventID  int `json:"event_id"`
	PaidType int `json:"paid_type"`
}

type LiveRoomStats struct {
	UserCount int `json:"userCount"`
}

type StreamData struct {
	PullData PullData `json:"pull_data"`
}

type PullData struct {
	Options    Options `json:"options"`
	StreamData string  `json:"stream_data"`
}

type Options struct {
	DefaultQuality    Quality   `json:"default_quality"`
	Qualities         []Quality `json:"qualities"`
	ShowQualityButton bool      `json:"show_quality_button"`
}

type Quality struct {
	IconType   int    `json:"icon_type"`
	Level      int    `json:"level"`
	Name       string `json:"name"`
	Resolution string `json:"resolution"`
	SdkKey     string `json:"sdk_key"`
	VCodec     string `json:"v_codec"`
}
