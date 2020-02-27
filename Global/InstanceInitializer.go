package Global

import (
	"github.com/jc01rho/ogamego_bot/Bot"
	"github.com/jc01rho/ogamego_bot/Queue"
)

func InitEssentials() {

	Queue.InitQueue()
	Bot.InitBot()

}
