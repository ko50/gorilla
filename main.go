package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"go_rilla/lib/communicate"
	"go_rilla/lib/parse"
	"go_rilla/lib/scraping"
	"go_rilla/lib/uho"
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

	testCh, _ := discord.Channel("711277442655977499")
	communicate.SendMessage(discord, testCh, "ウッホウッホホ(起動)")
	<-stopBot //プログラムが終了しないようロック
	return
}

// メッセージが送信されたときに実行される
func onMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	channel, err := session.State.Channel(message.ChannelID) //チャンネル取得
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	guildName := parse.GuildName(message.GuildID)
	channelName := parse.ChannelName(message.ChannelID)

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
		communicate.SendMessage(session, channel, "ｳｯﾎ↑ｳﾎ↓ｳﾎ？？ｳﾎ↓wwww")

	case strings.HasPrefix(message.Content, "uho"):
		callUhoCommands(session, message, channel)

	case strings.Contains(message.Content, "うほ"):
		communicate.SendMessage(session, channel, "ウホッ")

	case strings.Contains(message.Content, "バナナ") || strings.Contains(message.Content, "ばなな") || strings.Contains(message.Content, "banana") || strings.Contains(message.Content, "🍌"):
		communicate.SendMessage(session, channel, strings.Repeat("ウホ", banana))
		if banana < 101 {
			for _, sumpleBanana := range allBanana {
				banana += strings.Count(message.Content, sumpleBanana)
			}
		}
	case strings.Contains(message.Content, "とは"):
		fmt.Println("ウホ")
		url := "https://godoc.org/github.com/PuerkitoBio/goquery" // "https://www.google.com/search?q=ゴリラ"
		scraping.GetSearchResults(url)
	}
}

// PreFix "uho" によって呼び出されるゴリラコマンド
func callUhoCommands(session *discordgo.Session, message *discordgo.MessageCreate, channel *discordgo.Channel) {
	mainCommand, _, subCommands := uho.FormatOrder(message.Content)

	switch mainCommand {
	case "react":
		if len(subCommands) < 2 {
			communicate.SendMessage(session, channel, "ウホ(リアクションを付けたいメッセージのIDと絵文字を入力してください)")
			return
		}
		messageID := subCommands[0]
		emojiList := subCommands[1:]
		uho.AddReactions(session, channel, messageID, emojiList)

	default:
		communicate.SendMessage(session, channel, "ウホ")
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
	communicate.SendMessage(session, channel, sendText)
	vcsession, _ = session.ChannelVoiceJoin(channel.GuildID, "700920014475362348", false, false)
	vcsession.AddHandler(onVoiceReceived) //音声受信時のイベントハンドラ

case strings.HasPrefix(message.Content, fmt.Sprintf("%s %s", botID, "!disconnect")):
	vcsession.Disconnect() //今いる通話チャンネルから抜ける
*/
