package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	redigo "github.com/garyburd/redigo/redis"
	redis "gopkg.in/redis.v3"

	"samples"
)

func main() {

	samples.T79()
}

var (
	client *redis.Client
)

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return v.(string)
	case bool:
		return strconv.FormatBool(v.(bool))
	case int:
		return strconv.Itoa(v.(int))
	case int8, int16, int32, int64:
		return strconv.FormatInt(v.(int64), 10)
	case float32:
		return strconv.FormatFloat(v.(float64), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', 2, 64) //小数位保留两位数，且四舍五入
	case []byte:
		return string(v.([]byte))
	case uint8, uint16, uint32, uint64:
		return strconv.FormatUint(v.(uint64), 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	default:
		return fmt.Sprintf("%s", v)
	}
}

func redis_V() {
	client = redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:6379",
		Password:    "",
		DB:          0,
		PoolSize:    100,
		PoolTimeout: time.Second * 1000,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	var waitop sync.WaitGroup

	waitop.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func(index int) {
			defer waitop.Done()
			pip := client.Pipeline()
			defer pip.Close()

			key := "k" + ToString(index)
			field := "f" + ToString(index)
			value := "v" + ToString(index)
			pip.HSet(key, field, value).Result()
			pip.HGet(key, field).Result()

			key = "u" + ToString(index)
			field = "w" + ToString(index)
			value = "x" + ToString(index)
			pip.HSet(key, field, value).Result()
			pip.HGet(key, field).Result()

			_, e := pip.Exec()
			if e != nil {
				fmt.Println(e)
				return
			}

		}(i)
	}

	waitop.Wait()
}

func initRedis(poolName string) {

}

const (
	CONST_Redisgo_name = "test"
)

func redis_go() {
	rd := new(samples.RedisDriver)
	rd.Connect("127.0.0.1:6379", "", 0, 100, 100)
	samples.RedisRegister(CONST_Redisgo_name, rd)

	var waitop sync.WaitGroup

	waitop.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func(index int) {
			conn := samples.Rds[CONST_Redisgo_name].Pool.Get()
			if conn == nil {
				fmt.Println("Pool no client")
				return
			}
			defer conn.Close()
			defer waitop.Done()

			key := "k" + ToString(index)
			field := "f" + ToString(index)
			value := "v" + ToString(index)

			conn.Send("HSet", key, field, value)
			conn.Send("HGet", key, field)
			//_, err := conn.Do("HGet", key, field)
			//if err != nil {
			//	fmt.Println("have err:%v", err.Error())
			//}

			//time.Sleep(1 * time.Second)

			key = "u" + ToString(index)
			field = "w" + ToString(index)
			value = "x" + ToString(index)
			conn.Send("HSet", key, field, value)
			conn.Send("HGet", key, field)
			//_, err = conn.Do("HGet", key, field)
			//if err != nil {
			//	fmt.Println("have err:%v", err.Error())
			//}

			_, e := conn.Do("")
			if e != nil {
				fmt.Println(e.Error)
			}

		}(i)
	}

	waitop.Wait()

	//time.Sleep(3 * time.Second)

}

func redis_go1() {
	rd := new(samples.RedisDriver)
	rd.Connect("127.0.0.1:6379", "", 0, 100, 100)
	samples.RedisRegister(CONST_Redisgo_name, rd)

	var waitop sync.WaitGroup

	waitop.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func(index int) {
			for x := 0; x < 10; x++ {
				RediDo1(index)
			}
			defer waitop.Done()
		}(i)
	}

	waitop.Wait()

	//time.Sleep(3 * time.Second)

}
func redis_go2() {
	rd := new(samples.RedisDriver)
	rd.Connect("127.0.0.1:6379", "", 0, 100, 100)
	samples.RedisRegister(CONST_Redisgo_name, rd)

	var waitop sync.WaitGroup

	waitop.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func(index int) {
			conn := samples.Rds[CONST_Redisgo_name].Pool.Get()
			if conn == nil {
				fmt.Println("Pool no client")
				return
			}
			defer conn.Close()

			for x := 0; x < 10; x++ {
				RediDo2(index, conn)
			}
			defer waitop.Done()
		}(i)
	}

	waitop.Wait()

	//time.Sleep(3 * time.Second)

}

func RediDo1(index int) {
	conn := samples.Rds[CONST_Redisgo_name].Pool.Get()
	if conn == nil {
		fmt.Println("Pool no client")
		return
	}
	defer conn.Close()

	key := "k" + ToString(index)
	field := "f" + ToString(index)
	value := "v" + ToString(index)

	conn.Send("HSet", key, field, value)
	conn.Send("HGet", key, field)

	key = "u" + ToString(index)
	field = "w" + ToString(index)
	value = "x" + ToString(index)
	conn.Send("HSet", key, field, value)
	conn.Send("HGet", key, field)

	_, e := conn.Do("")
	if e != nil {
		fmt.Println(e.Error)
	}
}
func RediDo2(index int, conn redigo.Conn) {

	key := "k" + ToString(index)
	field := "f" + ToString(index)
	value := "v" + ToString(index)

	conn.Send("HSet", key, field, value)
	conn.Send("HGet", key, field)

	key = "u" + ToString(index)
	field = "w" + ToString(index)
	value = "x" + ToString(index)
	conn.Send("HSet", key, field, value)
	conn.Send("HGet", key, field)

	_, e := conn.Do("")
	if e != nil {
		fmt.Println(e.Error)
	}
}
func redis_go3() {
	rd := new(samples.RedisDriver)
	rd.Connect("mapping-prod.j5q6io.0001.cnn1.cache.amazonaws.com.cn:6379", "", 0, 100, 100)
	//rd.Connect("127.0.0.1:6379", "", 0, 100, 100)
	samples.RedisRegister(CONST_Redisgo_name, rd)

	conn := samples.Rds[CONST_Redisgo_name].Pool.Get()
	if conn == nil {
		fmt.Println("Pool no client")
		return
	}
	defer conn.Close()

	var index string = "0"
	var count int
	for {
		r, e := conn.Do("scan", index)
		if e != nil {
			fmt.Println(e.Error())
			conn.Close()
			conn = samples.Rds[CONST_Redisgo_name].Pool.Get()
			continue
		}
		if vr, ok := r.([]interface{}); ok {
			if len(vr) != 2 {
				fmt.Println("error: vr != 2")
				continue
			}

			if i, ok := vr[0].([]byte); ok {
				//if i, ok := vr[0].(int); ok {
				in := string(i)
				index = in
				//fmt.Println("bbb", in)
				if in == "0" {
					break
				}
			}

			if v, ok := vr[1].([]interface{}); ok {
				for _, vt := range v {
					if vk, ok := vt.([]byte); ok {
						svk := string(vk)
						if len(svk) < 4 || string(svk[0:4]) != "freq" {
							continue
						}

						r, e := redigo.Int(conn.Do("get", svk))
						if e != nil {
							fmt.Println("err get", svk, e)
						}

						if r == 1 {
							_, e := conn.Do("del", svk)
							if e != nil {
								fmt.Println("err del", svk, e)
							} else {
								count++
								fmt.Println("deleted key:", svk)
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("deleted count:", count)
}
