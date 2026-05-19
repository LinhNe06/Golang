package services

import (
	"context"
	config "mongo-server/Config"
	models "mongo-server/Models"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Users []models.User

func GetCollection() *mongo.Collection {
	return config.DB.Collection(os.Getenv("USER_COLLECTION"))
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	cursor, err := GetCollection().Find(context.TODO(), bson.M{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &users); err != nil {
		return users, err
	}

	return users, nil
}

func Create(newUser models.User) (models.User, error) {
	newUser.ID = bson.NewObjectID()
	if _, err := GetCollection().InsertOne(context.TODO(), newUser); err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func GetUserById(id string) (models.User, error) {
	var user models.User

	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return user, err
	}

	err = GetCollection().FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)

	return user, err
}

func UpdateUserById(id string, updatedUser models.User) (models.User, error) {
	var user models.User

	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return user, err
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = GetCollection().FindOneAndUpdate(
		context.TODO(),
		bson.M{"_id": objectId},
		bson.M{"$set": updatedUser},
		opts,
	).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUserById(id string) error {
	objectId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	err = GetCollection().FindOneAndDelete(context.TODO(), bson.M{"_id": objectId}).Err()
	return err
}

func UserExistByID(id string) bool {
	user, err := GetUserById(id)

	if err != nil || user.ID.IsZero() {
		return false
	}

	return true
}
