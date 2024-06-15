package main

import (
	"encoding/json"

	"log"

	FraudReportModel "snitch-spot/src/go/models/fraudreportmodel"
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

	_, err = UserModel.GetUserFromRequest(request)

	if err != nil {
		log.Printf("Error authenticating User %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	var body map[string]interface{}

	err = json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		log.Printf("Error unmarshaling JSON request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	hashedEmail := utils.HashString(utils.GetString(body, "email"))

	hashedGovtId := utils.HashString(utils.GetString(body, "govtId"))

	hashedAccNumber := utils.HashString(utils.GetString(body, "receivingBankAccountNumber"))

	isSuspiciousReport := false

	reportByEmail, err := FraudReportModel.FindRecord("encryptedEmail", hashedEmail)
	if err != nil {
		log.Printf("Error finding record by hashed email: %v\n", err)
		reportByEmail = nil
	} else if reportByEmail == nil {
		log.Println("No record found for hashed email")
		reportByEmail = nil
	}

	// Check if reportByEmail is not nil and set isSuspiciousReport to true
	if reportByEmail != nil {
		isSuspiciousReport = true
	}

	reportByGovtId, err := FraudReportModel.FindRecord("encryptedGovtId", hashedGovtId)
	if err != nil {
		log.Printf("Error finding record by hashed govt ID: %v\n", err)
		reportByGovtId = nil
	} else if reportByGovtId == nil {
		log.Println("No record found for hashed govt ID")
		reportByGovtId = nil
	}

	// Check if reportByGovtId is not nil and set isSuspiciousReport to true
	if reportByGovtId != nil {
		isSuspiciousReport = true
	}

	reportByAccNumber, err := FraudReportModel.FindRecord("bankAccountNumber", hashedAccNumber)
	if err != nil {
		log.Printf("Error finding record by hashed account number: %v\n", err)
		reportByAccNumber = nil
	} else if reportByAccNumber == nil {
		log.Println("No record found for hashed account number")
		reportByAccNumber = nil
	}

	// Check if reportByAccNumber is not nil and set isSuspiciousReport to true
	if reportByAccNumber != nil {
		isSuspiciousReport = true
	}

	var message string
	if isSuspiciousReport {
		message = "This user has been reported for fraud"
	} else {
		message = "This user has no reports with us"
	}

	// Create the response struct
	response := struct {
		Message            string                 `json:"message"`
		IsSuspiciousReport bool                   `json:"isSuspiciousReport"`
		Data               map[string]interface{} `json:"data,omitempty"`
	}{
		Message:            message,
		IsSuspiciousReport: isSuspiciousReport,
	}

	// Return the response
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling response: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
