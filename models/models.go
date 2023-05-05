package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Username     *string            `json:"user_name" validate:"required,min=2,max=30"`
	Email        *string            `json:"email" validate:"email,required"`
	Password     *string            `json:"password" validate:"required,min=6"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	Vaults       []Vault            `json:"vaults" bson:"vaults"`
	Friends      []User             `json:"friends" bson:"friends"`
}

type Vault struct {
	VaultID       primitive.ObjectID `bson:"_id"`
	VaultName     *string            `json:"vault_name"`
	Description   *string            `json:"description"`
	CreatedAt     time.Time          `json:"created_at"`
	Period        *int               `json:"period_days"`
	StatusOverall bool               `json:"status_overall"`
	FocusMode     bool               `json:"focus_mode"`
	EachDay       []Day              `json:"each_day" bson:"each_day"`
}

type Day struct {
	DayID          primitive.ObjectID `bson:"_id"`
	ToDos          []ToDo             `json:"to_dos" bson:"to_dos"`
	EverythingDone bool               `json:"everything_done"`
}

type ToDo struct {
	ToDoID   primitive.ObjectID `bson:"_id"`
	ToDoName string             `json:"to_do_name"`
	Flag     bool               `json:"flag"`
	Finished time.Time          `json:"finished"`
}
