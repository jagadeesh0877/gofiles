package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
	"reflect"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// kpnLogistics will implement the processes
type kpnLogistics struct {
}

type SalesOrderMaster struct {
	CRMOrderNumber				string `json:"crmOrderNumber"`

	CustomerNo					string `json:"customerNo"`
	CRMDashboardStatus			string `json:"crmDashboardStatus"`
	OracleDashboardStatus		string `json:"oracleDashboardStatus"`
	LSPDashboardStatus			string `json:"lspDashboardStatus"`
	TransporterDashboardStatus	string `json:"transporterDashboardStatus"`
	CustomerName				string `json:"customerName"`
	OrderTotalPrice				string `json:"orderTotalPrice"`
	ShippingMethod				string `json:"shippingMethod"`			
	OrderType					string `json:"orderType"`
	DeliveryAddress				string `json:"deliveryAddress"`
	Priority					string `json:"priority"`
	PaymentType					string `json:"paymentType"`
	AdditionalOrderInfo			string `json:"additionalOrderInfo"`
	DeliveryContactName			string `json:"deliveryContactName"`
	DeliveryContactPhone		string `json:"deliveryContactPhone"`
	SourceCode					string `json:"sourceCode"`
	CancelFlag					string `json:"cancelFlag"`
	CancelReason				string `json:"cancelReason"`
	CancelComments				string `json:"cancelComments"`
	TotalLineItems 				string `json:"totallineitems"`
	DemandClass					string `json:"demandClass"`
	CustomerPOref				string `json:"customerPOref"`
		
	OriginalSysDocNumber		string `json:"originalSysDocNumber"`
	OracleOrderNo				string `json:"oracleOrderNo"`
	CRMDateTimeDeliveredToEBS	string `json:"dateTimeDeliveredToEBS"`

	CustomerPhone				string `json:"customerPhone"`
	ShipmentLocation			string `json:"shipmentLocation"`
}
type SalesOrderDetails struct {
	SalesOrderDetailsID 		string `json:"salesOrderDetailsID"`

	CRMOrderNumber				string `json:"crmOrderNumber"`
	ProductID					string `json:"productID"`

	CRMOrderDate				string `json:"crmOrderDate"`
	ProductName					string `json:"productName"`
	ProductValue				string `json:"productValue"`
	ProductQuantity				string `json:"productQuantity"`
	ProductTotalPrice			string `json:"productTotalPrice"`
	ProductDescription			string `json:"ProductDescription"`
	ProductSubTotalPrice		string `json:"ProductSubTotalPrice"`
	Warehouse					string `json:"warehouse"`
	InventoryLocation			string `json:"inventoryLocation"`
	TrackingNumber				string `json:"trackingNumber"`
	QuantityReserved			string `json:"quantityReserved"`
	ScheduleShipDate			string `json:"scheduleShipDate"`
	ScheduleArrivalDate			string `json:"scheduleArrivalDate"`
	Appointmentdate				string `json:"appointmentdate"`
	AppointmentTimefromto		string `json:"appointmentTimefromto"`
	
}
type CRMmaster struct{
	CRMmasterID					string `json:"crmMasterID"`

	CRMOrderNumber				string `json:"crmOrderNumber"`
	CustomerNo					string `json:"customerNo"`
		
	CRMOrderStatus				string `json:"orderStatus"`
	CRMCreatedBy				string `json:"crmCreatedBy"`
	CRMOrderDate				string `json:"crmOrderDate"`
	CRMCustomerPhone			string `json:"crmCustomerPhone"`
	CRMCustomerEmailID			string `json:"crmCustomerEmailID"`
	CRMCarrierInformation		string `json:"crmCarrierInformation"`
	CRMShipToLocation			string `json:"crmShipToLocation"`
	CRMInvoiceLocation			string `json:"crmInvoiceLocation"`
	CRMRequestedDeliveryDate 	string `json:"crmRequestedDeliveryDate"`
	CRMBrand					string `json:"crmBrand"`
	CRMActivityDescription		string `json:"crmActivityDescription"`
	CRMProductsOrdered			string `json:"crmProductsOrdered"`
	CRMType						string `json:"crmType"`				
	CRMOrderSource				string `json:"crmOrderSource"`
	OrderTotalPrice				string `json:"orderTotalPrice"`
	Timestamp					string `json:"timeStamp"`
	FaultCode					string `json:"faultCode"`
	FaultDetail					string `json:"faultDetail"`
}
type CRMDetails struct{
	CRMDetailsID				string `json:"crmDetailsID"`
	SalesOrderDetailsID 		string `json:"salesOrderDetailsID"`
	CRMOrderNumber				string `json:"crmOrderNumber"`
	CRMmasterID					string `json:"crmMasterID"`
	
	ProductID					string `json:"productID"`
	CRMPromiseDate				string `json:"crmPromiseDate"`
	CRMUoM						string `json:"crmuoM"`
	CRMLineShippingInstructions	string `json:"crmLineShippingInstructions"`
	CRMRequestDate				string `json:"crmRequestDate"`
	CRMLinePackingInstructions	string `json:"crmLinePackingInstructions"`
	CRMContractType				string `json:"crmContractType"`
	CRMStatusReason				string `json:"crmStatusReason"`
	CRMStatusText				string `json:"crmStatusText"`
	CRMTestStatus				string `json:"crmtestStatus"`
}

type EBSMaster struct{
	EBSmasterID					string `json:"ebsMasterID"`
	CRMOrderNumber				string `json:"crmOrderNumber"`
		
	OracleOrderNo 				string `json:"oracleOrderNo"`
	OracleSalesOrderNoEBS		string `json:"orderSalesOrderNoEBS"`
	OracleOrderContactEmail		string `json:"oracleOrderContactEmail"`
	OracleWarehouse				string `json:"oracleWarehouse"`
	OracleAccountNumber 		string `json:"oracleAccountNumber"`
	OracleAccountName			string `json:"oracleAccountName"`
	OracleCurrency		 		string `json:"oracleCurrency"`
	OracleOrderAmount	 		string `json:"oracleOrderAmount"`
	OracleDate_TimeEBSCreated	string `json:"orcaleDate_TimeEBSCreated"`
	OracleOrderStatus			string `json:"oracleOrderStatus"`
	CancelFlag					string `json:"cancelFlag"`
	CancelReason				string `json:"cancelReason"`
	CancelComments				string `json:"cancelComments"`
	Onhold						string `json:"onHold"`
	Flag 						string `json:"flag"`
}
type EBSDetails struct{
	EBSDetailsID				string `json:"ebsDetailsID"`

	CRMOrderNumber				string `json:"crmOrderNumber"`
	SalesOrderDetailsID 		string `json:"salesOrderDetailsID"`
	EBSmasterID					string `json:"ebsMasterID"`

	ProductID					string `json:"productID"`
	EBSLineNumber				string `json:"ebsLineNumber"`
	EBSOnHold					string `json:"ebsOnHold"`
	EBSTaxValue					string `json:"ebsTaxValue"`
	EBSSubInventory				string `json:"ebsSubInventory"`
	EBSRequestDate				string `json:"ebsRequestDate"`
	EBSRequestTime				string `json:"ebsRequestTime"`
	EBSQuantityShipped			string `json:"ebsQuantityShipped"`
	EBSDeliveryLineID			string `json:"ebsDeliveryLineID"`
	EBSShipset					string `json:"ebsShipset"`
	EBSQuantityCancelled		string `json:"ebsQuantityCancelled"`
	EBSDate_TimeDeliveryCreated		string `json:"ebsDate_TimeDeliveryCreated"`

	EBSQuantityScheduled		string `json:"ebsQuantityScheduled"`
	EBSQuantityReleased			string `json:"ebsQuantityReleased"`
	EBSQuantityReserved			string `json:"ebsQuantityReserved"`
	DeliverOrderNo				string `json:"DeliverOrderNo"`
}
type LSPMaster struct{
	LSPmasterID					string `json:"lspMasterID"`
	LSPorderno				    string `json:"lspOrderno"`
	
	CRMOrderNumber				string `json:"crmOrderNumber"`
	
	OracleOrderNo				string `json:"oracleOrderNo"`
	CustomerName				string `json:"customerName"`
	ShipmentLocation			string `json:"shipmentLocation"`

	LSPOrderStatus				string `json:"lspOrderStatus"`
	LSPOrderDate				string `json:"lspOrderDate"`
	Deliveryno					string `json:"deliveryno"`
	LspFlag						string `json:"lspFlag"`

}
type LSPDetails struct{
	LSPDetailsID				string `json:"lspDetailsID"`

	CRMOrderNumber				string `json:"crmOrderNumber"`
	LSPmasterID					string `json:"lspMasterID"`

	ProductID					string `json:"productID"`
	LSPproductStatus 			string `json:"lspProductStatus"`
	DeliverOrderNo				string `json:"DeliverOrderNo"`
	
}
type Transporter struct{
	TransporterID				string `json:"transporterID"`
	CRMOrderNumber				string `json:"crmOrderNumber"`
	
	OracleOrderNo				string `json:"oracleOrderNo"`
	DeliverOrderNo				string `json:"DeliverOrderNo"`
	CustomerName				string `json:"customerName"`
	ShipmentLocation			string `json:"shipmentLocation"`
	TransporterOrderStatus		string `json:"transporterOrderStatus"`
	Lspstatus					string `json:"lspstatus"`

}
type SalesOrderTemp struct {
	SOMastertemp  			SalesOrderMaster    `json:"soMastertemp"`
	SODetailstemp 			[]SalesOrderDetails `json:"soDetailstemp"`
	CRMmastertemp 			CRMmaster		    `json:"crmMastertemp"`
	CRMDetailstemp			[]CRMDetails        `json:"crmDetailstemp"`
}
type EBSOrderTemp struct {
	SOMastertemp  			SalesOrderMaster    `json:"soMastertemp"`
	SODetailstemp 			[]SalesOrderDetails `json:"soDetailstemp"`
	EBSMastertemp 			EBSMaster		    `json:"ebsMastertemp"`
	EBSDetailstemp			[]EBSDetails        `json:"ebsDetailstemp"`
}
type LSPOrderTemp struct {
	SOMastertemp  			SalesOrderMaster    `json:"soMastertemp"`
	SODetailstemp 			[]SalesOrderDetails `json:"soDetailstemp"`
	LSPMastertemp 			LSPMaster		    `json:"lspMastertemp"`
	LSPDetailstemp			[]LSPDetails        `json:"lspDetailstemp"`
	EBSMastertemp 			EBSMaster		    `json:"ebsMastertemp"`
	EBSDetailstemp			[]EBSDetails        `json:"ebsDetailstemp"`
}
type Statusconfig struct{
	StatusId			string `json:"statusid"`
	System				string `json:"system"`
	Status 				string `json:"status"`
	Action				string `json:"action"`
	Thresholdtime		string `json:"thresholdtime"`
	StatusReason		string `json:"statusreason"`
	StatusField			string `json:"statusField"`
//modifications starts here
	FieldValue			string `json:"fieldvalue"`
	ConfigShippingMethod string `json:"configshippingmethod"`
	ConfigShippingTime 	string `json:"configshippingtime"`
	ThresholdUnit		string `json:"thresholdunit"`
	Equality		string `json:"equality"`

}
type KeyRecord struct {
	Key   				string   		 	`json:"string"`	
	Record 				LSPDetails    		`json:"Record"`
}
type KeyRecordSO struct {
	Key   				string   		 	`json:"string"`	
	Record				SalesOrderMaster 	`json:"Record"`
}
type KeyRecordCRM struct {
	Key   				string   		 	`json:"string"`	
	Record				CRMmaster		 	`json:"Record"`
}
type KeyRecordSC struct {
	Key   				string   		 	`json:"string"`	
	Record				Statusconfig 	`json:"Record"`
}

type WarningCRMdetails struct {
	Orderno   	 string `json:"ordernumber"`
	Customerno   string `json:"Customerno"`
	OrderStatus  string `json:"orderstatus"`
	Price        string `json:"price"`
	Reason        string `json:"reason"`
	//modifications
	Eventflag        string `json:"eventflag"`
}
type LogWarningCRMdetails struct {
	LogId   	 string `json:"logid"`
	Orderno   	 string `json:"ordernumber"`
	Customerno   string `json:"Customerno"`
	OrderStatus  string `json:"orderstatus"`
	Price        string `json:"price"`
	Reason        string `json:"reason"`
	//modifications
	Eventflag        string `json:"eventflag"`
}
//modifications
type KeyRecordEBS struct {
	Key   				string   		 	`json:"string"`	
	Record				EBSMaster		 	`json:"Record"`
}
//invoke methods CreateSalesOrder - CRM
func (t *kpnLogistics) CreateSalesOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var temp []SalesOrderTemp
	var err error
	
	fmt.Println("Entering Create Sales Order ")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	fmt.Println("Args [0] is : %v\n", args[0])
	
	//unmarshal SalesOrderMaster data from UI to "SalesOrderMaster" struct
	err = json.Unmarshal([]byte(args[0]), &temp)
	if err != nil {
		fmt.Printf("Unable to unmarshal CreateSalesOrder input SalesOrderMaster: %s\n", err)
		return shim.Error(err.Error())
		
	}
	// Data insertion for Couch DB starts here in SalesOrderMaster struct
	for _, order := range temp {

		transJSONasBytesSomaster, err := json.Marshal(order.SOMastertemp)
		err = stub.PutState(order.SOMastertemp.CRMOrderNumber, transJSONasBytesSomaster)

		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs code 001 : %v ", err)
			return shim.Error(err.Error())
		}
		for _, details := range order.SODetailstemp {
			transJSONasBytesSodetails, err := json.Marshal(details)
			err = stub.PutState(details.SalesOrderDetailsID, transJSONasBytesSodetails)

			if err != nil {
				fmt.Printf("\nUnable to make transevent inputs code 002 : %v ", err)
				return shim.Error(err.Error())
			}
		}
	}
		for _, crmMaster := range temp {
			transJSONasBytesCRMmaster, err := json.Marshal(crmMaster.CRMmastertemp)
			err = stub.PutState(crmMaster.CRMmastertemp.CRMmasterID, transJSONasBytesCRMmaster)

			if err != nil {
				fmt.Printf("\nUnable to make transevent inputs code 003 : %v ", err)
				return shim.Error(err.Error())
			}
		
		for _, crmDetails := range crmMaster.CRMDetailstemp {
			transJSONasBytesCRMdetails, err := json.Marshal(crmDetails)
			err = stub.PutState(crmDetails.CRMDetailsID, transJSONasBytesCRMdetails)

			if err != nil {
				fmt.Printf("\nUnable to make transevent inputs code 004 : %v ", err)
				return shim.Error(err.Error())
			}
		}
		}
	
	// Data insertion for Couch DB ends here

	fmt.Println("Create SalesOrderMaster Successfully Done")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//UpdateStatusOnProcessOrder - CRM,EBS
