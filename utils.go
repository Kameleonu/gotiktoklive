package gotiktoklive

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/erni27/imcache"
	pb "github.com/steampoweredtaco/gotiktoklive/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const (
	messageHistoryTimeout = 15 * time.Minute
)

var (
	msgIDCache imcache.Cache[int64, struct{}]
)

func getRandomDeviceID() string {
	const chars = "0123456789"
	b := make([]byte, 20)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func parseMsg(msg *pb.WebcastResponse_Message, warnHandler func(...interface{}), debugHandler func(...interface{}), enableExperimentalEvents bool) (out Event, err error) {
	tReflect, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(msg.Method))
	if err != nil {
		base := base64.RawStdEncoding.EncodeToString(msg.Payload)
		debugHandler(fmt.Sprintf("cannot find type %s:\n%s ", msg.Method, base))
		return nil, nil
	}
	m := tReflect.New().Interface()
	if err = proto.Unmarshal(msg.Payload, m); err != nil {
		base := base64.RawStdEncoding.EncodeToString(msg.Payload)
		err = fmt.Errorf("failed to unmarshal proto %T: %w\n%s", m, err, base)
		debugHandler(err)
		warnHandler(fmt.Errorf("failed to unmarshal proto %T: %w", m, err))
		return nil, nil
	}
	switch pt := m.(type) {
	case *pb.RoomMessage:
		return RoomEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Type:      pt.Common.Method,
			Message:   pt.Content,
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil
	case *pb.WebcastRoomPinMessage:
		{
			tReflect, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(pt.OriginalMsgType))
			if err != nil {
				base := base64.RawStdEncoding.EncodeToString(msg.Payload)
				debugHandler("cannot find proto type for pin message %s:\n%s ", msg.Method, base)
				return RoomEvent{
					MessageID: msg.MsgId,
					Timestamp: pt.Common.CreateTime,
					Type:      pt.OriginalMsgType,
					Message:   "<unknown>",
					isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
				}, nil
			}
			m := tReflect.New().Interface()
			if err = proto.Unmarshal(pt.PinnedMessage, m); err != nil {
				base := base64.RawStdEncoding.EncodeToString(msg.Payload)
				err = fmt.Errorf("failed to unmarshal proto %T: %w\n%s", m, err, base)
				debugHandler(err)
				warnHandler(fmt.Errorf("failed to unmarshal proto %T: %w", m, err))
				return RoomEvent{
					MessageID: msg.MsgId,
					Timestamp: pt.Common.CreateTime,
					Type:      pt.OriginalMsgType,
					Message:   "<unknown>",
					isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
				}, nil
			}

			typeStr := pt.OriginalMsgType
			msgPinned := "<unknown pinned type>"
			switch pt2 := m.(type) {
			// Todo make a pin return type
			case *pb.WebcastChatMessage:
				return ChatEvent{
					MessageID: pt.Common.MsgId,
					Timestamp: pt.Common.CreateTime,
					Comment:   "<pinned>: " + pt2.Content,
					User:      toUser(pt2.User),
					isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
				}, nil
			default:
				base := base64.RawStdEncoding.EncodeToString(pt.PinnedMessage)
				err = fmt.Errorf("unimplemented pinned message type %T\n%s", m, base)
				debugHandler(err)
				warnHandler(fmt.Sprintf("unimplemented pinned message type %T", m))

			}
			return RoomEvent{
				MessageID: pt.Common.MsgId,
				Timestamp: pt.Common.CreateTime,
				Type:      typeStr,
				Message:   msgPinned,
				isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
			}, nil
		}
	case *pb.WebcastChatMessage:
		return ChatEvent{
			MessageID:    pt.Common.MsgId,
			Comment:      pt.Content,
			User:         toUser(pt.User),
			UserIdentity: toUserIdentity(pt.UserIdentity),
			Timestamp:    pt.Common.CreateTime,
			isHistory:    msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil
	case *pb.WebcastMemberMessage:
		return UserEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Event:     toUserType(pt.Action.String()),
			User:      toUser(pt.User),
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil
	case *pb.WebcastLiveGameIntroMessage:
		return RoomEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Type:      pt.Common.Method,
			Message:   pt.GameText.DefaultPattern,
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil
	case *pb.WebcastRoomMessage:
		return RoomEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Type:      pt.Common.Method,
			// TODO: Make this actually use pieces list and fill out the format text correctly.
			Message:   pt.Common.DisplayText.DefaultPattern,
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil
	case *pb.WebcastRoomUserSeqMessage:
		return ViewersEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Viewers:   int(pt.Total),
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil
	case *pb.WebcastSocialMessage:
		return UserEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Event:     toUserType(pt.Common.DisplayText.Key),
			User:      toUser(pt.User),
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil
	case *pb.WebcastGiftMessage:
		if pt.GiftId == 0 && pt.User == nil {
			return nil, nil
		}

		return GiftEvent{
			MessageID:    pt.Common.MsgId,
			Timestamp:    pt.Common.CreateTime,
			ID:           pt.GiftId,
			GroupID:      pt.GroupId,
			Name:         pt.Gift.Name,
			Describe:     pt.Gift.Describe,
			Diamonds:     int(pt.Gift.DiamondCount),
			RepeatCount:  int(pt.RepeatCount),
			RepeatEnd:    pt.RepeatEnd == 1,
			Type:         int(pt.Gift.Type),
			ToUserID:     int64(pt.UserGiftReciever.UserId),
			User:         toUser(pt.User),
			UserIdentity: toUserIdentity(pt.UserIdentity),
			isHistory:    msg.IsHistory || cachedHistory(pt.Common.MsgId),
			IsComboGift:  pt.GroupId != 0,
		}, nil
	case *pb.WebcastLikeMessage:
		return LikeEvent{
			MessageID:   pt.Common.MsgId,
			Timestamp:   pt.Common.CreateTime,
			Likes:       int(pt.Count),
			TotalLikes:  int(pt.Total),
			User:        toUser(pt.User),
			DisplayType: pt.Common.Method,
			Label:       pt.Common.DisplayText.String(),
			isHistory:   msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil

	case *pb.WebcastQuestionNewMessage:
		return QuestionEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Quesion:   pt.Details.Text,
			User:      toUser(pt.Details.User),
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil

	case *pb.WebcastControlMessage:
		return ControlEvent{
			MessageID:   pt.Common.MsgId,
			Timestamp:   pt.Common.CreateTime,
			Action:      int(pt.Action),
			Description: pt.Action.String(),
			isHistory:   msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil

	case *pb.WebcastLinkMicBattle:
		users := []*User{}
		for _, u := range pt.HostTeam {
			groups := u.HostGroup
			for _, group := range groups {
				for _, user := range group.Host {
					// urls := make([]string, 5)
					// for _, img := range user.Images {
					// 	urls = append(urls, img.UrlList...)
					// }
					users = append(users, &User{
						ID:       int64(user.Id),
						Username: user.ProfileId,
						Nickname: user.Name,
						// ProfilePicture: &ProfilePicture{
						// 	Urls: urls,
						// },
					})

				}

			}
		}
		return MicBattleEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Users:     users,
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil

	case *pb.WebcastLinkMicArmies:
		battles := []*Battle{}
		for _, b := range pt.BattleItems {
			battle := &Battle{
				Host:   int64(b.HostUserId),
				Groups: []*BattleGroup{},
			}
			for _, g := range b.BattleGroups {
				group := BattleGroup{
					Points: int(g.Points),
					Users:  []*User{},
				}
				for _, u := range g.Users {
					group.Users = append(group.Users, toUser(u))
				}
				battle.Groups = append(battle.Groups, &group)
			}
			battles = append(battles, battle)
		}
		return BattlesEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			Status:    int(pt.BattleStatus),
			Battles:   battles,
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil
	case *pb.WebcastLiveIntroMessage:
		return IntroEvent{
			MessageID: pt.Common.MsgId,
			Timestamp: pt.Common.CreateTime,
			ID:        int(pt.RoomId),
			Title:     pt.Content,
			User:      toUser(pt.Host),
			isHistory: msg.IsHistory || cachedHistory(pt.Common.MsgId),
		}, nil

	case *pb.WebcastInRoomBannerMessage:
		var data interface{}
		// TODO: should we make a type for this instead of unmarshalling to see it is an error then feeding it up?
		err = json.Unmarshal([]byte(pt.GetJson()), &data)
		if err != nil {
			return nil, fmt.Errorf("WebcastInRoomBannerMessage: %w\n%s", err, data)
		}

		return RoomBannerEvent{
			MessageID: pt.Header.MsgId,
			Timestamp: pt.Header.CreateTime,
			Data:      data,
			isHistory: msg.IsHistory || cachedHistory(pt.Header.MsgId),
		}, nil

	default:
		base := base64.RawStdEncoding.EncodeToString(msg.Payload)
		err = fmt.Errorf("unimplemented type %T\n%s", m, base)
		debugHandler(err)
		warnHandler(fmt.Sprintf("unimplemented type %T", m))
		return nil, nil
	}
}

func toProfilePicture(pic *pb.Image) *ProfilePicture {
	if pic != nil && pic.UrlList != nil {
		return &ProfilePicture{
			Urls: pic.UrlList,
		}
	}
	return nil
}

func cachedHistory(id int64) bool {
	_, present := msgIDCache.GetOrSet(id, struct{}{}, imcache.WithExpiration(messageHistoryTimeout))
	return present
}

func defaultLogHandler(i ...interface{}) {
	slog.Debug(fmt.Sprint(i...), "logger", "gotiktoklive-default")
}

func routineErrHandler(err ...interface{}) {
	slog.Debug(fmt.Sprint(err...), "logger", "gotiktoklive-default")
}

func toUser(u *pb.User) *User {
	if u == nil {
		return &User{}
	}
	username := u.IdStr
	if u.IdStr == "" {
		username = u.Nickname
	}

	user := User{
		ID:          int64(u.Id),
		Username:    username,
		Nickname:    u.Nickname,
		AvatarThumb: toProfilePicture(u.AvatarThumb),
	}

	user.ExtraAttributes = &ExtraAttributes{
		FollowRole: int(u.UserRole),
	}

	if u.BadgeList != nil {
		var badges []*UserBadge
		for _, badge := range u.BadgeList {
			badges = append(badges, &UserBadge{
				Type: badge.DisplayType.String(),
				Name: badge.String(),
			})
		}
		user.Badge = &BadgeAttributes{
			Badges: badges,
		}
	}
	return &user
}

func toUserIdentity(uid *pb.UserIdentity) *UserIdentity {
	if uid == nil {
		return nil
	}
	return &UserIdentity{
		IsGiftGiver:       uid.IsGiftGiverOfAnchor,
		IsSubscriber:      uid.IsSubscriberOfAnchor,
		IsMutualFollowing: uid.IsMutualFollowingWithAnchor,
		IsFollower:        uid.IsFollowerOfAnchor,
		IsModerator:       uid.IsModeratorOfAnchor,
		IsAnchor:          uid.IsAnchor,
	}
}

func copyMap(m map[string]string) map[string]string {
	out := make(map[string]string)
	for key, value := range m {
		out[key] = value
	}
	return out
}

func toUserType(displayType string) userEventType {
	switch displayType {
	case "pm_main_follow_message_viewer_2":
		return USER_FOLLOW
	case "pm_mt_guidance_share":
		return USER_SHARE
	case "live_room_enter_toast":
		return USER_JOIN
	case "JOINED":
		return USER_JOIN
	}
	return userEventType(fmt.Sprintf("User type not implemented, please report: %s", displayType))
}
