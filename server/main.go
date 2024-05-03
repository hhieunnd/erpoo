package main

import (
	"context"
	dbproxy "erpoo/db/proxy"
	"fmt"
	"log"

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

	newParams := dbproxy.CreateTeamParams{
		ID:          pgtype.UUID{Bytes: uuid.New(), Valid: true},
		Name:        pgtype.Text{String: "Team1"},
		Description: pgtype.Text{String: "Des1"}}

	fmt.Printf("new params %v", newParams)

	// create new team
	team, err := queries.CreateTeam(ctx, newParams)

	if err != nil {
		log.Fatalf("err => %v", err)
	}

	fmt.Printf("team %v", team)
}
