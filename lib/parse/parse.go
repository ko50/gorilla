package parse

// GuildName サーバーIDをサーバー名に変換
func GuildName(guildID string) (name string) {
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

// ChannelName チャンネルIDを数字の羅列から既知のチャンネル名にパースする
func ChannelName(channelID string) (name string) {
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
