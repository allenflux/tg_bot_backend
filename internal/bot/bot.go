package bot

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"tg_bot_backend/internal/consts"
	"tg_bot_backend/internal/dao"
	"tg_bot_backend/internal/model/entity"
	"tg_bot_backend/utility/platform"
	"time"
)

var AwesomeBotApiChan = make(chan *tgbotapi.BotAPI, 100)

type GroupPayload struct {
	ID           string `json:"id"`
	Link         string `json:"link"`
	MembersCount int    `json:"members_count"`
	Title        string `json:"title"`
}

var AwesomeGroupChan = make(chan *GroupPayload, 1000)

func InitBotApiChanFromMysql(ctx context.Context, payload chan<- *tgbotapi.BotAPI) {
	dbQuery := dao.Bot.Ctx(ctx).
		Order("id desc")
	var bots []entity.Bot
	var totalCount int
	if err := dbQuery.ScanAndCount(&bots, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to query and count bots list: %v", err)
	}
	for _, bot := range bots {
		myBot, err := tgbotapi.NewBotAPI(bot.BotToken)
		if err != nil {
			g.Log().Errorf(ctx, "Failed to init bot: %v", err)
			continue
		}
		myBot.Debug = true
		payload <- myBot
	}
}

func MakeBotApiClientPipLine(ctx context.Context, payload <-chan *tgbotapi.BotAPI) {
	for {
		select {
		case bot, ok := <-payload:
			if !ok {
				g.Log().Error(ctx, "payload channel closed")
				return
			}
			go Program(ctx, bot)
		case <-ctx.Done():
			g.Log().Info(ctx, "BotApiClientPipLine closed")
			return
		}
	}
}

type CommandSession struct {
	Command string
	Step    int
	Answers []string
	Expires time.Time
}

var userSessions = make(map[string]*CommandSession)

func Program(ctx context.Context, bot *tgbotapi.BotAPI) {
	g.Log().Infof(ctx, "Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {

		// check bot status
		if ok, err := dao.Bot.Ctx(ctx).
			Where("account = ?", bot.Self.ID).
			Where("status = ?", consts.BotStatusUnAvailable).
			Exist(); err != nil {
			g.Log().Errorf(ctx, "Failed to query bot: %v", err)
			continue
		} else if ok {
			g.Log().Infof(ctx, "BotStatusUnAvailable  bot %s", bot.Self.UserName)
			continue
		}

		if update.Message != nil { // If we got a message
			g.Log().Infof(ctx, "[%s] %s", update.Message.From.UserName, update.Message.Text)
			g.Log().Infof(ctx, "From ID = [%d]", update.Message.From.ID)
			g.Log().Infof(ctx, "chatTyping = [%s]", update.Message.Chat.Type)
			g.Log().Infof(ctx, "chatTitle = [%s]", update.Message.Chat.Title)

			// 欢迎语
			go handleNewMembers(ctx, bot, update)

			// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
			chatTyping := update.Message.Chat.Type
			if chatTyping == "supergroup" {
				chatTyping = "group"
			}
			switch chatTyping {
			case "group":
				chatIdString := strconv.FormatInt(update.Message.Chat.ID, 10)
				if ok, err := g.Redis().Exists(ctx, chatIdString); err != nil {
					g.Log().Errorf(ctx, "Failed to check if [%s] exists: %v", chatIdString, err)
				} else if ok == 0 {
					g.Log().Infof(ctx, "[%s] exists: %v", chatIdString, ok)
					chatConfig := tgbotapi.ChatConfig{
						ChatID: update.Message.Chat.ID,
					}
					link, err := bot.GetInviteLink(tgbotapi.ChatInviteLinkConfig{
						ChatConfig: chatConfig,
					})
					g.Log().Infof(ctx, "Link = %s", link)
					if err != nil {
						g.Log().Errorf(ctx, "Failed to get invite link: %v", err)
						continue
					}
					membersCount, err := bot.GetChatMembersCount(tgbotapi.ChatMemberCountConfig{
						ChatConfig: chatConfig,
					})
					g.Log().Infof(ctx, "MembersCount = %d", membersCount)
					if err != nil {
						g.Log().Errorf(ctx, "Failed to get members count: %v", err)
						continue
					}

					groupData := GroupPayload{
						ID:           chatIdString,
						Link:         link,
						MembersCount: membersCount,
						Title:        update.Message.Chat.Title,
					}
					g.Log().Infof(ctx, "GroupData = %+v", groupData)
					AwesomeGroupChan <- &groupData
				} else {
					g.Log().Infof(ctx, "[%s] exists , will continue: %v", chatIdString, ok)
				}
			}
			go handleUpdate(ctx, bot, update)
		}
	}
}

