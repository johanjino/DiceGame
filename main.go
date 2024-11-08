package main


import (
	"Game/BaseDiceGame"
	"time"
)

func main() {
	serv := BaseDiceGame.CreateDiceServer(2, 5, 1, 12, 10*time.Millisecond, 100)
	//toggle message diagnostics
	serv.ReportMessagingDiagnostics()
	//begin simulation
	serv.Start()
}