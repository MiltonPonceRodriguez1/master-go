package models

type Man struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	alive     bool
}

func (this *Man) Breathe()       { this.breathing = true }
func (this *Man) Think()         { this.thinking = true }
func (this *Man) Eat()           { this.eating = true }
func (this *Man) Gender() string { return "Hombre" }
func (this *Man) IsAlive() bool  { return true }
