package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type Type struct {
	// Name of the type
	Name string `json:"name"`
	// The effective types, damage multiplize 2x
	EffectiveAgainst []string `json:"effectiveAgainst"`
	// The weak types that against, damage multiplize 0.5x
	WeakAgainst []string `json:"weakAgainst"`
}

type Pokemon struct {
	Number         string   `json:"Number"`
	Name           string   `json:"Name"`
	Classification string   `json:"Classification"`
	TypeI          []string `json:"Type I"`
	TypeII         []string `json:"Type II,omitempty"`
	Weaknesses     []string `json:"Weaknesses"`
	FastAttackS    []string `json:"Fast Attack(s)"`
	Weight         string   `json:"Weight"`
	Height         string   `json:"Height"`
	Candy          struct {
		Name     string `json:"Name"`
		FamilyID int    `json:"FamilyID"`
	} `json:"Candy"`
	NextEvolutionRequirements struct {
		Amount int    `json:"Amount"`
		Family int    `json:"Family"`
		Name   string `json:"Name"`
	} `json:"Next Evolution Requirements,omitempty"`
	NextEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Next evolution(s),omitempty"`
	PreviousEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Previous evolution(s),omitempty"`
	SpecialAttacks      []string `json:"Special Attack(s)"`
	BaseAttack          int      `json:"BaseAttack"`
	BaseDefense         int      `json:"BaseDefense"`
	BaseStamina         int      `json:"BaseStamina"`
	CaptureRate         float64  `json:"CaptureRate"`
	FleeRate            float64  `json:"FleeRate"`
	BuddyDistanceNeeded int      `json:"BuddyDistanceNeeded"`
}

// Move is an attack information. The
type Move struct {
	// The ID of the move
	ID int `json:"id"`
	// Name of the attack
	Name string `json:"name"`
	// Type of attack
	Type string `json:"type"`
	// The damage that enemy will take
	Damage int `json:"damage"`
	// Energy requirement of the attack
	Energy int `json:"energy"`
	// Dps is Damage Per Second
	Dps float64 `json:"dps"`
	// The duration
	Duration int `json:"duration"`
}

// BaseData is a struct for reading data.json
type BaseData struct {
	Types    []Type    `json:"types"`
	Pokemons []Pokemon `json:"pokemons"`
	Moves    []Move    `json:"moves"`
}

//getting errors in check function
func check(e error) {
	if e != nil {
		panic(e)
	}
} //end of the getting errors in check function

// sorting functions
type sortByBaseAttack []Pokemon

func (a sortByBaseAttack) Len() int           { return len(a) }
func (a sortByBaseAttack) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByBaseAttack) Less(i, j int) bool { return a[i].BaseAttack < a[j].BaseAttack }

// end of the sorting functions

func listHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("/list url:", r.URL)
	//reading json file
	dat, err := ioutil.ReadFile("data.json")
	check(err)
	//end of the reading json file

	//copying to json variables to BaseData Struct
	var basedata BaseData
	err2 := json.Unmarshal(dat, &basedata)
	check(err2)
	//end of the copy

	//parsing to url
	u, errUrl := url.Parse(r.URL.String())
	check(errUrl)
	m, _ := url.ParseQuery(u.RawQuery)
	urlparam := strings.Join(m["type"], "")  //converting this variable type []string to string in this function
	urlsort := strings.Join(m["sortby"], "") //converting this variable type []string to string in this function
	//end of the parsing url

	//listing all pokemons by type
	if len(urlparam) != 0 {

		if len(urlsort) != 0 && urlsort == "BaseAttack" { //if url's section is sortedby not NULL and equal to BaseAttack, then will work here

			sort.Sort(sortByBaseAttack(basedata.Pokemons)) //sorting array of struct

			fmt.Fprintln(w, "<a href='/list'>Type List</a> | <a href='/list?type="+urlparam+"'>Go to unsorted list</a><pre>")

			for _, a := range basedata.Pokemons {
				str := strings.Join(a.TypeI, "")
				str2 := strings.Join(a.TypeII, "")
				if str == urlparam || str2 == urlparam {
					fmt.Fprintln(w, a.Name)
					fmt.Fprintln(w, " ", "Number :", a.Number)
					fmt.Fprintln(w, " ", "Weight :", a.Weight)
					fmt.Fprintln(w, " ", "Height :", a.Height)
					fmt.Fprintln(w, " ", "Base Attack :", a.BaseAttack)
					fmt.Fprintln(w, " ", "BaseDefense :", a.BaseDefense)
					fmt.Fprintln(w, " ", "BaseStamina :", a.BaseStamina)
					fmt.Fprintln(w, " ", "Classification :", a.Classification)
					fmt.Fprintln(w, " ", "BuddyDistanceNeeded :", a.BuddyDistanceNeeded)
					fmt.Fprint(w, "  ", "TypeI :")
					for _, c := range a.TypeI {
						fmt.Fprintln(w, " ", c)
					}
					if len(a.TypeII) > 0 {
						fmt.Fprint(w, "  ", "TypeII :")
					}
					for _, c := range a.TypeII {
						fmt.Fprintln(w, " ", c)
					}
					fmt.Fprint(w, "  ", "Weaknesses :")
					for _, c := range a.Weaknesses {
						fmt.Fprintln(w, " ", c)
					}

					fmt.Fprintln(w, " ", "Special Attack(s) :")
					for _, c := range a.SpecialAttacks {
						fmt.Fprintln(w, "  ", c)
					}

					fmt.Fprintln(w, " ", "Fast Attack(s) :")
					for _, c := range a.FastAttackS {
						fmt.Fprintln(w, "  ", c)
					}

					fmt.Fprintln(w, " ", "Candy :")
					fmt.Fprintln(w, "  ", "FamilyID :", a.Candy.FamilyID)
					fmt.Fprintln(w, "  ", "Name :", a.Candy.Name)

					if len(a.NextEvolutions) > 0 {
						fmt.Fprintln(w, " ", "Next evolution(s) :")
						for _, c := range a.NextEvolutions {
							fmt.Fprintln(w, "   ", "Name :", c.Name)
							fmt.Fprintln(w, "   ", "Number :", c.Number)

						}
					}

					fmt.Fprintln(w, " ", "Next Evolution Requirements :")
					fmt.Fprintln(w, "  ", "Family :", a.NextEvolutionRequirements.Family)
					fmt.Fprintln(w, "  ", "Name :", a.NextEvolutionRequirements.Name)
					fmt.Fprintln(w, "  ", "Amount :", a.NextEvolutionRequirements.Amount)

					if len(a.PreviousEvolutions) > 0 {
						fmt.Fprintln(w, " ", "Previous evolution(s) :")
					}
					for _, c := range a.PreviousEvolutions {
						fmt.Fprintln(w, "  ", "Number :", c.Number)
						fmt.Fprintln(w, "  ", "Name :", c.Name)
					}

					fmt.Fprintln(w, "\n")
				}
			}
			fmt.Fprintln(w, "</pre>")
		} else { //if url's section is sortedby NULL, then will work here and will list all pokemons
			fmt.Fprintln(w, "<a href='/list'>Type List</a> | <a href='/list?type="+urlparam+"&sortby="+"BaseAttack"+"'>Sort by Base Attack</a><pre>")

			for _, a := range basedata.Pokemons {
				str := strings.Join(a.TypeI, "")
				str2 := strings.Join(a.TypeII, "")
				if str == urlparam || str2 == urlparam {
					fmt.Fprintln(w, a.Name)
					fmt.Fprintln(w, " ", "Number :", a.Number)
					fmt.Fprintln(w, " ", "Weight :", a.Weight)
					fmt.Fprintln(w, " ", "Height :", a.Height)
					fmt.Fprintln(w, " ", "Base Attack :", a.BaseAttack)
					fmt.Fprintln(w, " ", "BaseDefense :", a.BaseDefense)
					fmt.Fprintln(w, " ", "BaseStamina :", a.BaseStamina)
					fmt.Fprintln(w, " ", "Classification :", a.Classification)
					fmt.Fprintln(w, " ", "BuddyDistanceNeeded :", a.BuddyDistanceNeeded)
					fmt.Fprint(w, "  ", "TypeI :")
					for _, c := range a.TypeI {
						fmt.Fprintln(w, " ", c)
					}
					if len(a.TypeII) > 0 {
						fmt.Fprint(w, "  ", "TypeII :")
					}
					for _, c := range a.TypeII {
						fmt.Fprintln(w, " ", c)
					}
					fmt.Fprint(w, "  ", "Weaknesses :")
					for _, c := range a.Weaknesses {
						fmt.Fprintln(w, " ", c)
					}

					fmt.Fprintln(w, " ", "Special Attack(s) :")
					for _, c := range a.SpecialAttacks {
						fmt.Fprintln(w, "  ", c)
					}

					fmt.Fprintln(w, " ", "Fast Attack(s) :")
					for _, c := range a.FastAttackS {
						fmt.Fprintln(w, "  ", c)
					}

					fmt.Fprintln(w, " ", "Candy :")
					fmt.Fprintln(w, "  ", "FamilyID :", a.Candy.FamilyID)
					fmt.Fprintln(w, "  ", "Name :", a.Candy.Name)

					if len(a.NextEvolutions) > 0 {
						fmt.Fprintln(w, " ", "Next evolution(s) :")
						for _, c := range a.NextEvolutions {
							fmt.Fprintln(w, "   ", "Name :", c.Name)
							fmt.Fprintln(w, "   ", "Number :", c.Number)

						}
					}

					fmt.Fprintln(w, " ", "Next Evolution Requirements :")
					fmt.Fprintln(w, "  ", "Family :", a.NextEvolutionRequirements.Family)
					fmt.Fprintln(w, "  ", "Name :", a.NextEvolutionRequirements.Name)
					fmt.Fprintln(w, "  ", "Amount :", a.NextEvolutionRequirements.Amount)

					if len(a.PreviousEvolutions) > 0 {
						fmt.Fprintln(w, " ", "Previous evolution(s) :")
					}
					for _, c := range a.PreviousEvolutions {
						fmt.Fprintln(w, "  ", "Number :", c.Number)
						fmt.Fprintln(w, "  ", "Name :", c.Name)
					}

					fmt.Fprintln(w, "\n")
				}
			}
			fmt.Fprintln(w, "</pre>")
		}
		//end of the listing all pokemons by type

	} else { // if url's sections are empty, then will work here and will list all pokemon types

		fmt.Fprintln(w, "<html>")
		fmt.Fprintln(w, "<head></head><body>")
		fmt.Fprintln(w, "<a href='http://localhost:8080'>Main Section</a>")
		fmt.Fprintln(w, "<p>Please choose a pokemon type</p>")
		for _, a := range basedata.Types {

			fmt.Fprintln(w, "<a href='/list?type="+a.Name+"'> - "+a.Name+"</a><br/>")
		}
		fmt.Fprintln(w, "</body></html>")

	} //end of the section

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/get url:", r.URL)

	//reading json file
	dat, err := ioutil.ReadFile("data.json")
	check(err)
	//end of the reading json file

	//copying to json variables to BaseData Struct
	var basedata BaseData
	err2 := json.Unmarshal(dat, &basedata)
	check(err2)
	//end of the copy

	//parsing to url
	u, errUrl := url.Parse(r.URL.String())
	check(errUrl)
	m, _ := url.ParseQuery(u.RawQuery)
	urlparam := strings.Join(m["name"], "")
	//end of the parsing to url

	if len(urlparam) > 0 { //checking if url param's name is null or not
		fmt.Fprintln(w, "<a href='get'>Go to Pokemon List</a><pre>")

		for _, a := range basedata.Pokemons {
			//listing selected pokemon
			if a.Name == urlparam {
				//list pokemon's features
				fmt.Fprintln(w, a.Name)
				fmt.Fprintln(w, " ", "Number :", a.Number)
				fmt.Fprintln(w, " ", "Weight :", a.Weight)
				fmt.Fprintln(w, " ", "Height :", a.Height)
				fmt.Fprintln(w, " ", "Base Attack :", a.BaseAttack)
				fmt.Fprintln(w, " ", "BaseDefense :", a.BaseDefense)
				fmt.Fprintln(w, " ", "BaseStamina :", a.BaseStamina)
				fmt.Fprintln(w, " ", "Classification :", a.Classification)
				fmt.Fprintln(w, " ", "BuddyDistanceNeeded :", a.BuddyDistanceNeeded)
				fmt.Fprint(w, "  ", "TypeI :")
				for _, c := range a.TypeI {
					fmt.Fprintln(w, " ", c)
				}
				if len(a.TypeII) > 0 {
					fmt.Fprint(w, "  ", "TypeII :")
				}
				for _, c := range a.TypeII {
					fmt.Fprintln(w, " ", c)
				}
				fmt.Fprint(w, "  ", "Weaknesses :")
				for _, c := range a.Weaknesses {
					fmt.Fprintln(w, " ", c)
				}

				fmt.Fprintln(w, " ", "Special Attack(s) :")
				for _, c := range a.SpecialAttacks {
					fmt.Fprintln(w, "  ", c)
				}

				fmt.Fprintln(w, " ", "Fast Attack(s) :")
				for _, c := range a.FastAttackS {
					fmt.Fprintln(w, "  ", c)
				}

				fmt.Fprintln(w, " ", "Candy :")
				fmt.Fprintln(w, "  ", "FamilyID :", a.Candy.FamilyID)
				fmt.Fprintln(w, "  ", "Name :", a.Candy.Name)

				if len(a.NextEvolutions) > 0 {
					fmt.Fprintln(w, " ", "Next evolution(s) :")
					for _, c := range a.NextEvolutions {
						fmt.Fprintln(w, "   ", "Name :", c.Name)
						fmt.Fprintln(w, "   ", "Number :", c.Number)

					}
				}

				fmt.Fprintln(w, " ", "Next Evolution Requirements :")
				fmt.Fprintln(w, "  ", "Family :", a.NextEvolutionRequirements.Family)
				fmt.Fprintln(w, "  ", "Name :", a.NextEvolutionRequirements.Name)
				fmt.Fprintln(w, "  ", "Amount :", a.NextEvolutionRequirements.Amount)

				if len(a.PreviousEvolutions) > 0 {
					fmt.Fprintln(w, " ", "Previous evolution(s) :")
				}
				for _, c := range a.PreviousEvolutions {
					fmt.Fprintln(w, "  ", "Number :", c.Number)
					fmt.Fprintln(w, "  ", "Name :", c.Name)
				} // end of the listing features

				atyp1 := strings.Join(a.TypeI, "")  //converting this variable type []string to string in this function
				atyp2 := strings.Join(a.TypeII, "") //converting of this variable type []string to string in this function

				//listing pokemon's types
				for _, t := range basedata.Types {
					if atyp1 == t.Name || atyp2 == t.Name {
						fmt.Fprintln(w, "\n", " Pokemon Type :", t.Name)

						if len(t.EffectiveAgainst) > 0 {
							fmt.Fprintln(w, "   Effective Against :")
						}

						for _, c := range t.EffectiveAgainst {
							fmt.Fprintln(w, "    ", c)
						}

						if len(t.WeakAgainst) > 0 {
							fmt.Fprintln(w, "   Weak Against :")
						}
						for _, c := range t.WeakAgainst {
							fmt.Fprintln(w, "    ", c)

						}

					}
				} //end of the listing pokemon's types

				//listing example pokemons
				fmt.Fprintln(w, "  ", "Example Pokemons :")
				counter := 0
				for _, p := range basedata.Pokemons {
					ptyp1 := strings.Join(p.TypeI, "")  //converting this variable type []string to string in this function
					ptyp2 := strings.Join(p.TypeII, "") //converting this variable type []string to string in this function
					if ((ptyp1 == atyp1 || ptyp2 == atyp1) || (ptyp1 == atyp2 || ptyp2 == atyp2)) && (p.Name != a.Name) {
						fmt.Fprintln(w, "    ", p.Name)
						counter++

						if counter == 2 {
							break
						}
					}
				} //end of the listing example pokemons

				//listing moves
				fmt.Fprintln(w, "\n", " ", "Moves :")
				for _, moves := range basedata.Moves {
					if atyp1 == moves.Type || atyp2 == moves.Type {
						fmt.Fprintln(w, "    Moves ID:", moves.ID)
						fmt.Fprintln(w, "    Moves Name:", moves.Name)
						fmt.Fprintln(w, "    Moves Type:", moves.Type)
						fmt.Fprintln(w, "    Moves Damage:", moves.Damage)
						fmt.Fprintln(w, "    Moves Energy:", moves.Energy)
						fmt.Fprintln(w, "    Moves Dps:", moves.Dps)
						fmt.Fprintln(w, "    Moves Duration:", moves.Duration, "\n")
					}
				}
				//end of the listing moves

			}

		}
		fmt.Fprintln(w, "</pre>")
	} else {
		//listing all pokemon if its not selected on URL, i used HTML tag for use <a> tag
		fmt.Fprintln(w, "<html>")
		fmt.Fprintln(w, "<head></head><body>")
		fmt.Fprintln(w, "<a href='http://localhost:8080'>Go to main Section</a>")

		fmt.Fprintln(w, "<p>Please choose a pokemon</p>")

		for _, a := range basedata.Pokemons {

			fmt.Fprintln(w, "<a href='/get?name="+a.Name+"'> - "+a.Name+"</a><br/>")
		}
		fmt.Fprintln(w, "</body></html>")
	}

}

