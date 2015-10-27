package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetDataByQuery(collection *mgo.Collection, start, length int, fields string, query interface{}, val interface{}) {
	collection.Find(query).Limit(length).Skip(start).Sort(fields).All(val)
}

func Count(collection *mgo.Collection, query interface{}) int {
	cnt, err := collection.Find(query).Count()
	if err != nil {
		fmt.Println(err.Error())
	}

	return cnt
}

func Has(collection *mgo.Collection, query interface{}) bool {
	if Count(collection, query) > 0 {
		return true
	}

	return false
}

func Insert(collection *mgo.Collection, data interface{}) error {
	return collection.Insert(data)
}

func Update(collection *mgo.Collection, query, data interface{}) error {
	return collection.Update(query, data)
}

func GetOneByQuery(collection *mgo.Collection, query, val interface{}) {
	collection.Find(query).One(val)
}

func Tag(caption, slug string) error {
	var tag BlogTag

	tag.Id = bson.NewObjectId()
	tag.Caption = caption
	tag.Slug = slug

	err := Insert(DbTag, tag)

	return err
}
