package Http

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Api_fmt struct {
	AUrl string
	Parms url.Values
}

func (this Api_fmt) Get_api()(resp []byte,errs error){
	Url,_ := url.Parse(this.AUrl)
	Url.RawQuery = this.Parms.Encode()
	rsp,errs := http.Get(Url.String())
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}

func RespJson(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json; charset=utf-8")
}

func GetM3u8(url string)string{
	var m3u8 string
	var tmp []string = strings.Split(url, "$$$")
	if(len(tmp) > 1){
		if(strings.Index(tmp[0],"m3u8")>-1){
			m3u8 = tmp[0]
		}else{
			m3u8 = tmp[1]
		}
	}else{
		m3u8 = tmp[0]
	}
	return m3u8
}