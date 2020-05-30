package uho

import (
	"go_rilla/lib/communicate"

	"github.com/bwmarrin/discordgo"
)

// Call PreFix "uho" によって呼び出されるゴリラコマンド
func Call(session *discordgo.Session, message *discordgo.MessageCreate, channel *discordgo.Channel) {
	mainCommand, subCommands, _, err := FormatOrder(message.Content)

	if err != nil {
		communicate.SendMessage(session, channel, "ウホ(ネイティヴ)")
		return
	}

	switch mainCommand {
	case "react":
		if len(subCommands) < 2 {
			communicate.SendMessage(session, channel, "ウホ(リアクションを付けたいメッセージのIDと絵文字を入力してください)")
			return
		}
		messageID := subCommands[0]
		emojiList := subCommands[1:]
		AddReactions(session, channel, messageID, emojiList)

	default:
		communicate.SendMessage(session, channel, "ウホ")
	}
}
