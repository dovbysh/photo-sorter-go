package main

import "./Message"

func main() {
	m := Message.NewMessage(nil)

	m.SuccessCopied("/tmp/src", "/tmp/dst")
}
