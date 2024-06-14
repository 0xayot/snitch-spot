package FraudReportModel

import (
	"context"
	"fmt"
	"os"
	"snitch-spot/src/go/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FraudReport struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty"`
	Amount                     float64            `bson:"amount"`
	EncryptedOffenderEmail     string             `bson:"encryptedEmail" validate:"required"`
	EncryptedOffenderName      string             `bson:"encryptedName"`
	EncryptedOffenderGovtID    string             `bson:"encryptedGovtId"`
	EncryptedRecipientEmail    string             `bson:"encryptedRecipientEmail"`
	ReceivingBankAccountName   string             `bson:"bankAccountName"`
	ReceivingBankAccountNumber string             `bson:"bankAccountNumber"`
	ReceivingBankName          string             `bson:"bankName"`
	DeviceID                   string             `bson:"deviceId"`
	// Email                   string             `bson:"email"`
	// Name                    string             `bson:"name"`
	Offence                string             `bson:"offence"`
	MetaData               string             `bson:"metaData"`
	Resolved               string             `bson:"resolved" validate:"oneof=refunded loss"`
	EncryptData            bool               `bson:"encryptData"  default:"true" json:"-"`
	ShowVictim             bool               `bson:"showVictim" default:"false" json:"-"`
	VictimizedOrganization primitive.ObjectID `bson:"victimizedOrganization" validate:"required"`
	IncidentDate           time.Time          `bson:"incidentDate" validate:"required"`
	CreatedAt              time.Time          `bson:"createdAt"`
}

func SaveReport(incident FraudReport) error { // FIx line
	dbName := os.Getenv("DB_NAME")
	// Get the incident collection
	collection := utils.GetCollection(dbName, "fraudreports")

	// Insert the incident document
	_, err := collection.InsertOne(context.Background(), incident)
	if err != nil {
		return fmt.Errorf("error inserting fraud report document: %v", err)
	}
	return nil
}
