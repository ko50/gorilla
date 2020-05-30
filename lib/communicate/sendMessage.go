package communicate

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"

	"go_rilla/lib/parse"
)

// SendMessage メッセージを送信する関数 ゴリラちゃんでもお話ししたい！
func SendMessage(session *discordgo.Session, channel *discordgo.Channel, msg string) {
	_, err := session.ChannelMessageSend(channel.ID, msg)

	guildName := parse.GuildName(channel.GuildID)
	channelName := parse.ChannelName(channel.ID)

	log.Println(fmt.Sprintf("\n%s/%10s %20s\n    %s :  %s", guildName, channelName, time.Now().Format(time.Stamp), "go_rilla", msg))
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
