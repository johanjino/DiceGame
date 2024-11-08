package BaseDiceGame

import (
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/message"
)


type ScorePoolingMessage struct {
	message.BaseMessage
	Score int
	Team int
}

func (msg ScorePoolingMessage) InvokeMessageHandler (agent IDiceAgent){
	agent.HandleScorePoolingMessage(msg)
}


// Bottom ACL is irrelevant at the moment

// Agent Communication Language
// {
// 	"Type": "Message_Type",
// 	"Sender": "Agent_uid",
// 	"Receiver": "Agent_uid (can be null if broadcast)",
// 	"Content": "{Json}"
// }
