package interactive

import (
	"bufio"
	"fmt"
	"log"
	"os"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func Socket(inputURL string, Token string, clientID string) {

	// Prepare the connection URL
	connectionURL := fmt.Sprintf("%s?token=%s&userID=%s&apiType=INTERACTIVE", inputURL, Token, clientID)

	fmt.Println("Connection URL -->", connectionURL)

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["user"] = "user"
	opts.Query["pwd"] = "pass"
	uri := connectionURL

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})
	client.On("order", func(msg string) {
		log.Printf("on order:%v\n", msg)
	})
	client.On("position", func(msg string) {
		log.Printf("on position:%v\n", msg)
	})
	client.On("tradeConversion", func(msg string) {
		log.Printf("on tradeConversion:%v\n", msg)
	})
	client.On("trade", func(msg string) {
		log.Printf("on trade:%v\n", msg)
	})
	client.On("disconnect", func() {
		log.Printf("on disconnect\n")
	})
	client.On("logout", func(msg string) {
		log.Printf("on logout:%v\n", msg)
	})

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)
		client.Emit("message", command)
		log.Printf("send message:%v\n", command)
	}
}
