package main

import (
	"os"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
	 _ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
	
)

type App struct {
	Router *mux.Router
	hostname string
}

func (a *App) testing(w http.ResponseWriter, r *http.Request){
	//os.Setenv("SMTPMAIL","HELLLO")
	fmt.Fprintf(w, "You have visited %s!",os.Getenv("AUTH_PASSWORD"))
	log.Info("logrus")
}

func (a *App) initializeRoutes() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/v1/api/hosts", a.allHosts).Methods("GET")
	a.Router.HandleFunc("/v1/api/host",a.host).Methods("GET")
	a.Router.HandleFunc("/testing",a.testing).Methods("GET")
	a.Router.HandleFunc("/v1/api/service",a.getServiceNotOk).Methods("GET")
}


func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}


func (a *App) allHosts(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("ROOT_URL")+os.Getenv("HOSTS")
	log.Println("Invoking:"+url)
	a.listHosts(w, url)	
}

func (a *App) getServiceNotOk(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("ROOT_URL")+os.Getenv("SERVICE_NOTOK")
	log.Info("Invoking:"+url)
	a.listHosts(w, url)	
}

func (a *App) add(x,y int) int {
	return x + y

}

func (a *App) host(w http.ResponseWriter, r *http.Request) {
	v:= r.URL.Query()
	hostname:=v.Get("hostname")
	s:=App{hostname:hostname}
	url := os.Getenv("ROOT_URL")+"?host="+s.hostname
	log.Println("Invoking:"+url)
	 a.listHosts(w, url)
}

func (a *App) listHosts(w http.ResponseWriter, url string ) {	
	file,e := os.OpenFile("logs.txt",os.O_CREATE | os.O_WRONLY | os.O_APPEND,0666 )

	if e!=nil {
		log.Fatalln("Failed To Get log file")
	}

	log.SetOutput(file)
	log.Println(url)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("error")
	}

	req.SetBasicAuth(os.Getenv("AUTH_USERNAME"),os.Getenv("AUTH_PASSWORD"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, _ := client.Do(req)

	defer res.Body.Close()
	body, err:= ioutil.ReadAll(res.Body)

	// if err == nil{
	// 	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	// 	log.Println("Inside If")
	// 	return
	// }
	respondWithJSON(w, http.StatusOK,body)
	//response, _ := json.MarshalIndent(body,"","	")
	
	//fmt.Println(body)
	
}

func respondWithJSON(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
	log.Info(payload)
	//fmt.Fprintln(os.Std,w.Write(payload))

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, []byte{0,1})
}		
