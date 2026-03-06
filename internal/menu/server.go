package menu

import (
	pb "github.com/brd"
	pb "github.com/momoney3/cloud-feona"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu()
