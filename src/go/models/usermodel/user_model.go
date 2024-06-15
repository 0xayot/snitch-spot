package UserModel

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"snitch-spot/src/go/utils"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Email           string             `bson:"email" validate:"required"`
	Organisation    string             `bson:"organisation" validate:"required"`
	CognitoID       string             `bson:"cognitoId" validate:"required"`
	APIEnabled      bool               `bson:"apiEnabled" default:"false"`
	Name            string             `bson:"name" validate:"required"`
	EncryptedAPIKey string             `bson:"encryptedApiKey,omitempty" json:"-"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}

func GetUserByApiKey(key string) (*User, error) {
	dbName := os.Getenv("DB_NAME")
	usersCollection := utils.GetCollection(dbName, "users")
	inputBytes := []byte(key)

	hash := sha512.New()
	hash.Write(inputBytes)
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	filter := bson.M{"encryptedApiKey": hashString}

	var user User
	err := usersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		fmt.Println("No user found with encryptedApiKey:")
		return nil, err
	} else if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		return &user, nil
	}
}

func GetUserFromRequest(request events.APIGatewayV2HTTPRequest) (*User, error) {
	user, err := GetUserByApiKey(request.Headers["x-api-key"])
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No user found with the provided API key")
			return nil, fmt.Errorf("unauthorized: invalid API key")
		} else {
			log.Println("Error retrieving user:", err)
			return nil, fmt.Errorf("internal server error")
		}
	}
	return user, nil
}
