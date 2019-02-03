package main

import (	
	"fmt"
	"net/http"
	"time"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type Characters struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func main() {
	
	p := fmt.Println

	const publicKey = "20e1f7cf9171e70086f642f59e2a2df0"
	const privateKey = "6cd3d6dae69b6b3a07e54d16f2b9bd3861d16209"

	tsInt := time.Now().Unix()
	ts := strconv.FormatInt(tsInt, 10)
	
	hash := GetMD5Hash(ts, privateKey, publicKey)

	// Start
	var op string
	println("Seleccionar opcción:")
	println("[1] Buscar personaje")
	println("[2] Listar personajes")
	fmt.Scanln(&op)

	if op == "1" {
		p("Escribir el nombre del personaje")
		var name string
		fmt.Scanln(&name)
		URL := "http://gateway.marvel.com/v1/public/characters?name="+name+"&ts="+ts+"&apikey="+publicKey+"&hash="+hash
		res, err := http.Get(URL)
		if err != nil{
			fmt.Printf("Error %s", err)
		} else{
			body, _ := ioutil.ReadAll(res.Body)
			data := fmt.Sprintf("%s",body)
		}
	}
	if op == "2"{
		URL := "http://gateway.marvel.com/v1/public/characters?ts="+ts+"&apikey="+publicKey+"&hash="+hash
		res, err := http.Get(URL)
		if err != nil{
			fmt.Printf("Error %s", err)
		} else{
			data, _ := ioutil.ReadAll(res.Body)
			p(string(data))
		}
	} else if op != "1" && op != "2"{
		println("Opcción inválida")
	}
}

// Generate Hash
func GetMD5Hash(ts, privateKey, publicKey string) string {
    hasher := md5.New()
    hasher.Write([]byte(ts + privateKey + publicKey))
    return hex.EncodeToString(hasher.Sum(nil))
}