// Package dispatch
//
// @author: xwc1125
package dispatch

import (
	"github.com/chain5j/chain5j-pkg/types"
	"github.com/chain5j/chain5j-protocol/protocol"
	"sync"
)

var (
	_ protocol.AppContexts = new(appContexts)
)

type appContexts struct {
	appCtxs *sync.Map // types.TxType->AppContext
}

func (ctx *appContexts) Ctx(t types.TxType) protocol.AppContext {
	if appCtx, ok := ctx.appCtxs.Load(t); ok {
		return appCtx.(protocol.AppContext)
	}
	return nil
}
