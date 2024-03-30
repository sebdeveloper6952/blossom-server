package gin

import (
	"encoding/base64"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	goNostr "github.com/nbd-wtf/go-nostr"
)

func nostrAuthMiddleware(action string, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Debug("[nostrAuthMiddleware] missing Authorization header")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Nostr ") {
			log.Debug("[nostrAuthMiddleware] missing Nostr header prefix")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		eventBase64 := strings.TrimPrefix(authHeader, "Nostr ")

		eventBytes, err := base64.StdEncoding.DecodeString(eventBase64)
		if err != nil {
			log.Debug("[nostrAuthMiddleware] base64 decode event failed")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ev := &goNostr.Event{}
		if err := json.Unmarshal(eventBytes, ev); err != nil {
			log.Debug("[nostrAuthMiddleware] json decode failed")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if ok, err := ev.CheckSignature(); !ok || err != nil {
			log.Debug("[nostrAuthMiddleware] check event sig failed")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// ****************************** Blossom Auth logic from this point *******************************************

		// kind must be 24242
		if ev.Kind != 24242 {
			log.Debug("[nostrAuthMiddleware] invalid event kind")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// the created_at must be in the past
		if ev.CreatedAt.Time().Unix() > time.Now().Unix() {
			log.Debug("[nostrAuthMiddleware] invalid created_at")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expirationTagValue := ""
		tTagValue := ""
		sizeTagValue := ""
		xTagValue := ""

		for i := range ev.Tags {
			if ev.Tags[i][0] == "expiration" && len(ev.Tags[i]) == 2 {
				expirationTagValue = ev.Tags[i][1]
			} else if ev.Tags[i][0] == "t" && len(ev.Tags[i]) == 2 {
				tTagValue = ev.Tags[i][1]
			} else if ev.Tags[i][0] == "size" && len(ev.Tags[i]) == 2 {
				sizeTagValue = ev.Tags[i][1]
			} else if ev.Tags[i][0] == "x" && len(ev.Tags[i]) == 2 {
				xTagValue = ev.Tags[i][1]
			}
		}
		if expirationTagValue == "" || tTagValue == "" {
			log.Debug("[nostrAuthMiddleware] missing `expiration` or `t` tags")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// the expiration tag must be set to a Unix timestamp in the future
		n, err := strconv.Atoi(expirationTagValue)
		if time.Unix(int64(n), 0).Unix() < time.Now().Unix() {
			log.Debug("[nostrAuthMiddleware] invalid expiration")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// the t tag must have a verb matching the intended action of the endpoint
		if tTagValue != action {
			log.Debug("[nostrAuthMiddleware] invalid action")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// additional checks depending on action
		if action == "upload" {
			sizeBytes, err := strconv.Atoi(sizeTagValue)
			if err != nil {
				log.Debug("[nostrAuthMiddleware] upload requires `size` tag")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			if int64(sizeBytes) != c.Request.ContentLength {
				log.Debug("[nostrAuthMiddleware] upload size does not match")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else if action == "delete" {
			if xTagValue == "" {
				log.Debug("[nostrAuthMiddleware] delete requires `x` tag")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("x", xTagValue)
		}

		c.Set("pk", ev.PubKey)

		c.Next()
	}
}
