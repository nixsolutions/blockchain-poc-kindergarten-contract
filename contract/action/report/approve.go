package report

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc_kinder/contract/model"
	"poc_kinder/contract/service"
)

func ApproveRequest(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and report data")
	}

	var report model.Report
	reportService := service.NewReportService(stub)
	user, err := service.NewAuthService(stub).GetUser()
	if err != nil {
		return "", err
	}
	if !user.IsParent() && !user.IsPediatrician() {
		return "", errors.New("only parent or hospital worker can approve request")
	}

	err = reportService.FindAndUnmarshal(args[0], &report)
	if err != nil {
		return "", err
	}

	var reportData model.ReportData
	err = json.Unmarshal([]byte(args[1]), &reportData)
	if err != nil {
		return "", err
	}

	reportService.ApproveReport(&report, reportData)
	jsonBytes, err := json.Marshal(report)
	if err != nil {
		return "", fmt.Errorf("Failed to marshall report obj", args[0])
	}

	return string(jsonBytes), nil
}
