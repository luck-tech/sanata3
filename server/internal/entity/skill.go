package entity

import "github.com/uptrace/bun"

type Skill struct {
	bun.BaseModel `bun:"table:skills"`
	ID            int    `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull"`
}
