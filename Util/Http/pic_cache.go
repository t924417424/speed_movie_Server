package Http

import (
	"MovieServer/Others/Detail"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func GetPic(vid string,detail Detail.Detail)([]byte,error)  {
	var pic []byte
	if _,err := os.Stat("./Cache/Image/" + vid + ".jpg");os.IsNotExist(err){
		log.Println("文件不存在，网络获取",err)
		datas, _ := Api_fmt{detail.List[0].VodPic,nil}.Get_api()
		var s Detail.Detail
		json.Unmarshal(datas,&s)
		log.Println(s.List[0].VodPic)
		pic, _ = Api_fmt{s.List[0].VodPic,nil}.Get_api()
		file,_ := os.OpenFile("./Cache/Image/" + vid + ".jpg",os.O_RDWR|os.O_CREATE,0666)
		defer file.Close()
		file.Write(pic)
	}else{
		log.Println("文件存在，本地返回")
		file,_ := os.Open("./Cache/Image/" + vid + ".jpg")
		defer file.Close()
		pic,_ = ioutil.ReadAll(file)
	}
	return pic,nil
}