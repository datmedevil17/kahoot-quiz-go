package api

import (
	"github.com/datmedevil17/kahoot-quiz-go/internal/config"
	"github.com/datmedevil17/kahoot-quiz-go/internal/ws"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, hub *ws.Hub) *gin.Engine {
	r := gin.Default()

	// ... other routes ...

	// Protected routes
	protected := r.Group("/api/v1")
	// Middleware for auth would go here, e.g. protected.Use(AuthMiddleware(cfg.JWTSecret))

	// WebSocket endpoint
	// Note: WS usually needs to be accessible, but auth is handled inside HandleWS via query param
	// So it might not need the middleware if the middleware checks headers.
	// The user snippet said: protected.GET("/ws", ws.HandleWS(hub, jwtSecret))
	// If 'protected' has middleware that checks Authorization header, WS connecting from browser
	// (standard JS WebSocket) cannot set headers easily.
	// Usually WS uses query param 'token'.
	// So maybe 'protected' here just means "routes that require auth logic" but HandleWS does it manually?
	// Or maybe the user has a middleware that checks query params too.
	// For now, I will assume HandleWS does the check (as I implemented) and mount it.

	protected.GET("/ws", ws.HandleWS(hub, cfg.JWTSecret))

	return r
}
