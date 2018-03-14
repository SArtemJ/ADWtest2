package main

import (
	"context"
	"flag"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"html/template"
	"fmt"
	"github.com/google/go-querystring/query"
	"log"
	"net/url"
)

type Issue struct {
	ID               *int64            `json:"id,omitempty"`
	Number           *int              `json:"number,omitempty"`
	State            *string           `json:"state,omitempty"`
	Title            *string           `json:"title,omitempty"`
	URL              *string           `json:"url,omitempty"`

}

type StatusLabel struct {
	State string
	Labels []string
}

//

//глобальные переменные
var (
	//token
	PersonalToken = ""
	//client github
	MyClient = github.NewClient(nil)
	//html templates
	tpl *template.Template
	// status for issues

	//repo issues
	Repo = "ADWtest"
	Owner = "SArtemJ"
)

func init() {

	//параметр ключа можно задавать при запуске, если не указываем используется ключ по умолчанию
	tk := flag.String("token", "", "")
	flag.Parse()
	PersonalToken = *tk

	//создаем нового клиента github с персональным токеном
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: PersonalToken},
	)
	tokenClient := oauth2.NewClient(context.Background(), tokenService)
	MyClient = github.NewClient(tokenClient)


}

func main() {

	u := fmt.Sprintf("repos/%v/%v/issues", Owner, Repo)
	log.Println(u)

	sl := StatusLabel{
		State: "all",
		Labels: []string{"bug"},
	}

	//string to url
	k, _ := url.Parse(u)

	//params for url
	qs, _ := query.Values(sl)

	//encode to correct view
	k.RawQuery = qs.Encode()

	//get issues
	req, _ := MyClient.NewRequest("GET", k.String(), nil)

	var issues []*Issue
	//метод библиотеки по работе с получаемыми данными
	MyClient.Do(context.Background(), req, &issues)

	for _, k := range issues {
		log.Println(*k.ID, *k.Title, *k.State, *k.Number)
	}


}