package client

import (
	"github.com/A-ndrey/leetcode-client/internal/types"
	"strings"
)

type Request struct {
	Query string `json:"query"`
}

func NewRequest(operation types.Operation) Request {
	sb := strings.Builder{}
	operation.String(&sb)

	return Request{Query: sb.String()}
}
