package controller

import (
	"context"
	dbproxy "erpoo/db/proxy"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"
)

type Controller struct {
	ctx     context.Context
	queries *dbproxy.Queries
}

func New(ctx context.Context, queries *dbproxy.Queries) *Controller {
	return &Controller{ctx: ctx, queries: queries}
}

type TeamDto struct {
	Id          pgtype.UUID
	Name        string
	Description string
}

func (cl *Controller) GetListTeams(c *fiber.Ctx) error {
	teams, err := cl.queries.GetListTeams(cl.ctx)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	res := lo.Map(teams, func(x dbproxy.Team, index int) TeamDto {
		return TeamDto{
			Id:          x.ID,
			Name:        x.Name.String,
			Description: x.Description.String,
		}
	})

	return c.JSON(res)
}

type CreateTeamInputDto struct {
	Name        pgtype.Text `json:"name"`
	Description pgtype.Text `json:"description"`
}

func (cl *Controller) CreateTeam(c *fiber.Ctx) error {
	createTeamInput := new(CreateTeamInputDto)

	createTeam := new(dbproxy.CreateTeamParams)

	if err := c.BodyParser(createTeamInput); err != nil {
		return err
	}

	createTeam.ID = pgtype.UUID{Bytes: uuid.New(), Valid: true}
	createTeam.Name = createTeamInput.Name
	createTeam.Description = createTeamInput.Description

	team, err := cl.queries.CreateTeam(cl.ctx, *createTeam)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(team)
}
