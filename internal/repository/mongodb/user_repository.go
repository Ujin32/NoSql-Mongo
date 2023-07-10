package mongodb

import (
	"context"
	"mongodbrebe/internal/config"
	"mongodbrebe/internal/models"
	"mongodbrebe/internal/utils/wordz"

	"github.com/bxcodec/faker/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertUsers(ctx context.Context, count int) error
	AggregateUsersToSex(ctx context.Context) ([]bson.D, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewRepository(client *mongo.Client, cfg *config.Config) UserRepository {
	coll := client.Database(cfg.DBName).Collection(cfg.CollectionName)
	return &userRepository{
		collection: coll,
	}

}

func (u *userRepository) InsertUsers(ctx context.Context, count int) error {
	users := make([]interface{}, 0, count)
	for i := 1; i <= count; i++ {
		dateBirth, _ := faker.RandomInt(18, 70)
		user := &models.User{
			Name:        wordz.RandomWord(wordz.FakeNames),
			Sex:         wordz.RandomWord([]string{"male", "female"}),
			DateOfBirth: dateBirth[0],
			Smartphone: &models.Smartphone{
				Model: wordz.RandomWord([]string{
					"xiomi", "samsung", "iphone", "asus", "sony", "huawei", "tcl",
				}),
				Vendor: wordz.RandomWord([]string{
					"svyznoi", "ozon", "WB", "aliexpress", "SMM", "dns", "eldorado",
				}),
			},
		}

		users = append(users, user)
	}
	if _, err := u.collection.InsertMany(ctx, users); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) AggregateUsersToSex(ctx context.Context) ([]bson.D, error) {
	groupStage1 := bson.D{{"$group", bson.D{{"_id", bson.D{{"model", "$smartphone.model"}, {"sex", "$sex"}}}, {"count", bson.D{{"$sum", 1}}}}}}
	cursor, err := u.collection.Aggregate(ctx, mongo.Pipeline{groupStage1})
	if err != nil {
		return nil, err
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
