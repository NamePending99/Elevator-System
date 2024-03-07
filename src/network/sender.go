package network

import (
	"bytes"
	"elevator/types"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/libp2p/go-reuseport"
)

const MSG_TIMEOUT = 2000

/*
 *	Send message using UDP protocol to the sepcified address
 */
func Send(addr string, msgType types.MsgTypes, msgContent []byte) {
	if addr == "" {
		return
	}

	receiverPort := strings.Split(addr, ":")[1]
	packetConnection, err := reuseport.ListenPacket("udp4", fmt.Sprintf(":%s", receiverPort))
	if err != nil {
		panic(err)
	}
	defer packetConnection.Close()

	resolvedAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		panic(err)
	}

	/*
	 * Add type message type as an integer at the start of the byte array
	 */
	sizeofInt := 4
	msgTypeBuffer := make([]byte, sizeofInt)

	binary.BigEndian.PutUint32(msgTypeBuffer, uint32(msgType))

	msgAndTypeBuffer := [][]byte{msgTypeBuffer, msgContent}
	seperator := []byte("")

	msgWithType := bytes.Join(msgAndTypeBuffer, seperator)

	_, err = packetConnection.WriteTo(msgWithType, resolvedAddr)
	if err != nil {
		panic(err)
	}
}

/*
 * Send message to all nodes in the network
 * Resend if no reply is received within timeout
 */
func SecureSend(
	initialAddr string,
	msgType types.MsgTypes,
	msg []byte,
	replyReceived <-chan bool,
	updateAddr <-chan string,
) {
	addr := initialAddr

	Send(addr, msgType, msg)

	msgTimedOut := time.NewTicker(MSG_TIMEOUT * time.Millisecond)

	for {
		select {
		case newAddr := <-updateAddr:
			addr = newAddr

		case <-replyReceived:
			msgTimedOut.Stop()
			return

		case <-msgTimedOut.C:
			Send(addr, msgType, msg)

		default:
			/*
			 * Do nothing
			 */
			continue
		}
	}
}
