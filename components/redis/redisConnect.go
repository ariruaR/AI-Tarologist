package redis

import (
	"bot/components/models"
	"context"
	"encoding/json"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

func NewClient() *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &redisClient{
		client: rdb,
	}
}

func (r *redisClient) Setter(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
func (r *redisClient) Getter(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
func (r *redisClient) SetNewUser(
	ctx context.Context,
	ID int,
	Username string,
	isPremium bool,
	Info string,
	PaymentHistory map[string]string,
	expiration time.Duration,
) error {
	new_user, err := models.NewUser(ID, Username, isPremium, Info, nil)
	if err != nil {
		panic(err)
	}
	return r.client.Set(ctx, strconv.Itoa(ID), new_user, expiration).Err()
}

func (r *redisClient) UpdateFieldUser(
	ctx context.Context,
	ID int,
	field string,
	value interface{},
	expiration time.Duration,
) error {
	key := strconv.Itoa(ID)
	var user models.User
	userData, err := r.client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return err
	}

	switch field {
	case "username":
		if username, ok := value.(string); ok {
			user.Username = username
		} else {
			return err
		}
	case "isPremium":
		if isPremium, ok := value.(bool); ok {
			user.IsPremium = isPremium
		} else {
			return err
		}
	case "info":
		if info, ok := value.(string); ok {
			user.Info = info
		} else {
			return err
		}
	case "paymentHistory":
		if paymentHistory, ok := value.(map[string]string); ok {
			user.PaymentHistory = paymentHistory
		} else {
			return err
		}
	default:
		return err
	}

	updatedUserData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, updatedUserData, expiration).Err()
}
