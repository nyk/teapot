/**
 * Teapotd - an annotated media collation (AMC) server
 * User: Nyk Cowham <nyk@demotix.com>
 * Date: 5/25/13 AD
 */
package main

import (
    "os"
	"log"
	/* "teapot" */
	"net/http"
	"encoding/hex"
	"crypto/rand"
	"github.com/justinfx/go-socket.io/socketio" // only supports client 0.6.2 !!! (argghhh!)
)

func mediaHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1<<30)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		username := r.FormValue("username")
		uid := r.FormValue("uid")

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Username: " + username + "\n"))
		w.Write([]byte("uid:" + uid + "\n"))
	}
}

/**
 * Function that generates UUIDs based of the RFC 4122 standard.
 */
func genToken() (string, error) {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return "", err
	}
	// TODO: verify the two lines implement RFC 4122 correctly
	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u), nil
}

func main() {
	sio := socketio.NewSocketIO(nil)

	sio.OnConnect(func(c *socketio.Conn) {
		sio.Broadcast(struct{ announcement string }{"connected: " + c.String()})
	})

	sio.OnDisconnect(func(c *socketio.Conn) {
		sio.BroadcastExcept(c,
			struct{ announcement string }{"disconnected: " + c.String()})
	})

	sio.OnMessage(func(conn *socketio.Conn, msg socketio.Message) {
		// var col teapot.Collation

		if msg.Type() == socketio.MessageText {
			name := msg.Data()
			switch name {
			case "getToken":
				token, _ := genToken()
				conn.Send(token)
			case "annotateCollation":
				//var key, value string
				//msg.ReadArguments(&key, &value)
				//col.Annotate(key, value)
				conn.Send("collationAnnotated")
			case "addMedia":
				// Is it possible to get a file stream through js event message??
				// probably not - so will have to implement form upload and messaging.
				conn.Send("mediaAdded")
			case "annotateMedia":
				// var media, key, value string
				// msg.ReadArguments(&media, &key, &value)
				// mfile := c.GetMediaByKey(media)
				// mfile.Annotate(key, value)
				conn.Send("mediaAnnotated")
			}
		}
	})

	htmlDir, _ := os.Getwd()
	mux := sio.ServeMux()
	mux.Handle("/html/", http.FileServer(http.Dir(htmlDir)))
	mux.HandleFunc("/media", mediaHandler)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	/* Create a goroutine to watch the beanstalk queue for messages */
}