func (t *kpnLogistics) UpdateStatusOnProcessOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcCRMmaster CRMmaster
	var obj_uiCRMmaster CRMmaster
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var objEBSMaster EBSMaster

	var err error

	fmt.Println("Entering UpdateStatusOnProcessOrder")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}
	//unmarshal CRMmaster data from UI to "CRMmaster" struct
	err = json.Unmarshal([]byte(args[0]), &obj_uiCRMmaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateStatusOnProcessOrder input CRMMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Println("\n CRM Master ID value is : ", obj_uiCRMmaster.CRMmasterID)

	//unmarshal SalesOrderMAster data from UI to "SalesOrderMAster" struct
	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateStatusOnProcessOrder input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Println("\n CRM Order Number is : ", obj_uiSalesOrderMaster.CRMOrderNumber)

	//unmarshal EBSMaster data from UI to "EBSMaster" struct
	err = json.Unmarshal([]byte(args[0]), &objEBSMaster)
	if err != nil {
		fmt.Printf("Unable to unmarshal UpdateStatusOnProcessOrder input EBSMaster: %s\n", err)
		return shim.Error(err.Error())
		
	}
	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiCRMmaster.CRMmasterID)
	err = json.Unmarshal(bytesread, &obj_bcCRMmaster)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\noobj_bcCRMmaster : %s ", obj_bcCRMmaster)

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\nobj_bcSalesOrderMaster : %s ", obj_bcSalesOrderMaster)

	obj_bcCRMmaster.CRMOrderStatus = obj_uiCRMmaster.CRMOrderStatus

	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcCRMmaster)
	err = stub.PutState(obj_uiCRMmaster.CRMmasterID, transJSONasBytes)
	// Data insertion for Couch DB ends here

	obj_bcSalesOrderMaster.OracleOrderNo = obj_uiSalesOrderMaster.OracleOrderNo
	obj_bcSalesOrderMaster.CRMDateTimeDeliveredToEBS = obj_uiSalesOrderMaster.CRMDateTimeDeliveredToEBS
	//obj_bcSalesOrderMaster.OracleDashboardStatus = obj_uiSalesOrderMaster.OracleDashboardStatus
	//obj_bcSalesOrderMaster.CRMDashboardStatus = obj_uiSalesOrderMaster.CRMDashboardStatus

	// Data insertion for Couch DB starts here
	transJSONasBytesSales, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytesSales)
	// Data insertion for Couch DB ends here

	transJSONasBytesEBS, err := json.Marshal(objEBSMaster)
	err = stub.PutState(objEBSMaster.EBSmasterID, transJSONasBytesEBS)

	fmt.Println("UpdateStatusOnProcessOrder Successfully updated in SalesOrder master struct")

	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//updateCRMDashboardOnAcknowledge
func (t *kpnLogistics) updateCRMDashboardOnAcknowledge(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var obj_bcFlag EBSMaster
	var obj_uiFlag EBSMaster

	var err error

	fmt.Println("Entering updateCRMDashboardOnAcknowledge")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateCRMDashboardOnAcknowledge input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n CRM Order Number value is : ", obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal([]byte(args[0]), &obj_uiFlag)
	if err != nil {
		fmt.Printf("Unable to marshal  updateCRMDashboardOnAcknowledge input EbsMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcSalesOrderMaster.CRMDashboardStatus = obj_uiSalesOrderMaster.CRMDashboardStatus
	obj_bcSalesOrderMaster.OracleDashboardStatus = obj_uiSalesOrderMaster.OracleDashboardStatus
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytes)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiFlag.EBSmasterID)
	err = json.Unmarshal(bytesread, &obj_bcFlag)
	// code to get data from blockchain using dynamic key ends here

	obj_bcFlag.Flag = obj_uiFlag.Flag
	obj_bcFlag.OracleOrderStatus = obj_uiFlag.OracleOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytesFlag, err := json.Marshal(obj_bcFlag)
	err = stub.PutState(obj_uiFlag.EBSmasterID, transJSONasBytesFlag)
	// Data insertion for Couch DB ends here
	fmt.Println("updateCRMDashboardOnAcknowledge Successfully updated in SalesOrder master struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
		
	return shim.Success(nil)
}
//updateProcessAcknowledge
//updateProcessAcknowledge
func (t *kpnLogistics) updateProcessAcknowledgeXml(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcCRMmaster CRMmaster
	var obj_uiCRMmaster CRMmaster
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var objEBSMaster EBSMaster

	var err error

	fmt.Println("Entering UpdateStatusOnProcessOrder")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}
	//unmarshal CRMmaster data from UI to "CRMmaster" struct
	err = json.Unmarshal([]byte(args[0]), &obj_uiCRMmaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateStatusOnProcessOrder input CRMMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Println("\n CRM Master ID value is : ", obj_uiCRMmaster.CRMmasterID)
	fmt.Println("\n Fault code value is : ", obj_uiCRMmaster.FaultCode)
	fmt.Println("\n Fault detail value is : ", obj_uiCRMmaster.FaultDetail)
	
	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateStatusOnProcessOrder input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Println("\n CRM Order Number is : ", obj_uiSalesOrderMaster.CRMOrderNumber)

	//unmarshal EBSMaster data from UI to "EBSMaster" struct
	err = json.Unmarshal([]byte(args[0]), &objEBSMaster)
	if err != nil {
		fmt.Printf("Unable to unmarshal UpdateStatusOnProcessOrder input EBSMaster: %s\n", err)
		return shim.Error(err.Error())
		
	}
	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiCRMmaster.CRMmasterID)
	err = json.Unmarshal(bytesread, &obj_bcCRMmaster)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\noobj_bcCRMmaster : %s ", obj_bcCRMmaster)

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\nobj_bcSalesOrderMaster : %s ", obj_bcSalesOrderMaster)
	if obj_uiCRMmaster.FaultCode != "E" {
		obj_bcCRMmaster.CRMOrderStatus = obj_uiCRMmaster.CRMOrderStatus
	} else {
		obj_bcCRMmaster.CRMOrderStatus = "Error"
		obj_bcCRMmaster.FaultCode = "E"
		obj_bcCRMmaster.FaultDetail = "Inventory Item not defined in EBS"
	}
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcCRMmaster)
	err = stub.PutState(obj_uiCRMmaster.CRMmasterID, transJSONasBytes)
	// Data insertion for Couch DB ends here
	if obj_uiCRMmaster.FaultCode != "E" {
		obj_bcSalesOrderMaster.OracleOrderNo = obj_uiSalesOrderMaster.OracleOrderNo
		obj_bcSalesOrderMaster.CRMDateTimeDeliveredToEBS = obj_uiSalesOrderMaster.CRMDateTimeDeliveredToEBS
		obj_bcSalesOrderMaster.OracleDashboardStatus = "In-Progress"
		obj_bcSalesOrderMaster.CRMDashboardStatus = "Processed"
	}else{
		obj_bcSalesOrderMaster.CRMDashboardStatus = "Error"
	}

	// Data insertion for Couch DB starts here
	transJSONasBytesSales, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytesSales)
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	// Data insertion for Couch DB ends here
	if obj_uiCRMmaster.FaultCode != "E" {
		objEBSMaster.Flag = "1"
		objEBSMaster.OracleOrderStatus = "Order Placed"
		transJSONasBytesEBS, err := json.Marshal(objEBSMaster)
		err = stub.PutState(objEBSMaster.EBSmasterID, transJSONasBytesEBS)

		fmt.Println("UpdateStatusOnProcessOrder Successfully updated in SalesOrder master struct")

		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs : %v ", err)
			return shim.Error(err.Error())
		}
	}
	//unmarshal SalesOrderMAster data from UI to "SalesOrderMAster" struct
	
	return shim.Success(nil)
}
//updateCRMDashboardOnError
func (t *kpnLogistics) updateCRMDashboardOnError(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var obj_bcEbsMaster EBSMaster
	var obj_uiEbsMaster EBSMaster
	var obj_bcCRMmaster CRMmaster
	var obj_uiCRMmaster CRMmaster
	
	var err error

	fmt.Println("Entering updateCRMDashboardOnError")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateCRMDashboardOnError input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiEbsMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateCRMDashboardOnError input Ebsmaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiCRMmaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateCRMDashboardOnError input CRMmAster : %s\n", err)
		return shim.Error(err.Error())
	}
	
	fmt.Println("\n CRM Order Number value is : ", obj_uiSalesOrderMaster.CRMOrderNumber)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcSalesOrderMaster.CRMDashboardStatus = obj_uiSalesOrderMaster.CRMDashboardStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytes)
	// Data insertion for Couch DB ends here
	// code to get data from blockchain using dynamic key starts here
	var bytesread1 []byte 
	bytesread1, err = stub.GetState(obj_uiEbsMaster.EBSmasterID)
	err = json.Unmarshal(bytesread1, &obj_bcEbsMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcEbsMaster.OracleOrderStatus = obj_uiEbsMaster.OracleOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytesEbs, err := json.Marshal(obj_bcEbsMaster)
	err = stub.PutState(obj_uiEbsMaster.EBSmasterID, transJSONasBytesEbs)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	var bytesread2 []byte
	bytesread2, err = stub.GetState(obj_uiCRMmaster.CRMmasterID)
	err = json.Unmarshal(bytesread2, &obj_bcCRMmaster)
	// code to get data from blockchain using dynamic key ends here
	obj_bcCRMmaster.CRMOrderStatus = obj_uiCRMmaster.CRMOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytesCRM, err := json.Marshal(obj_bcCRMmaster)
	err = stub.PutState(obj_uiCRMmaster.CRMmasterID, transJSONasBytesCRM)
	// Data insertion for Couch DB ends here

	fmt.Println("updateCRMDashboardOnError Successfully updated in SalesOrder master struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
		
	return shim.Success(nil)
}

//updateEBSDashboardonError
func (t *kpnLogistics) updateEBSDashboardonError(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var obj_bcEbsMaster EBSMaster
	var obj_uiEbsMaster EBSMaster
	var obj_bcLSPMaster LSPMaster
	var obj_uiLSPMaster LSPMaster
	
	var err error

	fmt.Println("Entering updateEBSDashboardonError")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updatEBSDashboardOnError input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiEbsMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateEBSDashboardOnError input Ebsmaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiLSPMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateEBSDashboardOnError input LSPMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	
	fmt.Println("\n CRM Order Number value is : ", obj_uiSalesOrderMaster.CRMOrderNumber)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcSalesOrderMaster.OracleDashboardStatus = obj_uiSalesOrderMaster.OracleDashboardStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytes)
	// Data insertion for Couch DB ends here
	// code to get data from blockchain using dynamic key starts here
	var bytesread1 []byte 
	bytesread1, err = stub.GetState(obj_uiEbsMaster.EBSmasterID)
	err = json.Unmarshal(bytesread1, &obj_bcEbsMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcEbsMaster.OracleOrderStatus = obj_uiEbsMaster.OracleOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytesEbs, err := json.Marshal(obj_bcEbsMaster)
	err = stub.PutState(obj_uiEbsMaster.EBSmasterID, transJSONasBytesEbs)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	var bytesread2 []byte
	bytesread2, err = stub.GetState(obj_uiLSPMaster.LSPmasterID)
	err = json.Unmarshal(bytesread2, &obj_bcLSPMaster)
	// code to get data from blockchain using dynamic key ends here
	obj_bcLSPMaster.LSPOrderStatus = obj_uiLSPMaster.LSPOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytesLSP, err := json.Marshal(obj_bcLSPMaster)
	err = stub.PutState(obj_uiLSPMaster.LSPmasterID, transJSONasBytesLSP)
	// Data insertion for Couch DB ends here

	fmt.Println("updateCRMDashboardOnError Successfully updated in SalesOrder master struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
		
	return shim.Success(nil)
}
//UpdateStatusOnCancelOrder - CRM
func (t *kpnLogistics) UpdateStatusOnCancelOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcCRMmaster CRMmaster
	var obj_uiCRMmaster CRMmaster

	var err error

	fmt.Println("Entering UpdateStatusOnCancelOrder")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiCRMmaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateStatusOnCancelOrder input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n CRM Order Number value is : ", obj_uiCRMmaster.CRMmasterID)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiCRMmaster.CRMmasterID)
	err = json.Unmarshal(bytesread, &obj_bcCRMmaster)
	// code to get data from blockchain using dynamic key ends here

	if ((obj_bcCRMmaster.CRMOrderStatus=="Created")||(obj_bcCRMmaster.CRMOrderStatus=="In-Progress")) {
	obj_bcCRMmaster.CRMOrderStatus = obj_uiCRMmaster.CRMOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcCRMmaster)
	err = stub.PutState(obj_uiCRMmaster.CRMmasterID, transJSONasBytes)
	// Data insertion for Couch DB ends here
	fmt.Println("UpdateStatusOnCancelOrder Successfully updated in SalesOrder master struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	} else {
		fmt.Println("\nError while cancelling the Order. Cannot Cancel the Delivered Order")
	}
	
	return shim.Success(nil)
}

//UpdateProcessOrderInEBS - EBS
func (t *kpnLogistics) UpdateProcessOrderInEBS(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var temp []EBSOrderTemp
	var obj_bcEBSMaster EBSMaster
	var err error

	fmt.Println("Entering UpdateProcessOrderInEBS")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &temp)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateProcessOrderInEBS input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}

	for _, order := range temp {
	var bytesread []byte
	bytesread, err = stub.GetState(order.EBSMastertemp.EBSmasterID)
	err = json.Unmarshal(bytesread, &obj_bcEBSMaster)

	obj_bcEBSMaster.CRMOrderNumber = order.EBSMastertemp.CRMOrderNumber
	obj_bcEBSMaster.OracleOrderNo  = order.EBSMastertemp.OracleOrderNo 
	obj_bcEBSMaster.OracleSalesOrderNoEBS = order.EBSMastertemp.OracleSalesOrderNoEBS
	obj_bcEBSMaster.OracleOrderContactEmail = order.EBSMastertemp.OracleOrderContactEmail
	obj_bcEBSMaster.OracleWarehouse = order.EBSMastertemp.OracleWarehouse
	obj_bcEBSMaster.OracleAccountNumber = order.EBSMastertemp.OracleAccountNumber
	obj_bcEBSMaster.OracleAccountName = order.EBSMastertemp.OracleAccountName
	obj_bcEBSMaster.OracleCurrency = order.EBSMastertemp.OracleCurrency
	obj_bcEBSMaster.OracleOrderAmount = order.EBSMastertemp.OracleOrderAmount
	obj_bcEBSMaster.OracleDate_TimeEBSCreated = order.EBSMastertemp.OracleDate_TimeEBSCreated
	obj_bcEBSMaster.OracleOrderStatus = order.EBSMastertemp.OracleOrderStatus
	obj_bcEBSMaster.CancelFlag = order.EBSMastertemp.CancelFlag
	obj_bcEBSMaster.CancelReason = order.EBSMastertemp.CancelReason
	obj_bcEBSMaster.CancelComments = order.EBSMastertemp.CancelComments
	obj_bcEBSMaster.Onhold = order.EBSMastertemp.Onhold
	
	transJSONasBytes, err := json.Marshal(obj_bcEBSMaster)
	err = stub.PutState(order.EBSMastertemp.EBSmasterID, transJSONasBytes)

		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs code 001 : %v ", err)
			return shim.Error(err.Error())
		}
		for _, details := range order.EBSDetailstemp {
			transJSONasBytesSodetails, err := json.Marshal(details)
			err = stub.PutState(details.EBSDetailsID, transJSONasBytesSodetails)

			if err != nil {
				fmt.Printf("\nUnable to make transevent inputs code 002 : %v ", err)
				return shim.Error(err.Error())
			}
		}
	}
	
	fmt.Println("UpdateProcessOrderInEBS Successfully updated in EBS master struct")
		
	return shim.Success(nil)
}

