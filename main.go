package main

import "./helper"

func main() {
	m := helper.NewMessage(nil)

	m.SuccessCopied("/tmp/src", "/tmp/dst")
}