func sessionKey(botID, userID int64, chatID int64) string {
	return fmt.Sprintf("%d:%d:%d", botID, userID, chatID)
}

var validAnswersOfGroupType = map[string]bool{
	"1": true,
	"2": true,
}

func handleUpdate(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := update.Message
	chat := msg.Chat
	user := msg.From
	key := sessionKey(bot.Self.ID, user.ID, chat.ID)

	if chat.IsPrivate() {
		return
		//处理私聊命令
		//if msg.IsCommand() {
		//	cmd := msg.Command()
		//	switch cmd {
		//	// 群发所有内容到客户群
		//	case consts.BotCmdQF1:
		//		userSessions[key] = &CommandSession{
		//			Command: consts.BotCmdQF1,
		//			Step:    1,
		//			Answers: []string{},
		//			Expires: time.Now().Add(5 * time.Minute),
		//		}
		//		sendNextQf1Question(bot, chat.ID, key)
		//	//	群发所有内容到渠道群
		//	case consts.BotCmdQF2:
		//		userSessions[key] = &CommandSession{
		//			Command: consts.BotCmdQF2,
		//			Step:    1,
		//			Answers: []string{},
		//			Expires: time.Now().Add(5 * time.Minute),
		//		}
		//		sendNextQf1Question(bot, chat.ID, key)
		//	}
		//}
	} else {
		// 处理命令（只接受带 bot username 的 /bind@BotName）
		if msg.IsCommand() {
			cmd := msg.Command()
			cmdWithAt := msg.CommandWithAt()

			// 忽略不带 @bot 的命令
			if !strings.Contains(cmdWithAt, "@") || !strings.HasSuffix(cmdWithAt, "@"+bot.Self.UserName) {
				return
			}

			switch cmd {
			case consts.BotCmdUnbind:
				ok, newMsg := checkUserPermission(ctx, chat.ID, user.ID, consts.BotCmdUnbind)
				if !ok {
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					return
				}
				// 获取更新之前的platform id
				prePlatformId, _ := GetPrePlatformId(ctx, chat.ID)
				//	直接更新DB信息
				_, err := dao.Group.Ctx(ctx).Data(
					g.Map{
						"type":               0,
						"central_control_id": 0,
					}).
					Where("group_chat_id = ?", chat.ID).
					Update()
				if err != nil {
					g.Log().Errorf(ctx, "Failed to update group chat: %v", err)
					newMsg = fmt.Sprintf("Failed to update group chat: %v", err)
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					return
				}
				UpdateDbCentralControl(ctx, prePlatformId)
				UpdateDbCentralControl(ctx, "0")
				g.Log().Infof(ctx, "Updated group chat with id [%d]", chat.ID)
				newMsg = fmt.Sprintf("✅ 解绑成功 Updated group chat with id [%d]", chat.ID)
				bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))

			case consts.BotCmdBind:

				ok, newMsg := checkUserPermission(ctx, chat.ID, user.ID, consts.BotCmdBind)
				if !ok {
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					return
				}

				// 根据用户ID检查用户权限
				userSessions[key] = &CommandSession{
					Command: consts.BotCmdBind,
					Step:    1,
					Answers: []string{},
					Expires: time.Now().Add(5 * time.Minute),
				}

				sendNextQuestion(bot, chat.ID, key)
			case consts.BotCmdTopUp:
				ok, newMsg := checkUserPermission(ctx, chat.ID, user.ID, consts.BotCmdTopUp)
				if !ok {
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					return
				}

				userSessions[key] = &CommandSession{
					Command: consts.BotCmdTopUp,
					Step:    1,
					Answers: []string{},
					Expires: time.Now().Add(5 * time.Minute),
				}

				sendTopUpNextQuestion(bot, chat.ID, key)

			case "cancel":
				delete(userSessions, key)
				bot.Send(tgbotapi.NewMessage(chat.ID, "❌ 绑定已取消"))
			default:
				bot.Send(tgbotapi.NewMessage(chat.ID, "未知命令"))
			}
			return
		}
	}
	// 非命令：检查是否在会话中
	if session, ok := userSessions[key]; ok {
		if time.Now().After(session.Expires) {
			delete(userSessions, key)
			bot.Send(tgbotapi.NewMessage(chat.ID, "⏰ 会话超时，请重新输入 /bind@"+bot.Self.UserName))
			return
		}
		// 命令判定
		switch session.Command {
		//case consts.BotCmdQF1:
		//	session.Answers = append(session.Answers, msg.Text)
		//	session.Step++
		//	if session.Step <= 1 {
		//		sendNextQf1Question(bot, chat.ID, key)
		//	} else {
		//		var groups []entity.Group
		//		var totalCount int
		//		err := dao.Group.Ctx(ctx).Where("type = ?", consts.GroupTypeForCustomer).ScanAndCount(groups, &totalCount, false)
		//	}
		case consts.BotCmdTopUp:
			session.Answers = append(session.Answers, msg.Text)
			session.Step++
			if session.Step <= 2 {
				sendTopUpNextQuestion(bot, chat.ID, key)
			} else {
				//	valid amount
				amount, err := strconv.Atoi(session.Answers[1])
				if err != nil {
					g.Log().Infof(ctx, "[%s] amount Conver Error ", err.Error())
					newMsg := fmt.Sprintf("❌ 充值失败 [%s] Please make sure the amount you enter is a number %s ", session.Answers[1], err.Error())
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					delete(userSessions, key)
					return
				}

				if amount <= 0 {
					g.Log().Infof(ctx, " Please make sure the amount is an integer greater than 0 ")
					newMsg := fmt.Sprintf("❌ 充值失败 [%s] Please make sure the amount is an integer greater than 0 ", session.Answers[1])
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					delete(userSessions, key)
					return
				}

				//	Get group info
				var group entity.Group
				var totalCount int
				if err = dao.Group.Ctx(ctx).
					Where("group_chat_id = ?", chat.ID).
					Where("type = ?", consts.GroupTypeForCustomer).
					ScanAndCount(&group, &totalCount, false); err != nil {
					g.Log().Infof(ctx, "[%s] DB Scan Group Error", session.Answers[0])
					newMsg := fmt.Sprintf("❌ 充值失败 [%s] DB Scan Group Error %s ", session.Answers[1], err.Error())
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					delete(userSessions, key)
					return
				}
				if totalCount == 0 {
					g.Log().Infof(ctx, "[%s] DB Scan Group Success But totalCount == 0", session.Answers[0])
					newMsg := fmt.Sprintf("❌ 充值失败 [%s], 请确定当前群组是客户群 DB Scan Group Success But totalCount == 0 ", session.Answers[1])
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					delete(userSessions, key)
					return
				}

				// get platform info
				var PlatformInfo entity.CentralControl
				if err = dao.CentralControl.Ctx(ctx).
					Where("id = ?", group.CentralControlId).
					Scan(&PlatformInfo); err != nil {
					g.Log().Infof(ctx, "[%s] DB Scan CentralControl Error", session.Answers[0])
					newMsg := fmt.Sprintf("❌ 充值失败 [%s] DB Scan CentralControl Error %s ", session.Answers[1], err.Error())
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					delete(userSessions, key)
					return
				}

				apiToken, err := platform.GetPlatformToken(ctx, PlatformInfo.Domain, PlatformInfo.ApiUsername, PlatformInfo.SecretKey)

				if err != nil {
					g.Log().Infof(ctx, "[%s] Api Error ", err.Error())
					newMsg := fmt.Sprintf("❌ 充值失败 [%s] Api Error %s ", session.Answers[0], err.Error())
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					delete(userSessions, key)
					return
				}

				if ok, err = platform.AddCustomerFind(ctx, PlatformInfo.Domain, apiToken, group.CustomerId, amount); err != nil {
					g.Log().Errorf(ctx, "[%s] AddCustomerFind Error, Error %s ", session.Answers[1], err.Error())
					newMsg := fmt.Sprintf("❌ 充值失败 [%s] AddCustomerFind Errorr, Error %s ", session.Answers[1], err.Error())
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					delete(userSessions, key)
					return
				} else if ok {
					g.Log().Infof(ctx, "[%s] AddCustomerFind  Success", session.Answers[1])
					newMsg := fmt.Sprintf(" ✅ 充值成功 [%s] AddCustomerFind Successful ", session.Answers[1])
					bot.Send(tgbotapi.NewMessage(chat.ID, newMsg))
					delete(userSessions, key)
					return
				}

			}

		case consts.BotCmdBind:
			// 收集用户输入
			session.Answers = append(session.Answers, msg.Text)
			session.Step++
			result := ""
			bindSuccessful := true
			if session.Step <= 3 {
				sendNextQuestion(bot, chat.ID, key)
			} else {
				result = fmt.Sprintf("✅ 绑定成功：\n中控平台ID：%s\n群类型：%s",
					session.Answers[0], session.Answers[1])
				// Verify Group Type
				if _, ok := validAnswersOfGroupType[session.Answers[1]]; !ok {
					g.Log().Errorf(ctx, "[%s] Invalid group type", session.Answers[1])
					result = fmt.Sprintf("❌ 绑定失败 [%s] Invalid group type", session.Answers[1])
					bindSuccessful = false
				}
				// Verify And Action By DB
				if has, err := dao.CentralControl.Ctx(ctx).Where("id = ?", session.Answers[0]).Where("status = ?", consts.CentralControlStatusAvailable).Exist(); err != nil {
					g.Log().Errorf(ctx, "Failed to check if [%s] exists or Status is Available: %v", session.Answers[0], err)
					result = fmt.Sprintf("❌ 绑定失败 Failed to check if [%s] exists: %v", session.Answers[0], err)
					bindSuccessful = false
				} else if !has {
					g.Log().Infof(ctx, "[%s] not exists", session.Answers[0])
					result = fmt.Sprintf("❌ 绑定失败 [%s] not exists or Status is UnAvailable", session.Answers[0])
					bindSuccessful = false
				}

				// get platform info
				var PlatformInfo entity.CentralControl
				if err := dao.CentralControl.Ctx(ctx).
					Where("id = ?", session.Answers[0]).
					Scan(&PlatformInfo); err != nil {
					g.Log().Infof(ctx, "[%s] DB Scan CentralControl Error", session.Answers[0])
					result = fmt.Sprintf("❌ 绑定失败 [%s] DB Scan CentralControl Error", session.Answers[0])
					bindSuccessful = false
				}

				userInputId, err := strconv.Atoi(session.Answers[2])
				if err != nil {
					g.Log().Infof(ctx, "[%s] userInputId Conver Error ", err.Error())
					result = fmt.Sprintf("❌ 绑定失败 [%s] userInputId %s ", session.Answers[2], err.Error())
					bindSuccessful = false
				}

				if userInputId < 0 {
					g.Log().Infof(ctx, "[%s] userInputId Error ", session.Answers[2])
					result = fmt.Sprintf("❌ 绑定失败 输入userInputId不能小于0 [%s] userInputId  ", session.Answers[2])
					bindSuccessful = false
				}

				apiToken, err := platform.GetPlatformToken(ctx, PlatformInfo.Domain, PlatformInfo.ApiUsername, PlatformInfo.SecretKey)

				if err != nil {
					g.Log().Infof(ctx, "[%s] Api Error ", err.Error())
					result = fmt.Sprintf("❌ 绑定失败 [%s] Api Error %s ", session.Answers[0], err.Error())
					bindSuccessful = false
				}

				verifySuccessful := true
				if bindSuccessful {
					// 获取更新之前的platform id
					prePlatformId, _ := GetPrePlatformId(ctx, chat.ID)

					data := g.Map{
						"central_control_id": session.Answers[0],
						"type":               session.Answers[1],
					}
					switch cmd, _ := strconv.Atoi(session.Answers[1]); cmd {
					case consts.GroupTypeForCustomer:
						if ok, err = platform.VerifyCustomerId(ctx, PlatformInfo.Domain, apiToken, userInputId); err != nil {
							verifySuccessful = false
							g.Log().Errorf(ctx, "[%s] Verification Customer Error, Error %s ", session.Answers[2], err.Error())
							result = fmt.Sprintf("❌ 绑定失败 [%s] Verification Customer Error, Error %s ", session.Answers[2], err.Error())
						} else if ok {
							g.Log().Infof(ctx, "[%s] Customer Verification  Success", session.Answers[2])
							data["customer_id"] = session.Answers[2]
							data["business_id"] = 0
						}
					case consts.GroupTypeForBusiness:
						if ok, err = platform.VerifyBusinessId(ctx, PlatformInfo.Domain, apiToken, strconv.Itoa(userInputId)); err != nil {
							verifySuccessful = false
							g.Log().Errorf(ctx, "[%s] Verification Business Error, Error %s ", session.Answers[2], err.Error())
							result = fmt.Sprintf("❌ 绑定失败 [%s] Verification Business Error, Error %s ", session.Answers[2], err.Error())
						} else if ok {
							g.Log().Infof(ctx, "[%s] Business Verification  Success", session.Answers[2])
							data["customer_id"] = 0
							data["business_id"] = session.Answers[2]
						}
					default:
						panic("unhandled default case")

					}
					if verifySuccessful {

						_, err := dao.Group.Ctx(ctx).Where("group_chat_id = ?", chat.ID).Data(
							data,
						).Update()
						if err != nil {
							g.Log().Errorf(ctx, "Failed to update group")
							result = fmt.Sprintf("Failed to update group [%s] %s", session.Answers[0], err.Error())
						}
						UpdateDbCentralControl(ctx, prePlatformId)
						UpdateDbCentralControl(ctx, session.Answers[0])
					}
				}
				bot.Send(tgbotapi.NewMessage(chat.ID, result))
				delete(userSessions, key)
			}
		}

	}
}

