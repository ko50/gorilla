package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	token                = "Bot " + os.Getenv("GO_RILLA_TOKEN")
	botID                = "710666258831507476"
	genkaiVoiceChannelID = "683939861539192865"
	testVoiceChannelID   = "700920014475362348"
	stopBot              = make(chan bool)
	vcsession            *discordgo.VoiceConnection
	allBanana            = [4]string{"バナナ", "ばなな", "banana", "🍌"}
	allUnchi             = [10]string{"うんこ", "ウンコ", "うんち", "ウンチ", "クソ", "くそ", "unchi", "unnchi", "💩", "うーんこ"}
	banana               = 1
)

func main() {
	//Discordのセッションを作成
	discord, err := discordgo.New()
	discord.Token = token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(onMessageCreate) //全てのWSAPIイベントが発生した時のイベントハンドラを追加
	// websocketを開いてlistening開始
	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listening...")
	<-stopBot //プログラムが終了しないようロック
	return
}

// サーバーIDをサーバー名に変換
func parseGuildName(guildID string) (name string) {
	switch guildID {
	case "683939861539192860":
		name = "限界開発鯖"

	case "700920013439369257":
		name = "幣鯖"

	default:
		name = guildID
	}

	return
}

// チャンネルIDを数字の羅列から既知のチャンネル名にパースする
func parseChannelName(channelID string) (name string) {
	switch channelID {
	// カス
	case "683939861539192863":
		name = "一般"

	case "689819731531661339":
		name = "アイデア掃き溜め所"

	case "690909527461199922":
		name = "無法地帯"

	case "709810469493538858":
		name = "関連Tweet"

	case "683950833393467435":
		name = "Git Log"

	// 幣鯖
	case "700920014475362346":
		name = "一般"

	case "700920193580531712":
		name = "kobot"

	case "711277442655977499":
		name = "go_rilla"

	default:
		name = channelID
	}

	return
}

// メッセージが送信されたときに実行される
func onMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	channel, err := session.State.Channel(message.ChannelID) //チャンネル取得
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	guildName := parseGuildName(message.GuildID)
	channelName := parseChannelName(message.ChannelID)

	if message.Author.ID != botID || !(message.Author.Bot && channelName == "Git Log") {
		// log.Println(utf8.RuneCountInString(message.Content))
		fmt.Printf("\n%s/%10s %20s\n    %s :  %s\n", guildName, channelName, time.Now().Format(time.Stamp), message.Author.Username, message.Content)
	}

	throwUnko(session, message, channel)

	if message.Author.Bot {
		return
	}

	switch {
	case strings.HasPrefix(message.Content, "うほ"):
		sendMessage(session, channel, "ｳｯﾎ↑ｳﾎ↓ｳﾎ？？ｳﾎ↓wwww")

	case strings.HasPrefix(message.Content, "uho"):
		uhoCommands(session, message, channel)

	case strings.Contains(message.Content, "うほ"):
		sendMessage(session, channel, "ウホッ")

	case strings.Contains(message.Content, "バナナ") || strings.Contains(message.Content, "ばなな") || strings.Contains(message.Content, "banana") || strings.Contains(message.Content, "🍌"):
		sendMessage(session, channel, strings.Repeat("ウホ", banana))
		if banana < 101 {
			for _, sumpleBanana := range allBanana {
				banana += strings.Count(message.Content, sumpleBanana)
			}
		}
	}
}

// PreFix "uho" によって呼び出されるゴリラコマンド達
func uhoCommands(session *discordgo.Session, message *discordgo.MessageCreate, channel *discordgo.Channel) {
	var (
		orderList   = strings.Split(message.Content[4:], " ")
		mainCommand = orderList[0]
	)
	subCommands := make([]string, 0, 20)
	if len(orderList) > 1 {
		for _, order := range orderList[1:] {
			subCommands = append(subCommands, order)
		}
	}

	switch mainCommand {
	case "react":
		if len(subCommands) < 2 {
			sendMessage(session, channel, "ウホ(リアクションを付けたいメッセージのIDと絵文字を入力してください)")
			return
		}
		messageID := subCommands[0]
		emojiList := subCommands[1:]
		addReactions(session, channel, messageID, emojiList)

	default:
		sendMessage(session, channel, "ウホ")
	}
}

// 部分文字列 "うんこ" 等を含むメッセージにゴリラがウンコを投げつける関数
func throwUnko(session *discordgo.Session, message *discordgo.MessageCreate, channel *discordgo.Channel) {
	if channel.ID == "683950833393467435" {
		return
	}
	for _, unchi := range allUnchi {
		if strings.Contains(message.Content, unchi) {
			session.MessageReactionAdd(channel.ID, message.ID, "💩")
			return
		}
	}
}

// 多数のリアクションを付けていく関数 感情が豊かなゴリラ
func addReactions(session *discordgo.Session, channel *discordgo.Channel, messageID string, emojiList []string) {
	for _, emoji := range emojiList {
		defaultEmojiErr := session.MessageReactionAdd(channel.ID, messageID, emoji)
		if defaultEmojiErr != nil {
			err := session.MessageReactionAdd(channel.ID, messageID, strings.TrimRight(emoji[1:], ">"))
			if err != nil {
				log.Println(err)
				sendMessage(session, channel, "ちょっと何言ってるかわからないウホwwwwwwwwwwwww")
			}
		}
	}
}

//メッセージを送信する関数 ゴリラちゃんでもお話ししたい！
func sendMessage(session *discordgo.Session, channel *discordgo.Channel, msg string) {
	_, err := session.ChannelMessageSend(channel.ID, msg)

	guildName := parseGuildName(channel.GuildID)
	channelName := parseChannelName(channel.ID)

	log.Println(fmt.Sprintf("\n%s/%10s %20s\n    %s :  %s", guildName, channelName, time.Now().Format(time.Stamp), "go_rilla", msg))
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}

/*
VC関連のチュートリアル

//メッセージを受信した時の、声の初めと終わりにPrintされるようだ
func onVoiceReceived(vc *discordgo.VoiceConnection, vs *discordgo.VoiceSpeakingUpdate) {
	log.Print("しゃべったあああああ")
}

case strings.HasPrefix(message.Content, fmt.Sprintf("%s %s", botID, "!join")):
	guildChannels, _ := session.GuildChannels(channel.GuildID)
	var sendText string
	for _, a := range guildChannels {
		sendText += fmt.Sprintf("%vチャンネルの%v(IDは%v)\n", a.Type, a.Name, a.ID)
	}
	sendMessage(session, channel, sendText)
	vcsession, _ = session.ChannelVoiceJoin(channel.GuildID, "700920014475362348", false, false)
	vcsession.AddHandler(onVoiceReceived) //音声受信時のイベントハンドラ

case strings.HasPrefix(message.Content, fmt.Sprintf("%s %s", botID, "!disconnect")):
	vcsession.Disconnect() //今いる通話チャンネルから抜ける
*/
