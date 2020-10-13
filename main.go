package main

import (
	"MovieServer/Others/Detail"
	"MovieServer/Others/Search"
	"MovieServer/Util/Http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.Ldate | log.Lshortfile)
	router := mux.NewRouter()
	router.HandleFunc("/api/search/{name}", _search)
	router.HandleFunc("/api/detail/{vid}", _detail)
	http.ListenAndServe(":8091", router)
}



func _search(w http.ResponseWriter, r *http.Request) {
	Http.RespJson(w)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	//log.Print(vars)
	parms := url.Values{}
	parms.Set("wd", vars["name"])
	if r.FormValue("page") != "" {
		parms.Set("pg", r.FormValue("page"))
	}
	datas, _ := Http.Api_fmt{"https://api.okzy.tv/api.php/provide/vod/at/json/", parms}.Get_api()
	var s Search.Search
	json.Unmarshal(datas, &s)
	for _, v := range s.List {
		log.Println(v.TypeName + "||" + v.VodName + "||" + strconv.Itoa(v.VodID) + "||" + v.VodEn)
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  Search:  " + vars["name"])
	//index,err := template.ParseFiles("./View/search.html")
	//if err != nil{
	//	panic(err)
	//}
	//index.Execute(w,s)

	w.Write(datas)
}

func _detail(w http.ResponseWriter, r *http.Request) {
	Http.RespJson(w)
	var datas []byte
	var info Detail.Minfo
	var s Detail.Detail
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	datas, _ = Detail.Get(vars["vid"])
	if len(datas) == 0 {
		//fmt.Println("走远程")
		parms := url.Values{}
		parms.Set("ac", "detail")
		parms.Set("ids", vars["vid"])
		datas, _ = Http.Api_fmt{"https://api.okzy.tv/api.php/provide/vod/", parms}.Get_api()
		json.Unmarshal(datas, &s)
		info = Detail.Minfo{
			s.List[0].VodID,
			s.List[0].VodPic,
			s.List[0].VodName,
			s.List[0].VodClass,
			s.List[0].VodBlurb,
			s.List[0].VodContent,
			s.List[0].VodLang,
			s.List[0].VodRemarks,
			Http.GetM3u8(s.List[0].VodPlayURL),
		}
		go Detail.Set(vars["vid"], s.List[0].VodName, s.List[0].TypeID1, s.List[0].VodRemarks, datas)
	} else {
		//fmt.Println("走缓存")
		json.Unmarshal(datas, &s)
		info = Detail.Minfo{
			s.List[0].VodID,
			s.List[0].VodPic,
			s.List[0].VodName,
			s.List[0].VodClass,
			s.List[0].VodBlurb,
			s.List[0].VodContent,
			s.List[0].VodLang,
			s.List[0].VodRemarks,
			Http.GetM3u8(s.List[0].VodPlayURL),
		}
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "  Player:  " + info.VodName)
	datas, _ = json.Marshal(info)
	w.Write(datas)
}