package handler

import (
	"log"

	desc "github.com/mixdjoker/chat-server/pkg/chat_v1"
)

// ChatHandlerV1 is a struct that implements Chat_V1Server interface
type ChatHandlerV1 struct {
	desc.UnimplementedChat_V1Server
	log *log.Logger
}

// NewChatHandlerV1 returns a new ChatHandlerV1 instance
func NewChatHandlerV1(log *log.Logger) *ChatHandlerV1 {
	return &ChatHandlerV1{
		log: log,
	}
}
