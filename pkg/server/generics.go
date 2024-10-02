package server

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/store"
	"github.com/nats-io/nats.go"
)

type Server struct {
	App    *fiber.App
	Store  store.Store
	Broker *nats.Conn
}

func NewServer(store store.Store, natsStr string) *Server {
	nc, err := nats.Connect(natsStr)
	if err != nil {
		log.Fatal(err)
	}

	s := &Server{
		App:    fiber.New(),
		Store:  store,
		Broker: nc,
	}

	s.App.Use("/private", s.Authentication())
	s.App.Use(s.RequestID())
	s.App.Use(s.Log())
	return s
}

func (s *Server) ServeHTTP(addr string) error {
	return s.App.Listen(addr)
}

func (s *Server) Authentication() fiber.Handler {
	return func(req *fiber.Ctx) error {
		cookie := req.Cookies("auth")
		jwt, err := GetUserDataJWT(cookie)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		user, err := s.Store.User().Find(jwt.ID)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		req.Locals("user", user)
		return req.Next()
	}
}

// Log is a middleware function that logs the request details for every request.
//
// It takes a fiber.Ctx as input parameter and returns an error.
// The function captures the X-Request-ID header and logs the request details,
// including the status code, request ID, IP address, path, protocol, and duration.
func (c *Server) Log() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		// Передаем управление следующему обработчику
		err := ctx.Next()

		// Вычисляем длительность выполнения запроса
		duration := time.Since(start)

		// Получаем значение заголовка X-Request-ID
		requestID := ctx.Get("X-Request-ID")
		if requestID == "" {
			requestID = "unknown"
		}

		// Получаем информацию о пользователе из локальных данных контекста
		user, ok := ctx.Locals("user").(*model.User)
		if !ok {
			// Если пользователь не найден, создаем заглушку
			user = &model.User{
				Email: "unknown",
				Role: &model.Role{
					Name: "unknown",
				},
			}
		}

		// Логируем информацию о запросе
		log.Printf("[%d] => %s %s %s %s %s time>%v\n",
			ctx.Response().StatusCode(),
			user.Email,
			user.Role.Name,
			ctx.IP(),
			ctx.Path(),
			ctx.Method(),
			duration)

		return err
	}
}

// RequestID generates a unique request ID and sets it in the context.
//
// It takes a fiber.Ctx as input parameter and returns an error.
func (c *Server) RequestID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := uuid.New().String()
		ctx.Set("X-Request-ID", id)
		ctx.Locals("requestid", id)
		return ctx.Next()
	}
}
