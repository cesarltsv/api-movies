package helpers

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}
