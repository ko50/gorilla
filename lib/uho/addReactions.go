package uho

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"

	"go_rilla/lib/communicate"
)

// AddReactions メッセージのIDと絵文字のIDを受け取り、リアクション追加を行うウホ
func AddReactions(session *discordgo.Session, channel *discordgo.Channel, messageID string, emojiList []string) {
	for _, emoji := range emojiList {
		defaultEmojiErr := session.MessageReactionAdd(channel.ID, messageID, emoji)
		if defaultEmojiErr != nil {
			err := session.MessageReactionAdd(channel.ID, messageID, strings.TrimRight(emoji[1:], ">"))
			if err != nil {
				log.Println(err)
				communicate.SendMessage(session, channel, "ちょっと何言ってるかわからないウホwwwwwwwwwwwww")
			}
		}
	}
}