//UpdateReleaseOrderInEBS - EBS
func (t *kpnLogistics) UpdateReleaseOrderInEBS(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcEBSMaster EBSMaster
	var obj_uiEBSMaster EBSMaster
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var objLSPMaster LSPMaster
	var err error

	fmt.Println("Entering UpdateReleaseOrderInEBS")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiEBSMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateReleaseOrderInEBS input EBSMaster : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n CRM Order Number value is : ", obj_uiEBSMaster.EBSmasterID)

	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateReleaseOrderInEBS input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &objLSPMaster)
	if err != nil {
		fmt.Printf("Unable to unmarshal UpdateReleaseOrderInEBS input LSPMaster: %s\n", err)
		return shim.Error(err.Error())
		
	}
	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiEBSMaster.EBSmasterID)
	err = json.Unmarshal(bytesread, &obj_bcEBSMaster)
	// code to get data from blockchain using dynamic key ends here

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	if (obj_bcEBSMaster.OracleOrderStatus=="Pending"){
	obj_bcEBSMaster.OracleOrderStatus = obj_uiEBSMaster.OracleOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcEBSMaster)
	err = stub.PutState(obj_bcEBSMaster.EBSmasterID, transJSONasBytes)
	// Data insertion for Couch DB ends here
	fmt.Println("UpdateReleaseOrderInEBS Successfully updated in EBS master struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	} else {
		fmt.Println("\nError while cancelling the Order. Cannot Cancel the Delivered Order")
	}
	//obj_bcSalesOrderMaster.OracleDashboardStatus = obj_uiSalesOrderMaster.OracleDashboardStatus
	obj_bcSalesOrderMaster.LSPDashboardStatus = obj_uiSalesOrderMaster.LSPDashboardStatus

	// Data insertion for Couch DB starts here
	transJSONasBytesstatus, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytesstatus)
	// Data insertion for Couch DB ends here

	// Data insertion for Couch DB starts here
	transJSONasBytesLSP, err := json.Marshal(objLSPMaster)
	err = stub.PutState(objLSPMaster.LSPmasterID, transJSONasBytesLSP)
	// Data insertion for Couch DB ends here
	fmt.Println("UpdateReleaseOrderInEBS Successfully Done")


	return shim.Success(nil)
}




func (t *kpnLogistics) UpdateReleaseOrderAndGenerateLspDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	var temp []LSPOrderTemp
	//var obj_bcLSPMaster LSPMaster
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_bcEbsMaster EBSMaster
	var err error

	fmt.Println("Entering Release and Update ")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	fmt.Println("Args [0] is : %v\n", args[0])
	
	//unmarshal SalesOrderMaster data from UI to "SalesOrderMaster" struct
	err = json.Unmarshal([]byte(args[0]), &temp)
	if err != nil {
		fmt.Printf("Unable to unmarshal CreateSalesOrder input SalesOrderMaster: %s\n", err)
		return shim.Error(err.Error())
		
	}
	
	
	for _, order := range temp {
		
		fmt.Println("\n CRMOrderNumber Number value is : ",order.SOMastertemp.CRMOrderNumber)
		fmt.Println("Inside for loop ")
		var bytesread []byte
		bytesread, err = stub.GetState(order.SOMastertemp.CRMOrderNumber)
		err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
		
		obj_bcSalesOrderMaster.LSPDashboardStatus = order.SOMastertemp.LSPDashboardStatus

		transJSONasBytes, err := json.Marshal(obj_bcSalesOrderMaster)
		err = stub.PutState(order.SOMastertemp.CRMOrderNumber, transJSONasBytes)
		
		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs code 001 : %v ", err)
			return shim.Error(err.Error())
		}
		
	}


	for _, order := range temp {
		
		fmt.Println("\n EbsMstId Number value is : ",order.EBSMastertemp.EBSmasterID)
		fmt.Println("Inside 2nd for loop ")
		var bytesread []byte
		bytesread, err = stub.GetState(order.EBSMastertemp.EBSmasterID)
		err = json.Unmarshal(bytesread, &obj_bcEbsMaster)
		
		obj_bcEbsMaster.OracleOrderStatus = order.EBSMastertemp.OracleOrderStatus

		transJSONasBytes, err := json.Marshal(obj_bcEbsMaster)
		err = stub.PutState(order.EBSMastertemp.EBSmasterID, transJSONasBytes)
		
		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs code 001 : %v ", err)
			return shim.Error(err.Error())
		}
		
	}

	for _, order := range temp {
		
		//fmt.Println("\n EbsMstId Number value is : ",order.EBSMastertemp.EBSmasterID)
		fmt.Println("Inside 3rd for loop ")
		transJSONasBytesSomaster, err := json.Marshal(order.LSPMastertemp)
		err = stub.PutState(order.LSPMastertemp.LSPmasterID, transJSONasBytesSomaster)
        
		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs code 001 : %v ", err)
			return shim.Error(err.Error())
		}

		for _, details := range order.LSPDetailstemp {
			fmt.Println("Inside sub for loop ")
			transJSONasBytes, err := json.Marshal(details)
			err = stub.PutState(details.LSPDetailsID, transJSONasBytes)

			if err != nil {
				fmt.Printf("\nUnable to make transevent inputs code 002 : %v ", err)
				return shim.Error(err.Error())
			}
		}
		
	}
	return shim.Success(nil)
}

 

//GenerateLSPDetailsID
func (t *kpnLogistics) GenerateLSPDetailsID(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var temp []LSPOrderTemp
	var obj_bcLSPMaster LSPMaster
	var err error

	fmt.Println("Entering GenerateLSPDetailsID")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &temp)
	if err != nil {
		fmt.Printf("Unable to marshal  GenerateLSPDetailsID input LSPMaster : %s\n", err)
		return shim.Error(err.Error())
	}

	for _, order := range temp {
	var bytesread []byte
	bytesread, err = stub.GetState(order.LSPMastertemp.LSPmasterID)
	err = json.Unmarshal(bytesread, &obj_bcLSPMaster)

	obj_bcLSPMaster.LSPmasterID = order.LSPMastertemp.LSPmasterID
	
	transJSONasBytes, err := json.Marshal(obj_bcLSPMaster)
	err = stub.PutState(order.LSPMastertemp.LSPmasterID, transJSONasBytes)

		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs code 001 : %v ", err)
			return shim.Error(err.Error())
		}
		for _, details := range order.LSPDetailstemp {
			transJSONasBytesSodetails, err := json.Marshal(details)
			err = stub.PutState(details.LSPDetailsID, transJSONasBytesSodetails)

			if err != nil {
				fmt.Printf("\nUnable to make transevent inputs code 002 : %v ", err)
				return shim.Error(err.Error())
			}
		}
	}
	
	fmt.Println("UpdateProcessOrderInEBS Successfully updated in EBS master struct")
		
	return shim.Success(nil)
}

//updateLSPDashboardOnAcknowledge
func (t *kpnLogistics) updateLSPDashboardOnAcknowledge(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var obj_bcLSPMaster LSPMaster
	var obj_uiLSPMaster LSPMaster

	var err error

	fmt.Println("Entering updateLSPDashboardOnAcknowledge")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateLSPDashboardOnAcknowledge input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiLSPMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateLSPDashboardOnAcknowledge input LSPMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Println("\n CRM Order Number value is : ", obj_uiSalesOrderMaster.CRMOrderNumber)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcSalesOrderMaster.OracleDashboardStatus = obj_uiSalesOrderMaster.OracleDashboardStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytes)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	var bytesreads []byte
	bytesreads, err = stub.GetState(obj_uiLSPMaster.LSPmasterID)
	err = json.Unmarshal(bytesreads, &obj_bcLSPMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcLSPMaster.LspFlag = obj_uiLSPMaster.LspFlag
	obj_bcLSPMaster.LSPOrderStatus = obj_uiLSPMaster.LSPOrderStatus

	// Data insertion for Couch DB starts here
	transJSONasBytesLsp, err := json.Marshal(obj_bcLSPMaster)
	err = stub.PutState(obj_uiLSPMaster.LSPmasterID, transJSONasBytesLsp)
	// Data insertion for Couch DB ends here
	fmt.Println("updateCRMDashboardOnAcknowledge Successfully updated in SalesOrder master struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
		
	return shim.Success(nil)
}

//UpdateLSPStatusesForProductID
func (t *kpnLogistics) UpdateLSPStatusesForProductID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_uiLSPDetails LSPDetails
	var obj_bcLSPDetails LSPDetails
	var err error

	fmt.Println("Entering UpdateLSPStatusesForProductID")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiLSPDetails)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateLSPStatusesForProductID input EBSMaster : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n CRM Order Number value is : ", obj_uiLSPDetails.LSPDetailsID)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiLSPDetails.LSPDetailsID)
	err = json.Unmarshal(bytesread, &obj_bcLSPDetails)
	// code to get data from blockchain using dynamic key ends here

	obj_bcLSPDetails.LSPDetailsID = obj_uiLSPDetails.LSPDetailsID
	obj_bcLSPDetails.ProductID = obj_uiLSPDetails.ProductID
	obj_bcLSPDetails.CRMOrderNumber = obj_uiLSPDetails.CRMOrderNumber
	obj_bcLSPDetails.LSPproductStatus = obj_uiLSPDetails.LSPproductStatus

	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcLSPDetails)
	err = stub.PutState(obj_uiLSPDetails.LSPDetailsID, transJSONasBytes)
	// Data insertion for Couch DB ends here
	fmt.Println("UpdateLSPStatusesForProductID Successfully updated in LSPDetails struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	
	return shim.Success(nil)
}
//UpdateShipOrderInLSP
func (t *kpnLogistics) UpdateShipOrderInLSP(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_uiLSPDetails LSPDetails
	var obj_bcLSPDetails LSPDetails
	var obj_uiLSPMaster LSPMaster
	var obj_bcLSPMaster LSPMaster
	var objTransporter Transporter
	var obj_uiSalesOrderMaster SalesOrderMaster
	var obj_bcSalesOrderMaster SalesOrderMaster
	var err error

	fmt.Println("Entering UpdateShipOrderInLSP")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiLSPDetails)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateShipOrderInLSP input LSPDetails : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Println("\n CRM Order Number value is : ", obj_uiLSPDetails.LSPDetailsID)

	err = json.Unmarshal([]byte(args[0]), &obj_uiLSPMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateShipOrderInLSP input LSPMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Println("\n CRM Order Number value is : ", obj_uiLSPMaster.LSPmasterID)

	err = json.Unmarshal([]byte(args[0]), &objTransporter)
	if err != nil {
		fmt.Printf("Unable to unmarshal UpdateShipOrderInLSP input Transporter: %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to unmarshal UpdateShipOrderInLSP input SalesOrderMaster: %s\n", err)
		return shim.Error(err.Error())
	}
	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiLSPDetails.LSPDetailsID)
	err = json.Unmarshal(bytesread, &obj_bcLSPDetails)
	// code to get data from blockchain using dynamic key ends here

	obj_bcLSPDetails.LSPDetailsID = obj_uiLSPDetails.LSPDetailsID
	obj_bcLSPDetails.ProductID = obj_uiLSPDetails.ProductID
	obj_bcLSPDetails.CRMOrderNumber = obj_uiLSPDetails.CRMOrderNumber
	obj_bcLSPDetails.LSPproductStatus = "Ship Confirmed"
	obj_bcLSPDetails.DeliverOrderNo = obj_uiLSPDetails.DeliverOrderNo

	// Data insertion for Couch DB starts here
	transJSONasBytesLSPDetails, err := json.Marshal(obj_bcLSPDetails)
	err = stub.PutState(obj_uiLSPDetails.LSPDetailsID, transJSONasBytesLSPDetails)
	fmt.Println("\n----------- obj_bcLSPDetails------------ \n",obj_bcLSPDetails)
	// Data insertion for Couch DB ends here
//	var lspmasterID=obj_uiLSPMaster.LSPmasterID
	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiLSPMaster.LSPmasterID)
	err = json.Unmarshal(bytesread, &obj_bcLSPMaster)
	// code to get data from blockchain using dynamic key ends here

	//obj_bcLSPMaster.LSPmasterID = obj_uiLSPMaster.LSPmasterID
	//obj_bcLSPMaster.LSPOrderStatus = obj_uiLSPMaster.LSPOrderStatus
	
	// Data insertion for Couch DB starts here
	//transJSONasBytesLSPMastr, err := json.Marshal(obj_bcLSPMaster)
	//err = stub.PutState(obj_uiLSPMaster.LSPmasterID, transJSONasBytesLSPMastr)
	// Data insertion for Couch DB ends here

	// Data insertion for Couch DB starts here
	transJSONasBytesTransporter, err := json.Marshal(objTransporter)
	err = stub.PutState(objTransporter.TransporterID, transJSONasBytesTransporter)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic keyly ends here

	obj_bcSalesOrderMaster.TransporterDashboardStatus = obj_uiSalesOrderMaster.TransporterDashboardStatus
		
	// Data insertion for Couch DB starts here
	transJSONasBytesSales, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_bcSalesOrderMaster.CRMOrderNumber, transJSONasBytesSales)
	// Data insertion for Couch DB ends here

	fmt.Println("UpdateLSPStatusesForProductID Successfully updated in LSPDetails struct, LSPMaster struct & Transporter struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
//updateShipmentInLSPOrderStatus
func (t *kpnLogistics) updateShipmentInLSPOrderStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcLSPMaster LSPMaster
	var obj_uiLSPMaster LSPMaster
	var obj_uiTransporter Transporter
	var obj_bcTransporter Transporter
	var err error

	fmt.Println("Entering updateShipmentInLSPOrderStatus")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiLSPMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateShipmentInLSPOrderStatus input LSPMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiTransporter)
	if err != nil {
		fmt.Printf("Unable to marshal  updateShipmentInLSPOrderStatus input Transporter : %s\n", err)
		return shim.Error(err.Error())
	}

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiLSPMaster.LSPmasterID)
	err = json.Unmarshal(bytesread, &obj_bcLSPMaster)

	bytesread, err = stub.GetState(obj_uiTransporter.TransporterID)
	err = json.Unmarshal(bytesread, &obj_bcTransporter)
	// code to get data from blockchain using dynamic key ends here
 	var lspmasterID = obj_uiLSPMaster.LSPmasterID
	//if (obj_bcLSPDetails.LSPproductStatus == "Ship Confirmed") {
	queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"lspMasterID\":{\"$eq\":\"%s\"}},{\"lspProductStatus\":{\"$ne\":\"%s\"}}]}}", lspmasterID, "null")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	
	var arr_config []KeyRecord
		err = json.Unmarshal([]byte(queryResults), &arr_config)
		if err != nil {
			panic(err)
		}
	var flag_shipconfirmed = "true";
	for _, order := range arr_config {
		fmt.Println("\n Record -------------",order.Record)
		fmt.Println("\n LSP Order Stautus ----------",order.Record.LSPproductStatus)

		if (order.Record.LSPproductStatus == "Ship Confirmed") {
			fmt.Println("\ninside 1st else -------------\n")
		  	fmt.Println("\n Record -------------",order.Record)
		   	fmt.Println("\n LSP Order Stautus ----------",order.Record.LSPproductStatus)
			obj_bcLSPMaster.LSPOrderStatus = "Partially Shipped"

			transJSONasBytesLSPMast, err := json.Marshal(obj_bcLSPMaster)
			err = stub.PutState(obj_uiLSPMaster.LSPmasterID, transJSONasBytesLSPMast)
			if err != nil {
				fmt.Printf("\nUnable to make transevent inputs : %v ", err)
				return shim.Error(err.Error())
			}//done

		}

		if(order.Record.LSPproductStatus == "Released" || order.Record.LSPproductStatus == "Order Pick" || order.Record.LSPproductStatus == "Order Pack"){
			flag_shipconfirmed = "false"
			fmt.Printf("\n -------------- flag_shipconfirmed to 1 ---------------------")
		}

	}
	if(flag_shipconfirmed == "true")	{
		obj_bcLSPMaster.LSPOrderStatus = "Shipment In Transit"
		transJSONasBytesLSPMastr111, err := json.Marshal(obj_bcLSPMaster)
		err = stub.PutState(obj_uiLSPMaster.LSPmasterID, transJSONasBytesLSPMastr111)
		fmt.Printf("\n -------------- changing status to completely ---------------------")
		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs : %v ", err)
			return shim.Error(err.Error())
		}

		obj_bcTransporter.Lspstatus= "Shipment In Transit"
		transJSONasBytesLSPMastr22, err := json.Marshal(obj_bcTransporter)
		err = stub.PutState(obj_uiTransporter.TransporterID, transJSONasBytesLSPMastr22)
		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs : %v ", err)
			return shim.Error(err.Error())
		}
  }
	if(flag_shipconfirmed == "hi")	{
		obj_bcLSPMaster.LSPOrderStatus = "Completely Shipped"
		transJSONasBytesLSPMastr111, err := json.Marshal(obj_bcLSPMaster)
		err = stub.PutState(obj_uiLSPMaster.LSPmasterID, transJSONasBytesLSPMastr111)
		fmt.Printf("\n -------------- changing status to completely ---------------------")
		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs : %v ", err)
			return shim.Error(err.Error())
		}

		obj_bcTransporter.Lspstatus= "Completely Shipped"
		transJSONasBytesLSPMastr22, err := json.Marshal(obj_bcTransporter)
		err = stub.PutState(obj_uiTransporter.TransporterID, transJSONasBytesLSPMastr22)
		if err != nil {
			fmt.Printf("\nUnable to make transevent inputs : %v ", err)
			return shim.Error(err.Error())
		}
  }
	return shim.Success(nil)
}

 
//starts here
//UpdateReleaseOrderInEBS - EBS
//old func

