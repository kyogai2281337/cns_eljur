package service

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kyogai2281337/cns_eljur/internal/auth/structures"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/kyogai2281337/cns_eljur/pkg/sql/model"
)

type AuthController struct {
	Server *server.Server
}

func NewAuthController(s *server.Server) *AuthController {
	return &AuthController{
		Server: s,
	}
}

func (c *AuthController) Login(req *fiber.Ctx) error {
	usr := &structures.UserLoginRequest{}
	if err := req.BodyParser(usr); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	user, err := c.Server.Store.User().FindByEmail(usr.Email)
	if err != nil || !user.ComparePass(usr.Password) {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	token, err := toUserJWT(user).GenJWT("r")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	cookie := &fiber.Cookie{
		Name:     "auth",
		Value:    token,
		HTTPOnly: true,
	}
	req.Cookie(cookie)
	return req.JSON(fiber.Map{"status": fiber.StatusOK, "token": token})
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	req := new(structures.UserRegRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	uRole, err := c.Server.Store.Role().FindRoleByName("user")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	u := &model.User{
		Email:     req.Email,
		Pass:      req.Password,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Role:      uRole,
	}

	if err := c.Server.Store.User().Create(u); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	u.Sanitize()
	return ctx.Status(fiber.StatusCreated).JSON(u)
}

func (c *AuthController) Delete(req *fiber.Ctx) error {
	userData, err := GetUserDataJWT(req.Cookies("auth"))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if err := c.Server.Store.User().Delete(userData.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return req.JSON(fiber.Map{"status": fiber.StatusOK})
}

func (c *AuthController) User(req *fiber.Ctx) error {
	authCookie := req.Cookies("auth")
	if authCookie == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "missing auth cookie")
	}

	userData, err := GetUserDataJWT(authCookie)
	if err != nil {
		log.Println("Error decoding JWT:", err)
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	user, err := c.Server.Store.User().Find(userData.ID)
	if err != nil {
		log.Println("Error finding user:", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	resp := structures.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role.Name,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	log.Println("Successfully fetched user:", resp)
	return req.JSON(resp)
}

func (c *AuthController) Logout(req *fiber.Ctx) error {
	cookie := &fiber.Cookie{
		Name:     "auth",
		Value:    "",
		HTTPOnly: true,
	}
	req.Cookie(cookie)
	return req.JSON(fiber.Map{"status": fiber.StatusOK})
}

func (c *AuthController) Authentication() fiber.Handler {
	return func(req *fiber.Ctx) error {
		cookie := req.Cookies("auth")
		_, err := GetUserDataJWT(cookie)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		return req.Next()
	}
}

func (c *AuthController) Log() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		err := ctx.Next()
		duration := time.Since(start)

		log.Printf("%s %s %s %s completed in %v with response: %s\n",
			ctx.IP(), ctx.Method(), ctx.Path(), string(ctx.Body()), duration, string(ctx.Response().Body()))

		return err
	}
}

func (c *AuthController) RequestID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := uuid.New().String()
		ctx.Set("X-Request-ID", id)
		ctx.Locals("requestid", id)
		return ctx.Next()
	}
}
