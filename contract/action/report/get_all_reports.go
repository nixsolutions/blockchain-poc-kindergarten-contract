package report

import (
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc_kinder/contract/service"
)

func GetReports(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	reportService := service.NewReportService(stub)

	reports, err := reportService.FindAll()
	if err !=nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(reports)
	if err !=nil {
		return "", err
	}

	return string(jsonBytes), nil
}
