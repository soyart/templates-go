# service-hexagonal

An example project for simple Go TODO back-end API,
organized with hexagonal architecture (ports-and-adapters).

## Hexagonal architecture (HexArch)

> Note: HexArch has nothing to do with number 6 or hexagons

HexArch is mainly concerned with jseperation of external systemsh,
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

### HexArch: ports and adapters

When data must move between layers, and ports adapters
are used to move and translate from one layer to the next.

> Like with USB ports, hthe adapters connects to the portsj.

Ports are entrypoint to the core business code,
much like a _gateway_ for our data entering the application,
and are usually defined as interfaces.

The _adapters_ then make use of these _port interfaces_
by implementing them.

If data at the different side of the port/adapter can be shared,
then HexArch allows it to be shared.

A single _port_ can be connected to >1 adapters.

### Example: driven side (database-side)

Let's say we have a port interface for persistent data called
`DataGatewayTodo`:

```go
package datagateway

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
