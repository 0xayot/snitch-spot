package main

import (
	"encoding/json"
	"time"

	// "fmt"
	"log"

	FraudReportModel "snitch-spot/src/go/models/fraudreportmodel"
	UserModel "snitch-spot/src/go/models/usermodel"
	"snitch-spot/src/go/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	err := utils.ConnectDB()
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return events.APIGatewayProxyResponse{}, err
	}
	defer utils.CloseDBConnection()

	user, err := UserModel.GetUserFromRequest(request)

	// Validate and fail if some params are not supplied

	if err != nil {
		log.Printf("Error authenticating User %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	// log.Println("the found", user.CognitoID)
	var body map[string]interface{}

	err = json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		log.Printf("Error unmarshaling JSON request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	// Create an interface for the data object with arbitrary key-value pairs

	var report = FraudReportModel.FraudReport{
		ID:                         primitive.NewObjectID(),
		EncryptedOffenderEmail:     utils.HashString(utils.GetString(body, "email")),
		EncryptedOffenderName:      utils.HashString(utils.GetString(body, "name")),
		EncryptedOffenderGovtID:    utils.HashString(utils.GetString(body, "govtId")),
		EncryptedRecipientEmail:    utils.HashString(utils.GetString(body, "recipientEmail")),
		ReceivingBankAccountName:   utils.HashString(utils.GetString(body, "receivingBankAccountName")),
		ReceivingBankAccountNumber: utils.HashString(utils.GetString(body, "receivingBankAccountNumber")),
		ReceivingBankName:          utils.HashString(utils.GetString(body, "receivingBankName")),
		DeviceID:                   utils.HashString(utils.GetString(body, "deviceId")),
		Offence:                    utils.GetString(body, "offence"),
		MetaData:                   utils.GetString(body, "metabody"),
		Resolved:                   utils.GetString(body, "resolved"),
		ShowVictim:                 false,
		Amount:                     utils.GetFloat64(body, "amount"),
		VictimizedOrganization:     user.ID,
		IncidentDate:               time.Now(),
		CreatedAt:                  time.Now(),
	}

	err = FraudReportModel.SaveReport(report)

	if err != nil {
		log.Printf("Error saving report : %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	// Create a response object
	response := struct {
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data,omitempty"`
	}{
		Message: " Succcess Snitching",
	}

	// Convert the response object to JSON
	jsonResponse, err := json.Marshal(response)
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
