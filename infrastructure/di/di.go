package di

import (
	"sort"
)

type Di struct {
	preBuildServices map[ServiceKey]func() (any, error)
	errors           map[ServiceKey]error
	servicesPriority map[int]ServiceKey
}

func NewApp() Di {
	return Di{
		preBuildServices: make(map[ServiceKey]func() (any, error)),
		errors:           make(map[ServiceKey]error),
		servicesPriority: make(map[int]ServiceKey),
	}
}

func (di *Di) Add(name ServiceKey, initFunc func() (any, error), priority int) {
	di.preBuildServices[name] = initFunc
	for {
		if _, isSet := di.servicesPriority[priority]; isSet {
			priority++
		} else {
			break
		}
	}

	di.servicesPriority[priority] = name
}

func (di *Di) Build() {
	servicesCount := len(di.servicesPriority)
	var priorityList []int
	for k := range di.servicesPriority {
		priorityList = append(priorityList, k)
	}

	sort.Ints(priorityList)

	for i := servicesCount - 1; i >= 0; i-- {
		serviceName := di.servicesPriority[priorityList[i]]
		initFunc := di.preBuildServices[serviceName]
		serviceInstance, err := initFunc()
		if err != nil {
			di.errors[serviceName] = err
		}
		Set(serviceName, serviceInstance)
	}
}

func (di *Di) HasErrors() bool {
	return len(di.errors) > 0
}

func (di *Di) GetFirstError() (serviceName ServiceKey, err error) {
	for serviceName, err = range di.errors {
		break
	}

	return
}
