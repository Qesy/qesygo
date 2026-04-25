package qesygo

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type CacheRedis struct {
	Pool     *redis.Pool // redis connection pool
	Conninfo string
	Auth     string
}

func (cr *CacheRedis) newPool() {
	cr.Pool = &redis.Pool{
		MaxIdle:     30,
		MaxActive:   100, // 🔥 建议加上
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
	reply, err := cr.do("GET", key)
	if err != nil {
		return "", err
	}
	if reply == nil {
		return "", nil
	}
	return string(reply.([]byte)), err
}

func (cr *CacheRedis) Set(key string, value string) error {
	_, err := cr.do("SET", key, value)
	return err
}

func (cr *CacheRedis) SetNx(key string, value string) (int, error) {
	return redis.Int(cr.do("SETNX", key, value))
}

func (cr *CacheRedis) SetEx(key string, expire int, value string) (string, error) {
	return redis.String(cr.do("SETEX", key, expire, value))
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

func (cr *CacheRedis) HGet(key string, subKey string) (string, error) {
	return redis.String(cr.do("HGET", key, subKey))
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

func (cr *CacheRedis) SCard(Key string) (int, error) { // 返回集合 key 的基数(集合中元素的数量)。
	return redis.Int(cr.do("SCARD", Key))
}

func (cr *CacheRedis) SMembers(Key string) ([]string, error) { //返回集合 key 中的所有成员
	return redis.Strings(cr.do("SMEMBERS", Key))
}

func (cr *CacheRedis) SIsMembers(Key string, Val string) (int, error) { //是否集合 key 的成员
	return redis.Int(cr.do("SISMEMBER", Key, Val))
}

func (cr *CacheRedis) SRem(Key string, Val ...string) (int, error) { //移除集合 key 中的一个或多个 member 元素
	args := redis.Args{}.Add(Key)
	for _, v := range Val {
		args = args.Add(v)
	}
	return redis.Int(cr.do("SREM", args...))
}

// ZSET
func (cr *CacheRedis) ZAdd(Key string, Score int64, Name string) (int, error) {
	return redis.Int(cr.do("ZADD", Key, Score, Name))
}

func (cr *CacheRedis) ZRem(Key string, Name string) (int, error) {
	return redis.Int(cr.do("zRem", Key, Name))
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

// List
func (cr *CacheRedis) BLPop(keys ...string) (string, error) {
	return redis.String(cr.do("BLPOP", keys))
}

func (cr *CacheRedis) BRPop(keys ...string) (string, error) {
	return redis.String(cr.do("BRPOP", keys))
}

func (cr *CacheRedis) BRPopLPush(source string, destination string, timeout int) (string, error) {
	return redis.String(cr.do("BRPOPLPUSH", source, destination, timeout))
}

func (cr *CacheRedis) LIndex(key string, index int64) (string, error) {
	return redis.String(cr.do("LINDEX", key, index))
}

func (cr *CacheRedis) LInsert(key string, op string, pivot string, value string) (int, error) {
	return redis.Int(cr.do("LINSERT", key, op, pivot, value))
}

func (cr *CacheRedis) LLen(key string) (int64, error) {
	return redis.Int64(cr.do("LLEN", key))
}

func (cr *CacheRedis) LPop(key string) (string, error) {
	return redis.String(cr.do("LPOP", key))
}

func (cr *CacheRedis) LPush(key string, value string) (int, error) {
	return redis.Int(cr.do("LPUSH", key, value))
}

func (cr *CacheRedis) LPushX(key string, value string) (int, error) {
	return redis.Int(cr.do("LPUSHX", key, value))
}

func (cr *CacheRedis) LRange(key string, start int64, stop int64) ([]string, error) {
	return redis.Strings(cr.do("LRANGE", key, start, stop))
}

func (cr *CacheRedis) LRem(key string, count int64, value string) (int, error) {
	return redis.Int(cr.do("LREM", key, count, value))
}

func (cr *CacheRedis) LSet(key string, index int64, value string) error {
	_, err := cr.do("LSET", key, index, value)
	return err
}

func (cr *CacheRedis) LTRIM(key string, start int64, stop int64) error {
	_, err := cr.do("LTRIM", key, start, stop)
	return err
}

func (cr *CacheRedis) RPop(key string) (string, error) {
	return redis.String(cr.do("RPOP", key))
}

func (cr *CacheRedis) RPopLPush(source string, destination string) (string, error) {
	return redis.String(cr.do("RPOPLPUSH", source, destination))
}

func (cr *CacheRedis) RPush(key string, value string) (int, error) {
	return redis.Int(cr.do("RPUSH", key, value))
}

func (cr *CacheRedis) RPushX(key string, value string) (int, error) {
	return redis.Int(cr.do("RPUSHX", key, value))
}

// ================= 队列封装（🔥核心） =================

// 入队
func (cr *CacheRedis) QueuePush(queue string, data string) error {
	_, err := cr.LPush(queue, data)
	return err
}

// 安全消费
func (cr *CacheRedis) QueuePop(queue, processing string) (string, error) {
	return cr.BRPopLPush(queue, processing, 0)
}

// ACK
func (cr *CacheRedis) QueueAck(processing, data string) error {
	_, err := cr.LRem(processing, 1, data)
	return err
}

// 重试
func (cr *CacheRedis) QueueRetry(queue, processing, data string) error {
	_, _ = cr.LRem(processing, 1, data)
	_, err := cr.LPush(queue, data)
	return err
}

// 恢复（启动时）
func (cr *CacheRedis) QueueRecover(queue, processing string) error {
	tasks, err := cr.LRange(processing, 0, -1)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		cr.LPush(queue, task)
	}

	return cr.Del(processing)
}
