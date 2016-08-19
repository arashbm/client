// Auto-generated by avdl-compiler v1.3.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat_local.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	chat1 "github.com/keybase/gregor/protocol/chat1"
	context "golang.org/x/net/context"
)

type MessageText struct {
	Body string `codec:"body" json:"body"`
}

type MessageConversationMetadata struct {
	ConversationTitle string `codec:"conversationTitle" json:"conversationTitle"`
}

type MessageEdit struct {
	MessageID chat1.MessageID `codec:"messageID" json:"messageID"`
	Body      string          `codec:"body" json:"body"`
}

type MessageDelete struct {
	MessageID chat1.MessageID `codec:"messageID" json:"messageID"`
}

type MessageAttachment struct {
	Path string `codec:"path" json:"path"`
}

type MessageBody struct {
	Type                 chat1.MessageType            `codec:"type" json:"type"`
	Text                 *MessageText                 `codec:"text,omitempty" json:"text,omitempty"`
	Attachment           *MessageAttachment           `codec:"attachment,omitempty" json:"attachment,omitempty"`
	Edit                 *MessageEdit                 `codec:"edit,omitempty" json:"edit,omitempty"`
	Delete               *MessageDelete               `codec:"delete,omitempty" json:"delete,omitempty"`
	ConversationMetadata *MessageConversationMetadata `codec:"conversationMetadata,omitempty" json:"conversationMetadata,omitempty"`
}

type MessagePlaintext struct {
	ClientHeader  chat1.MessageClientHeader `codec:"clientHeader" json:"clientHeader"`
	MessageBodies []MessageBody             `codec:"messageBodies" json:"messageBodies"`
}

type Message struct {
	ServerHeader     chat1.MessageServerHeader `codec:"serverHeader" json:"serverHeader"`
	MessagePlaintext MessagePlaintext          `codec:"messagePlaintext" json:"messagePlaintext"`
}

type ThreadView struct {
	Messages   []Message         `codec:"messages" json:"messages"`
	Pagination *chat1.Pagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type MessageSelector struct {
	MessageTypes []chat1.MessageType `codec:"MessageTypes" json:"MessageTypes"`
	After        *Time               `codec:"After,omitempty" json:"After,omitempty"`
	Before       *Time               `codec:"Before,omitempty" json:"Before,omitempty"`
	OnlyNew      bool                `codec:"onlyNew" json:"onlyNew"`
	LimitNumber  int                 `codec:"limitNumber" json:"limitNumber"`
}

type GetInboxLocalArg struct {
	Pagination *chat1.Pagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetThreadLocalArg struct {
	ConversationID chat1.ConversationID `codec:"conversationID" json:"conversationID"`
	Pagination     *chat1.Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type PostLocalArg struct {
	ConversationID   chat1.ConversationID `codec:"conversationID" json:"conversationID"`
	MessagePlaintext MessagePlaintext     `codec:"messagePlaintext" json:"messagePlaintext"`
}

type NewConversationLocalArg struct {
	ConversationTriple chat1.ConversationIDTriple `codec:"conversationTriple" json:"conversationTriple"`
}

type GetOrCreateTextConversationLocalArg struct {
	TlfName   string          `codec:"tlfName" json:"tlfName"`
	TopicName string          `codec:"topicName" json:"topicName"`
	TopicType chat1.TopicType `codec:"topicType" json:"topicType"`
}

type GetMessagesLocalArg struct {
	Selector MessageSelector `codec:"selector" json:"selector"`
}

type ChatLocalInterface interface {
	GetInboxLocal(context.Context, *chat1.Pagination) (chat1.InboxView, error)
	GetThreadLocal(context.Context, GetThreadLocalArg) (ThreadView, error)
	PostLocal(context.Context, PostLocalArg) error
	NewConversationLocal(context.Context, chat1.ConversationIDTriple) (chat1.ConversationID, error)
	GetOrCreateTextConversationLocal(context.Context, GetOrCreateTextConversationLocalArg) (chat1.ConversationID, error)
	GetMessagesLocal(context.Context, MessageSelector) ([]Message, error)
}

func ChatLocalProtocol(i ChatLocalInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.chatLocal",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getInboxLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxLocalArg)(nil), args)
						return
					}
					ret, err = i.GetInboxLocal(ctx, (*typedArgs)[0].Pagination)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getThreadLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetThreadLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetThreadLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetThreadLocalArg)(nil), args)
						return
					}
					ret, err = i.GetThreadLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"postLocal": {
				MakeArg: func() interface{} {
					ret := make([]PostLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PostLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]PostLocalArg)(nil), args)
						return
					}
					err = i.PostLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationLocal": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationLocalArg)(nil), args)
						return
					}
					ret, err = i.NewConversationLocal(ctx, (*typedArgs)[0].ConversationTriple)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getOrCreateTextConversationLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetOrCreateTextConversationLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetOrCreateTextConversationLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetOrCreateTextConversationLocalArg)(nil), args)
						return
					}
					ret, err = i.GetOrCreateTextConversationLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getMessagesLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetMessagesLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetMessagesLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetMessagesLocalArg)(nil), args)
						return
					}
					ret, err = i.GetMessagesLocal(ctx, (*typedArgs)[0].Selector)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ChatLocalClient struct {
	Cli rpc.GenericClient
}

func (c ChatLocalClient) GetInboxLocal(ctx context.Context, pagination *chat1.Pagination) (res chat1.InboxView, err error) {
	__arg := GetInboxLocalArg{Pagination: pagination}
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.getInboxLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) GetThreadLocal(ctx context.Context, __arg GetThreadLocalArg) (res ThreadView, err error) {
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.getThreadLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) PostLocal(ctx context.Context, __arg PostLocalArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.postLocal", []interface{}{__arg}, nil)
	return
}

func (c ChatLocalClient) NewConversationLocal(ctx context.Context, conversationTriple chat1.ConversationIDTriple) (res chat1.ConversationID, err error) {
	__arg := NewConversationLocalArg{ConversationTriple: conversationTriple}
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.newConversationLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) GetOrCreateTextConversationLocal(ctx context.Context, __arg GetOrCreateTextConversationLocalArg) (res chat1.ConversationID, err error) {
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.getOrCreateTextConversationLocal", []interface{}{__arg}, &res)
	return
}

func (c ChatLocalClient) GetMessagesLocal(ctx context.Context, selector MessageSelector) (res []Message, err error) {
	__arg := GetMessagesLocalArg{Selector: selector}
	err = c.Cli.Call(ctx, "keybase.1.chatLocal.getMessagesLocal", []interface{}{__arg}, &res)
	return
}