//old func
//ends here
//updateDashboardOnTransporterAcknowledge
func (t *kpnLogistics) updateDashboardOnTransporterAcknowledge(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcSalesOrderMaster SalesOrderMaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var obj_bcTransport Transporter
	var obj_uiTransport Transporter
	var obj_bcLSPMaster LSPMaster
	var obj_uiLSPMaster LSPMaster

	var err error

	fmt.Println("Entering updateDashboardOnTransporterAcknowledge")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateDashboardOnTransporterAcknowledge input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiTransport)
	if err != nil {
		fmt.Printf("Unable to marshal  updateDashboardOnTransporterAcknowledge input LSPMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiLSPMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  updateDashboardOnTransporterAcknowledge input LSPMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Println("\n CRM Order Number value is : ", obj_uiSalesOrderMaster.CRMOrderNumber)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcSalesOrderMaster.LSPDashboardStatus = obj_uiSalesOrderMaster.LSPDashboardStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytes)
	// Data insertion for Couch DB ends here
	var bytesreads []byte
	bytesreads, err = stub.GetState(obj_uiTransport.TransporterID)
	err = json.Unmarshal(bytesreads, &obj_bcTransport)
	// code to get data from blockchain using dynamic key ends here

	obj_bcTransport.DeliverOrderNo = obj_uiTransport.DeliverOrderNo
	obj_bcTransport.TransporterOrderStatus = obj_uiTransport.TransporterOrderStatus
	
	
	// Data insertion for Couch DB starts here
	transJSONasBytesDeliver, err := json.Marshal(obj_bcTransport)
	err = stub.PutState(obj_uiTransport.TransporterID, transJSONasBytesDeliver)
	fmt.Println("updateTransportDashboardOnAcknowledge Successfully updated in struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	var bytesreadr []byte
	bytesreadr, err = stub.GetState(obj_uiLSPMaster.LSPmasterID)
	err = json.Unmarshal(bytesreadr, &obj_bcLSPMaster)
	// code to get data from blockchain using dynamic key ends here

	obj_bcLSPMaster.LSPOrderStatus = obj_uiLSPMaster.LSPOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytesLsp, err := json.Marshal(obj_bcLSPMaster)
	err = stub.PutState(obj_uiLSPMaster.LSPmasterID, transJSONasBytesLsp)
	fmt.Println("updateTransportDashboardOnAcknowledge Successfully updated in struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

//updateStatusInTransporter
func (t *kpnLogistics) updateStatusInTransporter(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcTransporter Transporter
	var obj_uiTransporter Transporter

	var err error

	fmt.Println("Entering updateStatusInTransporter")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiTransporter)
	if err != nil {
		fmt.Printf("Unable to marshal  updateStatusInTransporter input Transporter : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n CRM Order Number value is : ", obj_uiTransporter.TransporterID)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiTransporter.TransporterID)
	err = json.Unmarshal(bytesread, &obj_bcTransporter)
	// code to get data from blockchain using dynamic key ends here

	obj_bcTransporter.TransporterOrderStatus = obj_uiTransporter.TransporterOrderStatus
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcTransporter)
	err = stub.PutState(obj_uiTransporter.TransporterID, transJSONasBytes)
	// Data insertion for Couch DB ends here
	fmt.Println("updateStatusInTransporter Successfully updated in Transporter struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
		
	return shim.Success(nil)
}

//UpdateShipDeliveredInTransporter
func (t *kpnLogistics) UpdateShipDeliveredInTransporter(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcTransporter Transporter
	var obj_uiTransporter Transporter
	var obj_uiLSPMaster LSPMaster
	var obj_bcLSPMaster LSPMaster 
	var obj_bcEBSMaster EBSMaster
	var obj_uiEBSMaster EBSMaster
	var obj_bcCRMmaster CRMmaster
	var obj_uiCRMmaster CRMmaster
	var obj_uiSalesOrderMaster SalesOrderMaster
	var obj_bcSalesOrderMaster SalesOrderMaster

	var err error

	fmt.Println("Entering UpdateShipDeliveredInTransporter")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiTransporter)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateShipDeliveredInTransporter input Transporter : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiLSPMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateShipDeliveredInTransporter input LSPMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiEBSMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateShipDeliveredInTransporter input EBSMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiCRMmaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateShipDeliveredInTransporter input CRMMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	err = json.Unmarshal([]byte(args[0]), &obj_uiSalesOrderMaster)
	if err != nil {
		fmt.Printf("Unable to marshal  UpdateShipDeliveredInTransporter input SalesOrderMaster : %s\n", err)
		return shim.Error(err.Error())
	}
	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiTransporter.TransporterID)
	err = json.Unmarshal(bytesread, &obj_bcTransporter)
	// code to get data from blockchain using dynamic key ends here

	obj_bcTransporter.TransporterOrderStatus = obj_uiTransporter.TransporterOrderStatus
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcTransporter)
	err = stub.PutState(obj_uiTransporter.TransporterID, transJSONasBytes)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiLSPMaster.LSPmasterID)
	err = json.Unmarshal(bytesread, &obj_bcLSPMaster)
	// code to get data from blockchain using dynamic key ends here
	
	obj_bcLSPMaster.LSPOrderStatus = obj_uiLSPMaster.LSPOrderStatus
	// Data insertion for Couch DB starts here
	transJSONasBytesLSP, err := json.Marshal(obj_bcLSPMaster)
	err = stub.PutState(obj_uiLSPMaster.LSPmasterID, transJSONasBytesLSP)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiEBSMaster.EBSmasterID)
	err = json.Unmarshal(bytesread, &obj_bcEBSMaster)
	// code to get data from blockchain using dynamic key ends here
	
	obj_bcEBSMaster.OracleOrderStatus = obj_uiEBSMaster.OracleOrderStatus
	// Data insertion for Couch DB starts here
	transJSONasBytesEBS, err := json.Marshal(obj_bcEBSMaster)
	err = stub.PutState(obj_uiEBSMaster.EBSmasterID, transJSONasBytesEBS)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiCRMmaster.CRMmasterID)
	err = json.Unmarshal(bytesread, &obj_bcCRMmaster)
	// code to get data from blockchain using dynamic key ends here
	
	obj_bcCRMmaster.CRMOrderStatus = obj_uiCRMmaster.CRMOrderStatus
	// Data insertion for Couch DB starts here
	transJSONasBytesCRM, err := json.Marshal(obj_bcCRMmaster)
	err = stub.PutState(obj_uiCRMmaster.CRMmasterID, transJSONasBytesCRM)
	// Data insertion for Couch DB ends here

	// code to get data from blockchain using dynamic key starts here
	bytesread, err = stub.GetState(obj_uiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here
	
	obj_bcSalesOrderMaster.TransporterDashboardStatus = obj_uiSalesOrderMaster.TransporterDashboardStatus
	//modifcations starts here

	obj_bcSalesOrderMaster.CRMDashboardStatus = obj_uiSalesOrderMaster.CRMDashboardStatus
	
	//modifcations ends here
	// Data insertion for Couch DB starts here
	transJSONasBytesSales, err := json.Marshal(obj_bcSalesOrderMaster)
	err = stub.PutState(obj_uiSalesOrderMaster.CRMOrderNumber, transJSONasBytesSales)
	// Data insertion for Couch DB ends here
	fmt.Println("updateStatusInTransporter Successfully updated in Transporter struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
		
	return shim.Success(nil)
}

//get all Sales order
func (t *kpnLogistics) getAllSalesOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	fmt.Println("Entering get All Sales Orders")
	//crmMasterID
	//fetch data from couch db starts here
	queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmMasterID\":{\"$ne\":\"%s\"}},{\"customerNo\":{\"$ne\":\"%s\"}}]}}", "null", "null")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here
	if err != nil {
		fmt.Printf("Unable to read the list of Purchase orders : %s\n", err)
		return shim.Error(err.Error())
		//return nil, err
	}

	//fmt.Printf("list of all purchase order details : %v\n", queryResults)
	return shim.Success(queryResults)
}
//getAllEBSOrders
func (t *kpnLogistics) getAllEBSOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get All Sales Orders")
		//crmMasterID
		//fetch data from couch db starts here
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$ne\":\"%s\"}},{\"oracleOrderNo\":{\"$ne\":\"%s\"}}]}}", "null", "null")
		//queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmMasterID\":{\"$ne\":\"%s\"}},{\"customerNo\":{\"$ne\":\"%s\"}}]}}", "null", "null")
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read the list of getAllEBSOrders : %s\n", err)
			return shim.Error(err.Error())
			//return nil, err
		}
	
		//fmt.Printf("list of all getAllEBSOrders : %v\n", queryResults)
		return shim.Success(queryResults)
	}

//getAllLSPOrders
func (t *kpnLogistics) getAllLSPOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get All LSP Orders")
		//crmMasterID
		//fetch data from couch db starts here
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"lspMasterID\":{\"$ne\":\"%s\"}},{\"lspOrderStatus\":{\"$ne\":\"%s\"}}]}}", "null", "null")
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read the list of getAllLSPOrders : %s\n", err)
			return shim.Error(err.Error())
			//return nil, err
		}
	
		//fmt.Printf("list of all getAllEBSOrders : %v\n", queryResults)
		return shim.Success(queryResults)
	}

//getAllTransporterOrders
func (t *kpnLogistics) getAllTransporterOrders(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering getAllTransporterOrders")
		var stat ="Shipment In Transit"
		//fetch data from couch db starts here
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"transporterID\":{\"$ne\":\"%s\"}},{\"lspstatus\":{\"$eq\":\"%s\"}}]}}", "null", stat)
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read the list of getAllTransporterOrders : %s\n", err)
			return shim.Error(err.Error())
			//return nil, err
		}
	
		fmt.Printf("list of all getAllTransporterOrders : %v\n", queryResults)
		return shim.Success(queryResults)
	}


//getEbsMasterID
func (t *kpnLogistics) getEbsMasterID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get All Sales Orders")
		//crmMasterID
		//fetch data from couch db starts here
		var crmOrder = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"oracleAccountNumber\":{\"$ne\":\"%s\"}}]}}", crmOrder, "null")
		//queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmMasterID\":{\"$ne\":\"%s\"}},{\"customerNo\":{\"$ne\":\"%s\"}}]}}", "null", "null")
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read the list of getAllEBSOrders : %s\n", err)
			return shim.Error(err.Error())
			//return nil, err
		}
		return shim.Success(queryResults)
	}

//getEbsMaster
func (t *kpnLogistics) getEbsMaster(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get All getEbsMaster")
		//crmMasterID
		//fetch data from couch db starts here
		var ebsmast = args[0]
		//queryString := fmt.Sprintf("{\"selector\":{\"ebsMasterID\":{\"$eq\":\"%s\"}}}", ebsmast)
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"ebsMasterID\":{\"$eq\":\"%s\"}},{\"orderSalesOrderNoEBS\":{\"$ne\":\"%s\"}}]}}", ebsmast, "null")
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read the list of getEbsMaster : %s\n", err)
			return shim.Error(err.Error())
			//return nil, err
		}
		return shim.Success(queryResults)
	}

//getCRMMasterID
func (t *kpnLogistics) getCRMMasterID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering getCRMMasterID")
		//fetch data from couch db starts here
		var crmOrder = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"crmCreatedBy\":{\"$ne\":\"%s\"}}]}}", crmOrder, "null")
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read the list of getCRMMasterID : %s\n", err)
			return shim.Error(err.Error())
		}

		return shim.Success(queryResults)
	}

//getLSPDetailsStatusByProdID
func (t *kpnLogistics) getLSPDetailsStatusByProdID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get LSP Details Status By ProdID")
		//crmMasterID
		//fetch data from couch db starts here
		var proID = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"productID\":{\"$eq\":\"%s\"}},{\"lspProductStatus\":{\"$ne\":\"%s\"}}]}}", proID, "null")
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read the list of getLSPDetailsStatusByProdID : %s\n", err)
			return shim.Error(err.Error())
			//return nil, err
		}
	
		fmt.Printf("list of all getLSPDetailsStatusByProdID : %v\n", queryResults)
		return shim.Success(queryResults)
	}

//get all Sales order/ Details
func (t *kpnLogistics) getAllSalesOrderhd(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get All Sales Orders")
	
		//fetch data from couch db starts here
		queryString := fmt.Sprintf("{\"selector\":{\"$or\":[{\"crmOrderNumber\":{\"$ne\":\"%s\"}},{\"customerNo\":{\"$ne\":\"%s\"}}]}}", "null", "null")
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read the list of Purchase orders : %s\n", err)
			return shim.Error(err.Error())
			//return nil, err
		}
	
		fmt.Printf("list of all purchase order details : %v\n", queryResults)
		return shim.Success(queryResults)
	}



	//getstatusconfig
	func (t *kpnLogistics) getstatusconfig(stub shim.ChaincodeStubInterface, args []string) pb.Response {
		
			var err error
			fmt.Println("Entering  getstatusconfig")
		
			//fetch data from couch db starts here
			
			queryString := fmt.Sprintf("{\"selector\":{\"status\":{\"$ne\":\"%s\"}}}", "null")
			queryResults, err := getQueryResultForQueryString(stub, queryString)
			//fetch data from couch db ends here
			if err != nil {
				fmt.Printf("Unable to read Status config: %s\n", err)
				return shim.Error(err.Error())
			}
		
			fmt.Printf("list of Status config: %v\n", queryResults)
			return shim.Success(queryResults)
		}

