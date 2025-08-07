package context

import "github.com/saryginrodion/stackable"

type Context = stackable.Context[SharedState, LocalState]
