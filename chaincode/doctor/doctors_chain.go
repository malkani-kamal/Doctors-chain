package main

import (
	"fmt"
	"encoding/json"
	//"strconv"

	//"github.com/chaincode/doctor"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//SmartContract ... The SmartContract
type SmartContract struct {
}

//type DashboardCounts struct {
//	Doctors    int `json:"doctors"`
	//User     int `json:"user"`
	//Category int `json:"category"`
	//Request  int `json:"request"`
//}

type Doctors struct {
	DoctorName     string `json:"doctor_name"`
	DoctorId   	   string `json:"doctor_id"`
	Email   	   string `json:"doctor_email"`
	ContactNum     string `json:"contact_num"`
	Address        string `json:"address"`
}

type DoctorByIdResponse struct {
	ID      string `json:"id"`
	Request Doctors `json:"doctor"`
}


type Response struct {
	Status  string              `json:"status"`
	Message string            `json:"message"`
	Data    DoctorByIdResponse `json:"data"`
}



//var logger = shim.NewLogger("example_cc0")

//Init Function
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	
	_, args := stub.GetFunctionAndParameters()
	var doctors = Doctors{
		DoctorName:     args[0],
		DoctorId: 		args[1],
		Email:			args[2],
		ContactNum:   	args[3],
		Address:    	args[4]}

	doctorAsBytes, _ := json.Marshal(doctors)
	
	var uniqueID = args[1]	

	err := stub.PutState(uniqueID, doctorAsBytes)

	if err != nil {
		fmt.Println("Error in Init")
	}

	
	return shim.Success([]byte("Chaincode Successfully initialized"))
}

//CreateDoctor ... this function is used to create doctors
func CreateDoctor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var doctors = Doctors{
		DoctorName:     args[0],
		DoctorId: 		args[1],
		Email:			args[2],
		ContactNum:   	args[3],
		Address:    	args[4]}

	doctorAsBytes, _ := json.Marshal(doctors)
	
	var uniqueID = args[1]	

	err := stub.PutState(uniqueID, doctorAsBytes)

	if err != nil {
		fmt.Println("Erro in create doctor")
	}

	return shim.Success(nil)
}


//Invoke function
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()
	if fun == "CreateDoctor" {
		fmt.Println("Error occured ==> ")
		//logger.Info("########### create docs ###########")
		return CreateDoctor(stub, args)
	} else if fun == "GetDoctorByID" {
		return GetDoctorByID(stub, args)
	} 

	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", fun))
 	
//	return shim.Error(Response.CreateErrorResponse("Invalid function name = "+fun, 0, nil))
}


//GetDoctorByID ... This function will return a particular doctor
func GetDoctorByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.Expected 1 argument")
	}

	resultsIterator, err := stub.GetQueryResult("{\"selector\":{\"doc_type\":\"doctors\",\"_id\":\"" + args[0] + "\"}}")
	if err != nil {
		return shim.Error(err.Error())
	}

	if resultsIterator.HasNext() {
		resultQuery, resultError := resultsIterator.Next()

		if resultError != nil {
			return shim.Error(resultError.Error())
		}

		var doctor Doctors

		unMarshalError := json.Unmarshal(resultQuery.Value, &doctor)

		if unMarshalError != nil {
			return shim.Error(unMarshalError.Error())
		}

		doctorResponse := DoctorByIdResponse{ID: resultQuery.Key, Request: doctor}
		response := Response{Status: "1", Message: "success", Data: doctorResponse}

		responseByte, resConvErr := json.Marshal(response)

		if resConvErr != nil {
			return shim.Error(resConvErr.Error())
		}
		//logger.Infof("Query Response:%s\n", responseByte)
		return shim.Success(responseByte)

	}
	return shim.Error("Could not find any doctors.")

}

func main() {
	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Print(err)
	}
}