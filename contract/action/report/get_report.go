package report

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc_kinder/contract/model"
	"poc_kinder/contract/service"
)

func GetReport(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	var report model.Report
	reportService := service.NewReportService(stub)
	err := reportService.FindAndUnmarshal(args[0], &report)
	if err !=nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(report)
	if err !=nil {
		return "", err
	}

	return string(jsonBytes), nil
}
