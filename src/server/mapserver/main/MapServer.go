package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	port = 8001
)

type MapServer struct {

}

func getFormValue(hr *http.Request, name string) string {
	values := hr.Form[name]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func envoieRéponse(w http.ResponseWriter, out *MessageOut) {
	bout, err := json.Marshal(out)
	if err != nil {
		fmt.Println("Erreur encodage réponse : ", err) 
		return
	}
	fmt.Fprint(w, "receiveFromMapServer(")
	w.Write(bout)
	fmt.Fprint(w, ")")	
}

func (h *MapServer) ServeHTTP(w http.ResponseWriter, hr *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "GET")
	hr.ParseForm()

	in := new(MessageIn)
	out := new(MessageOut)
	defer envoieRéponse(w, out)
	fmt.Println(getFormValue(hr, "in"))
	err := json.Unmarshal(([]byte)(getFormValue(hr, "in")), in)
	if err != nil {
		out.Erreur = "Erreur décodage : " + err.Error()
		return
	}	
}

func (server *MapServer) Start() {
	http.Handle("/", server)
	fmt.Printf("mapserver démarre sur le port %d\n", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("Erreur au lancement : ", err)
	}
}

func main() {
	ms := new(MapServer)
	ms.Start()
}
