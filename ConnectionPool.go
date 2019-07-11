package main

import "fmt"

func main(){
	conp :=ConnectionPool{}

	id1 := "1"
	id2 := "2"
	id3 := "3"

	conp.getConnection(id1)
	fmt.Println(conp)
	conp.getConnection(id2)
	fmt.Println(conp)
	conp.getConnection(id3)
	fmt.Println(conp)

}

type ConnectionPool struct {
	connections map[Connection]string
}

type Connection struct {
	id      string
	timeout string
}

func (c *ConnectionPool) getConnection(id string)Connection{

}

