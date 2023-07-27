package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         Weapon
}

type Weapon struct {
	Attack int
}

func (h Hero) CountDamage() int {
	rand.Seed(time.Now().UnixNano())
	crit_dmg := 0
	if rand.Intn(2) == 0 {
		crit_dmg += h.CriticalDamage
	}
	total_damage := h.BaseAttack + h.Weapon.Attack + crit_dmg

	return total_damage
}

// soal 2
func (h *Hero) isAttackedBy(attacker Hero) {
	total_damage_received := attacker.CountDamage() - h.Defence

	if attacker.CountDamage() >= h.Defence {
		h.HealthPoint -= total_damage_received
	}

}

func Battle(attacker, defender Hero) {
	fmt.Printf("Battle between: %s and %s\n", attacker.Name, defender.Name)
	fmt.Printf("%s attacks %s\n", attacker.Name, defender.Name)
	defender.isAttackedBy(attacker)
	fmt.Printf("%s's HealthPoint after attacked: %d\n", defender.Name, defender.HealthPoint)
}

func main() {
	var weapon1 Weapon
	weapon1 = Weapon{
		Attack: 5,
	}

	var hero1 Hero
	hero1 = Hero{
		Name:           "Mortred",
		BaseAttack:     8,
		Defence:        5,
		CriticalDamage: 7,
		HealthPoint:    20,
		Weapon:         weapon1,
	}

	var hero2 Hero
	hero2 = Hero{
		Name:           "Magina",
		BaseAttack:     6,
		Defence:        8,
		CriticalDamage: 5,
		HealthPoint:    25,
		Weapon:         weapon1,
	}

	fmt.Println("Hero 1 : ", hero1)
	fmt.Println("Hero 2 : ", hero2)
	fmt.Println("")
	fmt.Println("Battle:")
	Battle(hero1, hero2)
}
