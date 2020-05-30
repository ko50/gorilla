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

	if message.Author.Bot {
		return
	}

	checkMessageContent(session, message, channel)

	// UhoCommandを呼び出し
	if strings.HasPrefix(message.Content, "uho") {
		uho.Call(session, message, channel)
	}
}

// メッセージの中に特定の要素が含まれているかチェック
func checkMessageContent(session *discordgo.Session, message *discordgo.MessageCreate, channel *discordgo.Channel) {
	messageContent := message.Content
	communicate.ThrowUnko(session, message, channel)
	communicate.CountBanana(session, messageContent, channel)

	if strings.HasPrefix(messageContent, "うほ") {
		communicate.SendMessage(session, channel, "ｳｯﾎ↑ｳﾎ↓ｳﾎ？？ｳﾎ↓wwww")
	}
	if strings.Contains(messageContent, "うほ") {
		communicate.SendMessage(session, channel, "ウホッ")
	}

	// とは が含まれていたらGoogle検索を実行し上位三件の結果を返します(製作中)
	if strings.Contains(messageContent, "とは") {
		fmt.Println("ウホ")
		url := "https://godoc.org/github.com/PuerkitoBio/goquery" // "https://www.google.com/search?q=ゴリラ"
		scraping.GetSearchResults(url)
	}

}
