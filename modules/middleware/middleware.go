package middlewareJWT

import (
	"context"
	"go-clean-arch/constants"
	"go-clean-arch/utils"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var tokenCtxKey = &contextKey{"token"}

type contextKey struct {
	name string
}

// Token struct to model token data in header
type Token struct {
	Value string
}

func TokenAuthorize(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenHeader := r.Header.Get(constants.Authorization)
		if tokenHeader == "" {
			err := constants.ErrTokenIsRequired
			log.Error(err)
			res := utils.Response{
				Code: http.StatusUnauthorized,
				Data: nil,
				Err:  utils.STATUS_UNAUTHORIZED,
				Msg:  err.Error(),
			}
			res.JSONResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), tokenCtxKey, &Token{Value: tokenHeader})
		r = r.WithContext(ctx)
		handlerFunc.ServeHTTP(w, r)
	})
}

// GetTokenFromContext return Token Object inside context data
func GetTokenFromContext(ctx context.Context) (*Token, error) {
	raw, err := ctx.Value(tokenCtxKey).(*Token)
	if err != true {
		return nil, constants.ErrTokenIsRequired
	}
	return raw, nil
}
