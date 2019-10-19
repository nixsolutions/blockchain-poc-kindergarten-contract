package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc_kinder/contract/model"
)

type ReportService struct {
	basicRepo *BasicRepository
	keyPrefix string
}

func NewReportService(stub shim.ChaincodeStubInterface) *ReportService {
	return &ReportService{basicRepo: &BasicRepository{Stub: stub}, keyPrefix: "RPR"}
}

func (service *ReportService) Find(key string) ([]byte, error) {
	return service.basicRepo.Find(service.keyPrefix + key)
}

func (service *ReportService) FindAndUnmarshal(key string, report *model.Report) error {
	return service.basicRepo.FindAndUnmarshal(service.keyPrefix+key, report)
}

func (service *ReportService) Exists(key string) (bool, error) {
	return service.basicRepo.Exists(service.keyPrefix + key)
}

func (service *ReportService) RequestReport(CardId string) model.Report {
	return model.Report{
		Status: model.CREATED_STATUS,
		CardId: CardId,
	}
}

func (service *ReportService) ApproveReport(report *model.Report, Data model.ReportData) {
	report.Data = Data
	report.Status = model.APPROVED_STATUS
}

func (service *ReportService) Put(key string, report []byte) error {
	return service.basicRepo.Stub.PutState(service.keyPrefix+key, report)
}
func (service *ReportService) FindReportsByQuery(queryString string) ([]model.Report, error) {
	resultsIterator, err := service.basicRepo.Stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	var reports []model.Report

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var card model.Report
		err = json.Unmarshal(queryResponse.Value, &card)
		if err != nil {
			return nil, err
		}
		reports = append(reports, card)
	}
	return reports, nil
}
func (service *ReportService) FindAll() ([]model.Report, error) {
	keysIter, err := service.basicRepo.Stub.GetStateByRange("", "")
	if err != nil {
		return nil, err
	}

	var reports []model.Report
	for keysIter.HasNext() {
		res, iterErr := keysIter.Next()
		if iterErr != nil {
			return nil, errors.New(fmt.Sprintf("keys operation failed. Error accessing state: %s", err))
		}

		var report model.Report
		err := json.Unmarshal(res.Value, &report)
		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}
	err = keysIter.Close()
	if err != nil {
		return nil, err
	}

	return reports, nil
}