func listHandlerType(w http.ResponseWriter, r *http.Request) {
	log.Println("/list/types url:", r.URL)

	//reading json file
	dat, err := ioutil.ReadFile("data.json")
	check(err)
	//end of the reading json file

	//copying to json variables to BaseData Struct
	var basedata BaseData
	err2 := json.Unmarshal(dat, &basedata)
	check(err2)
	//end of the copy

	//parsing to url
	u, errUrl := url.Parse(r.URL.String())
	check(errUrl)
	m, _ := url.ParseQuery(u.RawQuery)
	urlparam := strings.Join(m["name"], "")
	//end of the parsing url

	if len(urlparam) > 0 { //checking url parametre is null or not

		fmt.Fprintln(w, "<a href='/list/types'>Go to List of types</a><pre>")

		for _, a := range basedata.Types {
			if urlparam == a.Name {
				fmt.Fprintln(w, "Name :", a.Name)
				fmt.Fprintln(w, "Effective Against :")
				for _, c := range a.EffectiveAgainst {
					fmt.Fprintln(w, "-", c)
				}

				fmt.Fprintln(w, "Weak Against :")
				for _, c := range a.WeakAgainst {
					fmt.Fprintln(w, "-", c)

				}

			}
		}
		fmt.Fprintln(w, "</pre>")
	} else { // if url is null, then list all types
		fmt.Fprintln(w, "<a href='http://localhost:8080'>Go to main Section</a>")

		fmt.Fprint(w, "<p>The Types of List</p>")
		for _, a := range basedata.Types {
			fmt.Fprintln(w, "<a href='/list/types?name="+a.Name+"'>"+a.Name+"</a><br/>")
		}
	}

	fmt.Fprintln(w, "</body></html>")

}

func otherwise(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head></head><body>")
	fmt.Fprintln(w, "<p>Welcometo my pokemon task</p>")
	fmt.Fprintln(w, "<a href='/get'>See all pokemons</a><br/>")
	fmt.Fprintln(w, "<a href='/list'>See all pokemons in a type</a><br/>")
	fmt.Fprintln(w, "<a href='/list/types'>See all features of type </a><br/>")

	fmt.Fprintln(w, "</body></html>")
}

func main() {
	//TODO: read data.json to a BaseData

	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/list/types", listHandlerType)
	http.HandleFunc("/get", getHandler)
	//TODO: add more
	http.HandleFunc("/", otherwise)
	log.Println("starting server on :8080")
	http.ListenAndServe(":8080", nil)

}
