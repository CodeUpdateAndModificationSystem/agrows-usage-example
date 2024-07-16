package socket

import (
	"agrows-usage-example/internal/functions"
	"agrows-usage-example/web"
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/a-h/templ"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var (
	Send      = make(chan any, 100)
	connMutex sync.Mutex
	conn      net.Conn
	closeChan = make(chan struct{})
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	connMutex.Lock()
	defer connMutex.Unlock()

	log.Println("Upgrading HTTP request to WebSocket...")
	newConn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	if conn != nil {
		log.Println("Closing previous WebSocket connection...")
		_ = conn.Close()
	}

	conn = newConn
	close(closeChan)
	closeChan = make(chan struct{})

	log.Println("Starting message handlers...")

	go handleIncomingMessages()
	go handleOutgoingMessages()
}

var shouldReload = true

func handleIncomingMessages() {
	for {
		select {
		case <-closeChan:
			log.Println("Incoming message handler stopped.")
			return
		default:
			msg, op, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Printf("Error reading from WebSocket: %v", err)
				return
			}
			if op == ws.OpText {
				log.Printf("Received text message: %s", string(msg))
				switch string(msg) {
				case "debug":
					Send <- web.Debug("debug")
				case "reload?":
					if shouldReload {
						Send <- "reload!"
						shouldReload = false
					}
				}
			} else if op == ws.OpBinary {
				log.Printf("Received binary message:\n%s", functions.FormatXXD(msg))
				result, err := functions.AgrowsReceive(msg)
				if err != nil {
					log.Printf("Error processing message: %v", err)
				} else {
					Send <- result
				}
			}
		}
	}
}

func handleOutgoingMessages() {
	for msg := range Send {
		processed := ""

		switch m := msg.(type) {
		case templ.Component:
			ctx := context.Background()
			buf := new(bytes.Buffer)
			if err := m.Render(ctx, buf); err != nil {
				log.Printf("Error rendering component: %v", err)
				continue
			}
			processed = buf.String()
		case string:
			processed = m
		default:
			processed = fmt.Sprintf("%v", m)
		}

		if err := writeMessage(processed); err != nil {
			log.Printf("Error writing to WebSocket: %v", err)
			closeConnection()
		}
	}
}

func writeMessage(msg interface{}) error {
	connMutex.Lock()
	defer connMutex.Unlock()
	log.Printf("Writing message to WebSocket: %v", msg)

	var err error
	switch m := msg.(type) {
	case string:
		err = wsutil.WriteServerMessage(conn, ws.OpText, []byte(m))
	case []byte:
		err = wsutil.WriteServerMessage(conn, ws.OpBinary, m)
	default:
		log.Printf("Unknown message type: %T", m)
		return nil
	}
	return err
}

func closeConnection() {
	connMutex.Lock()
	defer connMutex.Unlock()

	if conn != nil {
		log.Println("Closing WebSocket connection...")
		_ = conn.Close()
		conn = nil
		close(closeChan)
	}
}
