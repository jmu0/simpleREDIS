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
