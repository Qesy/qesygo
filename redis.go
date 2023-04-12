package qesygo

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type CacheRedis struct {
	Pool     *redis.Pool // redis connection pool
	Conninfo string
	Auth     string
}

func (cr *CacheRedis) newPool() {
	cr.Pool = &redis.Pool{
		MaxIdle:     30,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cr.Conninfo)
			if err != nil {
				return nil, err
			}
			if cr.Auth != "" {
				if _, err := c.Do("AUTH", cr.Auth); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
	}
}

// var pool = newPool()
func (cr *CacheRedis) Connect() error {
	cr.newPool()
	c := cr.Pool.Get()
	defer c.Close()
	return c.Err()
}

func (cr *CacheRedis) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := cr.Pool.Get()
	defer c.Close()
	return c.Do(commandName, args...)
}

func (cr *CacheRedis) FlushAll() error {
	_, err := cr.do("FLUSHALL")
	return err
}

func (cr *CacheRedis) Get(key string) (string, error) {
	str, err := redis.String(cr.do("GET", key))
	return str, err
}

func (cr *CacheRedis) Set(key string, value string) error {
	_, err := cr.do("SET", key, value)
	return err
}

func (cr *CacheRedis) SetNx(key string, value string) error {
	return cr.SetNx(key, value)
}

func (cr *CacheRedis) Del(key ...string) error {
	ArgsArr := redis.Args{}
	for _, v := range key {
		ArgsArr = append(ArgsArr, v)
	}
	_, err := cr.do("DEL", ArgsArr...)
	return err
}

func (cr *CacheRedis) Exists(key string) (bool, error) {
	return redis.Bool(cr.do("EXISTS", key))
}

func (cr *CacheRedis) Expire(key string, second int) (bool, error) {
	return redis.Bool(cr.do("EXPIRE", key, second))
}

func (cr *CacheRedis) Keys(key string) ([]string, error) {
	return redis.Strings(cr.do("KEYS", key))
}

func (cr *CacheRedis) Ttl(key string) (interface{}, error) {
	return redis.Int(cr.do("TTL", key))
}

func (cr *CacheRedis) HMset(key string, arr interface{}) (interface{}, error) {
	return cr.do("HMSET", redis.Args{}.Add(key).AddFlat(arr)...)
}

func (cr *CacheRedis) HGet(key string, subKey string) (interface{}, error) {
	return cr.do("HGET", key, subKey)
}

func (cr *CacheRedis) HGetAll(key string) (map[string]string, error) {
	rsByte, err := cr.do("HGETALL", key)
	if err != nil {
		return nil, err
	}
	rs, err := redis.StringMap(rsByte, err)
	if err != nil {
		return nil, err
	}
	return rs, err
}

func (cr *CacheRedis) SAdd(Key string, Val string) (int, error) {
	return redis.Int(cr.do("SADD", Key, Val))
}

func (cr *CacheRedis) SCard(Key string) (interface{}, error) { // 返回集合 key 的基数(集合中元素的数量)。
	return redis.Int(cr.do("SCARD", Key))
}

func (cr *CacheRedis) SMembers(Key string, Val string) (interface{}, error) { //返回集合 key 中的所有成员
	return redis.Strings(cr.do("SMEMBERS", Key, Val))
}

func (cr *CacheRedis) SIsMembers(Key string, Val string) (int, error) { //是否集合 key 的成员
	return redis.Int(cr.do("SISMEMBER", Key, Val))
}

func (cr *CacheRedis) SRem(Key string, Val ...string) (interface{}, error) { //移除集合 key 中的一个或多个 member 元素
	return redis.Int(cr.do("SREM", Key, Val))
}

func (cr *CacheRedis) ZAdd(Key string, Score int64, Name string) (interface{}, error) {
	return cr.do("ZADD", Key, Score, Name)
}

func (cr *CacheRedis) ZRem(Key string, Name string) (interface{}, error) {
	return cr.do("zRem", Key, Name)
}

func (cr *CacheRedis) ZCard(Key string) (int, error) {
	return redis.Int(cr.do("ZCARD", Key))
}

func (cr *CacheRedis) ZRank(Key string, Name string) (int, error) {
	return redis.Int(cr.do("ZRANK", Key, Name))
}

func (cr *CacheRedis) ZRevRank(Key string, Name string) (int, error) {
	return redis.Int(cr.do("ZREVRANK", Key, Name))
}

func (cr *CacheRedis) ZRange(Key string, Start int32, End int32) ([][2]int64, error) {
	rank := [][2]int64{}
	if rsByte, err := cr.do("ZRANGE", Key, Start, End, "WITHSCORES"); err == nil {
		rsByteArr := rsByte.([]interface{})
		for i := 0; i < len(rsByteArr)/2; i++ {
			index := i * 2
			key := StrToInt64(string(rsByteArr[index].([]byte)))
			val := StrToInt64(string(rsByteArr[index+1].([]byte)))
			rank = append(rank, [2]int64{key, val})
		}
		return rank, err
	} else {
		return rank, err
	}
}

func (cr *CacheRedis) ZRevRange(Key string, Start int32, End int32) ([][2]int64, error) {
	rank := [][2]int64{}
	if rsByte, err := cr.do("ZREVRANGE", Key, Start, End, "WITHSCORES"); err == nil {
		rsByteArr := rsByte.([]interface{})
		for i := 0; i < len(rsByteArr)/2; i++ {
			index := i * 2
			key := StrToInt64(string(rsByteArr[index].([]byte)))
			val := StrToInt64(string(rsByteArr[index+1].([]byte)))
			rank = append(rank, [2]int64{key, val})
		}
		return rank, err
	} else {
		return rank, err
	}
}

func (cr *CacheRedis) ZCount(Key string, Min int32, Max int32) (int, error) {
	return redis.Int(cr.do("ZCOUNT", Key, Min, Max))
}
