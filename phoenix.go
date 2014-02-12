package main

import "time"
import "fmt"
import "os/exec"
import "net"
import "strconv"
import "log"

const MY_IP = "127.0.0.1"
const COM_PORT = "20013"


func send_ping(value int){
	udpAddr,err := net.ResolveUDPAddr("udp", MY_IP+":"+COM_PORT) //(*UDPAddr, error)
	if err != nil{
		fmt.Println(" error adr")	
	}
	sock,err := net.DialUDP("udp", nil, udpAddr) //(*UDPConn, error)
	if err != nil{
		fmt.Println(" error sock")	
	}
	ping := string(value)
	sock.Write([]byte(ping))
}
func read_ping(value_chan chan int){
	fmt.Println(" reading ping...")

	udpAddr,err := net.ResolveUDPAddr("udp", COM_PORT)
	if err != nil{
		fmt.Println(" error adr")	
	}
	ln,err := net.ListenUDP("udp", udpAddr)
	if err != nil{
		fmt.Println(" error ln")
	}
	
	b := make([]byte,1024)
	ln.Read(b)
	ping:=string(b)
	ping_int,_ := strconv.Atoi(ping)
	value_chan<-ping_int
	fmt.Println("UDP got: ")
	fmt.Println(ping_int)
}

func clone(){
	fmt.Println("Cloning ...")
	//exec.Command("go run", "/home/student/phoenix.go")	
	//exec.Command("echo","spawned")

	cmd := exec.Command("mate-terminal","-x","go" ,"run","/home/student/phoenix.go")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func main(){

	UDPChan := time.NewTicker(time.Millisecond * 40).C
	tickChan := time.NewTicker(time.Millisecond * 100).C
	pingChan := time.NewTicker(time.Millisecond * 100).C
	testChan := time.NewTicker(time.Millisecond * 500).C
	pongChan := make(chan int, 1)
	value_chan :=make(chan int, 1)
 
	primary := false
	take_over :=false	
	value := 0
	
	fmt.Println("init ok") //debug

	for{
		
		if primary{
			for {
		       select {
	 	       case <- pingChan:
					send_ping(value)
	    	   case <- tickChan:
					fmt.Println(value)
					value++
				}
			}		
		} 

		if !primary && take_over{
			fmt.Println("I've got this!")
			primary = true
			take_over = false			
			clone()
		}
						
		if !primary && !take_over{
			tick :=0
			ping :=0
			cont :=true
			fmt.Println("!prim && !take")
			for(cont) {
		       select {
			   case <-UDPChan:
					fmt.Println("UDP")
					read_ping(value_chan)
	 	       case value=<- pongChan:
					fmt.Println("ping")
					ping ++
	    	    case <- tickChan:
					fmt.Println("tick")
					tick ++
	    	    case <- testChan:
					if ping < tick-3{
						fmt.Println("OMG! Takeover!!!") //debug
						take_over = true
						cont = false
					}
					
					tick =0
					ping =0
							
	    	   }//select
	    	}//for
		}//if
	}//for
}//main

