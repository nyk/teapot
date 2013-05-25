/**
 * Created with IntelliJ IDEA.
 * User: Nyk Cowham <nyk@demotix.com>
 * Date: 5/25/13 AD
 * Time: 6:51 PM
 */
package teapot

import (
	"net/http"
	"log"
	"time"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1 << 20)
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

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is a teapot server.\n"))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/session", sessionHandler)

	server := &http.Server {
 		Addr: "127.0.0.1:6969",
	    ReadTimeout: 60*time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