func sendNextQuestion(bot *tgbotapi.BotAPI, chatID int64, key string) {
	session := userSessions[key]
	var question string

	switch session.Step {
	case 1:
		question = "请输入中控平台ID："
	case 2:
		question = "请输入群组类型(客户群输入->1 or 渠道群输入->2)："
	case 3:
		question = "请输入客户或者渠道ID"
	}

	if question != "" {
		bot.Send(tgbotapi.NewMessage(chatID, question))
	}
}

func sendNextQf1Question(bot *tgbotapi.BotAPI, chatID int64, key string) {
	session := userSessions[key]
	var question string

	switch session.Step {
	case 1:
		question = "请输入群发内容(改内容将会发送至所有已绑定'客户'身份的群聊)："
	}

	if question != "" {
		bot.Send(tgbotapi.NewMessage(chatID, question))
	}
}

func sendNextQf2Question(bot *tgbotapi.BotAPI, chatID int64, key string) {
	session := userSessions[key]
	var question string

	switch session.Step {
	case 1:
		question = "请输入群发内容(改内容将会发送至所有已绑定'渠道'身份的群聊)："
	}

	if question != "" {
		bot.Send(tgbotapi.NewMessage(chatID, question))
	}
}

func sendTopUpNextQuestion(bot *tgbotapi.BotAPI, chatID int64, key string) {
	session := userSessions[key]
	var question string

	switch session.Step {
	case 1:
		question = "请输入汇款银行名称:"
	case 2:
		question = "请输入加款金额:"
	}

	if question != "" {
		bot.Send(tgbotapi.NewMessage(chatID, question))
	}
}

