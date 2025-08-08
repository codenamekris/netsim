package main

import (
	"fmt"
	"strconv"
	"strings"
	//"go/printer"
)

/*
///////////////// TO-DO ////////////////

- Create a MAC Address generator for a unique assignment
- Implement the builder pattern or a function that creates a switch 
	- 
- Come up with a more meaningful way to represent a connection between two switches
- Status for switch interface is down by default. but if there is a an established connection, the status is up
- create a MAC Address generator 

*/
///// STRUCTS ////// 
type Topology struct{

}

type Device struct {
	DeviceType string 
	power bool
	name  string
	critical bool
	Iface *Interface
}

type Interface struct {
	PortId string
	IPAddress string
	MAC string
	ConnectedTo *Interface
	Status string 
	parentDevice *Device
}

type Switch struct{
	Device *Device
	MacTable MacTable
}

type Frame struct{
	srcMac string
	destMac string
	payload string
}

type MacEntry struct{
	srcMac string
	portId string
}

type MacTable []MacEntry

/////// METHODS ///////

func (intf *Interface) showConnection() string{
	port := intf.ConnectedTo // this should return another interface
	return port.PortId
}

func (intf *Interface) SetStatus(){ 
	// if a connection exists, set the status to Connected 
	if intf.ConnectedTo == nil {
		intf.Status = "Not Connected"
	} else  {
		intf.Status = "Connected"
	} 
}

func (intf *Interface) checkStaus() string{
	return intf.Status
}

func connect(intfA *Interface, intfB *Interface){
	intfA.ConnectedTo = intfB
	intfB.ConnectedTo = intfA
	intfA.SetStatus()
	intfB.SetStatus()

	/// return a map of connections????
}
func createFrame(intf *Interface, destMac string, payload string) Frame {
	frame := Frame {
		srcMac: intf.MAC,
		destMac: destMac,
		payload: payload,
	}
	
	return frame
}

func (intf *Interface) readFrame(frame *Frame) string{
	// update frame interface tag for the mac table 
	return "frame recieved" 
}

func (sw *Switch) processFrame(srcIntf *Interface, frame *Frame) {
	if frame.srcMac == "FF:FF:FF"{
		/// call ARP
	} else{
		//update MAC table 
		macEntry := MacEntry {srcMac: frame.srcMac, portId: srcIntf.PortId}
		sw.MacTable = append(sw.MacTable, macEntry)
	}
}

func (swi *Switch) updateMacTable(srcMac string, ){

}

func (intf *Interface) transmitFrame(frame *Frame, srcIntf *Interface, destIntf *Interface){
	if srcIntf.checkStaus() == "Connected" && destIntf.checkStaus() == "Connected"{
		destIntf.readFrame(frame)
	}
}


func isValidIp(ip string) bool {
	ipSlice := strings.Split(ip, ".")

	//splice should not contain more than 4 values
	if len(ipSlice) != 4 {
		fmt.Println("Ip Address invalid")
		return false
	}

	//an ipv4 octet range is 0-255
	for i := range ipSlice {
		octet, err := strconv.Atoi(ipSlice[i])

		if err != nil {
			fmt.Println("Invalid number:", err)
			return false
		}

		if octet > 255 {
			fmt.Println("IP Address invalid. Value cannot exceed 255")
			return false
		}
	}

	return true
}

/*func assignIp(device *Device, ip string) string {
	if isValidIp(ip){
		device.ip = ip
	} else {
		return "invalid ip"
	}
	
	return device.ip
} */

func assignSubnet(){
	// assigns subnet to a device. check to see if subnet is valid first
}

func isValidSubnet() {
	// checks to see if subnet is valid
}

var devices = map[string] Device{
	
}

func main() {
	sw1 := Device{
		name: "switch A", 
		critical: false, 
		Iface: &Interface{
					PortId: "f01", 
					IPAddress: "129.168.1.0",
					MAC: "AAA", 
					ConnectedTo: nil,
					Status: "Not Connected",
				},
	}

	sw2:= Device{
		name: "switch B", 
		critical: false, 
		Iface: &Interface{
					PortId: "f01", 
					IPAddress: "129.168.1.0",
					MAC: "BBB", 
					ConnectedTo: nil,
					Status: "Not Connected",
				},
	}

	var sw3 *Switch
	sw3 = new(Switch)

	sw1int := sw1.Iface
	sw2int := sw2.Iface
	connect(sw1int, sw2int)
	fmt.Println(sw1int.showConnection())
	sw1.Iface.SetStatus()
	fmt.Println(sw2int.showStaus())
}
