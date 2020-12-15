package model

// ServiceOrder struct
type ServiceOrder struct {
	IdOs       int `json: "idos"`
	IdCliente  int `json: "idcliente"`
	IdStatus   int `json: "idstatus"`
}

// ServiceOrders struct
type ServiceOrders struct {
	ServiceOrders []ServiceOrder `json: "serviceorders"`
}
