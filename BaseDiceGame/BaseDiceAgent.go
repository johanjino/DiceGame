package BaseDiceGame

import (
	"fmt"
	"math/rand"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

type IDiceAgent interface {
	agent.IAgent[IDiceAgent]

	// More Agent functions defined here
	SetTeam(int)
	ResetScore()
	RollDice()
	GetScore() int
	CreateScorePoolingMessage() *ScorePoolingMessage
	HandleScorePoolingMessage(ScorePoolingMessage)
}

type DiceAgent struct {
	*agent.BaseAgent[IDiceAgent]
	Team  int
	Score int
}

// More server functions implemented here
func (agent *DiceAgent) SetTeam(TeamNum int) {
	agent.Team = TeamNum
}

func (agent *DiceAgent) ResetScore() {
	agent.Score = 0
}

func (agent *DiceAgent) RollDice() {
	prev := 0
	total := 0
	stick := false
	bust := false

	for !stick && !bust {
		// Roll three dice
		r1, r2, r3 := (rand.Intn(6) + 1), (rand.Intn(6) + 1), (rand.Intn(6) + 1)
		score := r1 + r2 + r3

		if score > prev {
			total += score
			prev = score

			// Call do_I_stick to decide whether to stick or continue
			stick = do_I_stick(total, prev)
		} else {
			bust = true
			score = 0
		}
	}

	// Returned value is to be redistributed
	redistribute := 0                  // For now redistribute 0
	agent.Score = total - redistribute // Remaining kept by agent

	//Send Message
	msg := ScorePoolingMessage{
		BaseMessage: agent.CreateBaseMessage(),
		Score:       agent.Score,
		Team:        agent.Team,
	}
	agent.BroadcastMessage(&msg)
}

func (agent *DiceAgent) GetScore() int {
	return agent.Score
}

func (agent *DiceAgent) CreateScorePoolingMessage() *ScorePoolingMessage {
	return &ScorePoolingMessage{
		BaseMessage: agent.CreateBaseMessage(),
		Score:       agent.Score,
		Team:        agent.Team,
	}
}

func (agent *DiceAgent) HandleScorePoolingMessage(msg ScorePoolingMessage) {
	// Do nothing, no redistribution for now
	fmt.Printf("Received Msg from %v with score %v\n", msg.Team, msg.Score)
	agent.SignalMessagingComplete()
}

// func (msg ScorePoolingMessage) InvokeMessageHandler (agent IDiceAgent){
// 	agent.HandleScorePoolingMessage(msg)
// }

func do_I_stick(total, prev int) bool { // More parameters to be added later
	// Simple strategy for now
	// Re-role if there is probability of >50% to get higher
	// i.e. if roll>=9
	return prev >= 9
}

func CreateDiceAgent(serv agent.IExposedServerFunctions[IDiceAgent]) IDiceAgent {
	return &DiceAgent{
		BaseAgent: agent.CreateBaseAgent(serv),
		Team:      -1,
		Score:     0,
	}
}
