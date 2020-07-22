package member

import (
	"github.com/vmihailenco/msgpack"
)

type Member struct {
	ID    int64  `json:"id" db:"id"`
	Phone string `json:"phone" db:"phone" `
	Name  string `json:"name" dn:"name" `
	Age   int8   `json:"age" db:"age"`
}

func (m *Member) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(m)
}

// UnmarshalBinary use msgpack
func (m *Member) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, m)
}
