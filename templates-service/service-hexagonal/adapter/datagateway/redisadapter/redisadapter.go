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

type redisAdapter struct {
	rd *redis.Client
}

func New(conf config.RedisConf) datagateway.DataGatewayTodo {
	rd := redis.NewClient(&redis.Options{
		Addr: confAddress(conf),
	})

	if rd == nil {
		panic("nil redis client")
	}

	return &redisAdapter{rd: rd}
}

func keyTodo(todoId string) string {
	return fmt.Sprintf("todo:%s", todoId)
}

func confAddress(conf config.RedisConf) string {
	return fmt.Sprintf("localhost:%d", conf.Port)
}

func (r *redisAdapter) CreateTodo(
	ctx context.Context,
	todo entity.Todo,
) error {
	b, err := json.Marshal(todo)
	if err != nil {
		return errors.Wrap(err, "failed to marshal todo to json")
	}

	k := keyTodo(todo.Id)

	if err := r.rd.Set(ctx, k, b, 0).Err(); err != nil {
		return errors.Wrap(err, "failed to save todo to redis")
	}

	return nil
}

func (r *redisAdapter) GetTodo(ctx context.Context, id string) (entity.Todo, error) {
	k := keyTodo(id)

	result := r.rd.Get(ctx, k)
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

func (r *redisAdapter) UpdateTodo(
	ctx context.Context,
	id string,
	update entity.Todo,
) (
	entity.Todo,
	error,
) {
	if update.Id != id {
		return entity.Todo{}, fmt.Errorf("id and update.Id differ: %s vs %s", id, update.Id)
	}

	k := keyTodo(id)
	b, err := json.Marshal(update)
	if err != nil {
		return entity.Todo{}, errors.Wrap(err, "failed to marshal todo json")
	}

	result := r.rd.GetSet(ctx, k, b)
	if err := result.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return entity.Todo{}, errors.Wrapf(err, "no such existing todo id %s", id)
		}

		return entity.Todo{}, errors.Wrap(err, "failed to update todo")
	}

	oldTodoJson := result.Val()
	oldTodo := new(entity.Todo)

	if err := json.Unmarshal([]byte(oldTodoJson), oldTodo); err != nil {
		return entity.Todo{}, errors.Wrap(err, "failed to unmarshal old todo json from redis")
	}

	return *oldTodo, nil
}

func (r *redisAdapter) DeleteTodo(ctx context.Context, id string) error {
	k := keyTodo(id)

	c, err := r.rd.Del(ctx, k).Result()
	if err != nil {
		return errors.Wrapf(err, "failed to delete todo id %s", id)
	}
	if c == 0 {
		return fmt.Errorf("no such todo id %s to delete", id)
	}

	return nil
}
