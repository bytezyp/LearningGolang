package main

import (
	"github.com/go-redis/redis"
	"time"
	"BoomFilters"
	"os"
	"fmt"
	"math/rand"
)
func GetRandomString(leng int) string {
	allStr := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(allStr)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Println(r, "----", r.Intn(len(bytes)), "----",bytes)
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
func main()  {

	redisOption := &redis.Options{
		Addr:               "192.168.99.64:6379",
		DB:                 0,
		DialTimeout:        10 * time.Second,
		ReadTimeout:        30 * time.Second,
		WriteTimeout:       30 * time.Second,
		PoolSize:           10,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: 500 * time.Millisecond,
	}
	boom := BoomFilters.NewBoom("boomkey",2,redisOption)

	//boomTest := make([]string, 10)
	//boomTest := []string{"444", "555", "444", "837d8f1dc99b13938a91fce3c2544174"}
	//file,_ := os.Open("fileTest.txt")
	//inputReader := bufio.NewReader(file)
	delFile,_ := os.Create("delfile.txt")
	startTime := time.Now().Unix()
	for i := 1 ; i < 1000; i++ {
		//inputString, readerError := inputReader.ReadString('\n')
		//if readerError == io.EOF {
		//	return
		//}
		inputString := GetRandomString(32)
		flag := boom.BoomFilter(inputString)
		if !flag {
			delFile.Write([]byte(inputString))
		}
		//if offset == 789155204 || offset ==253343560 ||offset==1720160103 {
		//	fmt.Println(offset, "---", inputString)
		//}
	}
	defer delFile.Close()
	sum := time.Now().Unix() - startTime
	fmt.Println(boom.Count())
	fmt.Println(sum)
	return
	//for i := 0; i <= 3 ; i++ {
	//	//fmt.Println(strconv.Itoa(rand.Intn(100000)))
	//	//str := strconv.Itoa(rand.Intn(100000))
	//	//boomTest[i] = str
	//	fmt.Println(boom.BoomFilter(boomTest[i]), boomTest[i])
	//
	//	//boomTest = append(boomTest, strconv.Itoa(rand.Intn(100000)))
	//}
	//fmt.Println(boomTest)
	return



	//redisOption := &redis.Options{
	//	Addr:               "192.168.99.64:6379",
	//	DB:                 15,
	//	DialTimeout:        10 * time.Second,
	//	ReadTimeout:        30 * time.Second,
	//	WriteTimeout:       30 * time.Second,
	//	PoolSize:           10,
	//	PoolTimeout:        30 * time.Second,
	//	IdleTimeout:        500 * time.Millisecond,
	//	IdleCheckFrequency: 500 * time.Millisecond,
	//}
	//
	//red := redis.NewClient(redisOption)
	//ret,err := red.Set("rediskey", "redisValue",0).Result()
	//
	//fmt.Println(ret, err)
	//val,err := red.Get("rediskey").Result()
	//fmt.Println(val)
	return
	//boom := BoomFilters.Boom{}
	//fmt.Println(boom.M)
	//str := "seedTest"
	//seed := BoomFilters.GetBKDRHashSeed(len(str))
	//strbin := BoomFilters.BKDRHash(str,seed)
	//fmt.Println(strbin)
	//fmt.Println(int(rune(str)))

}