package adapters

import (
	"context"
	domain "github.com/fhyxzmkh/gorder/order/domain/order"
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

type MemoryOrderRepository struct {
	lock  *sync.RWMutex
	store []*domain.Order
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	s := make([]*domain.Order, 0)
	s = append(s, &domain.Order{
		ID:          "fake-ID",
		CustomerID:  "fake-customer-id",
		Status:      "fake-status",
		PaymentLink: "fake-payment-link",
		Items:       nil,
	})

	return &MemoryOrderRepository{
		lock:  &sync.RWMutex{},
		store: s,
	}
}

func (m *MemoryOrderRepository) Create(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	newOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  order.CustomerID,
		Status:      order.Status,
		PaymentLink: order.PaymentLink,
		Items:       order.Items,
	}

	m.store = append(m.store, newOrder)
	logrus.WithFields(logrus.Fields{
		"input_order":        order,
		"store_after_create": m.store,
	}).Infof("memory_order_repository.Create")
	return newOrder, nil
}

func (m *MemoryOrderRepository) Get(ctx context.Context, id, customerID string) (*domain.Order, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, order := range m.store {
		if order.ID == id {
			if order.CustomerID == customerID {
				logrus.Infof("memory_order_repository.Get")
				return order, nil
			}
		}
	}

	return nil, &domain.NotFoundError{
		OrderID: id,
	}
}

func (m *MemoryOrderRepository) Update(ctx context.Context, o *domain.Order, updateFn func(context.Context, *domain.Order) (*domain.Order, error)) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	found := false

	for i, order := range m.store {
		if order.ID == o.ID && order.CustomerID == o.CustomerID {
			found = true
			updatedOrder, err := updateFn(ctx, o)
			if err != nil {
				return err
			}
			m.store[i] = updatedOrder
		}
	}

	if !found {
		return &domain.NotFoundError{
			OrderID: o.ID,
		}
	}

	logrus.Debug("memory_order_repository.Update")
	return nil
}
