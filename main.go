package main

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)


type person struct{
	age int16
	name, lastname string
	school string
}

func main() {
	ctx := context.Background()
	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
	dbUri := "<URI for Neo4j database>"
	dbUser := "<Username>"
	dbPassword := "<Password>"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
}

type start struct{
	ctx context.Context
}

func (start int) Get(ctx context.Context) error{
	if start, err != nil{
		return err
	}
}