package main

import (
    "log"
    "net/http"
    "github.com/urfave/negroni"
    "github.com/abhilashsajeev/go-web/taskmanager/common"
    "github.com/abhilashsajeev/go-web/taskmanager/routers"
)

//Entry point of the program
func main() {

// Get the mux router object
    router := routers.InitRoutes()
// Create a negroni instance
    n := negroni.Classic()
    n.UseHandler(router)
    server := &http.Server{
        Addr: common.Appconfig.Server,
        Handler: n,
    }
    log.Println("Listening...")
    server.ListenAndServe()
}