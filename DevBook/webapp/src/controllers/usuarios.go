package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"cpf":  r.FormValue("cpf"),
	})

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(usuario))

	response, erro := http.Post("http://localhost:3000/usuarios", "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		log.Fatal(erro)
	}

	defer response.Body.Close()

}
