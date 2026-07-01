package service

import (
	"log"
	m "main/internal/models"
)

func (s *Service) GetUserCount(device string) (int, error) {

	var users []m.UserCard
	var err error

	nodeIds, err := s.rep.Device.GetActiveNode(device)
	if err != nil {
		log.Println(err)
	}

	if len(nodeIds) > 0 {
		users, err = s.rep.User.GetUserByNodes(nodeIds)
		if err != nil {
			log.Println(err)
		}
	}

	if len(users) > 0 {
		s.rep.User.DropCache(device)
		s.rep.User.CreateCache(device, users)
	}

	return len(users), nil
}
