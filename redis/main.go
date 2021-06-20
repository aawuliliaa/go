package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	value1 := make([]byte,100000)

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	for _,num:= range []int{1,2,3,4,5,6,7,8,9,10}{
		key := fmt.Sprintf("var10_%d",num)
		_, err = c.Do("SET", key, string(value1))
		if err != nil {
			fmt.Println("redis set failed:", err)
		}
	}


}