func MakeTgGroupPipLine(ctx context.Context, payload <-chan *GroupPayload) {
	for {
		select {
		case data, ok := <-payload:
			if !ok {
				g.Log().Error(ctx, "payload channel closed")
				return
			}
			TgGroupProgram(ctx, data)
		case <-ctx.Done():
			g.Log().Info(ctx, "MakeTgGroupPipLine closed")
			return
		}
	}
}

type TgUser struct {
	FirstName string `json:"first_name"`
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
}

func TgGroupProgram(ctx context.Context, data *GroupPayload) {
	//	Check Mysql
	if ok, err := dao.Group.Ctx(ctx).Where("group_chat_id = ?", data.ID).Exist(); err != nil {
		g.Log().Errorf(ctx, "Failed to check if [%s] exists: %v", data.ID, err)
	} else if ok {
		g.Log().Infof(ctx, "[%s] exists", data.ID)
		if _, err = g.Redis().Set(ctx, data.ID, "true"); err != nil {
			g.Log().Errorf(ctx, "Failed to set true: %v", err)
		}
		return
	}
	//	Get Group Members Info
	value, err := g.Cfg().Get(ctx, "tg_bot_assistant.address")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get tg_bot_assistant.address: %v", err)
		return
	}
	fullLink := value.String() + url.QueryEscape(data.Link)
	g.Log().Infof(ctx, "tg_bot_assistant address : [%s]", fullLink)
	resp, err := http.Get(fullLink)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get http resp [%s]: %v", fullLink, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to read http resp [%s]: %v", fullLink, err)
		return
	}

	kind, out, err := parseAPIResponseFlexible(ctx, body)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to parse http resp [%s]: %v", fullLink, err)
		return
	}
	switch kind {
	case paresKindInvalid:
		g.Log().Errorf(ctx, "Invalid kind: %s", kind)
	case paresKindError:
		g.Log().Errorf(ctx, "Invalid kind: %s", kind)
		g.Log().Infof(ctx, "Invalid out: %s", out)
	case paresKindUnknown:
		g.Log().Errorf(ctx, "Unknown kind: %s", kind)
		g.Log().Infof(ctx, "Invalid out: %s", out)
	case paresKindUserArray:
		g.Log().Infof(ctx, "Normal kind: %s", kind)
		var tgUsers []TgUser
		if err := json.Unmarshal(body, &tgUsers); err != nil {
			g.Log().Warning(ctx, "Failed to parse user array structure:", err)
			return
		}
		err = SaveGroupAndUsers(ctx, data, tgUsers)
		if err != nil {
			g.Log().Error(ctx, "Failed to save group and users:", err)
			return
		}
		if _, err = g.Redis().Set(ctx, data.ID, "true"); err != nil {
			g.Log().Errorf(ctx, "Failed to set true: %v", err)
		}
	}

}

