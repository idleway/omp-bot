package answer

import "errors"

const answerIDNotFound = "answerID not found"

type DummyAnswerService struct{}

func NewDummyAnswerService() *DummyAnswerService {
	return &DummyAnswerService{}
}

func (*DummyAnswerService) Describe(answerID uint64) (*Answer, error) {
	for i := range allEntities {
		if allEntities[i].ID == answerID {
			return &allEntities[i], nil
		}
	}
	return nil, errors.New(answerIDNotFound)
}

func (*DummyAnswerService) List(cursor uint64, limit uint64) ([]Answer, error) {
	//TODO в каком случае тут может быть ошибка?
	start := cursor
	start = min(start, uint64(len(allEntities))-1)
	finish := min(start+limit, uint64(len(allEntities)))

	return allEntities[start:finish], nil
}

func (*DummyAnswerService) Create(answer Answer) (uint64, error) {
	//TODO в каком случае тут может быть ошибка?
	answer.ID = allEntities[len(allEntities)-1].ID + 1
	allEntities = append(allEntities, answer)
	return answer.ID, nil
}

func (*DummyAnswerService) Update(answer Answer) error {
	for i := range allEntities {
		if allEntities[i].ID == answer.ID {
			allEntities[i] = answer
			return nil
		}
	}
	return errors.New(answerIDNotFound)
}

func (*DummyAnswerService) Remove(answerID uint64) (bool, error) {
	for i := range allEntities {
		if allEntities[i].ID == answerID {
			allEntities = append(allEntities[:i], allEntities[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (*DummyAnswerService) Count() uint64 {
	return uint64(len(allEntities))
}

func min(x uint64, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}
