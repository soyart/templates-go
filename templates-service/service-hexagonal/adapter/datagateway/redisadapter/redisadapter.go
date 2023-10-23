package redisadapter

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"example.com/servicehex/config"
	"example.com/servicehex/domain/datagateway"
	"example.com/servicehex/domain/entity"
)

type adapterRedis struct {
	rd *redis.Client
}

func New(conf config.RedisConf) datagateway.DataGatewayTodo {
	rd := redis.NewClient(&redis.Options{
		Addr: confAddress(conf),
	})

	if rd == nil {
		panic("nil redis client")
	}

	return &adapterRedis{rd: rd}
}

func baseKeyTodo(userID string) string {
	return "todo:" + userID
}

func confAddress(conf config.RedisConf) string {
	return fmt.Sprintf("localhost:%d", conf.Port)
}

func (r *adapterRedis) CreateTodo(
	ctx context.Context,
	todo entity.Todo,
) error {
	b, err := json.Marshal(todo)
	if err != nil {
		return errors.Wrap(err, "failed to marshal todo to json")
	}

	if err := r.rd.HSet(ctx, baseKeyTodo(todo.UserID), todo.ID, b, 0).Err(); err != nil {
		return errors.Wrap(err, "failed to save todo to redis")
	}

	return nil
}

func (r *adapterRedis) GetTodo(ctx context.Context, userID, todoID string) (entity.Todo, error) {
	result := r.rd.HGet(ctx, baseKeyTodo(userID), todoID)
	if err := result.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return entity.Todo{}, errors.Wrapf(err, "no such todo id")
		}

		return entity.Todo{}, errors.Wrap(err, "failed to get todo")
	}

	val := result.Val()
	todo := new(entity.Todo)

	if err := json.Unmarshal([]byte(val), todo); err != nil {
		return entity.Todo{}, errors.Wrap(err, "failed to unmarshal todo json from redis")
	}

	return *todo, nil
}

func (r *adapterRedis) GetTodos(ctx context.Context, userID string) ([]entity.Todo, error) {
	return nil, errors.New("not implemented")
}

func (r *adapterRedis) UpdateTodo(
	ctx context.Context,
	userID string,
	id string,
	update entity.Todo,
) error {
	key := baseKeyTodo(userID)

	exists, err := r.rd.HExists(ctx, key, id).Result()
	if !exists || errors.Is(err, redis.Nil) {
		return errors.Wrapf(err, "no such existing todo id %s, and failed to rollback (delete) updated todo", id)
	}

	b, err := json.Marshal(update)
	if err != nil {
		return errors.Wrap(err, "failed to marshal todo json")
	}

	if err := r.rd.HSet(ctx, key, id, b).Err(); err != nil {
		return errors.Wrap(err, "failed to update todo")
	}

	return nil
}

func (r *adapterRedis) DeleteTodo(ctx context.Context, userID, id string) error {
	c, err := r.rd.HDel(ctx, baseKeyTodo(userID), id).Result()
	if err != nil {
		return errors.Wrapf(err, "failed to delete todo id %s", id)
	}
	if c == 0 {
		return fmt.Errorf("no such todo id %s to delete", id)
	}

	return nil
}

func (r *adapterRedis) DeleteTodos(ctx context.Context, userID string) error {
	err := r.rd.Del(ctx, baseKeyTodo(userID)).Err()
	if err != nil {
		return errors.Wrapf(err, "failed to delete todos for user %s", userID)
	}

	return nil
}
