package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type All []struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	Dates        []string
}

type Date struct {
	Index []struct {
		Id    int
		Dates []string
	}
}
type Location struct {
	Index []struct {
		Id        int
		Locations []string
	}
}

type Artist []struct {
	Id           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string
}

var t *template.Template
var art All = Alls()

func init() {
	tmp, err := template.ParseGlob("assets/*.html")
	if err != nil {
		fmt.Println("owibka pri parsinge")
		return
	}
	t = template.Must(tmp, err)
}

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", grupieBackGround)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func grupieBackGround(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {

	}
	switch {
	case r.Method == "GET":

		t.ExecuteTemplate(w, "udemy.html", art)

	case r.Method == "POST":
		id := r.FormValue("id")
		idInt, _ := strconv.Atoi(id)
		switch {
		case idInt != 0:
			for _, k := range art {

				if k.Id == idInt {
					t.ExecuteTemplate(w, "index.html", k)
				}
			}
		case r.FormValue("back") != "":

			t.ExecuteTemplate(w, "udemy.html", art)
			fmt.Println(r.FormValue("search"))
		case r.FormValue("search") != "":
			search := r.FormValue("search")
			for _, i := range art {
				fmt.Println(i, search)
				// case i.CreationDate == search:
				switch {
				case i.Name == search:
					t.ExecuteTemplate(w, "search.html", i)
				case i.FirstAlbum == search:
					t.ExecuteTemplate(w, "search.html", i)
					for _, j := range i.Locations {
						if j == search {
							t.ExecuteTemplate(w, "search.html", i)
						}
					}
					for _, j := range i.Dates {
						if j == search {
							t.ExecuteTemplate(w, "search.html", i)
						}
					}
				}
			}
		}

	}
}

// func check(a, b string) bool {

// }
func Alls() All {

	dat := Dates()
	loc := Locations()
	art := Artists()
	t := 0
	for range art {
		t++
	}
	all := make(All, t)
	for _, i := range art {
		for _, j := range loc.Index {
			for _, k := range dat.Index {
				if i.Id == k.Id && i.Id == j.Id {
					all[i.Id-1].Id = i.Id
					all[i.Id-1].Name = i.Name
					all[i.Id-1].Image = i.Image
					all[i.Id-1].Members = i.Members
					all[i.Id-1].CreationDate = i.CreationDate
					all[i.Id-1].FirstAlbum = i.FirstAlbum
					all[i.Id-1].Locations = j.Locations
					all[i.Id-1].Dates = k.Dates

				}
			}
		}
	}
	return all
}

func Artists() Artist {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	// read json http response
	jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var jsonData Artist

	err = json.Unmarshal(jsonDataFromHttp, &jsonData) // here!

	if err != nil {
		panic(err)
	}


	return jsonData
}
func Locations() Location {

	url := "https://groupietrackers.herokuapp.com/api/locations"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("owibka pri location")

	}
	defer resp.Body.Close()
	locationByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("owibka pri readall")

	}
	var locationJson Location
	err = json.Unmarshal(locationByte, &locationJson)
	fmt.Println(err)
	return locationJson
}
func Dates() Date {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("owibka pri GET")

	}
	defer resp.Body.Close()
	dateByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("owibka pri READALL")
	}
	var DatesJson Date
	err = json.Unmarshal(dateByte, &DatesJson)
	if err != nil {
		fmt.Println("owibka pir UNMARSHAL")

	}
	return DatesJson
}