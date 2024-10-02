package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kyogai2281337/cns_eljur/internal/constructor/structures"
	"github.com/kyogai2281337/cns_eljur/internal/constructor/xlsx"
	constructor "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/logic"
	constructor_logic_entrypoint "github.com/kyogai2281337/cns_eljur/internal/constructor_logic/scd"
	"github.com/kyogai2281337/cns_eljur/internal/mongo/primitives"
	"github.com/kyogai2281337/cns_eljur/pkg/server"
	"github.com/nats-io/nats.go"
)

type ConstructorController struct {
	Server *server.Server
}

func NewConstructorController(server *server.Server) *ConstructorController {
	return &ConstructorController{
		Server: server,
	}
}

// Create handles the creation of a schedule by parsing the request body, constructing a new schedule from the request,
// and saving the schedule to the database. It returns a JSON response with the status of the operation.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, constructing the schedule, saving the schedule to the database,
//     or returning the JSON response. Otherwise, nil.
func (c *ConstructorController) Create(ctx *fiber.Ctx) error {
	request := &structures.CreateRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	groups, cabs, teachers, plans, err := c.makeLists(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	schedule := constructor.MakeSchedule(request.Name, request.Limits.Days, request.Limits.Pairs, groups, teachers, cabs, plans, request.Limits.MaxDays, request.Limits.MaxWeeks)
	schedule.Make()
	err = primitives.NewMongoConn().Schedule().Make(schedule)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "ok"})
}

// Get handles the get by id request by finding the schedule by ID and returning a JSON response
// with the schedule.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue finding the schedule or returning the JSON response.
//     Otherwise, nil.
func (c *ConstructorController) Get(ctx *fiber.Ctx) error {
	request := &structures.GetByIDRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	mongoSchedule, err := primitives.NewMongoConn().Schedule().Find(request.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok",
		"schedule": mongoSchedule})
}

// GetList handles the get list request by finding all schedules in the database and returning a JSON response
// with the list of schedules.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue finding the schedules or returning the JSON response.
//     Otherwise, nil.
func (c *ConstructorController) GetList(ctx *fiber.Ctx) error {
	response, err := primitives.NewMongoConn().Schedule().GetList()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "schedules": response})
}

// Update handles the updating of a schedule by parsing the request body, finding the schedule by ID, performing the
// requested operation (insert or delete), and returning a JSON response with the status of the operation.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the schedule, performing the operation,
//     or returning the JSON response. Otherwise, nil.
func (c *ConstructorController) Update(ctx *fiber.Ctx) error {
	request := &structures.UpdateRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	directive := constructor_logic_entrypoint.Directive{
		Type: constructor_logic_entrypoint.DirTX,
		ID:   uuid.New().String(),
		Data: request,
	}

	marshaledDirective, err := json.Marshal(directive)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.Server.Broker.Publish("constructor", marshaledDirective); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sub := new(nats.Subscription)
	sub, err = c.Server.Broker.Subscribe("constructor_resp", func(msg *nats.Msg) {
		var response constructor_logic_entrypoint.DirResp
		if err := json.Unmarshal(msg.Data, &response); err != nil {
			return
		}
		if response.Data.(string) == directive.ID {
			ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
		}
	})
	defer sub.Unsubscribe()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return nil
}

// Delete handles the deletion of a schedule by parsing the request body, finding the schedule by ID, deleting the schedule from the database,
// and returning a JSON response with the status of the operation.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the schedule, deleting the schedule from the database,
//     or returning the JSON response. Otherwise, nil.
func (c *ConstructorController) Delete(ctx *fiber.Ctx) error {
	request := &structures.GetByIDRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := primitives.NewMongoConn().Schedule().Delete(request.ID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}

// Rename handles the renaming of a schedule by parsing the request body, finding the schedule by ID, renaming the schedule,
// and returning a JSON response with the status of the operation.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the schedule, renaming the schedule,
//     or returning the JSON response. Otherwise, nil.
func (c *ConstructorController) Rename(ctx *fiber.Ctx) error {
	request := &structures.RenameRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := primitives.NewMongoConn().Schedule().Rename(request.ID, request.Name); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}

// SaveXLSX handles the saving of a schedule to an xlsx file request by parsing the request body, finding the schedule by ID, converting it to a full schedule,
// saving the xlsx file, and returning a JSON response with the status of the operation.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue parsing the request, finding the schedule, converting it to a full schedule, saving the xlsx file,
//     or returning the JSON response. Otherwise, nil.
func (c *ConstructorController) SaveXLSX(ctx *fiber.Ctx) error {
	request := &structures.SaveXLSXRequest{}
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	mongoSchedule, err := primitives.NewMongoConn().Schedule().Find(request.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	schedule, err := c.RecoverToFull(mongoSchedule)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err := xlsx.LoadFile(schedule, fmt.Sprintf("./uploads/%s.xlsx", request.ID)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}

// ExportByID downloads a schedule with the given id as an Excel file named after the id.
//
// Parameters:
//   - ctx: a pointer to a fiber.Ctx object representing the HTTP request context.
//
// Returns:
//   - error: an error if there was an issue downloading the file. Otherwise, nil.
func (c *ConstructorController) ExportByID(ctx *fiber.Ctx) error {
	filename := ctx.Params("id")
	return ctx.Download(fmt.Sprintf("./uploads/%s.xlsx", filename), filename+".xlsx")
}
