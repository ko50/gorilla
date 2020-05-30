package communicate

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// ThrowUnko 部分文字列 "うんこ" 等を含むメッセージにゴリラがウンコを投げつける関数
func ThrowUnko(session *discordgo.Session, message *discordgo.MessageCreate, channel *discordgo.Channel) {
	if channel.ID == "683950833393467435" {
		return
	}
	allUnchi := [10]string{"うんこ", "ウンコ", "うんち", "ウンチ", "クソ", "くそ", "unchi", "unnchi", "💩", "うーんこ"}
	for _, unchi := range allUnchi {
		if strings.Contains(message.Content, unchi) {
			session.MessageReactionAdd(channel.ID, message.ID, "💩")
			return
		}
	}
}
