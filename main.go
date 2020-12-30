package main

import (
	"bufio"
	"fmt"
	"github.com/galehuang/chatroom-project/login"
	"os"
)

func main() {
	fmt.Print("welcome to chatroom, log in (1) or create new account (2) or quit (0):")
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	choice = choice[0 : len(choice)-1]
	if err != nil {
		panic("error when reading user input")
	}
	switch choice {
	case "0":
		fmt.Println("bye")
		return
	case "1":
		err := login.Login()
		if err != nil {
			panic("error when logging")
		}
		fmt.Println("login successfully!")
	case "2":
		panic("not ready")
	}


}
