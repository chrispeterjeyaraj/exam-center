package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/chrispeterjeyaraj/exam-center/backend/configs"
	models "github.com/chrispeterjeyaraj/exam-center/backend/pkg/models"
)

func HandleDatabaseInsert(DBname string, CollectionName string, email string, phone int, password string, fname string, lname string, uid string, created time.Time, updated time.Time, token string, code int, agent interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, errInsert := collection.InsertOne(ctx, bson.M{
		"email":      email,
		"phone":      phone,
		"password":   password,
		"first_name": fname,
		"last_name":  lname,
		"user_id":    uid,
		"created_at": created,
		"updated_at": updated,
		"UserAgent":  agent,
	})

	HandleInsertToken(DBname, "tokens", token, code, created)

	if errInsert != nil {
		return false
	}
	defer cancel()
	return true

}

func HandleInsertToken(DBName string, CollectionName string, token string, code int, created time.Time) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collectionToken := configs.GetCollection(configs.DB, CollectionName)

	if code != 0 {
		data := bson.M{"token": token, "code": code, "created_at": created}

		_, errInsert := collectionToken.InsertOne(ctx, data)

		if errInsert != nil {
			return false
		}
	} else if code == 0 {
		data := bson.M{"token": token, "created_at": created}

		_, errInsert := collectionToken.InsertOne(ctx, data)

		if errInsert != nil {
			return false
		}
	}

	return true

}

func HandleAuthentication(email string, password string, DBname string, CollectionName string) (bool, string, string, string, string) {

	var user models.AuthenticationModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	errFind := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	decryptPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if decryptPassword != nil {
		return false, "", "", "", ""
	}

	if errFind != nil {
		return false, "", "", "", ""

	}
	defer cancel()
	return true, user.Email, user.First_name, user.Last_name, user.User_id
}

func HandleTokenAuthentication(DBname string, CollectionName string, token string, code int) bool {

	var result models.ResponseModel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	if code != 0 {
		errFind := collection.FindOne(ctx, bson.M{"token": token, "code": code}).Decode(&result)

		if errFind != nil {
			return false
		}
	} else if code == 0 {
		errFind := collection.FindOne(ctx, bson.M{"token": token}).Decode(&result)

		if errFind != nil {
			return false
		}
	}

	defer cancel()
	return true
}

func HandleForgotPass(email string, DBName string, CollectionName string) bool {
	var result models.ResponseModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	errFind := collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)

	if errFind != nil {
		return false
	}

	defer cancel()
	return true
}

func HandleRemoveCode(DBName string, CollectionName string, code int, token string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, err := collection.DeleteOne(ctx, bson.M{"code": code})
	if err != nil {
		return false
	}

	defer cancel()
	return true
}

func HandleUpdatePassword(DBName string, CollectionName string, email string, password string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	collection := configs.GetCollection(configs.DB, CollectionName)

	_, err := collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"password": password}})

	if err != nil {
		return false
	}

	defer cancel()
	return true
}
