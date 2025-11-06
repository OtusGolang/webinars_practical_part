package typeswitch

import "testing"

type MsgUserBalanceChanged struct {
	userID  string
	balance string
}

type MsgEventChanged struct {
	eventID string
}

func processMessage(msg interface{}) {
	switch m := msg.(type) {
	case MsgUserBalanceChanged:
		println("user", m.userID, "balance was changed to", m.balance)
	case MsgEventChanged:
		println("event", m.eventID, "was changed")
	default:
		println("unknown message:", msg)
	}
}

/*
Required output:
	user "user-1" balance was changed to "1000"
	event "event-1" was changed
	unknown message: unknown
*/
func TestProcessMessage(t *testing.T) {
	processMessage(MsgUserBalanceChanged{"user-1", "1000"})
	processMessage(MsgEventChanged{"event-1"})
	processMessage("unknown")
}
