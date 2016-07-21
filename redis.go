package QesyGo

import(
	"time"
	"github.com/garyburd/redigo/redis"
	//"fmt"

)

type CacheRedis struct{
	Pool 		*redis.Pool // redis connection pool
	conninfo 	string
	key      		string
}

func newPool() *redis.Pool {
    return &redis.Pool{
        MaxIdle: 30,
        IdleTimeout: 240 * time.Second,
        Dial: func () (redis.Conn, error) {
            c, err := redis.Dial("tcp", "127.0.0.1:6379")
            if err != nil {
                return nil, err
            }
            return c, err
        },
    }
}

var pool = newPool()

func RedisGet(key string) (string, error){
	conn := pool.Get()
    	defer conn.Close()
	str, err := redis.String(conn.Do("GET", key))
	return str, err
}

func RedisSet(key string, value string) error {
	conn := pool.Get()
    	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	return err
}

func RedisDel(key string) error{
	conn := pool.Get()
    	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}
 
func RedisExists(key string) (bool, error){
	conn := pool.Get()
    	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}

func RedisExpire(key string, second int) (bool, error){
	conn := pool.Get()
    	defer conn.Close()
	return redis.Bool(conn.Do("EXPIRE", key, second))
}

func RedisKeys(key string) (interface{}, error){
	conn := pool.Get()
    	defer conn.Close()
	return conn.Do("KEYS", key)
}

func RedisTtl(key string) (interface{}, error){
	conn := pool.Get()
    	defer conn.Close()
	return redis.Int(conn.Do("TTL", key))
}

func RedisHMset(key string, arr interface{})(interface{}, error){
	conn := pool.Get()
    	defer conn.Close()
	return conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(arr)...)
}

func RedisHGet(key string, subKey string)(interface{}, error){
	conn := pool.Get()
    	defer conn.Close()
	return conn.Do("HGET", key, subKey)
}

func RedisHGetAll(key string)(interface{}, error){
	conn := pool.Get()
    	defer conn.Close()
	return conn.Do("HGETALL", key)
}


//-- 暂时还缺少zadd --