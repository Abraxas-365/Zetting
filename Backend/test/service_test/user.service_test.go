package service_test

import (
	// m "mongoCrud/modeles"
	service "mongoCrud/services"
	"testing"
	//"time"
)

// func TestCreate(t *testing.T) {
// 	user := m.User{
// 		Name:    "Luisf",
// 		Email:   "luisf@gmail.com",
// 		Created: time.Now(),
// 		Updated: time.Now(),
// 	}
// 	if err := service.Create(user); err != nil {
// 		t.Error("Fallo la prueba de crear")
// 	} else {
// 		t.Log("Finaliso con exito")
// 	}
// }

// func TestRead(t *testing.T) {
// 	users, err := service.Read()

// 	if err != nil {
// 		t.Error("Fallo la prueba de crear")
// 		t.Fail()
// 	}
// 	if len(users) == 0 {
// 		t.Error("No retorno datos")
// 		t.Fail()
// 	}
// 	t.Log("Funco")

// }

// func TestUpdate(t *testing.T) {

// 	user := m.User{
// 		Name:    "Nuevo",
// 		Email:   "nuevo@gmail.com",
// 		Created: time.Now(),
// 		Updated: time.Now(),
// 	}
// 	if err := service.Update(user, "000000000000000000000000"); err != nil {
// 		t.Error("Fallo la prueba la actualizar")
// 		t.Fail()
// 	} else {
// 		t.Log("todo guicci")
// 	}
// }

func TestDelete(t *testing.T) {
	if err := service.DeleteUser("000000000000000000000000"); err != nil {
		t.Error("Fallo borrando")
		t.Fail()
	} else {
		t.Log("fucno")
	}
}