const (
	paresKindInvalid   = "invalid"
	paresKindError     = "error"
	paresKindUnknown   = "unknown"
	paresKindUserArray = "user_array"
)

func parseAPIResponseFlexible(ctx context.Context, data []byte) (kind string, output string, err error) {
	var result any
	if err := json.Unmarshal(data, &result); err != nil {
		g.Log().Error(ctx, "Failed to unmarshal JSON:", err)
		return paresKindInvalid, "", err
	}

	switch v := result.(type) {
	case map[string]any:
		// Check if it's an error structure
		if errMsg, ok := v["error"]; ok {
			g.Log().Info(ctx, "Detected error structure.")
			out, _ := json.MarshalIndent(map[string]any{"error": errMsg}, "", "  ")
			return paresKindError, string(out), nil
		}
		g.Log().Warning(ctx, "Unknown map structure detected.")
		out, _ := json.MarshalIndent(v, "", "  ")
		return paresKindUnknown, string(out), nil

	case []any:
		// Assume it's a user array structure
		g.Log().Info(ctx, "Detected user array structure.")
		out, _ := json.MarshalIndent(v, "", "  ")
		return paresKindUserArray, string(out), nil

	default:
		g.Log().Warning(ctx, "Unrecognized JSON root type.")
		return paresKindUnknown, "", nil
	}
}

