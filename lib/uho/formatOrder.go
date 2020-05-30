package uho

import (
	"errors"
	"strings"
)

// FormatOrder uhoコマンドが呼ばれたときのメッセージを orderList, mainCommand, subCommands に整えて返します
func FormatOrder(rawOrder string) (mainCommand string, orderList, subCommands []string, err error) {
	if len(rawOrder) < 5 {
		err = errors.New("コマンドとして不十分な入力")
		return
	}
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
