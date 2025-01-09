package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("product_view_%s", strconv.Itoa(int(id)))
}
