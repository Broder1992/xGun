package main



import (

	"errors"

	"fmt"

	"strconv"

	"strings"

	//Import the shim for the ledger
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"encoding/json"

	"regexp"

)

//==============================================================================================================================
//
//		Represents the participants in the specific guns life cycle. Participants are the authorities and/or
// 		buissnies owners that directly interacts with the life cycle of the weapon.
//
//==============================================================================================================================


var logger = shim.NewLogger("CDLChaincode")

//The authority regulating gun control and gun requirements
const REGULATORAUTHORITY = "regulatorauthority"

//The manufacturer of the gun who has to oblige the regulations set by the regulatoratuhority
const GUNMANUFACTURER = "gunmanufacturer"

//The gun merchant distributes the guns manufactured by the gunmanufacturer
const GUNMERCHANT = "gunmerhcant"

//The private gun owner purhcases the gun from the gun merchant
const PRIVATEGUNOWNER = "privategunowner"

//The gun destructor destroys the gun, thus ending the life cycle of the weapon.
const GUNDESTRUCTOR = "gundestructor"

//==============================================================================================================================

//	 Represents the different states that the asset, here by called "The Gun" can be in.

//==============================================================================================================================

//The gun is in the "Template mode", meaning that the regulator has pointed out requirements for "a gun".
const   STATE_GUN_TEMPLATE  		=  0

//The gun is manufactured/being manufactured by a gun manufacturer.
const   STATE_MANUFACTURE  		=  1

//The gun is merchandise for sale at a gun merchant.
const   STATE_MERCHANDISE  		=  2

//The gun is currently owned by a private person.
const   STATE_PRIVATELY_OWNED		=  3

//The weapons life cycle has ended and is now being destructed by a licensed weapon destructor.
const   STATE_BEING_DESTRUCTED 		=  4

type  SimpleChaincode struct {

}

//==============================================================================================================================
//
// 				The gun struct defines the structure for a gun object.
//
//==============================================================================================================================


type Guns struct{

	Make		String `json:"make"`

	Model		String `json:"model"`

	LicensNm	String `json:"licensNm"`

	Type 		String `json:"type"`

	Owner		String `json:"owner"`

	Destructed	bool `json:"destructed"`

	Status		int `json:"status"`

	OwnerContractID	string `json"ownerContractID"`

	GunLogBookID    string `json:"gunLogBookID"`

}

//==============================================================================================================================
//	GunLogBookID holder - The structure that keeps a record of all "LogBookID:S" for guns that have been created
//==============================================================================================================================



type GunLogBookID_holder struct {

	GunLogBookIds 	[]string `json:"GunLogBookID"`
}

//==============================================================================================================================
//	User_and_eCert - Struct for storing the JSON of a user and their ecert - Generic
//==============================================================================================================================


type User_and_eCert struct {

	Identity string `json:"identity"`

	eCert string `json:"ecert"`

}

//==============================================================================================================================
//	The init function that is called on deplpoyment of the chain code.
//==============================================================================================================================

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var gunLogBookIds GunLogBookID_holder

	bytes, err := json.Marshal(gunLogBookIds)

	if err !nil { return nil, errors.New("Error creating GunLogBookId_holder record")}

	err = stub.PutState("gunLogBookIds", bytes)

	for i:=0 i < len(args); i=1+2{
		t.add_ecert(stub, args[i], args[i + 1])
	}

	return nil, nil;
}

//==============================================================================================================================

//	 General Functions

//==============================================================================================================================

//	 get_ecert - Takes the name passed and calls out to the REST API for HyperLedger to retrieve the ecert

//				 for that user. Returns the ecert as retrived including html encoding.

//==============================================================================================================================

func (t *SimpleChaincode) get_ecert(stub shim.ChaincodeStubInterface, name string) ([]byte, error) {

	ecert, err := stub.GetState(name)

	if err != nil { return nil, errors.New("Couldn't retrieve ecert for user " + name) }

	return ecert, nil

}


