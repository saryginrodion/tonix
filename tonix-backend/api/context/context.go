package context

import "github.com/radyshenkya/stackable"

type Context = stackable.Context[SharedState, LocalState]