// SaveGroupAndUsers inserts group info and unique users into the database.
func SaveGroupAndUsers(ctx context.Context, data *GroupPayload, tgUsers []TgUser) error {
	// Start transaction
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// Check if group already exists
		exists, err := dao.Group.Ctx(ctx).Where(g.Map{
			"group_chat_id": data.ID,
		}).One()
		if err != nil {
			g.Log().Errorf(ctx, "Failed to check group existence: %v", err)
			return err
		}

		var groupId int64
		if exists != nil {
			groupId = exists["id"].Int64()
			g.Log().Infof(ctx, "Group already exists, ID: %d", groupId)
		} else {
			// Insert new group
			groupId, err = dao.Group.Ctx(ctx).Data(entity.Group{
				Name:             data.Title,
				CentralControlId: 0,
				TgLink:           data.Link,
				Type:             0,
				Size:             data.MembersCount,
				BotSize:          0,
				RoleSize:         0,
				GroupChatId:      data.ID,
			}).InsertAndGetId()
			if err != nil {
				g.Log().Errorf(ctx, "Failed to insert group: %v", err)
				return err
			}
			g.Log().Infof(ctx, "Inserted new group, ID: %d", groupId)
		}

		// Fetch existing users by TgId + GroupId
		tgIds := make([]string, 0, len(tgUsers))
		for _, user := range tgUsers {
			tgIds = append(tgIds, strconv.FormatInt(user.ID, 10))
		}

		// Get already inserted users
		existingMap := make(map[string]struct{})
		if len(tgIds) > 0 {
			list, err := dao.TgUsers.Ctx(ctx).Where("tg_id IN (?) AND group_id = ?", tgIds, groupId).Fields("tg_id").Array()
			if err != nil {
				g.Log().Errorf(ctx, "Failed to query existing users: %v", err)
				return err
			}
			for _, v := range list {
				existingMap[v.String()] = struct{}{}
			}
		}

		// Filter and insert users
		for _, user := range tgUsers {
			tgIdStr := strconv.FormatInt(user.ID, 10)
			if _, exists := existingMap[tgIdStr]; exists {
				g.Log().Infof(ctx, "User already exists: %s", tgIdStr)
				continue
			}

			isBot := 0
			if user.IsBot {
				isBot = 1
			}

			_, err := dao.TgUsers.Ctx(ctx).Data(entity.TgUsers{
				TgAccount: tgIdStr,
				GroupId:   int(groupId),
				TgName:    user.Username,
				RoleId:    0,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				IsBot:     isBot,
				Phone:     user.Phone,
				TgId:      tgIdStr,
			}).Insert()
			if err != nil {
				g.Log().Errorf(ctx, "Failed to insert user %s: %v", tgIdStr, err)
				return err
			}
			g.Log().Infof(ctx, "Inserted user: %s (%s)", tgIdStr, user.Username)
		}

		return nil
	})
}

