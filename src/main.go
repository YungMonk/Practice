package main

import (
	"fmt"
	"time"
	"train"

	"github.com/go-redis/redis"
)

func main() {
	train.XpathParserEngine()
	// standalonRedisTest()
}

func errHandler(err error) {
	fmt.Printf("errHandler, error:%s\n", err.Error())
	panic(err.Error())
}

func standalonRedisTest() {
	fmt.Printf("standalon_redis_test")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
		errHandler(err)
	}
	fmt.Printf("ping result: %s\n", pong)

	fmt.Printf("----------------------------------------\n")

	// set / get test
	fmt.Printf("set/get test\n")
	err = client.Set("foo", "bar", 0).Err()
	if err != nil {
		fmt.Printf("try set key[foo] to value[bar] error[%s]\n",
			err.Error())
		errHandler(err)
	}

	err = client.Set("foo1", "bar1", time.Second*5).Err()
	if err != nil {
		fmt.Printf("try set key[foo1] to value[bar1] error[%s]\n",
			err.Error())
		errHandler(err)
	}

	// get value
	value, err := client.Get("foo").Result()
	if err != nil {
		fmt.Printf("try get key[foo] error[%s]\n", err.Error())
		// errHandler(err)
	}

	fmt.Printf("key[foo]'s value is %s\n", value)

	value, err = client.Get("foo1").Result()
	if err != nil {
		fmt.Printf("try get key[foo1] error[%s]\n", err.Error())
		// errHandler(err)
	}

	fmt.Printf("key[foo1]'s value is %s\n", value)

	value, err = client.Get("foo2").Result()
	if err != nil {
		fmt.Printf("try get key[foo2] error[%s]\n", err.Error())
		// errHandler(err)
	}

	fmt.Printf("key[foo2]'s value is %s\n", value)

	// get ttl
	duration, err := client.TTL("foo").Result()
	if err != nil {
		fmt.Printf("try get ttl of key[foo] error[%s]\n", err.Error())
		errHandler(err)
	}

	fmt.Printf("key[foo]' ttl is [%s] %fs\n",
		duration.String(), duration.Seconds())

	duration, err = client.TTL("foo1").Result()
	if err != nil {
		fmt.Printf("try get ttl of key[foo1] error[%s]\n", err.Error())
		errHandler(err)
	}

	fmt.Printf("key[foo1]' ttl is [%s] %ds\n",
		duration.String(), int64(duration.Seconds()))

	fmt.Printf("----------------------------------------\n")

	// list test
	fmt.Printf("list test\n")

	err = client.RPush("tqueue", "tmsg1").Err()
	if err != nil {
		fmt.Printf("rpush list[tqueue] error[%s]\n", err.Error())
		errHandler(err)
	}

	listLen, err := client.LLen("tqueue").Result()
	if err != nil {
		fmt.Printf("try get len of list[tqueue] error[%s]\n",
			err.Error())
		errHandler(err)
	}

	fmt.Printf("len of list[tqueue] is %d\n", listLen)

	result, err := client.BLPop(time.Second*1, "tqueue").Result()
	if err != nil {
		fmt.Printf("blpop list[tqueue] error[%s]\n", err.Error())
		errHandler(err)
	}
	fmt.Printf("blpop list[tqueue], value[%s]\n", result[1])

	fmt.Printf("----------------------------------------\n")

	fmt.Printf("hmap test\n")

	err = client.HSet("tmap", "1", "f1").Err()
	if err != nil {
		fmt.Printf("try hset map[tmap] field[1] error[%s]\n",
			err.Error())
		errHandler(err)
	}

	err = client.HSet("tmap", "2", "f2").Err()
	if err != nil {
		fmt.Printf("try hset map[tmap] field[2] error[%s]\n",
			err.Error())
		errHandler(err)
	}

	kvMap := make(map[string]interface{})
	kvMap["3"] = "f3"
	kvMap["4"] = "f4"

	err = client.HMSet("tmap", kvMap).Err()
	if err != nil {
		fmt.Printf("try mset map[tmap] field[3] field[4] error[%s]\n",
			err.Error())
		errHandler(err)
	}

	mapLen, err := client.HLen("tmap").Result()
	if err != nil {
		fmt.Printf("try get len of map[tmap] error[%s]\n", err.Error())
		errHandler(err)
	}
	fmt.Printf("len of map[tmap] is %d\n", mapLen)

	// get map value
	value, err = client.HGet("tmap", "1").Result()
	if err != nil {
		fmt.Printf("try get field[1] value of map[tmap] error[%s]\n",
			err.Error())
		errHandler(err)
	}

	fmt.Printf("field[1] value of map[tmap] is %s\n", value)

	// hgetall
	resultKv, err := client.HGetAll("tmap").Result()
	if err != nil {
		fmt.Printf("try hgetall map[tmap] error[%s]\n", err.Error())
		errHandler(err)
	}

	for k, v := range resultKv {
		fmt.Printf("map[tmap] %s = %s\n", k, v)
	}

	fmt.Printf("----------------------------------------\n")

	fmt.Printf("pubsub test\n")

	pubsub := client.Subscribe("test_channel")

	_, err = pubsub.Receive()
	if err != nil {
		fmt.Printf("try subscribe channel[test_channel] error[%s]\n", err.Error())
		errHandler(err)
	}

	// go channel to used to receive message
	ch := pubsub.Channel()

	// publish a message
	err = client.Publish("test_channel", "hello").Err()
	if err != nil {
		fmt.Printf("try publish message to channel[test_channel] error[%s]\n", err.Error())
		errHandler(err)
	}

	time.AfterFunc(time.Second*2, func() {
		_ = pubsub.Close()
	})

	// consume message
	for {
		msg, ok := <-ch
		if !ok {
			break
		}

		fmt.Printf("recv message[%s] from channel[%s]\n", msg.Payload, msg.Channel)
	}
}
