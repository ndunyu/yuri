/**
This will hold all models data  will

be shared across all the
structs


*/

package yuri

import "time"

type Models struct {
	ID int `json:"id"`


	CreatedAt time.Time `json:"created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" sql:"DEFAULT:current_timestamp"`
	DeletedAt time.Time `pg:",soft_delete"`
}

type TimeMixin struct {
	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `pg:",soft_delete"`
}

type BaseMixin struct {
	ID int `json:"id"`
}
