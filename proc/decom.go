package proc

import (
	"fmt"
	"net"
)

func PacketListener(packet TelemetryPacket, channel chan []byte) {
	packetSize := GetPacketSize(packet)
	fmt.Printf("Packet size for port %d: %d\n", packet.Port, packetSize)

	// Listen over UDP
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", packet.Port))
	if err != nil {
		fmt.Printf("Error resolving UDP address: %v\n", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("Error listening on UDP: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Listening on port %d for telemetry packet...\n", packet.Port)

	// Receive data
	buffer := make([]byte, packetSize)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading from UDP: %v\n", err)
			continue
		}

		if n == packetSize {

			channel <- buffer[:n] // Send data over channel
		} else {
			fmt.Printf("Received packet of incorrect size. Expected: %d, Received: %d\n", packetSize, n)
		}
	}
}

// TODO: Placeholder consumer function. Remove once feature for publishing data is complete
func TestReceiver(channel chan []byte) {
	i := 0
	for {
		data := <-channel
		fmt.Printf("Packet %d: %v\n", i, data)
		i++
	}
}
