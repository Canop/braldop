package main

/*
Objet exportable en json

*/

type Vue struct {
	Time     int64 // secondes depuis 1970. Une date à 0 signifie que l'objet est vide ou invalide
	Voyeur   uint  // id du braldun. Un id à 0 signifie que l'objet est vide ou invalide
	XMin     int16
	XMax     int16
	YMin     int16
	YMax     int16
	Bralduns []*VueBraldun
}

func NewVue() *Vue {
	vue := new(Vue)
	vue.Bralduns = make([]*VueBraldun, 0, 5)
	return vue
}
