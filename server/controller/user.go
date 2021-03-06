package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var users [10]UserData

func init() {
	for i := range users {
		users[i].Id = uint32(i + 1)
		users[i].Login = "User " + strconv.Itoa(i)
	}
}

type CUser struct {
	BaseController
}

type UserData struct {
	Id    uint32 `json:"id"`
	Login string `json:"login"`
}

func User() *CUser {
	return &CUser{
		*Default("/user", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				r.ParseForm()
				id, err := strconv.Atoi(r.Form.Get("id"))
				if err != nil {
					http.Error(w, "User id required", http.StatusBadRequest)
					return
				}

				found := false
				for _, user := range users {
					if uint32(id) == user.Id {
						found = true
						data, err := json.Marshal(user)

						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}

						w.Header().Set("Content-Type", "application/json")
						w.Write(data)
					}
				}
				if !found {
					http.Error(w, "Not found", http.StatusNotFound)
					return
				}
			}
		})}
}
