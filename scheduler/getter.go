package scheduler

import (
	"HTTProxyPool/model"
)

type Getter interface {
	Run (result chan <- *model.Proxy)
}
