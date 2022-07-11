package main

import "fmt"

func main() {
	msg := "some message"

	var msgPointer *string = &msg //* => ใส่เพื่อว่าจะใส่ค่าใน pointer, & => pointer ตัวแปรที่เราจะใส่
	fmt.Println(msg)
	fmt.Println(msgPointer)
	fmt.Println(*msgPointer) // * เพื่อจะนำค่า value ออกมาแสดง ex. = some message

	changeMessage(&msg, "new message 1")
	fmt.Println(msg)

	changeMessage(msgPointer, "new message 2")
	fmt.Println(msg)

	changeMessage(msgPointer, "new message 3")
	fmt.Println(*msgPointer)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
