package handler

type OrdersGrpcHandler struct {
	//service injection
	// uniplemented UnimplementedOrderServiceServer
}

func NewGrpcOrderService() {
	gRPCHandler := &OrdersGrpcHandler{}

	// register the OrderServiceServer
}
