package tapcontext

import (
	"context"
	"net/http"
	"sync"
)

type customContextType string

var (
	mutex sync.RWMutex
	data  = make(map[*http.Request]map[interface{}]interface{})
	datat = make(map[*http.Request]int64)
)

const (
	//TAPCtx - defining a separate type to avoid colliding with basic type
	TAPCtx customContextType = "tapCtx"
)

type TContext struct {
	context.Context
	TapContext
}

// TapContext contains context of client
type TapContext struct {
	UserEmail      string              // Email of the user
	PermissionsMap map[string][]string // this map will help in flagging
	TapApiToken    string              // TapApiToken - used to authenticate the session/request
	Application    string              // application for dynamic application auth
	Locale         string              // Locale for language
	RequestID      string              // RequestID - used to track logs across a request-response cycle
}

// GetTapCtx returns the tap context from the context provided
func GetTapCtx(ctx context.Context) (TapContext, bool) {
	if ctx == nil {
		return TapContext{}, false
	}
	tapCtx, exists := ctx.Value(TAPCtx).(TapContext)
	return tapCtx, exists
}

// UpgradeCtx embeds native context and TapContext
func UpgradeCtx(ctx context.Context) TContext {
	var tContext TContext
	tapCtx, _ := GetTapCtx(ctx)

	tContext.Context = ctx
	tContext.TapContext = tapCtx
	return tContext
}

func NewTapContext() TContext {
	return TContext{
		Context:    context.Background(),
		TapContext: TapContext{},
	}
}
