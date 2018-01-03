package simpleREDIS

import (
	"github.com/go-redis/redis"
)

//NewRedis gets new Redis, host=host:port, default port=6379
func NewRedis(host string) (Redis, error) {
	var r = Redis{
		serverURL: host,
	}
	err := r.connect()
	return r, err
}

//Redis simple redis get/set
type Redis struct {
	serverURL string
	client    *redis.Client
}

func (r *Redis) connect() error {
	r.client = redis.NewClient(&redis.Options{
		Addr:     r.serverURL,
		Password: "",
		DB:       0,
	})
	_, err := r.client.Ping().Result()
	return err
}

//Set sets key/value
func (r *Redis) Set(key, value string) error {
	return r.client.Set(key, value, 0).Err()
}

//Get gets value of key
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

//Del deletes key
func (r *Redis) Del(key string) (int64, error) {
	return r.client.Del(key).Result()
}

//Scan returns all keys
func (r *Redis) Scan(match string) ([]string, error) {
	var ret, keys []string
	var err error
	var cursor uint64
	for {
		keys, cursor, err = r.client.Scan(cursor, match, 10).Result()
		if err != nil {
			return ret, err
		}
		ret = append(ret, keys...)
		if cursor == 0 {
			break
		}
	}
	return ret, nil
}

//Rpush adds to list
func (r *Redis) Rpush(key, value string) int64 {
	return r.client.RPush(key, value).Val()
}

//GetList returns all items in list
func (r *Redis) GetList(key string) ([]string, error) {
	return r.GetRange(key, 0, -1)
}

//GetRange returns all items in list
func (r *Redis) GetRange(key string, from, to int64) ([]string, error) {
	res := r.client.LRange(key, from, to)
	if res.Err() != nil {
		return make([]string, 0), res.Err()
	}
	return res.Val(), nil
}

//GetType gets type for keys
func (r *Redis) GetType(key string) string {
	return r.client.Type(key).Val()
}
