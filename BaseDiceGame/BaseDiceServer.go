package BaseDiceGame

import (
	"fmt"
	"time"
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
)

type IDiceServer interface{
	server.IServer[IDiceAgent]
	
	// More server functions defined here
	SelfSelection()
	// NegotiateRules() // Not Implemented Yet
	PlayGame()
}

type DiceServer struct {
	*server.BaseServer[IDiceAgent]
	Threshold int
}

// More server functions implemented here
func (serv *DiceServer) SelfSelection(){
	for _, agent := range serv.GetAgentMap() {
		agent.SetTeam(1) //Set all to Team 1 for now
	}
}

// Not Implemented yet
// func (serv *DiceServer) NegotiateRules(){
// 	for _, agent := range serv.GetAgentMap() {
// 		// What goes here?
// 	}
// }

func (serv *DiceServer) PlayGame(){
	for _, agent := range serv.GetAgentMap() {
		agent.RollDice()
	}
}

func (serv *DiceServer) RunTurn(i, j int) {
	fmt.Printf("Running iteration %v, turn  %v\n", i+1, j+1)
	// Run Game
	serv.SelfSelection()
	serv.PlayGame()
	
	// How to ensure agents have redistributed scores by now?

	// Eliminate agents below threshold and also Reset Score for next round
	
	// Create a slice to hold agents that need to be removed
	var agentsToRemove []IDiceAgent

	// Loop through each agent
	for _, agent := range serv.GetAgentMap() {
		// Check if agent's score is below the threshold
		if agent.GetScore() <= serv.Threshold {
			// Add agent to the slice for removal later
			agentsToRemove = append(agentsToRemove, agent)
		} else {
			// Reset agent's score if it's above the threshold
			agent.ResetScore()
		}
	}

	// Now remove the agents from the map after the loop finishes
	time.Sleep(2 * time.Second) // This added because seems like goroutine of agent still running when attempt to remove. Temporary fix?
	for _, agent := range agentsToRemove {
		serv.RemoveAgent(agent)
	}
}

func (serv *DiceServer) RunStartOfIteration(iteration int) {
	fmt.Printf("Starting iteration %v\n", iteration+1)
	fmt.Printf("Number of agents %v\n", len(serv.GetAgentMap()))
	fmt.Println()
}

func (serv *DiceServer) RunEndOfIteration(iteration int){
	fmt.Println()
	fmt.Printf("Number of agents remaining %v\n", len(serv.GetAgentMap()))
	fmt.Printf("Ending iteration %v\n", iteration+1)
}


// Iteration vs Turns ?

func CreateDiceServer(numAgents, iterations, turns int, threshold int, maxDuration time.Duration, agentBandwidth int) *DiceServer{
	serv := &DiceServer{
		BaseServer: server.CreateBaseServer[IDiceAgent](
			iterations, turns, maxDuration, agentBandwidth,
		),
		Threshold: threshold,

	}
	for i:=0; i<numAgents; i++ { 
		serv.AddAgent(CreateDiceAgent(serv))
	}
	serv.SetGameRunner(serv)
	return serv
}