package Detail

import (
	"MovieServer/Util/Redis"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Set(id,name string,mtype int,remark string,detail []byte){
	rcli := Redis.R_pool.Get()
	defer rcli.Close()
	var key = id + "_" + name
	isset, _ := redis.Bool(rcli.Do("EXISTS", key))
	//fmt.Println(!isset && (mtype == 11 || remark == "完结"))
	if(!isset && (mtype == 11 || remark == "完结")){		//数据未缓存且当类型为电影或已完结则写入缓存
		rcli.Do("SET", key, detail)
	}else{
		rcli.Do("SETEX", key, 60 * 60 * 10, detail)	//其他影视资源缓存10小时
	}
}

func Get(id string) (detail []byte,err error){
	var key = id + "_*"
	fmt.Println(key)
	rcli := Redis.R_pool.Get()
	defer rcli.Close()
	details, err := redis.Values(rcli.Do("keys", key))
	if(len(details) == 0 || err != nil){
		return nil,errors.New("资源未缓存")
	}
	detail, err = redis.Bytes(rcli.Do("GET", details[0]))
	if err != nil {
		return
	}
	return detail,nil
}