package handlers

import (
	. "bots/internal/cache"
	"net/http"
	"strconv"
)

func GetBotsCount(w http.ResponseWriter, _ *http.Request) {
	sum := CacheApp.GetBotsCount()
	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(sum)))
}

func UpdateBotsCount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userId := r.Form.Get("user_id")
	CacheApp.Set(userId)
	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(CacheApp.Get(userId))))
}