func JsonArrayContains(jsonStr, target string) bool {
	var list []string
	err := json.Unmarshal([]byte(jsonStr), &list)
	if err != nil {
		return false // 如果解析失败，认为不包含
	}

	for _, item := range list {
		if item == target {
			return true
		}
	}

	return false
}

// 检查用户是否拥有某个命令的权限，比如 "bind"
func checkUserPermission(ctx context.Context, chatID int64, userID int64, command string) (bool, string) {
	// 获取Group ID
	var groups []entity.Group
	var totalCount int
	if err := dao.Group.Ctx(ctx).Where("group_chat_id = ?", chatID).ScanAndCount(&groups, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "查询群组失败: %v", err)
		return false, fmt.Sprintf("查询群组失败: %v", err)
	}
	if totalCount == 0 {
		return false, "群组不存在"
	}
	groupId := groups[0].Id

	// 获取用户在该群组的角色
	var tgUsers []entity.TgUsers
	if err := dao.TgUsers.Ctx(ctx).Where("group_id = ?", groupId).Where("tg_id = ?", userID).ScanAndCount(&tgUsers, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "查询群用户失败: %v", err)
		return false, fmt.Sprintf("查询群用户失败: %v", err)
	}
	if totalCount == 0 {
		return false, "无效的用户"
	}
	if tgUsers[0].RoleId == 0 {
		return false, "未被授权的用户"
	}

	// 获取角色权限
	var roles []entity.Role
	if err := dao.Role.Ctx(ctx).Where("id = ?", tgUsers[0].RoleId).ScanAndCount(&roles, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "查询角色失败: %v", err)
		return false, fmt.Sprintf("查询角色失败: %v", err)
	}
	if totalCount == 0 {
		return false, "无效的Role信息"
	}

	// 判断是否包含对应权限
	if !JsonArrayContains(roles[0].Cmd, command) {
		return false, fmt.Sprintf("没有%s权限", command)
	}

	return true, ""
}

