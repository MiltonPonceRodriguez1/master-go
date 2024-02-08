package interfaces_exercises

import (
	"fmt"
	"master-go/interfaces"
)

func HumansBreathing(human interfaces.Human) {
	human.Breathe()
	fmt.Printf("Soy un/a %s, y estoy respirando\n", human.Gender())
}

func LivingBeingBreathing(human interfaces.Human) {
	human.Breathe()
	fmt.Printf("Soy un/a %s, y estoy respirando; Esta vivo %t\n", human.Gender(), human.IsAlive())
}
