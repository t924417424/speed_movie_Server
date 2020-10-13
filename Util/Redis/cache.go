package Redis

import (
	"time"
	redigo "github.com/garyburd/redigo/redis"
)

var R_pool *redigo.Pool

func init(){
	var addr = "127.0.0.1:6379"
	var password = ""
	R_pool = poolInitRedis(addr,password)
}

func poolInitRedis(server string, password string) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     2,//空闲数
		IdleTimeout: 240 * time.Second,
		MaxActive:   3,//最大数
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}