func UpdateDbCentralControl(ctx context.Context, centralControlId string) {
	businessCount, err := dao.Group.Ctx(ctx).
		Where("central_control_id = ?", centralControlId).
		Where("type = ?", consts.GroupTypeForBusiness).
		Count()
	if err != nil {
		g.Log().Errorf(ctx, "Count Group GroupTypeForBusiness Error: %v", err)
	}
	customerCount, err := dao.Group.Ctx(ctx).
		Where("central_control_id = ?", centralControlId).
		Where("type = ?", consts.GroupTypeForCustomer).
		Count()
	if err != nil {
		g.Log().Errorf(ctx, "Count Group GroupTypeForCustomer Error: %v", err)
	}
	_, err = dao.CentralControl.Ctx(ctx).Data(
		g.Map{
			"number_of_customers": customerCount,
			"number_of_business":  businessCount,
		}).
		Where("id = ?", centralControlId).
		Update()
	if err != nil {
		g.Log().Errorf(ctx, "Update CentralControl Error: %v", err)
	}

}

func GetPrePlatformId(ctx context.Context, chatId int64) (string, error) {
	var groups []entity.Group
	var totalCount int
	prePlatformId := 0

	if err := dao.Group.Ctx(ctx).Where("group_chat_id = ?", chatId).ScanAndCount(&groups, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Failed to get group count: %v", err)
		return "0", err
	}

	if totalCount != 0 {
		prePlatformId = groups[0].CentralControlId
	}

	return strconv.Itoa(prePlatformId), nil
}

func handleNewMembers(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.NewChatMembers == nil {
		return
	}

	// chet Greeting Status
	var myBot []entity.Bot
	var totalCount int
	if err := dao.Bot.Ctx(ctx).Where("account = ?", bot.Self.ID).
		ScanAndCount(&myBot, &totalCount, false); err != nil {
		g.Log().Errorf(ctx, "Unkone Bot %d: %v", bot.Self.ID, err)
		return
	}
	if totalCount == 0 {
		g.Log().Errorf(ctx, "Invalid bot ID: %d", bot.Self.ID)
		return
	}
	if myBot[0].GreetingStatus == consts.GreetingStatusUnAvailable {
		g.Log().Infof(ctx, "GreetingStatusUnAvailable %d", myBot[0].GreetingStatus)
		return
	}
	welcomeText := myBot[0].Greeting

	for _, newUser := range update.Message.NewChatMembers {
		g.Log().Infof(ctx, "New chat member %s", newUser.UserName)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, welcomeText)
		msg.ReplyToMessageID = update.Message.MessageID
		g.Log().Infof(ctx, "msg %v", msg)
		if _, err := bot.Send(msg); err != nil {
			g.Log().Errorf(ctx, "handleNewMembers Send Msg Error: %v", err)
		}
	}
}
