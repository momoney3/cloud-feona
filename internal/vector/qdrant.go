import (
	"context"

	"github.com/momoney3/cloud-feona/internal/rag"
	"github.com/qdrant/go-client/qdrant"
)

func (s *Store) Search(ctx context.Context, vec []float32, limit int) ([]rag.Document, error) {
	resp, err := s.client.SearchPoints(ctx, &qdrant.SearchPoints{
		CollectionName: "docs",
		Vector:         vec,
		Limit:          uint64(limit),
	})
	if err != nil {
		return nil, err
	}

	var docs []rag.Document
	for _, r := range resp.Result {
		val, ok := r.Payload["text"]
		if !ok {
			continue
		}

		text := val.GetStringValue()

		docs = append(docs, rag.Document{
			Text: text,
		})
	}

	return docs, nil
}
