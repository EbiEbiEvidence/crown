package query

import (
	"database/sql"
)

type CardQueryModel struct {
	CardID     int64
	IsChild    bool
	ScenarioID sql.NullInt64
	CardText   string
	CardImage  string
}
