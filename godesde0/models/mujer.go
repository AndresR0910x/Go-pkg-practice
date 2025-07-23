package models

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() {m.respirando = true}
func (m *Mujer) Comiendo() {m.comiendo = true}
func (m *Mujer) Pensando() {m.pensando = true}
func (m *Mujer) Sexo() string {return "Mujer"}