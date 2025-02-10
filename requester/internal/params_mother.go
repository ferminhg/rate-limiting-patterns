package internal

import (
	"github.com/go-faker/faker/v4"
)

func ParamsMotherRamdon() *Params {
	numRequests, _ := faker.RandomInt(1)
	host := faker.URL()

	return NewParams(numRequests[0], host)
}
