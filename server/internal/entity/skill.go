package entity

import "github.com/uptrace/bun"

type Skill struct {
	bun.BaseModel `bun:"table:skills"`
	ID            int    `bun:"id,pk,autoincrement" json:"id"`
	Name          string `bun:"name,unique" json:"name"`
}
