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

const REGULATORAUTHORITY = "regulatorauthority"
const GUNMANUFACTURER = "gunmanufacturer"
const GUNMERCHANT = "gunmerhcant"
const GUNDESTRUCTOR = "gundestructor"







var logger = shim.NewLogger("CLDChaincode")



const   AUTHORITY      =  "regulator"

const   MANUFACTURER   =  "manufacturer"

const   PRIVATE_ENTITY =  "private"

const   LEASE_COMPANY  =  "lease_company"

const   SCRAP_MERCHANT =  "scrap_merchant"