//getstatusconfigByStatusID
func (t *kpnLogistics) getstatusconfigByStatusID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
		
			var err error
			fmt.Println("Entering  getstatusconfigByStatusID")
		
			//fetch data from couch db starts here
			var statId = args[0]
			queryString := fmt.Sprintf("{\"selector\":{\"statusid\":{\"$eq\":\"%s\"}}}", statId)
			queryResults, err := getQueryResultForQueryString(stub, queryString)
			//fetch data from couch db ends here
			if err != nil {
				fmt.Printf("Unable to read Status config for statusid: %s\n", err)
				return shim.Error(err.Error())
			}
		
			fmt.Printf("list of Status config by statusid: %v\n", queryResults)
			return shim.Success(queryResults)
		}

	//getOrderDetailsByCRMOrderNo
func (t *kpnLogistics) getOrderDetailsByCRMOrderNo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get OrderDetails By CRMOrder No")
	
		//fetch data from couch db starts here
		var crrmOrderNo = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"crmOrderNumber\":{\"$eq\":\"%s\"}}}", crrmOrderNo)
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read Sales orders details for the given CRM Order No: %s\n", err)
			return shim.Error(err.Error())
		}
	
		//fmt.Printf("list of order details for given CRMOrderNo : %v\n", queryResults)
		return shim.Success(queryResults)
	} 
//Invoke CreateStatusConfig
func (t *kpnLogistics) CreateStatusConfig(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var statusconfig Statusconfig
		var err error
	
		if len(args) < 1 {
			fmt.Println("At least 1 value required for Statusconfig function")
			return shim.Error(err.Error())
		}
	
		err = json.Unmarshal([]byte(args[0]), &statusconfig)
		if err != nil {
			fmt.Printf("Unable to unmarshal data in Statusconfig : %s\n", err)
			return shim.Error(err.Error())
		}
	
		// Start - Put into Couch DB
		JSONBytes, err := json.Marshal(statusconfig)
		err = stub.PutState(statusconfig.StatusId, JSONBytes)
		// End
	
		if err != nil {
			fmt.Printf("\nUnable to make transaction for statusconfig : %v ", err)
			return shim.Error(err.Error())
		}
	
		return shim.Success(nil)
	}

//updateStatusConfig
func (t *kpnLogistics) updateStatusConfig(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcStatusconfig Statusconfig
	var obj_uiStatusconfig Statusconfig
	var err error

	fmt.Println("Entering updateStatusConfig")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiStatusconfig)
	if err != nil {
		fmt.Printf("Unable to marshal  updateStatusConfig input Statusconfig : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n CRM Order Number value is : ", obj_uiStatusconfig.StatusId)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiStatusconfig.StatusId)
	err = json.Unmarshal(bytesread, &obj_bcStatusconfig)
	// code to get data from blockchain using dynamic key ends here

	if (obj_uiStatusconfig.FieldValue=="thresholdvalue"){

	obj_bcStatusconfig.System = obj_uiStatusconfig.System
	obj_bcStatusconfig.Status = obj_uiStatusconfig.Status
	obj_bcStatusconfig.Action = obj_uiStatusconfig.Action
	obj_bcStatusconfig.Thresholdtime = obj_uiStatusconfig.Thresholdtime
	obj_bcStatusconfig.FieldValue = obj_uiStatusconfig.FieldValue
	obj_bcStatusconfig.ThresholdUnit = obj_uiStatusconfig.ThresholdUnit
	obj_bcStatusconfig.Equality = obj_uiStatusconfig.Equality

	}else{
	obj_bcStatusconfig.System = obj_uiStatusconfig.System
	obj_bcStatusconfig.Status = obj_uiStatusconfig.Status
	obj_bcStatusconfig.Action = obj_uiStatusconfig.Action
	obj_bcStatusconfig.ConfigShippingTime = obj_uiStatusconfig.ConfigShippingTime
	obj_bcStatusconfig.ConfigShippingMethod = obj_uiStatusconfig.ConfigShippingMethod
	obj_bcStatusconfig.FieldValue = obj_uiStatusconfig.FieldValue
	obj_bcStatusconfig.Equality = obj_uiStatusconfig.Equality

	}
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcStatusconfig)
	err = stub.PutState(obj_uiStatusconfig.StatusId, transJSONasBytes)
	// Data insertion for Couch DB ends here
	fmt.Println("updateStatusConfig Successfully updated in Statusconfig struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
		
	return shim.Success(nil)
}

//deleteStatusConfig
func (t *kpnLogistics) deleteStatusConfig(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var obj_bcStatusconfig Statusconfig
	var obj_uiStatusconfig Statusconfig
	var err error

	fmt.Println("Entering deleteStatusConfig")

	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &obj_uiStatusconfig)
	if err != nil {
		fmt.Printf("Unable to marshal  deleteStatusConfig input Statusconfig : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n CRM Order Number value is : ", obj_uiStatusconfig.StatusId)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(obj_uiStatusconfig.StatusId)
	err = json.Unmarshal(bytesread, &obj_bcStatusconfig)
	// code to get data from blockchain using dynamic key ends here

	obj_bcStatusconfig.StatusField = obj_uiStatusconfig.StatusField
	
	// Data insertion for Couch DB starts here
	transJSONasBytes, err := json.Marshal(obj_bcStatusconfig)
	err = stub.PutState(obj_uiStatusconfig.StatusId, transJSONasBytes)
	// Data insertion for Couch DB ends here
	fmt.Println("deleteStatusConfig Successfully updated in Statusconfig struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
	}
		
	return shim.Success(nil)
}

//getOrdersByCRMDashboardStatus
func (t *kpnLogistics) getOrdersByCRMDashboardStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get OrderDetails By CRMDAshboard Status")
	
		//fetch data from couch db starts here
		var status = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"crmDashboardStatus\":{\"$eq\":\"%s\"}}}", status)
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read Sales orders details for the given CRM DAshsboard Status: %s\n", err)
			return shim.Error(err.Error())
		}
		
		fmt.Printf("list of order details for given CRMDashboard Status: %v\n", queryResults)
		return shim.Success(queryResults)
	}
//getOrdersByEBSDashboardStatus
func (t *kpnLogistics) getOrdersByEBSDashboardStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering get OrderDetails By EBSDAshboard Status")
	
		//fetch data from couch db starts here
		var status = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"oracleDashboardStatus\":{\"$eq\":\"%s\"}}}", status)

		//queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"oracleDashboardStatus\":{\"$ne\":\"%s\"}},{\"customerNo\":{\"$ne\":\"%s\"}}]}}", "null", "null")
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to read Sales orders details for the given EBS DAshsboard Status: %s\n", err)
			return shim.Error(err.Error())
		}
	
		fmt.Printf("list of order details for given EBSDashboard Status: %v\n", queryResults)
		return shim.Success(queryResults)
	}

//getOrdersByLSPDashboardStatus
func (t *kpnLogistics) getOrdersByLSPDashboardStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering getOrders By LSP Dashboard Status")
	
		//fetch data from couch db starts here
		var status = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"lspDashboardStatus\":{\"$eq\":\"%s\"}}}", status)
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to get Orders By LSP DashboardStatus details for the given LSPDAshsboard Status: %s\n", err)
			return shim.Error(err.Error())
		}
	
		fmt.Printf("list of order details for given EBSDashboard Status: %v\n", queryResults)
		return shim.Success(queryResults)
	}


//getOrdersByTransportDashboardStatus
func (t *kpnLogistics) getOrdersByTransportDashboardStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering getOrders By transport Dashboard Status")
	
		//fetch data from couch db starts here
		var status = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"transporterDashboardStatus\":{\"$eq\":\"%s\"}}}", status)
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to get Orders By transport DashboardStatus details for the given transport Status: %s\n", err)
			return shim.Error(err.Error())
		}
	
		fmt.Printf("list of order details for given transportdashboard Status: %v\n", queryResults)
		return shim.Success(queryResults)
	}
	func (t *kpnLogistics) LogCRMWarnings(stub shim.ChaincodeStubInterface, args []string) pb.Response {
		
			var logwarnings_Obj LogWarningCRMdetails
			var err error
		
			if len(args) < 1 {
				fmt.Println("At least 1 value required for LogCRMWarnings function")
				return shim.Error(err.Error())
			}
		
			err = json.Unmarshal([]byte(args[0]), &logwarnings_Obj)
			if err != nil {
				fmt.Printf("Unable to unmarshal data in CreateSubscriberRecord : %s\n", err)
				return shim.Error(err.Error())
			}
		
			// Start - Put into Couch DB
			JSONBytes, err := json.Marshal(logwarnings_Obj)
			err = stub.PutState(logwarnings_Obj.LogId, JSONBytes)
			// End
		
			if err != nil {
				fmt.Printf("\nUnable to make transaction for LogCRMWarnings : %v ", err)
				return shim.Error(err.Error())
			}
		
			return shim.Success(nil)
		}

func (t *kpnLogistics) checkcrmwarning(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("Entering checkcrmwarning")
	var orderno = args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"ordernumber\":{\"$eq\":\"%s\"}},{\"logid\":{\"$ne\":\"%s\"}}]}}", orderno, "null")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	var arrcrmwarnings []LogWarningCRMdetails
	err = json.Unmarshal([]byte(queryResults), &arrcrmwarnings)
	if err != nil {
		panic(err)
	}
	var arrlength = len(arrcrmwarnings)
	jsonbytesarrlength, err := json.Marshal(arrlength)
	
			if err != nil {
				fmt.Printf("\nUnable marshal arrlength : %v ", err)
				return shim.Error(err.Error())
			}
	return shim.Success(jsonbytesarrlength)
}		
		
