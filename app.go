package main

import (	
	"fmt"
	"net/http"
	"time"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"io/ioutil"
	// "encoding/json"
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
	println("[1] Buscar personajes")
	println("[2] Listar personajes")
	print("> ")
	fmt.Scanln(&op)

	if op == "1" {
		p("[1] Buscar un personaje")
		p("[2] Empiezan por")
		var charOp string
		print("> ")
		fmt.Scanln(&charOp)

		if charOp == "1" {
			// Un solo personaje
			p("Escribir el nombre del personaje")
			var name string
			fmt.Scanln(&name)
			// URL
			URL := "http://gateway.marvel.com/v1/public/characters?name="+name+"&ts="+ts+"&apikey="+publicKey+"&hash="+hash
			// Peticion GET
			res, err := http.Get(URL)
			if err != nil{
			fmt.Printf("Error %s", err)
			} else{
			data, _ := ioutil.ReadAll(res.Body)
			p(string(data))
			}
		} else if charOp == "2"{
			// Empeizan por
			print("> ")
			startWith := StartWith()
			p("Ordenado por:")
			p("[1] Nombre (Ascendente)")
			p("[2] Nombre (Descendente)")
			p("[3] Modificación (Ascendente)")
			p("[4] Modificación (Descendente)")
			var orderOp string
			print("> ")
			fmt.Scanln(&orderOp)
	
			if orderOp == "1" {
				order := "name"
				URL := "http://gateway.marvel.com/v1/public/characters?nameStartsWith="+startWith+"&orderBy="+order+"&ts="+ts+"&apikey="+publicKey+"&hash="+hash
				// Peticion GET
				res, err := http.Get(URL)
				if err != nil{
				fmt.Printf("Error %s", err)
				} else{
				data, _ := ioutil.ReadAll(res.Body)
				p(string(data))
				}
			} else if orderOp == "2"{
				order := "-name"
				URL := "http://gateway.marvel.com/v1/public/characters?nameStartsWith="+startWith+"&orderBy="+order+"&ts="+ts+"&apikey="+publicKey+"&hash="+hash
				// Peticion GET
				res, err := http.Get(URL)
				if err != nil{
				fmt.Printf("Error %s", err)
				} else{
				data, _ := ioutil.ReadAll(res.Body)
				p(string(data))
				}
			} else if orderOp == "3"{
				order := "modified"
				URL := "http://gateway.marvel.com/v1/public/characters?nameStartsWith="+startWith+"&orderBy="+order+"&ts="+ts+"&apikey="+publicKey+"&hash="+hash
				// Peticion GET
				res, err := http.Get(URL)
				if err != nil{
				fmt.Printf("Error %s", err)
				} else{
				data, _ := ioutil.ReadAll(res.Body)
				p(string(data))
				}
			} else if orderOp == "4"{
				order := "-modified"
				URL := "http://gateway.marvel.com/v1/public/characters?nameStartsWith="+startWith+"&orderBy="+order+"&ts="+ts+"&apikey="+publicKey+"&hash="+hash
				// Peticion GET
				res, err := http.Get(URL)
				if err != nil{
				fmt.Printf("Error %s", err)
				} else{
				data, _ := ioutil.ReadAll(res.Body)
				p(string(data))
				}
			} else if orderOp != "1" && orderOp != "2" && orderOp != "3" && orderOp != "4"{	
				p("Opcción inválida")}
		} else if charOp != "1" && charOp != "2"{
			p("Opcción inválida")}		
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

func StartWith () string {
	var text string
	fmt.Scanln(&text)
	return text
}