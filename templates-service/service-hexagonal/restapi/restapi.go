package restapi

import "example.com/servicehex/domain/core"

// A HTTP REST `view/presentation` layer of our program
// Adapter `serviceTodo` connects to core port PortTodo,
// likewise, `serviceUser` connects to PortUser.
type restApi struct {
	serviceTodo core.PortTodo
	serviceUser core.PortUser
}
