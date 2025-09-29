package interactive

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
    "net/url"
    "strings"
	"github.com/gorilla/websocket"
	"regexp"
)

func getPortAndWSType(rawurl string) (port, wsType, socketiopath string, err error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", "", "", err
	}

	// Host + port
	port = u.Host

	// ws or wss
	wsType = "ws" // use =, not := 
	if u.Scheme == "https" {
		wsType = "wss"
	}

	// Determine socket.io path
	if u.Path != "" && u.Path != "/" {
		socketiopath = strings.TrimRight(u.Path, "/") + "/socket.io"
	} else {
		if strings.Contains(u.Host, ":") {
			socketiopath = "/socket.io"
		} else {
			socketiopath = "/interactive/socket.io"
		}
	}

	return port, wsType, socketiopath, nil
}



type EventHandler func(data string)

// global registry of handlers
var eventHandlers = make(map[string]EventHandler)

// Allow users to register handlers
func On(eventName string, handler EventHandler) {
	eventHandlers[eventName] = handler
}

func dispatchEvent(message string) {
	if !strings.HasPrefix(message, "42") {
		return
	}

	payload := strings.TrimPrefix(message, "42")

	re := regexp.MustCompile(`^\["([^"]+)",\s*"(.+)"\]$`)
	matches := re.FindStringSubmatch(payload)
	if len(matches) < 3 {
		log.Println("Invalid event format:", payload)
		return
	}

	eventName := matches[1]
	eventData := matches[2]

	if handler, ok := eventHandlers[eventName]; ok {
		handler(eventData)
	} else {
		log.Println("No handler registered for event:", eventName)
	}
}

func Socket(inputURL string, Token string, clientID string) {

	if isHostlookupDone == true {
		inputURL = socketio_url
	}

	port, wsType, handshakepath,  err := getPortAndWSType(inputURL)
    if err != nil {
            fmt.Println("Error:", err)
        }
	// fmt.Println("s URL:", socketio_url)
	// fmt.Println("Port:", port)
	// fmt.Println("Socket.IO Path:", handshakepath)
	//Socket
	connectionURL := fmt.Sprintf("%s://%s%s/?token=%s&userID=%s&apiType=INTERACTIVE&transport=websocket&EIO=3",
		wsType, port, handshakepath, Token, clientID)

	// connectionURL := fmt.Sprintf("%s?token=%s&userID=%s&apiType=INTERACTIVE", inputURL, Token, clientID)

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
			// handleWebSocketMessage(string(message))
			// fmt.Printf("Received message: %s\n", message)
			dispatchEvent(string(message))

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

