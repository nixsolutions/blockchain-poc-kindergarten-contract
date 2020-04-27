# KinderContract chaincode(smart contract)

## Actors

There are 2 actors:
 - Hospital organization
 - Parent organization
 - Kindergarten organization

## Purpose

The main purpose of KinderContract is to store medical reports of children in the blockchain system. There are some rules related to approve and create actions for reports:
 - only hospital worker can create report
 - only parent can approve a report
 - kindergarten can read reports after parent's approve

## Where to start

Every program in Golang starts with a function main(). You can find main() function in  main.go.
In main() function we start the chaincode. We pass a reference of KinderContract object to shim.Start() method to start the chaincode.

## Contract methods

You can find KinderContract struct and implementation of 2 required methods `Invoke` and `Init` in `contract/MedicalContract.go`.
`Init` is called during chaincode instantiation to initialize any data. Note that chaincode upgrade also calls this function to reset or to migrate data.
`Invoke` is called per transaction on the chaincode. Each transaction is either a 'get' or a 'set' on the asset created by `Init` function. The Set method may create a new asset by specifying a new key-value pair.


## Action handlers
You can find actions in `contract/action` folder. Just like controllers in web apps with MVC pattern, actions are connectors between business logic layer of smart contract and transport layer.
There are actions for report:
- createReport - hospital worker can create a report for kindergarten
- getReport - parent, hospital worker and kindergarten can get report from the system
- approveReport - parent can approve report to get kindergarten access to report.

## Models
You can find models in `contract/model` folder. This folder contains entities of smart contract.
There are 3 models:
- Report
- User

## Services
You can find services in `contract/service` folder. This  folder contains business logic of smart contract
There are `ReportService` to work with report. `AuthService` is a wrapper to define who invokes the chaincode.
 
## License
The project is developed by [NIX Solutions](http://nixsolutions.com) Go team and distributed under [MIT LICENSE](https://github.com/nixsolutions/blockchain-poc-kindergarten-contract/master/LICENSE)