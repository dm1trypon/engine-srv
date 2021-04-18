package manager

import (
	logger "github.com/dm1trypon/easy-logger"
	"github.com/dm1trypon/engine-srv/engine"
	"github.com/dm1trypon/engine-srv/jsonvalidator"
	"github.com/dm1trypon/engine-srv/netsrv"
)

// LC - logging's category
const LC = "MANAGER"

type Manager struct {
	schemasPaths map[string]string
	engineInst   *engine.Engine
	netSRVInst   *netsrv.NetSRV
	msgs         chan []byte
	interConn    chan netsrv.ConnStatus
	stateConn    chan netsrv.ConnStatus
}

func (m *Manager) Create() *Manager {
	logger.Info(LC, "Creating Manager module")

	m.initSchemas()

	m.netSRVInst = new(netsrv.NetSRV).Create()
	m.engineInst = new(engine.Engine).Create()

	m.msgs = m.netSRVInst.GetMsgsChan()
	m.interConn = m.netSRVInst.GetIntChan()
	m.stateConn = m.netSRVInst.GetStateChan()

	m.startConnEvents()
	return m
}

func (m *Manager) initSchemas() {
	m.schemasPaths = map[string]string{
		"move":         "./manager/msg_schemas/move.schema.json",
		"mouse":        "./manager/msg_schemas/mouse.schema.json",
		"choiceweapon": "./manager/msg_schemas/choiceweapon.schema.json",
	}
}

func (m *Manager) startConnEvents() {
	go func() {
		for {
			m.intConnEvent(<-m.interConn)
		}
	}()

	go func() {
		for {
			m.stateConnEvent(<-m.stateConn)
		}
	}()

	go func() {
		for {
			m.msgsConnEvent(<-m.msgs)
		}
	}()
}

func (m *Manager) intConnEvent(connStatus netsrv.ConnStatus) {
	if connStatus.IsConnected {
		m.engineInst.NewPlayer(connStatus.Nickname)
		return
	}

	m.engineInst.RemovePlayer(connStatus.Nickname)
}

func (m *Manager) stateConnEvent(connStatus netsrv.ConnStatus) {
	// coming soon
}

func (m *Manager) msgsConnEvent(data []byte) {
	for key, path := range m.schemasPaths {
		if !jsonvalidator.Check(data, path) {
			continue
		}

		if key == "move" {
			m.onMove(data)
		} else if key == "mouse" {
			m.onMouse(data)
		} else if key == "choiceweapon" {
			m.onChoiceWeapon(data)
		}
	}
}

func (m *Manager) onMove(data []byte) {

}

func (m *Manager) onMouse(data []byte) {

}

func (m *Manager) onChoiceWeapon(data []byte) {

}
