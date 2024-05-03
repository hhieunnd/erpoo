package main

import (
	"context"
	dbproxy "erpoo/db/proxy"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://root:It123456@@localhost:5432/erpoo?sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(ctx)

	queries := dbproxy.New(conn)

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "ERP",
		AppName:       "ERP v1.0.1",
	})

	app.Get("/api/teams", func(c *fiber.Ctx) error {
		teams, err := queries.GetListTeams(ctx)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(teams)
	})

	app.Post("/api/teams", func(c *fiber.Ctx) error {

		type Params struct {
			// ID          pgtype.UUID `json:"id"`
			Name        pgtype.Text `json:"name"`
			Description pgtype.Text `json:"description"`
		}

		params := new(Params)

		p := new(dbproxy.CreateTeamParams)

		if err := c.BodyParser(params); err != nil {
			return err
		}

		p.ID = pgtype.UUID{Bytes: uuid.New(), Valid: true}
		p.Name = params.Name
		p.Description = params.Description

		team, err := queries.CreateTeam(ctx, *p)

		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(team)
	})

	app.Listen(":3000")
}
