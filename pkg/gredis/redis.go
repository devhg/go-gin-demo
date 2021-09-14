package gredis

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/devhg/go-gin-demo/pkg/config"
)

var RedisConn *redis.Pool

func New(serviceName string) error {
	servicer := config.GetServicer(serviceName)
	redisConf := servicer.(*config.Redis)

	RedisConn = &redis.Pool{
		MaxIdle:     redisConf.Config.MaxIdle,
		MaxActive:   redisConf.Config.MaxActive,
		IdleTimeout: time.Duration(redisConf.Config.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConf.Config.Host)
			if err != nil {
				return nil, err
			}
			if redisConf.Config.Password != "" {
				if _, err := c.Do("AUTH", redisConf.Config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	reply, err := redis.String(conn.Do("SET", key, value))
	if err != nil {
		return reply, err
	}

	_, _ = conn.Do("EXPIRE", key, time)
	return reply, err
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

// Hash store
// 利用redis库自带的Args 和 AddFlat对结构体进行转换。然后以hash类型存储。
// 该方式实现简单，但存在最大的问题是不支持数组结构（如：结构体中内嵌结构体、数组等）。
func DoHashStore(key string, src interface{}) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.String(conn.Do("hmset", redis.Args{key}.AddFlat(src)...))
}

func DoHashGet(key string, dest interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := redis.Values(conn.Do("hgetall", key))
	if err != nil {
		return err
	}
	return redis.ScanStruct(value, dest)
}

// Gob Encoding
func DoGobStore(key string, src interface{}) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(src)

	return redis.String(conn.Do("set", key, buffer.Bytes()))
}

func DoGobGet(key string, dest interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	reBytes, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return err
	}

	reader := bytes.NewReader(reBytes)
	decoder := gob.NewDecoder(reader)
	return decoder.Decode(dest)
}

// JSON Encoding
func DoJSONStore(key string, src interface{}) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	datas, err := json.Marshal(src)
	if err != nil {
		return "", err
	}
	return redis.String(conn.Do("set", key, datas))
}

func DoJSONGet(key string, dest interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	datas, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return err
	}
	return json.Unmarshal(datas, dest)
}
