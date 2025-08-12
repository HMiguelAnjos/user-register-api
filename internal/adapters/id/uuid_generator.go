package id

import (
	"userregisterapi/internal/app/ports"

	"github.com/google/uuid"
)

// Ensure implementation
var _ ports.IDGenerator = (*UUIDGenerator)(nil)

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator { return &UUIDGenerator{} }

func (g *UUIDGenerator) NewID() (string, error) {
	return uuid.NewString(), nil
}
