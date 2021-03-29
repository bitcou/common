package auth

import (
	"context"
	"log"

	"errors"
	"net/http"

	"github.com/bitcou/common/dbmodels/graph/model"
	"gorm.io/gorm"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// APIKeyInfo struct containing all elements required to validate the API requests and set permissions/write access.
type APIKeyInfo struct {
	APIKey    string
	ClientID  int64
	IsAdmin   bool
	IsPremium bool
	IsDev     bool
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := r.Header.Get("X-API-Key")
			// Allow unauthenticated users in, we enable this only to allow the graphql playground to work,
			// for a production release this may be not the case, blocking non auth users at this point prevents
			// access to playground.
			if c == "" {
				next.ServeHTTP(w, r)
				return
			}

			keyInfo, _ := getUserByAPIKey(c, db)
			// if err != nil {
			// 	http.Error(w, "Auth Not Set", http.StatusForbidden)
			// 	return
			// }

			/*key := APIKeyInfo{
				APIKey:    keyInfo.Key,
				ClientID:  int64(keyInfo.Client.ID),
				IsAdmin:   keyInfo.IsAdmin,
				IsPremium: keyInfo.Client.IsPremium,
				IsDev:     keyInfo.IsDev,
			}*/ //TODO add missing info to the model

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, keyInfo.Client)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.Client {
	raw, _ := ctx.Value(userCtxKey).(*model.Client)
	if raw == nil {
		log.Println("could not find value in context")
	}
	return raw
}

func getUserByAPIKey(apiKey string, db *gorm.DB) (model.APIKey, error) {
	var keyInfo model.APIKey
	db.Where("key = ?", apiKey).Preload("Client").First(&keyInfo)
	if keyInfo.ID == 0 {
		return keyInfo, errors.New("ApiKey not found")
	}
	return keyInfo, nil
}
