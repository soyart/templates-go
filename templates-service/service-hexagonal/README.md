# service-hexagonal

An example project for simple Go TODO back-end API,
organized with hexagonal architecture (ports-and-adapters).

## Hexagonal architecture (HexArch)

> Note: HexArch has nothing to do with number 6 or hexagons

HexArch is mainly concerned with seperation of external systems,
that is, code for connecting to other systems, e.g. databases,
messengers, etc must be pushed to the edge (outter ring) and
must not taint business-related code.

This leaves business code 0 knowledge about the external systems,
i.e. **the business-level knows nothing about databases being used,
and access to external systems is only done via interfaces**.

This gives developers freedom to change their 3rd party
systems at any times, preserving the business logic.

Data flow in HexArch always flows from _driving_ parts
(e.g. presentation layer like HTTP REST and gRPC)
through _core_ to _driven_ parts.

Both sides of core can have ports and adapters.

### My convention

- No technical details is known in the business level,
  and that every implemenations and their fields are private.

- For database:

  - Port interfaces are prefixed with `DataGateway`, e.g. `DataGatewayUser`
    is the interface for writing/querying users in our systems.

  - Adapters are prefixed with `adapter` (unexported),
    or `dataStore`, e.g. `adapterSql` and `adapterRedis`,
    `dataStoreSql`, `dataStoreRedis`.

  - Field names of type `DataGateway*` are prefixed with `repo`,
    e.g:

    ```go
    type DataGatewayUser interface {
      // code
    }

    type adapterRedis struct {
      // implements DataGatewayUser
    }

    type serviceUser struct {
      repo DataGatewayUser
    }
    ```

- For service:

  - Port interfaces are prefixed with `Port`, e.g. `PortUser`

  - Adapters (implementations of ports) are prefixed with `service`,
    e.g. `serviceUser` implements `PortUser`

### HexArch: ports and adapters

> Like with USB ports, the adapters connects to the ports.

When data must move between layers, ports and adapters
are used to move and translate data from one layer to the next.

Ports are entrypoint to the core business code,
much like a _gateway_ for our data entering the application,
and are usually defined as interfaces.

The _adapters_ then make use of these _port interfaces_
by implementing them.

If data at the different side of the port/adapter can be shared,
then HexArch allows it to be shared.

A single _port_ can be connected to >1 adapters, i.e. a single
interface and be implemented by many classes of objects.

### Example: driven side (database-side)

Let's say we have a port interface for persistent data called
`DataGatewayTodo`:

```go
package datagateway

// code

// *Port*
type DataGatewayTodo interface {
  CreateTodo(entity.Todo) error
  GetTodo(string) (entity.Todo, error)
  UpdateTodo(string, entity.Todo) error
  DeleteTodo(string) (entity.Todo, error)
}
```

Then we can have multiple connectors implementing `DataGatewayTodo`,
first we have Redis:

```go
package rediswrapper

import (
  "example.com/servicehex/domain/entity"
  "example.com/servicehex/domain/datagateway"

  "github.com/redis/redis-go"
)

// *Adapter*
type dataStoreRedis struct {
  rd *redis.Client
}

func new(conf *Config) (datagateway.DataGatewayTodo, error) {..}
func (d *dataStoreRedis) CreateTodo(todo entity.Todo) {..}
func (d *dataStoreRedis) GetTodo(id string) {..}
func (d *dataStoreRedis) UpdateTodo(id string, todo entity.Todo) {..}
func (d *dataStoreRedis) DeleteTodo(string) {..}

```

And if we also need a Postgres data store:

```go
package sqlwrapper

import (
  "example.com/servicehex/domain/entity"
  "example.com/servicehex/domain/datagateway"

  "github.com/jinxi/gorm"
)

// *Adapter*
type dataStoreSql struct {
  db *gorm.DB
}

func (d *dataStoreSql) CreateTodo(todo entity.Todo) {..}
func (d *dataStoreSql) GetTodo(id string) {..}
func (d *dataStoreSql) UpdateTodo(id string, todo entity.Todo) {..}
func (d *dataStoreSql) DeleteTodo(string) {..}
```

And now, let's say, the business logic of creating a new to-do
is in this pure function `createTodo`:

```go
package todo

import (
  "example.com/servicehex/domain/entity"
  "example.com/servicehex/domain/datagateway"
)

func CreateTodo(
  todo entity.Todo,
  dgTodo datagateway.DataGatewayTodo,
) error {
  if err := validateTodo(todo); err != nil {
    return errors.Wrap(err, "failed to validate todo")
  }

  if err := dgTodo.CreateTodo(todo); err != nil P
    return errors.Wrap(err, "failed to create todo")
  }

  return nil
```

Now, it doesn't matter to `createTodo` above if `dgTodo` was
a `dataStoreSql` or `dataStoreRedis`, all it knows is that
`dgTodo` can do `dgTodo.CreateTodo(entity.Todo)`.

If the data representation in Redis and SQL tables are different
from `entity.Todo`, it is the adapter's responsibilities to
translate the low-level data representation to the one used in core.

Since `createTodo` is tech-agnostic, it can be called from
gRPC or HTTP REST handlers, and the result would have been the same.

### Example: driving side

Let's say our core business ports are (1) create todo, (2)
get all user's unexpired todos, and (3) expire todo:

```go
package core

// code

/* Port */
type PortTodo interface {
  CreateTodo(entity.Todo) error
  GetLiveTodos(userId string) ([]entity.Todo, error)
  ExpireTodo(userId string, todoId string) error
}
```

And this port interface is implemented by `serviceTodo`:

```go
package servicetodo

// code

/* Adapter */
type serviceTodo struct {
  // serviceTodo.repo is just an interface
  repo datagateway.DataGatewayTodo
}

func (s *serviceTodo) CreateTodo(todo entity.Todo) error {
  return s.repo.CreateTodo(todo)
}
func (s *serviceTodo) GetLiveTodos(userId string) {..}
func (s *serviceTodo) ExpireTodo(userId, todoId string) {..}
```

If our code happens to be run as either a HTTP REST server or gRPC server,
then all we need to do is to make these server implementations call our `service`
wrapped behind `core.PortTodo`.


### All together


Now, if we combine both _driving_ and _driven_ side of our code,
we have a nice, clean code that connects to other pieces via interfaces.

```text
HTTP REST / gRPC

       |
       v

    PortTodo
    (Implemented by serviceTodo)
                        |
                        v

                  DataGatewayTodo (`via serviceTodo.repo`)
                  (Implemented by dataStoreRedis)
                        |
                        v

                  External Systems (Redis server)
```

And gRPC server does not know how `PortTodo` does its thing, all it knows is
to trust `PortTodo` implementation to do the right thing.

`PortTodo` and its implementations does not know if it's being called
from gRPC or REST callers. `PortTodo` doesn't even know if its implementations
persists data, and most probably does not even know about `DataGatewayTodo`.

`serviceTodo` does now know exactly what `serviceTodo.repo` does,
all it knows is that it can use it as a todo repository via `DataGatewayTodo`.

`DataGatewayTodo` does not know what databases, if at all, are used
to implement its methods. When this interface is actually implemented by
`dataStoreRedis`, only `dataStoreRedis` knows that this implementation
of `DataGatewayTodo` uses Redis.

Now we can mess up one part of our code without affecting the other.