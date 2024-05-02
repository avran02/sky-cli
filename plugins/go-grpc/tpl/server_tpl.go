package tpl

var Server = `package server

import (
	"context"
	// proto "{{ .ProtoPath }}" you need to import your proto files here
)

type Server struct {
	proto.Unimplemented_Server // TODO: correct name
}
`
