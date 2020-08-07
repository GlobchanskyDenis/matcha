package fakeSql

import (
	"MatchaServer/config"
)

func (conn ConnFake) AddInterests(unknownInterests []config.Interest) error {
	// var interests []config.Interest{}
	var interest config.Interest
	var lastId int

	for _, knownInterest := range conn.interests {
		lastId = knownInterest.Id
	}
	
	for _, unknownInterest := range unknownInterests {
		interest.Name = unknownInterest.Name
		lastId++
		interest.Id = lastId
		conn.interests[lastId] = interest
		// interests = append(interests, interest)
	}
	return nil
	
}

func (conn ConnFake) GetInterests() ([]config.Interest, error) {
	var interests = []config.Interest{}
	for _, interest := range conn.interests {
		interests = append(interests, interest)
	}
	return interests, nil
}