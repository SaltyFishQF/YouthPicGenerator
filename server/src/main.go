package main

import (
	"./utils"
	"fmt"
	"github.com/robfig/cron"
	"log"
	"net/http"
)

func main() {
	utils.Spider()
	c := cron.New()
	c.AddFunc("0 0 0 * * ?", utils.Spider)
	c.Start()
	defer c.Stop()
	fmt.Println("server starting")
	http.HandleFunc("/api/getLessonList", LessonHandler)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func LessonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling")
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	fmt.Fprintln(w, utils.LESSON_JSON)
}
