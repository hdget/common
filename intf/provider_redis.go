package intf

import (
	"github.com/hdget/common/protobuf"
)

type RedisCommand struct {
	Name string
	Args []interface{}
}

type RedisProvider interface {
	Provider
	My() RedisClient
	By(string) RedisClient
}

type RedisClient interface {
	// general purpose
	Del(key string) error
	Dels(keys []string) error
	Exists(key string) (bool, error)
	Expire(key string, expire int) error
	Incr(key string) error
	IncrBy(key string, number int) error
	DecrBy(key string, number int) error
	Ttl(key string) (int64, error)
	Pipeline(commands []*RedisCommand) (reply interface{}, err error)
	Ping() error

	// Set operations
	Set(key string, value interface{}) error
	SetEx(key string, value interface{}, expire int) error

	// Get operations
	Get(key string) ([]byte, error)
	GetInt(key string) (int, error)
	GetInt64(key string) (int64, error)
	GetFloat64(key string) (float64, error)
	GetString(key string) (string, error)

	// HashMap operations
	HGetAll(key string) (map[string]string, error)
	HGet(key string, field any) ([]byte, error)
	HGetInt(key string, field string) (int, error)
	HGetInt64(key string, field string) (int64, error)
	HGetFloat64(key string, field string) (float64, error)
	HGetString(key string, field string) (string, error)
	HMGet(key string, fields []string) ([][]byte, error)
	HSet(key string, field interface{}, value interface{}) (int, error)
	HMSet(key string, args map[string]interface{}) error
	HDel(key string, field interface{}) (int, error)
	HDels(key string, fields []interface{}) (int, error)
	HLen(key string) (int, error)

	// set
	SIsMember(key string, member interface{}) (bool, error)
	SAdd(key string, members interface{}) error
	SRem(key string, members interface{}) error
	SInter(keys []string) ([]string, error)
	SUnion(keys []string) ([]string, error)
	SDiff(keys []string) ([]string, error)
	SMembers(key string) ([]string, error)

	// zset
	ZAdd(key string, score int64, member interface{}) error
	ZCard(key string) (int, error)
	ZRange(key string, min, max int64) ([]string, error)
	ZRemRangeByScore(key string, min, max interface{}) error
	ZRangeByScore(key string, min, max interface{}, withScores bool, list *protobuf.ListParam) ([]string, error)
	ZScore(key string, member interface{}) (int64, error)
	ZInterstore(newKey string, keys ...interface{}) (int64, error)
	ZIncrBy(key string, increment int64, member interface{}) error
	ZRem(destKey string, members ...interface{}) (int64, error)

	// list
	LPush(key string, values ...any) error
	RPush(key string, values ...any) error
	RPop(key string) ([]byte, error)
	LRangeInt64(key string, start, end int64) ([]int64, error)
	LRangeString(key string, start, end int64) ([]string, error)
	LLen(key string) (int64, error)
	Eval(scriptContent string, keys []interface{}, args []interface{}) (interface{}, error)

	// redis bloom
	BfExists(key string, item string) (exists bool, err error)
	BfAdd(key string, item string) (exists bool, err error)
	BfReserve(key string, errorRate float64, capacity uint64) (err error)
	BfAddMulti(key string, items []interface{}) ([]int64, error)
	BfExistsMulti(key string, items []interface{}) ([]int64, error)
}
