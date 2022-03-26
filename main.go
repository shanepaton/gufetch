package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/viper"
)

type UserConfig struct {
	Token      string
	GithubUser string
	GitlabID   string
	site       string
}

type GHUser struct {
	Login       string `json:"login"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Blog        string `json:"blog"`
	Location    string `json:"location"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	Message     string `json:"message"`
}

type GLUser struct {
	Login       string `json:"username"`
	Name        string `json:"name"`
	Blog        string `json:"website_url"`
	Pronouns    string `json:"pronouns"`
	Location    string `json:"location"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	Message     string `json:"message"`
}

func GetUserConfig() UserConfig {
	var viper = viper.New()
	viper.AddConfigPath("~/.config/gitfetch/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return UserConfig{viper.GetString("token.gitlab"), viper.GetString("defualt.githubUsername"), viper.GetString("defualt.gitlabID"), viper.GetString("defualt.site")}
}

func ApiRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func EmptyCheckFMT(string string) string {
	if string == "" {
		return "N/A"
	} else {
		return strings.TrimSpace(string)
	}
}

func main() {
	isGithub := flag.Bool("github", false, "select Github as your platform")
	isGitlab := flag.Bool("gitlab", false, "select Gitlab as your platform")
	isMono := flag.Bool("mono", false, "display with no colours")

	flag.Parse()
	additionalCLA := flag.Args()

	var userConfig UserConfig = GetUserConfig()

	if *isMono {
		color.Disable()
	}

	if *isGithub {
		if len(additionalCLA) < 1 && userConfig.GithubUser == "" {
			fmt.Println("Please enter a Github username or set it in the config file")
			return
		}

		var result GHUser

		if len(additionalCLA) > 0 {
			var fetch = ApiRequest("https://api.github.com/users/" + additionalCLA[0])
			if err := json.Unmarshal(fetch, &result); err != nil {
				fmt.Println("Can not unmarshal Github JSON")
			}
		} else {
			var fetch = ApiRequest("https://api.github.com/users/" + userConfig.GithubUser)
			if err := json.Unmarshal(fetch, &result); err != nil {
				fmt.Println("Can not unmarshal Github JSON")
			}
		}

		if result.Message == "Not Found" {
			fmt.Println("Github username not found")
			return
		}

		fmt.Println(`        @@@@@@@@@@@@@@@                  gitfetch - Github`)
		fmt.Println(`     @@@@@@@@@@@@@@@@@@@@@               ~~~~~~~~~`)
		fmt.Println(`   @@@@    @@@@&@@@@    @@@@             Username:  `, result.Login)
		fmt.Println(`  @@@@@                 @@@@@            Name:      `, EmptyCheckFMT(result.Name))
		fmt.Println(` @@@@@                   @@@@@           Bio:       `, EmptyCheckFMT(result.Bio))
		fmt.Println(` @@@@@                   @@@@@           Website:   `, EmptyCheckFMT(result.Blog))
		fmt.Println(` @@@@@                   @@@@@           Location:  `, EmptyCheckFMT(result.Location))
		fmt.Println(` @@@@@@@               @@@@@@@           Followers: `, result.Followers)
		fmt.Println(`  @@@@ @@@@@       @@@@@@@@@@            Following: `, result.Following)
		fmt.Println(`    @@@            @@@@@@@@              Repos:     `, result.PublicRepos)
		fmt.Println(`       @@@@@       @@@@@                 `)

	} else if *isGitlab {
		if len(additionalCLA) < 1 && userConfig.GitlabID == "" {
			fmt.Println("Please enter a Gitlab ID or set it in the config file")
			return
		}
		if userConfig.Token == "" {
			print("You need a token to fetch a Gitlab Profile!")
		} else {

			var result GLUser

			if len(additionalCLA) > 0 {
				var fetch = ApiRequest("https://gitlab.com/api/v4/users/" + additionalCLA[0] + "?private_token=" + userConfig.Token)
				if err := json.Unmarshal(fetch, &result); err != nil {
					fmt.Println("Can not unmarshal Gitlab JSON")
				}
			} else {
				var fetch = ApiRequest("https://gitlab.com/api/v4/users/" + userConfig.GitlabID + "?private_token=" + userConfig.Token)
				if err := json.Unmarshal(fetch, &result); err != nil {
					fmt.Println("Can not unmarshal Gitlab JSON")
				}
			}

			if result.Message == "404 User Not Found" {
				fmt.Println("Gitlab ID not found")
				return
			}

			color.Println(`<fg=fca326>     ((                  ((</>              gitfetch - Gitlab`)
			color.Println(`<fg=fca326>    ((((                ((((</>             ~~~~~~~~~~~~~~~~~`)
			color.Println(`<fg=fca326>   ((((((              (((((</>             Username:  `, EmptyCheckFMT(result.Login))
			color.Println(`<fg=fca326>  ((((((((            (((((((</>            Name:      `, EmptyCheckFMT(result.Name))
			color.Println(`<fg=fca326> //</><fg=FA6C25>///////</><fg=E04228>(((((((((((</><fg=FA6C25>///////</><fg=fca326>//</>           Bio:       `, EmptyCheckFMT(result.Bio))
			color.Println(`<fg=fca326> ////</><fg=FA6C25>//////</><fg=E04228>((((((((((</><fg=FA6C25>//////</><fg=fca326>////</>          Pronouns:  `, EmptyCheckFMT(result.Pronouns))
			color.Println(`<fg=fca326>//////</><fg=FA6C25>//////</><fg=E04228>((((((((</><fg=FA6C25>/////</><fg=fca326>//////</>          Website:   `, EmptyCheckFMT(result.Blog))
			color.Println(`<fg=fca326>   /////</><fg=FA6C25>/////</><fg=E04228>((((((</><fg=FA6C25>////</><fg=fca326>/////</>             Location:  `, EmptyCheckFMT(result.Location))
			color.Println(`<fg=fca326>       ////</><fg=FA6C25>///</><fg=E04228>(((((</><fg=FA6C25>///</><fg=fca326>////</>               Followers: `, result.Followers)
			color.Println(`<fg=fca326>          //</><fg=FA6C25>//</><fg=E04228>(((</><fg=FA6C25>//</><fg=fca326>//</>                    Following: `, result.Following)
			color.Println(`<fg=fca326>             /</><fg=FA6C25>/</><fg=E04228>(</><fg=FA6C25>/</><fg=fca326>/</>`)
		}
	}
}
