package common

import (
    "os"
    "log"
    "encoding/json"
    "net/http"
)


type (
    configuration struct {
        Server, MongoDBHost, DBUser, DBPass, Database string
    }

    appError struct {
        Error string `json:"error"`
        Message string `json:"message"`
        HttpStatus int `json:"status"`
    }

    errorResource struct {
        Data appError `json:"data"`
    }
)

// to hold values from config.json
var Appconfig configuration

func initConfig() {
    loadAppConfig()
}

func loadAppConfig() {
    file, err := os.Open("common/config.json")
    defer file.Close()
    if err != nil {
        log.Fatalf("[loadConfig]: %s\n", err)
    }
    decoder := json.NewDecoder(file)
    Appconfig = configuration{}
    err = decoder.Decode(&Appconfig)
    if err != nil {
        log.Fatalf("[loadAppConfig]: %s\n", err)
    }
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
    errorObj := appError {
        Error : handlerError.Error(),
        Message: message,
        HttpStatus: code,
    }
    log.Printf("AppError]: %s\n", handlerError)

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(code)    

    if j,err := json.Marshal(errorResource{Data:errorObj}); err == nil {
        w.Write(j)
    }
}
