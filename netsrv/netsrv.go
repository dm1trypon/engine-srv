package netsrv

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	logger "github.com/dm1trypon/easy-logger"
	"github.com/dm1trypon/engine-srv/config"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// LC - logging's category
const LC = "NETSRV"

type wsConnData struct {
	wsStates      *websocket.Conn
	wsInteractive *websocket.Conn
}

type ConnStatus struct {
	Nickname    string
	IsConnected bool
}

type NetSRV struct {
	wsCodes   map[string]int
	server    *http.Server
	router    *mux.Router
	upgrader  websocket.Upgrader
	clients   map[string]*wsConnData
	msgs      chan []byte
	interConn chan ConnStatus
	stateConn chan ConnStatus
}

func (n *NetSRV) startServer() {
	logger.Debug(LC, "WS Server setup")
	port := config.ConfigInst.GetConfigData().Net.Port

	n.server = &http.Server{
		Addr:              ":" + strconv.Itoa(port),
		Handler:           n.router,
		TLSConfig:         nil,
		ReadTimeout:       1000 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      1000 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      map[string]func(*http.Server, *tls.Conn, http.Handler){},
		ConnState: func(net.Conn, http.ConnState) {
		},
		ErrorLog:    &log.Logger{},
		BaseContext: func(net.Listener) context.Context { return context.Background() },
		ConnContext: func(ctx context.Context, c net.Conn) context.Context { return context.Background() },
	}

	logger.Info(LC, fmt.Sprint("WS Server started at port ", port))

	if err := n.server.ListenAndServe(); err != nil {
		logger.Critical(LC, fmt.Sprint("Failed starting WS Server: ", err.Error()))
		os.Exit(-1)
	}
}

func (n *NetSRV) setUpgrader() {
	logger.Debug(LC, "Upgrader setup")
	n.upgrader = websocket.Upgrader{
		HandshakeTimeout: 0,
		ReadBufferSize:   0,
		WriteBufferSize:  0,
		WriteBufferPool:  nil,
		Subprotocols:     []string{},
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: false,
	}
}

func (n *NetSRV) setWSCodes() {
	n.wsCodes = map[string]int{
		"CloseNormalClosure":           1000,
		"CloseGoingAway":               1001,
		"CloseProtocolError":           1002,
		"CloseUnsupportedData":         1003,
		"CloseNoStatusReceived":        1005,
		"CloseAbnormalClosure":         1006,
		"CloseInvalidFramePayloadData": 1007,
		"ClosePolicyViolation":         1008,
		"CloseMessageTooBig":           1009,
		"CloseMandatoryExtension":      1010,
		"CloseInternalServerErr":       1011,
		"CloseServiceRestart":          1012,
		"CloseTryAgainLater":           1013,
		"CloseTLSHandshake":            1015,
	}
}

func (n *NetSRV) setRouter() {
	logger.Debug(LC, "Router setup")
	n.router = mux.NewRouter()
	n.router.HandleFunc("/states", n.onStates)
	n.router.HandleFunc("/interactive", n.onInteractive)
}

func (n *NetSRV) GetMsgsChan() chan []byte {
	return n.msgs
}

func (n *NetSRV) GetStateChan() chan ConnStatus {
	return n.stateConn
}

func (n *NetSRV) GetIntChan() chan ConnStatus {
	return n.interConn
}

func (n *NetSRV) onInteractive(w http.ResponseWriter, r *http.Request) {
	loggerMask := fmt.Sprint("[INTERACTIVE][", r.RemoteAddr, "]")
	wsConn, err := n.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(LC, fmt.Sprint(loggerMask, " Upgrader error: ", err.Error()))
		return
	}

	defer wsConn.Close()

	nickname := r.URL.Query().Get("nickname")
	loggerMask = fmt.Sprint(loggerMask, "[", nickname, "]")
	if len(nickname) < 1 {
		logger.Error(LC, fmt.Sprint(loggerMask, " Empty param 'nickname'"))
		return
	}

	if _, ok := n.clients[nickname]; ok {
		logger.Error(LC, fmt.Sprint(loggerMask, " WS client is already exist"))
		return
	}

	logger.Info(LC, fmt.Sprint(loggerMask, " Client connected"))

	wsConnData := &wsConnData{
		wsInteractive: wsConn,
		wsStates:      nil,
	}

	n.clients[nickname] = wsConnData
	n.interConn <- ConnStatus{Nickname: nickname, IsConnected: true}
	n.reader(wsConn, loggerMask)
}

func (n *NetSRV) onStates(w http.ResponseWriter, r *http.Request) {
	loggerMask := fmt.Sprint("[STATES][", r.RemoteAddr, "]")
	wsConn, err := n.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(LC, fmt.Sprint(loggerMask, " Upgrader error: ", err.Error()))
		return
	}

	defer wsConn.Close()

	nickname := r.URL.Query().Get("nickname")
	loggerMask = fmt.Sprint(loggerMask, "[", nickname, "]")
	if len(nickname) < 1 {
		logger.Error(LC, fmt.Sprint(loggerMask, " Empty param 'nickname'"))
		return
	}

	if _, ok := n.clients[nickname]; !ok {
		logger.Error(LC, fmt.Sprint(loggerMask, " WS client is not authorized"))
		return
	}

	if n.clients[nickname].wsStates != nil {
		logger.Error(LC, fmt.Sprint(loggerMask, " WS client is already connected"))
		return
	}

	logger.Info(LC, fmt.Sprint(loggerMask, " Client connected"))
	n.clients[nickname].wsStates = wsConn
	n.stateConn <- ConnStatus{Nickname: nickname, IsConnected: true}
	n.reader(wsConn, loggerMask)
}

func (n *NetSRV) deleteClient(wsConn *websocket.Conn, loggerMask string) {
	for nickname, wsConnData := range n.clients {
		if wsConnData.wsInteractive == wsConn || wsConnData.wsStates == wsConn {
			logger.Error(LC, fmt.Sprint(loggerMask,
				" There was a client disconnection, removing all child connections."))

			delete(n.clients, nickname)
			n.interConn <- ConnStatus{Nickname: nickname, IsConnected: false}
			n.stateConn <- ConnStatus{Nickname: nickname, IsConnected: false}
			return
		}
	}
}

func (n *NetSRV) reader(wsConn *websocket.Conn, loggerMask string) {
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			logger.Error(LC, fmt.Sprint(loggerMask, " WS reading message error: ", err.Error()))
			n.deleteClient(wsConn, loggerMask)
			break
		}

		if n.isWSError(err) {
			logger.Warning(LC, fmt.Sprint(loggerMask, " WS handle error: ", err.Error()))
			n.deleteClient(wsConn, loggerMask)
			break
		}

		logger.Debug(LC, fmt.Sprint(loggerMask, " RECV: ", message))
		n.GetMsgsChan() <- message
	}
}

func (n *NetSRV) isWSError(err error) bool {
	for _, code := range n.wsCodes {
		if websocket.IsUnexpectedCloseError(err, code) {
			return true
		}
	}

	return false
}

func (n *NetSRV) Create() *NetSRV {
	logger.Info(LC, "Creating NetSRV module")

	n.msgs = make(chan []byte)
	n.interConn = make(chan ConnStatus)
	n.stateConn = make(chan ConnStatus)
	n.clients = make(map[string]*wsConnData)

	n.setWSCodes()
	n.setUpgrader()
	n.setRouter()
	n.startServer()

	return n
}