//getWarnings
func (t *kpnLogistics) getWarnings(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(" ----------------------- Entering checkcrmwarning-------------------------")
		var err error
		//var obj_bcSalesOrderMaster SalesOrderMaster
		var crmorders []KeyRecordCRM

		fmt.Println("Entering get getWarnings")
		currentTime := time.Now()
		fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
		//fetch data from couch db starts here
		//var status = "Order Created"
		//queryString := fmt.Sprintf("{\"selector\":{\"status\":{\"$eq\":\"%s\"}}}", status)
		//queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"status\":{\"$eq\":\"%s\"}},{\"statusField\":{\"$eq\":\"%s\"}}]}}", "Order Created", "Active")
		//queryString := fmt.Sprintf("{\"selector\":{\"statusid\":{\"$gt\":null}}}")
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"statusid\":{\"$gt\":null}},{\"statusField\":{\"$eq\":\"%s\"}}]}}", "Active")
		queryResults2, err := getQueryResultForQueryString(stub, queryString)

		var arr_config []KeyRecordSC
		err = json.Unmarshal([]byte(queryResults2), &arr_config)
		if err != nil {
			fmt.Println(" -------Panic 1----------------")
			panic(err)
		}
		//fmt.println("arr_config",arr_config[0].Record);
		var errorcount int = 0 
		var errorcount1 int =0
		var errorcount2 int = 0 
		var errorcount3 int =0
		var totallineItems int = 0
		var totallineItems1 int =0
		var totallineItems2 int = 0
		var totallineItems3 int =0
		var totallineItems4 int =0
		var errorcount4 int=0
		var arrerrorcount [2]int
			for _,statusconfigdetails := range arr_config{
				if (statusconfigdetails.Record.System=="CRM" && statusconfigdetails.Record.Status=="Created" && statusconfigdetails.Record.Equality=="equal"){
					queryString2 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"orderStatus\":{\"$eq\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "Created", "null")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults2, err := getQueryResultForQueryString(stub, queryString2)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults2), &crmorders)
						if err != nil {
							fmt.Println(" -------Panic 2----------------")
							panic(err)
						}
					if (statusconfigdetails.Record.FieldValue=="thresholdvalue"){
						//get crm orders of order created
						var thresholdvaluetime=statusconfigdetails.Record.Thresholdtime 
						//fmt.Println("thresholdtime -----",arr_config[0].Record.Thresholdtime)
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample2 := gettimedifference(thresholdvaluetime,pasttimestr)
						fmt.Println("query results for sample2 -----",sample2)
						fmt.Println("thresholdtime inside ebs -----",statusconfigdetails.Record.Thresholdtime)
						if (sample2>0){errorcount++
						// get LineItems here
						 //totallineItems= getLineItems(stub,orderNo)
					queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults3, err := getQueryResultForQueryString(stub, queryString3)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						fmt.Println(" -------Panic 3----------------")
						panic(err)
					}
					//CRMDashboardStatus update
					//var CRMDashboardStatus="In-Progress_Warning"
					/*transJSONasBytesLSPDetails, err := json.Marshal(obj_bcLSPDetails)
					err = stub.PutState(obj_uiLSPDetails.LSPDetailsID, transJSONasBytesLSPDetails)
*/					fmt.Println("\n ***********************************************************************")
					

				/*	for _, salesorderrecords := range soMaster {
						fmt.Println("\n salesorderrecords.Record before : ",salesorderrecords)
						fmt.Println("\n salesorderrecords.Record.CRMDashboardStatus before",salesorderrecords.Record.CRMDashboardStatus)
						fmt.Println("\n salesorderrecords.Record.CRMOrderNumber ",salesorderrecords.Record.CRMOrderNumber)
						salesorderrecords.Record.CRMDashboardStatus = "In-Progress_Warning"
					//	fmt.Println("\n salesorderrecords.Record.CRMDashboardStatus",salesorderrecords)
						fmt.Println("\n salesorderrecords.Record.CRMDashboardStatus after",salesorderrecords.Record.CRMDashboardStatus)
	
		var bytesread []byte
		bytesread, err = stub.GetState(salesorderrecords.Record.CRMOrderNumber)
		err = json.Unmarshal(bytesread, &obj_bcSalesOrderMaster)
		fmt.Println("\n obj_bcSalesOrderMaster before update : ", obj_bcSalesOrderMaster)
		
			obj_bcSalesOrderMaster.CRMDashboardStatus = salesorderrecords.Record.CRMDashboardStatus

					transJSONasBytesSodetails, err := json.Marshal(obj_bcSalesOrderMaster)
					err = stub.PutState(obj_bcSalesOrderMaster.CRMOrderNumber, transJSONasBytesSodetails)

					fmt.Println("\n obj_bcSalesOrderMaster full struct before udate  :  ",obj_bcSalesOrderMaster)
					
					//fmt.Println("\n obj_bcSalesOrderMaster full struct after update :  ",obj_bcSalesOrderMaster)
					fmt.Println("\n ***********************************************************************")

						if err != nil {
							fmt.Printf("\nUnable to make transevent inputs code 002 : %v ", err)
							return shim.Error(err.Error())
							}
					}*/
					var Dashboardinprogressstatus =soMaster[0].Record.CRMDashboardStatus
					if (Dashboardinprogressstatus=="In-Progress"){

						soMaster[0].Record.CRMDashboardStatus="warning"
						transJSONasBytesSodetails, err := json.Marshal(soMaster)
						fmt.Println("\n crm order number  :  ",soMaster[0].Record.CRMOrderNumber)
						err = stub.PutState(soMaster[0].Record.CRMOrderNumber, transJSONasBytesSodetails)
						if err != nil {
							fmt.Printf("\nUnable to make transevent inputs code 002 : %v ", err)
							return shim.Error(err.Error())
							}
					}
						bytesread, err := stub.GetState(soMaster[0].Record.CRMOrderNumber)
							err = json.Unmarshal(bytesread, &soMaster)
									fmt.Println("\n obj_bcSalesOrderMaster full struct after udate  test:  ",soMaster)
					
					// Data insertion for Couch DB starts here
										// Data insertion for Couch DB ends here


					//CRMDashboardStatus update
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						fmt.Println(" -------Panic 4----------------")
						panic(err)
					}
					totallineItems += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems)
					fmt.Printf("satisfied")
							}
							}				
						}	
								//shippingMethod starts here

						if (statusconfigdetails.Record.FieldValue=="shippingmethod" ){
						//get crm orders of order created
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							//get shipping method starts herere
							
							pasttimestr := order.Record.Timestamp
							sample2 := getdifference(pasttimestr)
						fmt.Println("query results for sample2 -----",sample2)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						if (sample2==1){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
								queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
								queryResults3, err := getQueryResultForQueryString(stub, queryString3)
								var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults3), &soMaster)
								if err != nil {
									fmt.Println(" -------Panic 5----------------")
									panic(err)
								}
								if (soMaster[0].Record.ShippingMethod=="Delivery Next Day"){
								errorcount1++	
								 //totallineItems1= getLineItems(stub,orderNo)				
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						fmt.Println(" -------Panic 6----------------")
						panic(err)
					}
					totallineItems1 += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems)
					fmt.Printf("satisfied")
								}
											}
									}else if (sample2>1){
								// if order date is greater than one day here
							 	//configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							//if (configtimecheckvalue>0){
								queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
								queryResults3, err := getQueryResultForQueryString(stub, queryString3)
								var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults3), &soMaster)
								if err != nil {
									fmt.Println(" -------Panic 5----------------")
									panic(err)
								}
								if (soMaster[0].Record.ShippingMethod=="Delivery Next Day"){
								errorcount1++	
								 //totallineItems1= getLineItems(stub,orderNo)				
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						fmt.Println(" -------Panic 6----------------")
						panic(err)
					}
					totallineItems1 += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems)
					fmt.Printf("satisfied")
								}
										//	}
									}

								// if order date is greater than one day ends here					
							}
						}
				}
						//inprogress orders starts herere
						if (statusconfigdetails.Record.System=="CRM" && statusconfigdetails.Record.Status=="In Progress" && statusconfigdetails.Record.Equality=="equal"){
					queryString4 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"orderStatus\":{\"$eq\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "In-Progress", "null")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults4, err := getQueryResultForQueryString(stub, queryString4)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults4), &crmorders)
						if err != nil {
							panic(err)
						}
					if (statusconfigdetails.Record.FieldValue=="thresholdvalue"){
						//get crm orders of order created
						var thresholdvaluetime=statusconfigdetails.Record.Thresholdtime 
						//fmt.Println("thresholdtime -----",arr_config[0].Record.Thresholdtime)
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample3 := gettimedifference(thresholdvaluetime,pasttimestr)
						fmt.Println("query results for sample2 -----",sample3)
						fmt.Println("thresholdtime inside ebs -----",statusconfigdetails.Record.Thresholdtime)
						if (sample3>0){errorcount2++
						// get LineItems here
						// totallineItems2= getLineItems(stub,orderNo)
							queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults3, err := getQueryResultForQueryString(stub, queryString3)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						panic(err)
					}
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems2 += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems)
					fmt.Printf("satisfied")
							}
							}				
						}	
								//shippingMethod starts here

						if (statusconfigdetails.Record.FieldValue=="shippingmethod"){
						//get crm orders of order created
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample4 := getdifference(pasttimestr)
						fmt.Println("query results for sample2 -----",sample4)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						//query to get the shipping method is Delivery Next Day
						if (sample4==1){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
								queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
								queryResults3, err := getQueryResultForQueryString(stub, queryString3)
								var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults3), &soMaster)
								if err != nil {
									panic(err)
								}
								if(soMaster[0].Record.ShippingMethod=="Delivery Next Day"){
								errorcount3++	
								 //totallineItems3= getLineItems(stub,orderNo)			
								
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems3 += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems)
					fmt.Printf("satisfied")
								}
											}
									}else if (sample4>1){
							//configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							//if (configtimecheckvalue>0){
								queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
								queryResults3, err := getQueryResultForQueryString(stub, queryString3)
								var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults3), &soMaster)
								if err != nil {
									panic(err)
								}
								if(soMaster[0].Record.ShippingMethod=="Delivery Next Day"){
								errorcount3++	
								 //totallineItems3= getLineItems(stub,orderNo)			
								
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems3 += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems)
					fmt.Printf("satisfied")
								}
								//			}
									}


													
							}
						}
						//inprogress ends here
				}

				if (statusconfigdetails.Record.System=="ALL" && statusconfigdetails.Record.Status=="Delivered" && statusconfigdetails.Record.Equality=="notequal"){
					queryString2 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"orderStatus\":{\"$ne\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "null", "null")
							//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults2, err := getQueryResultForQueryString(stub, queryString2)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults2), &crmorders)
						if err != nil {
							fmt.Println(" -------Panic 7----------------")
							panic(err)
						}

						if (statusconfigdetails.Record.FieldValue=="shippingmethod"){
						//get crm orders of order created
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderstatus = order.Record.CRMOrderStatus
							if (orderstatus!="Delivered"){
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample7 := getdifference(pasttimestr)
						fmt.Println("query results for sample2 -----",sample7)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						//query to get the shipping method is Delivery Next Day
						if (sample7==1){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
								queryString10 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
								queryResults10, err := getQueryResultForQueryString(stub, queryString10)
								var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults10), &soMaster)
								if err != nil {
									panic(err)
								}
								if(soMaster[0].Record.ShippingMethod=="Delivery Next Day"){
								errorcount4++	
								 //totallineItems3= getLineItems(stub,orderNo)			
								
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems4 += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems4)
					fmt.Printf("satisfied")
								}
											}
									}else if (sample7>1){
							//configtimecheckvalue:= configtimecheck(ConfigShippingTime)
						///	if (configtimecheckvalue>0){
								queryString10 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
								queryResults10, err := getQueryResultForQueryString(stub, queryString10)
								var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults10), &soMaster)
								if err != nil {
									panic(err)
								}
								if(soMaster[0].Record.ShippingMethod=="Delivery Next Day"){
								errorcount4++	
								 //totallineItems3= getLineItems(stub,orderNo)			
								
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems4 += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems4)
					fmt.Printf("satisfied")
								}
										//	}
									}
													
							}
							}
						}
						
				}

			}
		//arr_errorcount[0] = errorcount
		arrerrorcount[0]= errorcount + errorcount1+errorcount3+errorcount2+errorcount4
		arrerrorcount[1]= totallineItems+totallineItems1+totallineItems3+totallineItems2+totallineItems4
		
		//jsonbytes_errorcount, err := json.Marshal(arr_errorcount)
		jsonbytes_errorcount, err := json.Marshal(arrerrorcount)

		if err != nil {
			fmt.Printf("\nUnable marshal arr_errorcount : %v ", err)
			return shim.Error(err.Error())
		}
		//fetch data from couch db ends here	
		if err != nil {
			fmt.Printf("Unable to read Sales orders details for the given CRM Order No: %s\n", err)
			return shim.Error(err.Error())
		}
	
		//fmt.Printf("list of order details for given CRMOrderNo : %v\n", queryResults)
		return shim.Success(jsonbytes_errorcount)
	} 
func getdifference(x string) int{
	
	layout := "2006-01-02 15:04"
    var ordertimestamp =x
    timeFormat:="2006/01/02"
    currentTime := time.Now()
    t, err := time.Parse(layout, ordertimestamp)
    if err != nil {
        fmt.Println(err)
        
    }
	var mydate=t.Format("2006/01/02")
	FF_date, _ := time.Parse(timeFormat, mydate)
	//fmt.Println(t)
	//fmt.Println(mydate)
	//fmt.Println("order date",FF_date)
	//fmt.Println(reflect.TypeOf(FF_date))
	currentdate := currentTime.Format("2006/01/02")
	current_date, _ := time.Parse(timeFormat, currentdate)
	fmt.Println("currentdate ",currentdate)
	diff := current_date.Sub(FF_date)
	var days=int(diff.Hours()/24)
	fmt.Println("diff",diff.Hours()/24)
	
	return days
	}
func configtimecheck(configtime string) int{
	 configshippingcount:=0
	var timeFormat = "15:04"
	//var configtime= "10:25"
	givenconfigtime, err := time.Parse(timeFormat, configtime)
	if err != nil {
		panic(err)
		
		}
		fmt.Println(reflect.TypeOf(givenconfigtime))
	utc := time.Now().UTC().Format("15:04")
	fmt.Println("utc",reflect.TypeOf(utc ))
	//var ArriveTime_string = "23:47"
	currentutctime, err := time.Parse(timeFormat, utc)
	fmt.Println("utc arrive",reflect.TypeOf(currentutctime))
	duration := currentutctime.Sub(givenconfigtime)
	fmt.Printf("\nduration : %s", duration)
	calcDuration := duration.Minutes()
	//calcDuration_string := fmt.Sprintf("%.2f", calcDuration)
	//fmt.Printf("\ncalculated duratifuncon : %s ", calcDuration_string)
	if (calcDuration >0){
		configshippingcount=configshippingcount+1
	}else
	{ configshippingcount=configshippingcount+0}
	return configshippingcount
}
func getLineItems(stub shim.ChaincodeStubInterface,orderNo string) int{
totallineItems:= 0
					queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults3, err := getQueryResultForQueryString(stub, queryString3)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						panic(err)
					}
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems)
					fmt.Printf("satisfied")

return totallineItems

}
func (t *kpnLogistics) getWarningsDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		var arrCRMDetails []WarningCRMdetails
		var crmorders []KeyRecordCRM

		fmt.Println("Entering get getWarningsDetails")
		currentTime := time.Now()
		fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"statusid\":{\"$gt\":null}},{\"statusField\":{\"$eq\":\"%s\"}}]}}", "Active")
		queryResults, err := getQueryResultForQueryString(stub, queryString)

		var arr_config []KeyRecordSC
		err = json.Unmarshal([]byte(queryResults), &arr_config)
		if err != nil {
			panic(err)
		}
		//fmt.println("arr_config",arr_config[0].Record);
			for _,statusconfigdetails := range arr_config{
				if (statusconfigdetails.Record.System=="CRM" && statusconfigdetails.Record.Status=="Created" && statusconfigdetails.Record.Equality=="equal"){
					
					queryString2 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"orderStatus\":{\"$eq\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "Created", "null")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults2, err := getQueryResultForQueryString(stub, queryString2)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults2), &crmorders)
						if err != nil {
							panic(err)
						}
					if (statusconfigdetails.Record.FieldValue=="thresholdvalue"){
						//get crm orders of order created
						var thresholdvaluetime=statusconfigdetails.Record.Thresholdtime 
						var statusreason = statusconfigdetails.Record.StatusReason
						//fmt.Println("thresholdtime -----",arr_config[0].Record.Thresholdtime)
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							//var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample2 := gettimedifference(thresholdvaluetime,pasttimestr)
						fmt.Println("query results for sample2 -----",sample2)
						fmt.Println("thresholdtime inside ebs -----",statusconfigdetails.Record.Thresholdtime)
						if (sample2>0){
							arrCRMDetails = append(arrCRMDetails, WarningCRMdetails{
					Orderno: fmt.Sprintf("%v", order.Record.CRMOrderNumber),
					Customerno: fmt.Sprintf("%v", order.Record.CustomerNo),
					OrderStatus: fmt.Sprintf("%v", order.Record.CRMOrderStatus),
					Price: fmt.Sprintf("%v", order.Record.OrderTotalPrice),
					Reason: fmt.Sprintf("%v", statusreason),
					Eventflag: fmt.Sprintf("%v", "CRM_thr_created"),
										})
							}
							}				
						}	
								//shippingMethod starts here

						if (statusconfigdetails.Record.FieldValue=="shippingmethod"){
						//get crm orders of order created
						var statusreason = statusconfigdetails.Record.StatusReason
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample2 := getdifference(pasttimestr)
						//	var salesorder SalesOrderMaster
							//query to get the shipping method if(soMaster[0].Record.ShippingMethod=="Delivery Next Day"){
						queryString8 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"shippingMethod\":{\"$eq\":\"%s\"}}]}}",orderNo , "Delivery Next Day")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults8, err := getQueryResultForQueryString(stub, queryString8)
						//fmt.Println("query results for crm orders -----",queryResults2)
						var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults8), &soMaster)
								if err != nil {
									panic(err)
								}
								for _, salesorderrecords := range soMaster {
								if (salesorderrecords.Record.ShippingMethod=="Delivery Next Day"){
								
						//fmt.Println("salesorder ",salesorder)
						fmt.Println("salesorder shiping method -----",salesorderrecords.Record.ShippingMethod)
					//	if(salesorder.shippingmethod=="Delivery Next Day"){
						fmt.Println("query results for sample2 -----",sample2)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						if (sample2>0){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
							arrCRMDetails = append(arrCRMDetails, WarningCRMdetails{
					Orderno: fmt.Sprintf("%v", order.Record.CRMOrderNumber),
					Customerno: fmt.Sprintf("%v", order.Record.CustomerNo),
					OrderStatus: fmt.Sprintf("%v", order.Record.CRMOrderStatus),
					Price: fmt.Sprintf("%v", order.Record.OrderTotalPrice),
					Reason: fmt.Sprintf("%v", statusreason),
					Eventflag: fmt.Sprintf("%v", "CRM_ship_created"),
				})
							}
											}
						}
								}
									}
													
							}
						}
										//inprogress orders starts herere
						if (statusconfigdetails.Record.System=="CRM" && statusconfigdetails.Record.Status=="In Progress" && statusconfigdetails.Record.Equality=="equal"){
					queryString4 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"orderStatus\":{\"$eq\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "In-Progress", "null")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults4, err := getQueryResultForQueryString(stub, queryString4)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults4), &crmorders)
						if err != nil {
							panic(err)
						}
					if (statusconfigdetails.Record.FieldValue=="thresholdvalue"){
						//get crm orders of order created
						var thresholdvaluetime=statusconfigdetails.Record.Thresholdtime 
						var statusreason = statusconfigdetails.Record.StatusReason
						//fmt.Println("thresholdtime -----",arr_config[0].Record.Thresholdtime)
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							//var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample3 := gettimedifference(thresholdvaluetime,pasttimestr)
						fmt.Println("query results for sample2 -----",sample3)
						fmt.Println("thresholdtime inside ebs -----",statusconfigdetails.Record.Thresholdtime)
						if (sample3>0){arrCRMDetails = append(arrCRMDetails, WarningCRMdetails{
					Orderno: fmt.Sprintf("%v", order.Record.CRMOrderNumber),
					Customerno: fmt.Sprintf("%v", order.Record.CustomerNo),
					OrderStatus: fmt.Sprintf("%v", order.Record.CRMOrderStatus),
					Price: fmt.Sprintf("%v", order.Record.OrderTotalPrice),
					Reason: fmt.Sprintf("%v", statusreason),
					Eventflag: fmt.Sprintf("%v", "CRM_thr_inprogress"),
				})
							}
							}				
						}	
								//shippingMethod starts here

						if (statusconfigdetails.Record.FieldValue=="shippingmethod"){
						//get crm orders of order created
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
						var statusreason = statusconfigdetails.Record.StatusReason
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample4 := getdifference(pasttimestr)
						fmt.Println("query results for sample2 -----",sample4)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						queryString9 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"shippingMethod\":{\"$eq\":\"%s\"}}]}}",orderNo , "Delivery Next Day")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults9, err := getQueryResultForQueryString(stub, queryString9)
						//fmt.Println("query results for crm orders -----",queryResults2)
						var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults9), &soMaster)
								if err != nil {
									panic(err)
								}
								for _, salesorderrecords := range soMaster {
								if (salesorderrecords.Record.ShippingMethod=="Delivery Next Day"){
								
						if (sample4>0){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
								arrCRMDetails = append(arrCRMDetails, WarningCRMdetails{
					Orderno: fmt.Sprintf("%v", order.Record.CRMOrderNumber),
					Customerno: fmt.Sprintf("%v", order.Record.CustomerNo),
					OrderStatus: fmt.Sprintf("%v", order.Record.CRMOrderStatus),
					Price: fmt.Sprintf("%v", order.Record.OrderTotalPrice),
					Reason: fmt.Sprintf("%v", statusreason),
					Eventflag: fmt.Sprintf("%v", "CRM_ship_inprogress"),
				})
											}
									}
								}
								}					
							}
						}
						//inprogress ends here
				}
