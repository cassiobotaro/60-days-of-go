package main

import "fmt"

// Inspired by this excellent video: https://www.youtube.com/watch?v=wfMtDGfHWpA

// Barker is something that bark
type Barker struct {
	Name string
}

// Bark is a noise
func (b Barker) Bark() {
	fmt.Printf("Woof, I'm am %s\n", b.Name)
}

// Meower is something that meow
type Meower struct {
	Name string
}

// Meow is a noise
func (b Meower) Meow() {
	fmt.Printf("Meow, I'm am %s\n", b.Name)
}

// Pooper is something that poop
type Pooper struct{}

// Poop shit happens!
func (p Pooper) Poop() {
	fmt.Println("💩💩")
}

// Driver is something that drive
type Driver struct {
	Position int
	Speed    int
}

//Drive is a move
// Your receiver is a pointer because have to change the value
func (d *Driver) Drive() {
	d.Position += d.Speed
	fmt.Printf("New position: %d\n", d.Position)
}

// Cleaner is something that clean
type Cleaner struct{}

// Clean is an action
func (c Cleaner) Clean() {
	fmt.Println("Clean")
}

// Killer is something that kill
type Killer struct{}

// Kill is an action
func (k Killer) Kill() {
	fmt.Println("I kill you!")
}

// Dog is a barker that poop
type Dog struct {
	Barker
	Pooper
}

// Cat is meower that poop
type Cat struct {
	Meower
	Pooper
}

// CleaningRobot is a driver that clean
type CleaningRobot struct {
	Driver
	Cleaner
}

// MurderRobot is a driver that kill
type MurderRobot struct {
	Driver
	Killer
}

// MurderRobotDog WAT?
type MurderRobotDog struct {
	MurderRobot
	Barker
}

func main() {
	// a barker can bark
	b := Barker{Name: "Rex"}
	b.Bark()
	// a meower can meow
	m := Meower{Name: "Kitty"}
	m.Meow()
	// a pooper can poop
	p := Pooper{}
	p.Poop()
	// a driver can drive
	d := Driver{Position: 0, Speed: 10}
	d.Drive()
	// a cleaner can clean
	c := Cleaner{}
	c.Clean()
	// a killer can kil
	k := Killer{}
	k.Kill()
	// a dog can bark and poop
	dog := Dog{}
	dog.Name = "Fox"
	dog.Bark()
	dog.Poop()
	// a Murder Robot Dog
	mrd := MurderRobotDog{
		MurderRobot{Driver{Position: 0, Speed: 10}, Killer{}},
		Barker{Name: "Murder robot dog!"}}
	// can bark
	mrd.Bark()
	// can drive
	mrd.Drive()
	// can kill and don't poop!
	mrd.Kill()
}
