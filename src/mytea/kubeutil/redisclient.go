package kubeutil

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"regexp"
	"strings"
)

type MyRedisClient struct {
	rdb *redis.Client
}

type redisConnector interface {
	getHost() string
	getPort() string
	getPass() string
}

func NewRedisClient(connector redisConnector) MyRedisClient {
	myRedisClient := MyRedisClient{
		redis.NewClient(&redis.Options{
			Addr:     connector.getHost() + ":" + connector.getPort(),
			Password: connector.getPass(),
		}),
	}
	return myRedisClient
}

func (myRedisClient MyRedisClient) exec(cmdText string) string {
	strs := splitBySpace(cmdText)
	args := make([]interface{}, len(strs))
	for i := range strs {
		args[i] = strs[i]
	}

	cmd := myRedisClient.rdb.Do(context.TODO(), args...)
	results, err := cmd.StringSlice()
	if err == nil {
		builder := strings.Builder{}
		for _, result := range results {
			builder.WriteString(result)
			builder.WriteString("\n")
		}
		return builder.String()
	}
	return cmd.String()
}

func splitBySpace(cmd string) []string {
	cmd = strings.TrimSpace(cmd)
	reg, err := regexp.Compile("\\s+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.Split(cmd, -1)
}

type manualConnector struct {
}

func (manual manualConnector) getHost() string {
	return "r-bp15upyxyulbg2y4drpd.redis.rds.aliyuncs.com"
}

func (manual manualConnector) getPort() string {
	return "6379"
}

func (manual manualConnector) getPass() string {
	return "Asp123456"
}
