// socket code using websocket lib

package interactive

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func Socket(inputURL string, Token string, clientID string) {

	//Socket
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		log.Println("Error while creating parsed URL:", err)

	}

	// Remove the scheme (http or https) from the URL
	domain := strings.TrimPrefix(inputURL, parsedURL.Scheme+"://")

	// Determine WebSocket scheme (ws or wss)
	var wsScheme string
	if parsedURL.Scheme == "http" {
		wsScheme = "ws"
	} else if parsedURL.Scheme == "https" {
		wsScheme = "wss"
	} else {
		log.Println("Unknown URL scheme:", parsedURL.Scheme)

	}

	// Construct the WebSocket connection URL
	connectionURL := fmt.Sprintf("%s://%s/interactive/socket.io/?token=%s&userID=%s&apiType=INTERACTIVE&transport=websocket&EIO=3", wsScheme, domain, Token, clientID)

	fmt.Println("Connection URL -->", connectionURL)

	u, _, err := websocket.DefaultDialer.Dial(connectionURL, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
		return
	}
	defer u.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := u.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			fmt.Printf("Received message: %s\n", message)

		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := u.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("Error writing message:", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupt received. Closing connection...")
			err := u.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error writing close message:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
