package queries

import (
	"context"
	"fmt"
	"github.com/AnirudhAgnihotri2902/jwt-using-go-and-postgres/ent"
	"github.com/AnirudhAgnihotri2902/jwt-using-go-and-postgres/ent/registers"
	"log"
)

// creating user..
func CreateUser(ctx context.Context, client *ent.Client, username string, userpassword string) (*ent.Registers, error) {
	u, err := client.Registers.
		Create().
		SetUsername(username).
		SetPassword(userpassword).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

// query checking..
func QueryUser(ctx context.Context, client *ent.Client, username string) (*ent.Registers, error) {
	u, err := client.Registers.
		Query().
		Where(registers.Username(username)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
