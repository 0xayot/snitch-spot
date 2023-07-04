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
	EncryptedAPIKey string             `bson:"encryptedApiKey" validate:"required"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}

func GetUserByApiKey(key string) (*User, error) {
	dbName := os.Getenv("DB_NAME")
	usersCollection := utils.GetCollection(dbName, "users")
	// Convert the input string to byte array
	inputBytes := []byte(key)

	// Create a new SHA256 hash object
	hash := sha512.New()
	// Write the input bytes to the hash object
	hash.Write(inputBytes)

	// Calculate the SHA256 hash
	hashBytes := hash.Sum(nil)

	// Convert the hash bytes to a hexadecimal string
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
	return GetUserByApiKey(request.Headers["x-api-key"])
}
