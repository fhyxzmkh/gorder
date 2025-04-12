package stock

import (
	"context"
	"fmt"
	"github.com/fhyxzmkh/gorder/common/genproto/orderpb"
)

type Repository interface {
	GetItems(ctx context.Context, ids []string) ([]*orderpb.Item, error)
}

type NotFoundError struct {
	Missing []string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("thses items are missing: %v", e.Missing)
}
