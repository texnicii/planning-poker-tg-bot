package messaging

type ResponseBag struct {
	Responses []any
}

type ChatResponse struct {
	chatId       int64
	text         string
	markup       any
	callbackText string
}

type CallbackResponse struct {
	queryId string
	text    string
}

func (rb *ResponseBag) AddChatResponse(chatId int64, text string) {
	rb.Responses = append(rb.Responses, ChatResponse{
		chatId: chatId,
		text:   text,
	})
}

func (rb *ResponseBag) AddChatResponseWithMarkup(chatId int64, text string, markup any) {
	rb.Responses = append(rb.Responses, ChatResponse{
		chatId: chatId,
		text:   text,
		markup: markup,
	})
}

func (rb *ResponseBag) AddCallbackResponse(queryId string, text string) {
	rb.Responses = append(rb.Responses, CallbackResponse{
		queryId: queryId,
		text:    text,
	})
}

func (r ChatResponse) Text() string {
	return r.text
}

func (r ChatResponse) Markup() any {
	return r.markup
}

func (r ChatResponse) ChatId() int64 {
	return r.chatId
}

func (r CallbackResponse) Text() string {
	return r.text
}

func (r CallbackResponse) QueryId() string {
	return r.queryId
}
