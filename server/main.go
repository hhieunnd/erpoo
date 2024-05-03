package main

import (
	"context"
	"erpoo/controller"
	dbproxy "erpoo/db/proxy"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://root:It123456@@localhost:5432/erpoo?sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(ctx)

	queries := dbproxy.New(conn)

	cl := controller.New(ctx, queries)

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "ERP",
		AppName:       "ERP v1.0.1",
	})

	app.Get("/api/teams", cl.GetListTeams)
	app.Post("/api/teams", cl.CreateTeam)

	app.Listen(":3000")
}
