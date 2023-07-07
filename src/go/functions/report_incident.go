package main

import (
	"encoding/json"
	// "fmt"
	"log"

	// FraudReportModel  "snitch-spot/src/go/models/fraudreportmodel"
	UserModel "snitch-spot/src/go/models/usermodel"
	"snitch-spot/src/go/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	err := utils.ConnectDB()
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return events.APIGatewayProxyResponse{}, err
	}
	defer utils.CloseDBConnection()

	user, err := UserModel.GetUserFromRequest(request)

	if err != nil {
		log.Printf("Error authenticating User %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}
	// body :=
	// save the data into a struct
	log.Println("the found", user)

	log.Println("the found", user.CognitoID)

	var body interface{} // Add the appropriate type

	err = json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		log.Printf("Error unmarshaling JSON request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	log.Println()

	// Create an interface for the data object with arbitrary key-value pairs

	// var report = interface{
	// 	ID:                    primitive.NewObjectID(),
	// 	EncryptedEmail:        utils.hashString ,
	// 	EncryptedName:         "John Doe",
	// 	EncryptedGovtID:       "encrypted_govt_id",
	// 	EncryptedRecipientEmail: "encrypted_recipient@example.com",
	// 	BankAccountName:       "John Doe",
	// 	BankAccountNumber:     "1234567890",
	// 	BankName:              "Bank XYZ",
	// 	DeviceID:              "device_id",
	// 	Email:                 "john.doe@example.com",
	// 	Name:                  "John Doe",
	// 	Offence:               "Fraud",
	// 	MetaData:              "additional metadata",
	// 	Resolved:              "refunded",
	// 	ShowVictim:            false,
	// 	Amount:                1000.50,
	// 	VictimizedOrganization: primitive.NewObjectID(),
	// 	IncidentDate:          time.Now(),
	// 	CreatedAt:             time.Now(),
	// }

	utils.HashString(body.email)

	// Create a response object
	// response := struct {
	// 	Message string                 `json:"message"`
	// 	Data    map[string]interface{} `json:"data,omitempty"`
	// }{
	// 	Message: user,
	// }

	// Convert the response object to JSON
	jsonResponse, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshaling JSON response: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// Return the JSON response
	return events.APIGatewayProxyResponse{
		Body:       string(jsonResponse),
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	defer utils.CloseDBConnection()
	lambda.Start(Handler)
}
