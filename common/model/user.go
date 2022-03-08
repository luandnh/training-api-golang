package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,type:varchar(50),notnull"`
	Email         string `bun:"email,type:varchar(100),nullzero"`
}
