package interfaces

type Sender interface {
	Send(chatId int64, text string)
}
