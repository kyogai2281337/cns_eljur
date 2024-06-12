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

// Login handles the login request by parsing the request body, finding the user by email,
// comparing the password, generating a JWT token, setting the "auth" cookie, and returning
// a JSON response with the status code fiber.StatusOK and the token.
//
// Parameters:
//   - req: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the user, comparing the password,
//     generating the JWT token, setting the cookie, or creating the JSON response.
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
	response := &structures.UserLoginResponse{
		Token:  token,
		Status: fiber.StatusOK,
	}
	return req.JSON(response)
}

// Register handles the registration request by parsing the request body, finding the user role,
// creating a new user with the provided information, and returning a JSON response with the user data.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the user role, creating the user,
//     or returning the JSON response. Otherwise, nil.
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

	response := &structures.UserRegResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  u.Role.Name,
	}
	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// Delete handles the deletion of a user by retrieving the user data from the JWT token in the "auth" cookie,
// deleting the user from the database, and returning a JSON response with the status code fiber.StatusOK.
//
// Parameters:
//   - req: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error object if there was an issue decoding the JWT token, finding the user,
//     or deleting the user from the database. Otherwise, nil.
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

// User handles the user request by retrieving the "auth" cookie, decoding the JWT token,
// finding the user in the database, and returning the user data in a JSON response.
//
// Parameters:
//   - req: a pointer to a fiber.Ctx object representing the HTTP request.
//
// Returns:
//   - error: an error object if there was an issue decoding the JWT token, finding the user,
//     or returning the JSON response. Otherwise, nil.
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

	response := structures.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role.Name,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return req.JSON(response)
}

// Logout handles the logout request by clearing the "auth" cookie and returning a JSON response with the status code fiber.StatusOK.
//
// Parameters:
//   - req: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was a problem setting the cookie or creating the JSON response.
func (c *AuthController) Logout(req *fiber.Ctx) error {
	cookie := &fiber.Cookie{
		Name:     "auth",
		Value:    "",
		HTTPOnly: true,
	}
	req.Cookie(cookie)
	return req.JSON(fiber.Map{"status": fiber.StatusOK})
}

// Authentication is a middleware function that checks if the user is authenticated by verifying the JWT token in the "auth" cookie.
//
// It takes a fiber.Ctx as input parameter and returns an error.
// The function retrieves the "auth" cookie from the request and decodes the JWT token using the GetUserDataJWT function.
// If the token is invalid or expired, it returns a fiber.NewError with the status code fiber.StatusUnauthorized and the error message.
// Otherwise, it calls req.Next() to pass the request to the next handler in the middleware chain.
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

// Log is a middleware function that logs the request details for every request.
//
// It takes a fiber.Ctx as input parameter and returns an error.
// The function captures the X-Request-ID header and logs the request details,
// including the status code, request ID, IP address, path, protocol, and duration.
func (c *AuthController) Log() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		err := ctx.Next()

		duration := time.Since(start)

		// Ensure X-Request-ID header is captured
		requestID := ctx.Get("X-Request-ID")
		if requestID == "" {
			requestID = "unknown"
		}

		log.Printf("[%d] => %s %s %s time>%v\n",
			ctx.Response().StatusCode(),
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
func (c *AuthController) RequestID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := uuid.New().String()
		ctx.Set("X-Request-ID", id)
		ctx.Locals("requestid", id)
		return ctx.Next()
	}
}
