// Package server
package server

import (
	"context"

	"github.com/momoney3/cloud-feona/internal/rag"
	pb "github.com/momoney3/cloud-feona/proto"
)

type Server struct {
	pb.UnimplementedRagServiceServer
	pipeline *rag.Pipeline
}

func New(p *rag.Pipeline) *Server {
	return &Server{pipeline: p}
}

func (s *Server) Ask(reg *pb.AskRequest, stream pb.RagService_AskServer) error {
	ctx := context.Background()

	var sendErr error

	err := s.pipeline.Ask(ctx, reg.Question, func(token string) {
		if sendErr != nil {
			return
		}

		sendErr = stream.Send(&pb.AskResponse{
			Token: token,
		})
	})

	if sendErr != nil {
		return sendErr
	}
	return err
}
