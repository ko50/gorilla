package communicate

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	// Banana is just banana.
	//         --- Gorilla-Gorilla-Gorilla
	Banana    = 0
	allBanana = [4]string{"バナナ", "ばなな", "banana", "🍌"}
)

// CountBanana 渡されたメッセージの中にバナナ関係の言葉があるかチェックします
func CountBanana(session *discordgo.Session, messageContent string, channel *discordgo.Channel) {
	if Banana >= 100 {
		return
	}

	c := 0
	for _, b := range allBanana {
		c += strings.Count(messageContent, b)
	}
	if c != 0 {
		Banana += c
		SendMessage(session, channel, strings.Repeat("ウホ", Banana))
	}
}
