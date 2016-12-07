package routers

import (
   "github.com/urfave/negroni"
   "github.com/gorilla/mux"
   "github.com/abhilashsajeev/go-web/taskmanager/common"
   "github.com/abhilashsajeev/go-web/taskmanager/controllers"
)


func SetNoteRoutes(router *mux.Router) *mux.Router {
    noteRouter := mux.NewRouter()
    noteRouter.HandleFunc("/notes", controllers.CreateTask).Methods("POST")
    noteRouter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
    noteRouter.HandleFunc("/notes/{id}", controllers.UpdateTask).Methods("PUT")
    noteRouter.HandleFunc("/notes/{id}", controllers.GetTaskById).Methods("GET")
    noteRouter.HandleFunc("/notes/{id}", controllers.DeleteTask).Methods("DELETE")
    router.PathPrefix("/notes").Handler(negroni.New(
        negroni.HandlerFunc(common.Authorize),
        negroni.Wrap(noteRouter),
    ))
    return router
}