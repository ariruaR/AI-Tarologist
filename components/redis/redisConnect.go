package redis

import (
	"bot/components/models"
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	configReader "bot/config"

	redis "github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

var config = configReader.Readconfig()

func NewClient() *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_ADDR,
		Password: config.REDIS_PASSWORD,
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
		log.Printf("Ошибка при создании пользователя, %s", err)
		return err
	}
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return err
	}
	switch field {
	case "username":
		if username, ok := value.(string); ok {
			user.Username = username
		} else {
			return errors.New("неверный тип данных для добавления в поле")
		}
	case "isPremium":
		if isPremium, ok := value.(bool); ok {
			user.IsPremium = isPremium
		} else {
			return errors.New("неверный тип данных для добавления в поле")
		}
	case "info":
		if info, ok := value.(string); ok {
			user.Info = info
		} else {
			return errors.New("неверный тип данных для добавления в поле")
		}
	case "LastPayment":
		if lastPayment, ok := value.(map[string]string); ok {
			user.LastPayment = lastPayment
		} else {
			return errors.New("неверный тип данных для добавления в поле")
		}
	default:
		return errors.New("такого поля не существует")
	}

	updatedUserData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, updatedUserData, expiration).Err()
}

func (r *redisClient) ReadUser(ctx context.Context, id int) (models.User, error) {
	key := strconv.Itoa(id)
	var user models.User
	userData, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return models.User{}, err
	}
	if err := json.Unmarshal([]byte(userData), &user); err != nil {
		return models.User{}, err
	}
	return user, nil
}
