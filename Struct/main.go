package main

import (
	"fmt"
	"reflect"
)

// việc tận dụng struct và con trỏ sẽ tạo ra 1 hiệu suất rất tốt

type A struct {
	Name string
	Age  int
}

// update profile
func updateProfile(profile A, newAge int) A {
	profile.Age = newAge
	return profile
}

func (p *A) UpdateAge(newAge int) {
	p.Age = newAge
}

// update for pointer
func updateProfileOnline(p *A, newAge int) {
	p.Age = newAge
}

func main() {
	//1.profileGao

	profileGao := A{
		Name: "Gao",
		Age:  30,
	}

	fmt.Println(profileGao.Age)
	fmt.Println("Type profileGao: ", reflect.TypeOf(profileGao))

	// 2.profileKen

	var profileKen A
	profileKen.Name = "Ken"
	profileKen.Age = 28

	fmt.Println(profileKen.Name)
	fmt.Println("Type profileKen: ", reflect.TypeOf(profileKen))

	// 3. Update profileGao's ag
	profileGao = updateProfile(profileGao, 31)
	fmt.Println("Updated profileGao's age:", profileGao.Age)

	profileGao.UpdateAge(32)
	fmt.Println("Updated profileGao's age using method:", profileGao.Age)

	// 4. profileGao online
	profileGaoOnline := new(A)
	profileGaoOnline.Name = "Gao Online"
	profileGaoOnline.Age = 25

	fmt.Println(profileGaoOnline.Name)
	fmt.Println("Type profileGaoOnline: ", reflect.TypeOf(profileGaoOnline))

	// 5. Update profileGaoOnline's age
	updateProfileOnline(profileGaoOnline, 26)
	fmt.Println("Updated profileGaoOnline's age:", profileGaoOnline.Age)

	profileGaoOnline.UpdateAge(27)
	fmt.Println("Updated profileGaoOnline's age using method:", profileGaoOnline.Age)

}
