package controllers

import (
    "net/http"
    "encoding/json"
    u "lens/utils"
)

var FindRoute = func(w http.ResponseWriter, r *http.Request) {

    err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
    if err != nil {
        u.Respond(w, u.Message(false, "Invalid request"))
        return
    }

    resp := account.Create() //Create account
    u.Respond(w, resp)
}