//All delivered not equal starts here
					
				if (statusconfigdetails.Record.System=="ALL" && statusconfigdetails.Record.Status=="Delivered" && statusconfigdetails.Record.Equality=="notequal"){
							queryString2 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"orderStatus\":{\"$ne\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "null", "null")
					//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults2, err := getQueryResultForQueryString(stub, queryString2)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults2), &crmorders)
						if err != nil {
							fmt.Println(" -------Panic 7----------------")
							panic(err)
						}

						if (statusconfigdetails.Record.FieldValue=="shippingmethod"){
						//get crm orders of order created
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
						var statusreason = statusconfigdetails.Record.StatusReason
						
							for _, order := range crmorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderstatus = order.Record.CRMOrderStatus
							if (orderstatus!="Delivered"){
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample7 := getdifference(pasttimestr)
						fmt.Println("query results for sample2 -----",sample7)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						//query to get the shipping method is Delivery Next Day
						if (sample7>0){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
								queryString10 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
								queryResults10, err := getQueryResultForQueryString(stub, queryString10)
								var soMaster []KeyRecordSO
								err = json.Unmarshal([]byte(queryResults10), &soMaster)
								if err != nil {
									panic(err)
								}
								if(soMaster[0].Record.ShippingMethod=="Delivery Next Day"){
								//append
									arrCRMDetails = append(arrCRMDetails, WarningCRMdetails{
					Orderno: fmt.Sprintf("%v", order.Record.CRMOrderNumber),
					Customerno: fmt.Sprintf("%v", order.Record.CustomerNo),
					OrderStatus: fmt.Sprintf("%v", order.Record.CRMOrderStatus),
					Price: fmt.Sprintf("%v", order.Record.OrderTotalPrice),
					Reason: fmt.Sprintf("%v", statusreason),
					Eventflag: fmt.Sprintf("%v", "ALL_ship_inprogress"),
				})
								//append
								}
											}
									}
													
							}
							}
						}
						
				}

//All delivered not eaqual ends herere 

			}
		//arr_errorcount[0] = errorcount
		jsonbytes_crmorderdetails, err := json.Marshal(arrCRMDetails)
	fmt.Printf(string(jsonbytes_crmorderdetails))

		
		//fetch data from couch db ends here	
		if err != nil {
			fmt.Printf("Unable to read Sales orders details for the given CRM Order No: %s\n", err)
			return shim.Error(err.Error())
		}
	
		fmt.Printf("list of order details for given CRMOrderNo : %v\n", queryResults)
		return shim.Success(jsonbytes_crmorderdetails)
	}
//EBS created warnings starts here
/*func (t *kpnLogistics) getWarningsEBS(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		var ebsorders []KeyRecordEBS

		fmt.Println("Entering get getWarnings for ebs")
		currentTime := time.Now()
		fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
		//fetch data from couch db starts here
		//var status = "Order Created"
		//queryString := fmt.Sprintf("{\"selector\":{\"status\":{\"$eq\":\"%s\"}}}", status)
		//queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"status\":{\"$eq\":\"%s\"}},{\"statusField\":{\"$eq\":\"%s\"}}]}}", "Order Created", "Active")
		//queryString := fmt.Sprintf("{\"selector\":{\"statusid\":{\"$gt\":null}}}")
		queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"statusid\":{\"$gt\":null}},{\"statusField\":{\"$eq\":\"%s\"}}]}}", "Active")
		queryResults, err := getQueryResultForQueryString(stub, queryString)

		var arr_config []KeyRecordSC
		err = json.Unmarshal([]byte(queryResults), &arr_config)
		if err != nil {
			panic(err)
		}
		//fmt.println("arr_config",arr_config[0].Record);
		var errorcount_ebs_op_threshold,errorcount_ebs_op_shipping,errorcount_ebs_pen_threshold,errorcount_ebs_pen_shipping, errorcount_ebs_wh_threshold int
		var  errorcount_ebs_wh_shipping,totallineItems_op_threshold,totallineItems_op_shipping,totallineItems_pen_threshold,totallineItems_pen_shipping int 
		var totallineItems_wh_shipping,totallineItems_wh_threshold int
		errorcount_ebs_op_threshold  =0
		
		var arrerrorcount [2]int
			for _,statusconfigdetails := range arr_config{
				if (statusconfigdetails.Record.System=="EBS" && statusconfigdetails.Record.Status=="Order Placed"){
					queryString2 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"oracleOrderStatus\":{\"$gt\":\"%s\"}},{\"ebsMasterID\":{\"$ne\":\"%s\"}}]}}", "Order Placed", "null")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults2, err := getQueryResultForQueryString(stub, queryString2)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults2), &ebsorders)
						if err != nil {
							panic(err)
						}
					if (statusconfigdetails.Record.FieldValue=="thresholdvalue"){
						//get crm orders of order created
						var thresholdvaluetime=statusconfigdetails.Record.Thresholdtime 
						//fmt.Println("thresholdtime -----",arr_config[0].Record.Thresholdtime)
							for _, order := range ebsorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample2 := gettimedifference(thresholdvaluetime,pasttimestr)
						fmt.Println("query results for sample2 -----",sample2)
						fmt.Println("thresholdtime inside ebs -----",statusconfigdetails.Record.Thresholdtime)
						if (sample2>0){errorcount_ebs_op_threshold++
						// get LineItems here
						 //totallineItems= getLineItems(stub,orderNo)
						 queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults3, err := getQueryResultForQueryString(stub, queryString3)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						panic(err)
					}
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems_op_threshold += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems_op_threshold)
					fmt.Printf("satisfied")
							}
							}				
						}	
								//shippingMethod starts here

						if (statusconfigdetails.Record.FieldValue=="shippingmethod"){
						//get crm orders of order created
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
							for _, order := range ebsorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample2 := getdifference(pasttimestr)
						fmt.Println("query results for sample2 -----",sample2)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						if (sample2>0){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
								errorcount_ebs_op_shipping++	
								 //totallineItems1= getLineItems(stub,orderNo)			
								queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults3, err := getQueryResultForQueryString(stub, queryString3)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						panic(err)
					}
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems_op_shipping += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems_op_shipping)
					fmt.Printf("satisfied")
											}
									}
													
							}
						}
				}
						//inprogress orders starts herere
						if (statusconfigdetails.Record.System=="EBS" && statusconfigdetails.Record.Status=="Pending"){
					queryString4 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"oracleOrderStatus\":{\"$eq\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "Pending", "null")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults4, err := getQueryResultForQueryString(stub, queryString4)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults4), &crmorders)
						if err != nil {
							panic(err)
						}
					if (statusconfigdetails.Record.FieldValue=="thresholdvalue"){
						//get crm orders of order created
						var thresholdvaluetime=statusconfigdetails.Record.Thresholdtime 
						//fmt.Println("thresholdtime -----",arr_config[0].Record.Thresholdtime)
							for _, order := range ebsorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample3 := gettimedifference(thresholdvaluetime,pasttimestr)
						fmt.Println("query results for sample2 -----",sample3)
						fmt.Println("thresholdtime inside ebs -----",statusconfigdetails.Record.Thresholdtime)
						if (sample3>0){errorcount_ebs_pen_threshold++
						// get LineItems here
						// totallineItems2= getLineItems(stub,orderNo)
							queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults3, err := getQueryResultForQueryString(stub, queryString3)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						panic(err)
					}
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems_pen_threshold += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems_pen_threshold)
					fmt.Printf("satisfied")
							}
							}				
						}	
								//shippingMethod starts here

						if (statusconfigdetails.Record.FieldValue=="shippingmethod"){
						//get crm orders of order created
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
							for _, order := range ebsorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample4 := getdifference(pasttimestr)
						fmt.Println("query results for sample2 -----",sample4)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						if (sample4>0){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
								errorcount_ebs_pen_shipping++	
								 //totallineItems3= getLineItems(stub,orderNo)			
								queryString3 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults3, err := getQueryResultForQueryString(stub, queryString3)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						panic(err)
					}
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems_pen_shipping += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems_pen_shipping)
					fmt.Printf("satisfied")
											}
									}
													
							}
						}
						//inprogress ends here

						//send to warehouse starts here
						if (statusconfigdetails.Record.System=="EBS" && statusconfigdetails.Record.Status=="Order sent to warehouse"){
					queryString5 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"oracleOrderStatus\":{\"$eq\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "Order sent to warehouse", "null")
						//queryString2 := fmt.Sprintf("{\"selector\":{\"timeStamp\":{\"$ne\":\"%s\"}}}", "null")
						queryResults5, err := getQueryResultForQueryString(stub, queryString4)
						//fmt.Println("query results for crm orders -----",queryResults2)
						err = json.Unmarshal([]byte(queryResults5), &crmorders)
						if err != nil {
							panic(err)
						}
					if (statusconfigdetails.Record.FieldValue=="thresholdvalue"){
						//get crm orders of order created
						var thresholdvaluetime=statusconfigdetails.Record.Thresholdtime 
						//fmt.Println("thresholdtime -----",arr_config[0].Record.Thresholdtime)
							for _, order := range ebsorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample5 := gettimedifference(thresholdvaluetime,pasttimestr)
						fmt.Println("query results for sample5 -----",sample5)
						fmt.Println("thresholdtime inside ebs -----",statusconfigdetails.Record.Thresholdtime)
						if (sample5>0){errorcount_ebs_wh_threshold++
						// get LineItems here
						// totallineItems2= getLineItems(stub,orderNo)
							queryString6 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults6, err := getQueryResultForQueryString(stub, queryString6)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						panic(err)
					}
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems_wh_threshold += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems_wh_threshold)
					fmt.Printf("satisfied")
							}
							}				
						}	
								//shippingMethod starts here

						if (statusconfigdetails.Record.FieldValue=="shippingmethod"){
						//get crm orders of order created
						var ConfigShippingTime=statusconfigdetails.Record.ConfigShippingTime 
							for _, order := range ebsorders {
							fmt.Println("timestamp  -----",order.Record.Timestamp)
							var orderNo = order.Record.CRMOrderNumber
							pasttimestr := order.Record.Timestamp
							sample6 := getdifference(pasttimestr)
						fmt.Println("query results for sample6 -----",sample6)
						fmt.Println("ConfigShippingTime inside ebs -----",statusconfigdetails.Record.ConfigShippingTime)
						if (sample6>0){
							configtimecheckvalue:= configtimecheck(ConfigShippingTime)
							if (configtimecheckvalue>0){
								errorcount_ebs_wh_shipping++	
								 //totallineItems3= getLineItems(stub,orderNo)			
								queryString7 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"crmOrderNumber\":{\"$eq\":\"%s\"}},{\"totallineitems\":{\"$ne\":\"%s\"}}]}}", orderNo, "null")
					queryResults7, err := getQueryResultForQueryString(stub, queryString7)
					var soMaster []KeyRecordSO
					err = json.Unmarshal([]byte(queryResults3), &soMaster)
					if err != nil {
						panic(err)
					}
					
					fmt.Println("\n Total LineItems************************",soMaster[0].Record.TotalLineItems)
					var totallineItems_string = soMaster[0].Record.TotalLineItems
					totallineItems_int, err := strconv.Atoi(totallineItems_string);
					if err != nil {
						panic(err)
					}
					totallineItems_wh_shipping += totallineItems_int
					fmt.Println("\n totallineItems addded************************",totallineItems_wh_shipping)
					fmt.Printf("satisfied")
											}
									}
													
							}
						}

						//sed to ware House endshere
				}
			}
		//arr_errorcount[0] = errorcount
		arrerrorcount[0]= errorcount + errorcount1+errorcount3+errorcount2
		arrerrorcount[1]= totallineItems+totallineItems1+totallineItems3+totallineItems2+totallineItems6+totallineItems7
		
		//jsonbytes_errorcount, err := json.Marshal(arr_errorcount)
		jsonbytes_errorcount, err := json.Marshal(arrerrorcount)

		if err != nil {
			fmt.Printf("\nUnable marshal arr_errorcount : %v ", err)
			return shim.Error(err.Error())
		}
		//fetch data from couch db ends here	
		if err != nil {
			fmt.Printf("Unable to read Sales orders details for the given CRM Order No: %s\n", err)
			return shim.Error(err.Error())
		}
	
		fmt.Printf("list of order details for given CRMOrderNo : %v\n", queryResults)
		return shim.Success(jsonbytes_errorcount)
	} 



}
*/

