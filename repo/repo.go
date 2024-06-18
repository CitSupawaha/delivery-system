package repo

import (
	db "delivery-system/database"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetManyStatement(resource *db.Resource, collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	ctx, cancel := db.InitContext()
	option := options.Find()
	option.SetSort(filterOption)
	defer cancel()
	fmt.Println("obj ===> ", resource)
	obj, err := resource.DB.Collection(collection).Find(ctx, filter, option)
	fmt.Println("obj ===> ", obj)
	if err != nil {
		return err
	}
	err = obj.All(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
