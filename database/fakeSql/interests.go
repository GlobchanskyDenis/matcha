package fakeSql

import (
	"MatchaServer/common"
)

func (conn ConnFake) AddInterests(unknownInterests []common.Interest) error {
	var interest common.Interest
	var lastId int

	for _, knownInterest := range conn.interests {
		lastId = knownInterest.Id
	}

	for _, unknownInterest := range unknownInterests {
		interest.Name = unknownInterest.Name
		lastId++
		interest.Id = lastId
		conn.interests[lastId] = interest
	}
	return nil

}

func (conn ConnFake) GetInterests() ([]common.Interest, error) {
	var interests = []common.Interest{}
	for _, interest := range conn.interests {
		interests = append(interests, interest)
	}
	return interests, nil
}
