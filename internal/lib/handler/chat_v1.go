package handler

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/fatih/color"
	desc "github.com/mixdjoker/chat-server/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

// Create is a method that implements the Create method of the Chat_V1Server interface
func (h *ChatHandlerV1) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if dline, ok := ctx.Deadline(); ok {
		h.log.Println(color.BlueString("Deadline: %v", dline))
	}

	buf := strings.Builder{}
	buf.WriteString("Received Create:\n")

	for i, user := range req.Usernames {
		usrStr := fmt.Sprintf("\t#%d Username: %s\n", i, user)
		buf.WriteString(usrStr)
	}

	h.log.Println(color.BlueString(buf.String()))

	randInt64, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<63-1))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	id := randInt64.Int64()

	respStr := fmt.Sprintf("Response Create:\n\tID: %v\n", id)

	h.log.Println(color.GreenString(respStr))

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// Delete is a method that implements the Delete method of the Chat_V1Server interface
func (h *ChatHandlerV1) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	id := req.Id

	if dline, ok := ctx.Deadline(); ok {
		h.log.Println(color.BlueString("Deadline: %v", dline))
	}

	h.log.Println(color.BlueString("Received Delete:\n\tID: %v", id))

	return &emptypb.Empty{}, nil
}

// SendMessage is a method that implements the SendMessage method of the Chat_V1Server interface
func (h *ChatHandlerV1) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {
	from := req.From
	text := req.Text
	when := req.Timestamp

	if dline, ok := ctx.Deadline(); ok {
		h.log.Println(color.BlueString("Deadline: %v", dline))
	}

	h.log.Println(color.BlueString("Received SendMessage:\n\tFrom: %v\n\tText: %v\n\tTimestamp: %v", from, text, when))

	return &emptypb.Empty{}, nil
}
