package service

import (
	"github.com/gofiber/fiber/v2"
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

	uRole, err := c.Server.Store.Role().FindByName("user")
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
	user := req.Locals("user").(*model.User)
	if err := c.Server.Store.User().Delete(user.ID); err != nil {
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
	user := req.Locals("user").(*model.User)

	response := &structures.UserResponse{
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
