package text

type myTextUsecases struct {
	DB TextMessageRepository
}

func NewTextUsecases(
	db TextMessageRepository) TextMessageUsecases {
	return &myTextUsecases{
		DB: db,
	}
}

func (u *myTextUsecases) GetRandTextMessage() (randMsg string, err error) {
	randMsg, err = u.DB.GetRandTextMessage()
	if err != nil {
		return
	}

	return
}

func (u *myTextUsecases) AddTextMessage(message string) (err error) {
	err = u.DB.AddTextMessage(message)
	if err != nil {
		return
	}

	return
}

func (u *myTextUsecases) GetTextMessagesCount() (text int, err error) {
	text, err = u.DB.GetTextMessagesCount()
	if err != nil {
		return
	}
	return
}
