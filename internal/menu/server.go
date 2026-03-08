// Package menu
package menu

import (
	"context"

	pb "github.com/momoney3/cloud-feona/proto/ai"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(*pb.MenuRequest, pd.ServerStreamingServer) error {
	item := []*pb.Item
}

func (s *server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
}

func (s *server) GetOrderStatus(context.Context, *pb.Receipt) (*pb.OrderStatus, error) {
}
