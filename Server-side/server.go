package main

import (
	"errors"
	"fmt"
	"math"
	"net"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}

type Response struct {
	Quo, Res int
}

type Math byte

func (m *Math) Add(args *Args, res *int) error {
	*res = args.A + args.B
	return nil
}

func (m *Math) Divide(args *Args, res *Response) error {
	if args.B == 0 {
		return errors.New("u're trying divide by zero")
	}
	res.Quo = args.A / args.B
	res.Res = args.A % args.B
	return nil
}

func (m *Math) Major(slice *[]int, res *int) error {
	var major = math.MinInt32
	for _, v := range *slice {
		if v > major {
			major = v
		}
	}
	*res = major
	return nil
}

func (m *Math) Minor(slice *[]int, res *int) error {
	var minor = math.MaxInt32
	for _, v := range *slice {
		if v < minor {
			minor = v
		}
	}
	*res = minor
	return nil
}

func checkingError(err error) {
	if err != nil {
		fmt.Printf("Error!! %v", err.Error())
		os.Exit(1)
	}
}

func main() {
	math := new(Math)
	rpc.Register(math)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":3233")
	checkingError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	defer listener.Close()
	checkingError(err)
	fmt.Println("Running in port 3233")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error!! %v", err.Error())
			continue
		}
		fmt.Printf("Connection stablished from %v\n", conn.RemoteAddr())
		go rpc.ServeConn(conn)
	}
}