//EBS created warning ends here

	/*
func gettimedifferenceshippingmethod(pasttimestring string, ConfigShippingTime string) int{
	total :=0
	
	currentTime := time.Now()
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
		layout := "2006-01-02 15:04"
				loc := currentTime.Location()
				fmt.Println("location  -----",currentTime )
				fmt.Println("x-----",x)
				fmt.Println("y-----",y)

					//converting string to date
				pastdatestring, err := time.ParseInLocation(layout, y, loc)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Past date: ", pastdatestring)
				fmt.Println("Current date: ", currentdatestring)
				//bth are same currentTime== currentdatestring
				//utc format date convert into date format 2006-01-02 
				currentdate := currentTime.Format("2006-01-02")
				fmt.Println("current date",currentdate)
				//convert orderdate utc format to date 
				pastdate := pastdatestring.Format("2006-01-02")
				fmt.Println("current date",pastdate)
				//diff := currentTime.Sub(pasttime)
				diff := currentdate.Sub(pastdate)
				fmt.Printf("time difference is %v or %v in minutes\n", diff)
				//fmt.Printf("differnce",diff.Minutes() ," given time" ,x)
				if (diff==0)
				{
					fmt.Printf("still the order date and current date are same .. so no warnings yet")
				}
				else{

				}
			//if (diff>)

}*/
func gettimedifference(x string, y string) int {
    total := 0
	currentTime := time.Now()
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
		
				layout := "2006-01-02 15:04"
				loc := currentTime.Location()
				fmt.Println("location  -----",currentTime )
				fmt.Println("x-----",x)
				fmt.Println("y-----",y)
	
			//converting string to date
				pasttime, err := time.ParseInLocation(layout, y, loc)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Past Time: ", pasttime)
				fmt.Println("Current Time: ", currentTime)
				diff := currentTime.Sub(pasttime)
				fmt.Printf("time difference is %v or %v in minutes\n", diff, diff.Minutes())
				fmt.Printf("differnce",diff.Minutes() ," given time" ,x)
				var mintimediff int = int(diff.Minutes())
				var minconfig = x
				var mininttimeconfig int
				mininttimeconfig, err = strconv.Atoi(minconfig); 
				if err == nil {
					fmt.Printf("i=%d, type: %T\n", mininttimeconfig, mininttimeconfig)
				}
				fmt.Printf("time diff============", mintimediff," config minutes------------", mininttimeconfig)
				if (mintimediff > mininttimeconfig) {
					total=total+1
				}else{
					total=total+0
				}
    return total
}

//getWarningsDetails


//getAllWarningLogs
func (t *kpnLogistics) getAllWarningLogs(stub shim.ChaincodeStubInterface, args []string) pb.Response {
                
	var err error
	fmt.Println("Entering getAllWarningLogs")
	//fetch data from couch db starts here 
	queryString := fmt.Sprintf("{\"selector\":{\"logid\":{\"$ne\":\"%s\"}}}", "null")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here
	if err != nil {
					fmt.Printf("Unable to read the list of getAllWarningLogs : %s\n", err)
					return shim.Error(err.Error())
	}

	fmt.Printf("list of all getAllWarningLogs : %v\n", queryResults)
	return shim.Success(queryResults)
}

// change dashboard status on warning starts here
func (t *kpnLogistics) updateDashboardStatusBasedOnWarings(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  	var objBcSalesOrderMaster SalesOrderMaster
	var objUiSalesOrderMaster SalesOrderMaster
	var err error

	fmt.Println("Entering updateDashboardStatusBasedOnWarings")
	if len(args) < 1 {
		fmt.Println("Invalid number of args")
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &objUiSalesOrderMaster)
	fmt.Println("args[0]", args[0])
	if err != nil {
		fmt.Printf("Unable to marshal  updateDashboardStatusBasedOnWarings : %s\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("\n refno SalesOrderMaster id value is : ", objUiSalesOrderMaster.CRMOrderNumber)

	// code to get data from blockchain using dynamic key starts here
	var bytesread []byte
	bytesread, err = stub.GetState(objUiSalesOrderMaster.CRMOrderNumber)
	err = json.Unmarshal(bytesread, &objBcSalesOrderMaster)
	// code to get data from blockchain using dynamic key ends here

	fmt.Printf("\nobjBcSalesOrderMaster in updateDashboardStatusBasedOnWarings : %s ", objBcSalesOrderMaster)

	objBcSalesOrderMaster.CRMDashboardStatus = "Warning"
		
	// Data insertion for Couch DB starts here
	transJSONasBytesContract, err := json.Marshal(objBcSalesOrderMaster)
	err = stub.PutState(objUiSalesOrderMaster.CRMOrderNumber, transJSONasBytesContract)
	// Data insertion for Couch DB ends here

	fmt.Println("updateDashboardStatusBasedOnWarings Successfully added in CRMDashboard struct")
	if err != nil {
		fmt.Printf("\nUnable to make transevent inputs : %v ", err)
		return shim.Error(err.Error())
		//return nil,nil
	}
	return shim.Success(nil)
}

//change dashboard status on warning ends here
// get the count of warning based on the dashboard status starts here

func (t *kpnLogistics) getWarningsCount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 	var err error
	fmt.Println("Entering getWarningsCount")

	//var CRMDashboardStatus = args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"crmDashboardStatus\":{\"$eq\": \"%s\"}}}", "warning")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	//fetch data from couch db ends here

	if err != nil {
		fmt.Printf("Unable to read the PurchaseOrder Details : %s\n", err)
		return shim.Error(err.Error())
	}
	fmt.Printf("DEtails od PurchaseOrder No : %v\n", queryResults)
	return shim.Success(queryResults)
}


//get the count of warning based on the dashboard status ends here

// Get log WarningCRMdetails starts here
func (t *kpnLogistics) getLogwarningDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
		var err error
		fmt.Println("Entering getLogwarningDetails By OrderNUmber for warnings")
	
		//fetch data from couch db starts here
		var waringOrdernumber = args[0]
		queryString := fmt.Sprintf("{\"selector\":{\"ordernumber\":{\"$eq\":\"%s\"}}}", waringOrdernumber)
		queryResults, err := getQueryResultForQueryString(stub, queryString)
		//fetch data from couch db ends here
		if err != nil {
			fmt.Printf("Unable to get Orders By getLogwarningDetails for the given waringOrdernumber: %s\n", err)
			return shim.Error(err.Error())
		}
	
		fmt.Printf("list of get Orders By getLogwarningDetails for the given waringOrdernumber: %v\n", queryResults)
		return shim.Success(queryResults)
	}




//Get log WarningCRMdetails ends here

// func (t *kpnLogistics) logWarningsDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	fmt.Println("Entering get getWarnings details")
// 	var err error
// 	var arrCRMDetails []WarningCRMdetails
// 	var newlogid = args[0]
// 	currentTime := time.Now()
// 	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
// 	//fetch data from couch db starts here
// 	//var status = "Order Created"
// 	queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"status\":{\"$eq\":\"%s\"}},{\"statusField\":{\"$eq\":\"%s\"}}]}}", "Order Created", "Active")
// 	queryResults, err := getQueryResultForQueryString(stub, queryString)

// 	var arr_config []KeyRecordSC
// 	err = json.Unmarshal([]byte(queryResults), &arr_config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("thresholdtime -----",arr_config[0].Record.Thresholdtime)
// 	queryString2 := fmt.Sprintf("{\"selector\":{\"$and\":[{\"orderStatus\":{\"$eq\":\"%s\"}},{\"timeStamp\":{\"$ne\":\"%s\"}}]}}", "Created", "null")
// 	queryResults2, err := getQueryResultForQueryString(stub, queryString2)
// 	fmt.Println("query results for crm orders -----",queryResults2)
// 	var crmorders []KeyRecordCRM
// 	err = json.Unmarshal([]byte(queryResults2), &crmorders)
// 	if err != nil {
// 		panic(err)
// 	}

// 	currentTime = time.Now()
// 	var statusreason = arr_config[0].Record.StatusReason
// 	for _, order := range crmorders {
// 			fmt.Println("timestamp  -----",order.Record.Timestamp)
// 			//past time comes in as string
// 			pasttimestr := order.Record.Timestamp
// 			layout := "2006-01-02 15:04"
// 			loc := currentTime.Location()
// 			fmt.Println("location  -----",loc)
// 			//converting string to date
// 			pasttime, err := time.ParseInLocation(layout, pasttimestr, loc)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println("Past Time: ", pasttime)
// 			fmt.Println("Current Time: ", currentTime)
// 			//differnce between pastdate and current date
// 			diff := currentTime.Sub(pasttime)
// 			fmt.Printf("time difference is %v or %v in minutes\n", diff, diff.Minutes())
// 			//fmt.Printf("differnce",diff.Minutes() ," given time" ,arr_config[0].Record.Thresholdtime)
// 			var mintimediff int = int(diff.Minutes())
// 			var minconfig = arr_config[0].Record.Thresholdtime
// 			var mininttimeconfig int
// 			if mininttimeconfig, err := strconv.Atoi(minconfig); err == nil {
// 				fmt.Printf("i=%d, type: %T\n", mininttimeconfig, mininttimeconfig)
// 			}

// 			if mintimediff > mininttimeconfig {
// 				//Log to blockchain
				





// 				arrCRMDetails =  WarningCRMdetails{
// 					LogId: fmt.Sprintf("%v",newlogid),
// 					Orderno: fmt.Sprintf("%v", order.Record.CRMOrderNumber),
// 					Customerno: fmt.Sprintf("%v", order.Record.CustomerNo),
// 					OrderStatus: fmt.Sprintf("%v", order.Record.CRMOrderStatus),
// 					Price: fmt.Sprintf("%v", order.Record.OrderTotalPrice),
// 					Reason: fmt.Sprintf("%v", statusreason),
// 				}
// 				jsonbytes_crmorderwarningdetails, err := json.Marshal(arrCRMDetails)
// 				if err != nil {
// 					fmt.Printf("\nUnable marshal arrCRMDetails : %v ", err)
// 					return shim.Error(err.Error())
// 				}
// 				err = stub.PutState(newlogid, jsonbytes_crmorderwarningdetails)
// 				if err != nil {
// 					fmt.Printf("\nUnable to log crm warning details to blockchain : %v ", err)
// 					return shim.Error(err.Error())
// 				}
// 				fmt.Printf("satisfied")
// 			}
// 	}

	
// 	//fmt.Printf(string(jsonbytes_crmorderdetails))
	

// 	return shim.Success(jsonbytes_crmorderdetails)
// }  
// getQueryResultForQueryString
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	//fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// Init sets up the chaincode
func (t *kpnLogistics) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Initiate the chaincde")
	return shim.Success(nil)
	//	return nil,nil
}

// Invoke the function in the chaincode
func (t *kpnLogistics) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "CreateSalesOrder" {
		return t.CreateSalesOrder(stub, args)
	}//UpdateReleaseOrderAndGenerateLspDetailsSecond
	if function == "UpdateStatusOnProcessOrder" {
		return t.UpdateStatusOnProcessOrder(stub, args)
	}
	if function == "updateProcessAcknowledgeXml" {
		return t.updateProcessAcknowledgeXml(stub, args)
	}
	if function == "UpdateStatusOnCancelOrder" {
		return t.UpdateStatusOnCancelOrder(stub, args)
	}
	if function == "getAllSalesOrder" {
		return t.getAllSalesOrder(stub, args)
	}
	if function == "getAllSalesOrderhd" {
		return t.getAllSalesOrderhd(stub, args)
	}
	if function == "UpdateProcessOrderInEBS" {
		return t.UpdateProcessOrderInEBS(stub, args)
	}
	if function == "updateCRMDashboardOnAcknowledge" {
		return t.updateCRMDashboardOnAcknowledge(stub, args)
	}
	if function == "UpdateReleaseOrderInEBS" {
		return t.UpdateReleaseOrderInEBS(stub, args)
	}
	if function == "UpdateLSPStatusesForProductID" {
		return t.UpdateLSPStatusesForProductID(stub, args)
	}
	if function == "UpdateShipOrderInLSP" {
		return t.UpdateShipOrderInLSP(stub, args)
	}
	if function == "updateShipmentInLSPOrderStatus" {
		return t.updateShipmentInLSPOrderStatus(stub, args)
	}
	if function == "updateLSPDashboardOnAcknowledge" {
		return t.updateLSPDashboardOnAcknowledge(stub, args)
	}
	if function == "GenerateLSPDetailsID" {
		return t.GenerateLSPDetailsID(stub, args)
	}
	if function == "getOrdersByCRMDashboardStatus" {
		return t.getOrdersByCRMDashboardStatus(stub, args)
	}
	if function == "getOrdersByEBSDashboardStatus" {
		return t.getOrdersByEBSDashboardStatus(stub, args)
	}
	if function == "getOrdersByLSPDashboardStatus" {
		return t.getOrdersByLSPDashboardStatus(stub, args)
	}
	
	if function == "getOrdersByTransportDashboardStatus" {
		return t.getOrdersByTransportDashboardStatus(stub, args)
	}
	if function == "getEbsMasterID" {
		return t.getEbsMasterID(stub, args)
	}
	if function == "getAllEBSOrders" {
		return t.getAllEBSOrders(stub, args)
	}
	if function == "getOrderDetailsByCRMOrderNo" {
		return t.getOrderDetailsByCRMOrderNo(stub, args)
	}
	if function == "getAllLSPOrders" {
		return t.getAllLSPOrders(stub, args)
	}
	if function == "getAllTransporterOrders" {
		return t.getAllTransporterOrders(stub, args)
	}
	if function == "getstatusconfig" {
		return t.getstatusconfig(stub, args)
	}
	if function == "CreateStatusConfig" {
		return t.CreateStatusConfig(stub, args)
	}
	if function == "updateStatusInTransporter" {
		return t.updateStatusInTransporter(stub, args)
	}
	if function == "UpdateShipDeliveredInTransporter" {
		return t.UpdateShipDeliveredInTransporter(stub, args)
	}
	if function == "updateDashboardOnTransporterAcknowledge" {
		return t.updateDashboardOnTransporterAcknowledge(stub, args)
	}
	if function == "getLSPDetailsStatusByProdID" {
		return t.getLSPDetailsStatusByProdID(stub, args)
	}
	if function == "updateCRMDashboardOnError" {
		return t.updateCRMDashboardOnError(stub, args)
	}
	if function == "updateEBSDashboardonError" {
		return t.updateEBSDashboardonError(stub, args)
	}
	if function == "getCRMMasterID" {
		return t.getCRMMasterID(stub, args)
	}
	if function == "getEbsMaster" {
		return t.getEbsMaster(stub, args)
	}
	if function == "getWarnings" {
		return t.getWarnings(stub, args)
	}
	if function == "getWarningsDetails" {
		return t.getWarningsDetails(stub, args)
	}
	if function == "updateStatusConfig" {
		return t.updateStatusConfig(stub, args)
	}
	if function == "getstatusconfigByStatusID" {
		return t.getstatusconfigByStatusID(stub, args)
	}
	if function == "deleteStatusConfig" {
		return t.deleteStatusConfig(stub, args)
	}
	 if function == "checkcrmwarning" {
	 	return t.checkcrmwarning(stub, args)
	 }
	 if function == "LogCRMWarnings" {
		return t.LogCRMWarnings(stub, args)
	}
	if function == "getAllWarningLogs" {
		return t.getAllWarningLogs(stub, args)
	}
	if function == "UpdateReleaseOrderAndGenerateLspDetails" {
		return t.UpdateReleaseOrderAndGenerateLspDetails(stub, args)
	}
	
	//modifications 
	
	if function == "updateDashboardStatusBasedOnWarings" {
		return t.updateDashboardStatusBasedOnWarings(stub, args)
	}
	if function == "getLogwarningDetails" {
		return t.getLogwarningDetails(stub, args)
	}
	
	
	fmt.Println("Function not found")
	return shim.Error("Received unknown function invocation")
	//return nil, nil
}

func main() {
	err := shim.Start(new(kpnLogistics))
	if err != nil {
		fmt.Println("Could not start Chaincode")
	} else {
		fmt.Println("Chaincode successfully started")
	}

}
