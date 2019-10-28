package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"poc/contract"
)

// main function starts up the chaincode in the container during instantiate
func main() {
	//vac := []model.VaccinationItem{
	//	model.VaccinationItem{
	//		Name: "measles",
	//		Timestamp: 1571590625,
	//	},
	//}
	//card := model.Card{
	//	Id: "card-uniq-id",
	//	Name: "John Doe",
	//	BirthDate: "1999-05-24",
	//	Height: 197,
	//	Weight: 85,
	//	Vaccination: vac,
	//	Parents: []model.Parent{ {Id: "parent-1"}, {Id: "parent-2"}},
	//}

	//report := model.Report{
	//	Id: "report-unique-id",
	//	Status: "created",
	//	CardId: "card-uniq-id",
	//	Data: model.ReportData{
	//		Name: "John",
	//		Surname: "Doe",
	//		Vaccination: vac,
	//	},
	//}
	//bytes, err := json.Marshal(report)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(bytes))
	if err := shim.Start(new(contract.MedicalContract)); err != nil {
		fmt.Printf("Error starting MedicalContract chaincode: %s", err)
	}
}
