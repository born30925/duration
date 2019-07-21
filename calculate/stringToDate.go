package calculate

import (
	"fmt"
	"time"
)

func StringToDateTime(stringDate string) time.Time {
	t, _ := time.Parse("02-01-2006", stringDate)
	fmt.Println(t)
	return t
}
