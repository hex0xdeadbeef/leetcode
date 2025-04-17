package main

import (
	"unsafe"
)

const (
	cacheLineSize = 128
)

type Person struct {
	person
	_ cacheLinePad
}

type person struct {
	name    string
	age     int
	isAdult bool
}

func (p *Person) BalanceAge() {
	p.balanceAge()
}

func (p *person) balanceAge() {
	p.age++
	p.age--
}

type cacheLinePad [cacheLineSize - unsafe.Sizeof(person{})]byte

type PlainPerson struct {
	person
}

func (p *PlainPerson) BalanceAge() {
	p.balanceAge()
}
