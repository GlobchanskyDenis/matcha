package fakeSql

import (
	"MatchaServer/config"
)

func (conn ConnDB) AddInterests(unknownInterests []config.Interest) error {
	// var interests []config.Interest{}
	var interest config.Interest
	var lastId int

	for _, knownInterest := range conn.Interests {
		lastId = knownInterest.Id
	}
	
	for _, unknownInteres := range unknownInterests {
		interest.Name = unknownInterest.Name
		lastId++
		interest.Id = lastId
		conn.Interests[lastId] = interest
		// interests = append(interests, interest)
	}
	return nil
	
}

func (conn ConnDB) GetInterests() ([]config.Interest, error) {
	var interests = []config.Intersts{}
	for _, interest := range conn.Interests {
		interests = append(interests, interest)
	}
	return interests, nil
}