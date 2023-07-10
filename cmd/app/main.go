package main

import (
	"context"
	"log"
	"mongodbrebe/internal/config"
	mdb "mongodbrebe/internal/repository/mongodb"
	uc "mongodbrebe/internal/usecase"
	"mongodbrebe/internal/utils/fileutil"
)

func main() {
	var ctx context.Context = context.TODO()

	configApp, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	mongoClient, err := mdb.NewMongoDBCLient(ctx, configApp)
	if err != nil {
		panic(err)
	}

	defer func() {

		if err := mdb.DisconnectMongoDBClient(mongoClient, ctx); err != nil {
			log.Fatal(err)
		}
	}()

	userRepository := mdb.NewRepository(mongoClient, configApp)
	userUsecase := uc.NewUserUsecase(userRepository)

	if err := userUsecase.InsertUsers(ctx, 15); err != nil {
		panic(err)
	}

	users, err := userUsecase.AggregateUsers(ctx)
	if err != nil {
		log.Println(err)
	}
	err = fileutil.WriteToJSONFile(users)
	if err != nil {
		panic(err)
	}
}
