package gim

import (
	"errors"
	"strings"

	"github.com/flywave/topotypes/gim/ec"
	"github.com/flywave/topotypes/gim/gs"
	"github.com/flywave/topotypes/gim/gt"
)

const Major = "GIM"

type Shape interface {
	GetType() string
}

func Unmarshal(ty string, bt []byte) (Shape, error) {
	tys := strings.Split(ty, "/")
	if len(tys) != 3 {
		return nil, errors.New("invalid type")
	}
	if tys[0] != Major {
		return nil, errors.New("invalid major")
	}
	if tys[1] == ec.Major {
		return ec.Unmarshal(ty, bt)
	} else if tys[1] == gs.Major {
		return gs.Unmarshal(ty, bt)
	} else if tys[1] == gt.Major {
		return gt.Unmarshal(ty, bt)
	}
	return nil, errors.New("invalid type")
}
