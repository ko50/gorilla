package uho

import (
	"strings"
)

// FormatOrder uhoコマンドが呼ばれたときのメッセージを orderList, mainCommand, subCommands に整えて返します
func FormatOrder(rawOrder string) (mainCommand string, orderList, subCommands []string) {
	orderList = strings.Split(rawOrder[4:], " ")
	mainCommand = orderList[0]
	subCommands = make([]string, 0, 20)
	if len(orderList) > 1 {
		for _, order := range orderList[1:] {
			subCommands = append(subCommands, order)
		}
	}

	return
}
