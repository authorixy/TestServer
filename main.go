package main

import (
	"encoding/json"
	"fmt"
	"github.com/TestServer/db"
	"github.com/TestServer/idl"
	"io/ioutil"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hippo")
}

func UploadScore(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var req idl.UpdateScoreRequest
	_ = json.Unmarshal(body, &req)
	err := db.GetDB().UPDATE(req.Name, req.Score)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "上传成绩成功")
	}
}

func RankListHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, db.GetDB().GetRankList())
}

func main () {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/rank_list", RankListHandler)
	http.HandleFunc("/upload_score", UploadScore)
	http.ListenAndServe(":10030", nil)
}
