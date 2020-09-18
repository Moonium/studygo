package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Println("connect redis failed, err:", err)
	}

	fmt.Println("连接成功！")

	// zset
	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "PHP"},
		redis.Z{Score: 96, Member: "Golang"},
		redis.Z{Score: 97, Member: "Python"},
		redis.Z{Score: 99, Member: "Java"},
	}

	rdb.ZAdd(key, items...)

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

}
