package models

type Women struct {
	Man
	isMother bool
}

func (this *Women) Breathe()       { this.breathing = true }
func (this *Women) Think()         { this.thinking = true }
func (this *Women) Eat()           { this.eating = true }
func (this *Women) Gender() string { return "Mujer" }
