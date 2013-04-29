package main

import "flag"
import "net"
import "bufio"
import "fmt"
import "os"


func main(){
	number_of_lines_string := flag.String("lines", "85", "Number of lines to request") 
	number_of_loops := flag.Int("loops", 10000, "Number of times to loop")
	network_interface := flag.String("interface", ":8080", "Interface to attach to")
	flag.Parse()

//	number_of_lines_string = *number_of_lines_string + "\n"
	number_of_lines := []byte(*number_of_lines_string + "\n")

	conn, err := net.Dial("tcp6", *network_interface)

	if err != nil {
		fmt.Println("Had Error", err)
		os.Exit(1)
		}

	reader := bufio.NewReader(conn)

	for i := 0; i <= *number_of_loops; i++ {
		_, err = conn.Write(number_of_lines)
	}

	for i := 0; i <= *number_of_loops; i++ {
		_, _ = reader.ReadString('\n')
		//fmt.Println(line)
	}

}
