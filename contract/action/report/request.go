package report

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc_kinder/contract/service"
)

func RequestReport(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. 1 - report key, 2 - med card id")
	}

	reportService := service.NewReportService(stub)
	report := reportService.RequestReport(args[1])

	jsonBytes, err := json.Marshal(report)
	if err != nil {
		return "", err
	}

	err = reportService.Put(args[0], jsonBytes)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
