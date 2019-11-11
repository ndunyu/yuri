package yuri

import (
	"log"
	"reflect"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

////it takes in  a request object reads its data
func InsertItem(item interface{}, DB *pg.DB) *ErrResponse {
	////it returns error if any
	//occurred
	if reflect.ValueOf(item).Kind() != reflect.Ptr {
		log.Println("not a pointer")

		return ErrInvalidRequest

	}
	err := DB.Insert(item)
	if err != nil {
		log.Println(err)
		return ErrInternalServerError

	}

	return nil

}

//update items using id
func UpdateItem(item interface{},q *orm.Query, DB *pg.DB) *ErrResponse {
	//test presence of object
	if reflect.ValueOf(item).Kind() != reflect.Ptr {
		log.Println("not a pointer")

		return ErrInvalidRequest

	}
	results, err := q.Returning("*").UpdateNotZero()
	if err != nil && err == pg.ErrNoRows {
		log.Println(err)
		///raven.CaptureError(err, nil)
		return ErrNotFound

	} else if err != nil {
		log.Println(err)
		return ErrInternalServerError
	}
	if results.RowsAffected() < 1 {
		log.Println(err)
		return ErrNotFound
	}

	return nil

}

///this will delete an item using
///an id
func DeleteItem(item interface{}, id int, DB *pg.DB) *ErrResponse {
	if reflect.ValueOf(item).Kind() != reflect.Ptr {
		log.Println("not a pointer")

		return ErrInvalidRequest

	}

	results, err := DB.Model(item).Where("id=?", id).Returning("*").Delete()
	if err != nil && err == pg.ErrNoRows {
		log.Println(err)
		///raven.CaptureError(err, nil)
		return ErrNotFound

	} else if err != nil {
		log.Println(err)
		return ErrInternalServerError
	}
	if results.RowsAffected() < 1 {
		log.Println(err)
		return ErrNotFound
	}

	return nil
}

///func get many items
func GetItemsHandler(item interface{}, q *orm.Query, p *Pagination) (*ResponseData, *ErrResponse) {
	if p != nil {
		q.Limit(p.Size)
		q.Offset(p.Offset)
	}

	count, err := q.SelectAndCount()
	if err != nil {
		log.Println(err)
		return nil, ErrInternalServerError

	}

	response := ResponseData{Items: item, TotalItems: count}

	return &response, nil
}

///get a single item and its relations if any
func GetItemHandler(q *orm.Query) *ErrResponse {

	err := q.First()

	if err != nil && err == pg.ErrNoRows {
		log.Println(err)
		return ErrNotFound

	} else if err != nil {
		log.Println(err)
		return ErrInternalServerError

	}

	//pages:= int(math.Ceil(float64(count/p.Size)))

	return nil
}
