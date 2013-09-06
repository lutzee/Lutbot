package main

import (
	"flag"
	"fmt"
	"github.com/jdiez17/irc-go"
	"os"
	"strings"
	"time"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	configFile := fs.String("config", "", "config.json")
	fs.Parse(os.Args[1:])

	err := loadConfig(*configFile)
	if err != nil {
		fmt.Println("Error reading the configuration: " + err.Error())
		return
	}

	conn, err := irc.NewConnection(Config.IRC.Server, int(Config.IRC.Port))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()

	conn.LogIn(irc.Identity{Nick: Config.Nick})

	conn.AddHandler(irc.MOTD_END, func(c *irc.Connection, e *irc.Event) {
		if Config.NickServPassword != "" {
			c.Privmsg("NickServ", "identify "+Config.NickServPassword)
		}

		for _, channel := range Config.Channels {
			c.Join(channel)
		}
	})

	bot := irc.NewBot(conn)
	bot.AddCommand("join", func(c *irc.Connection, e *irc.Event) {
		for _, admin := range Config.Admins {
			if e.Payload["sender"] == admin {
				c.Join(e.Params[0] + e.Params[1])
			}
		}
	})
	bot.AddCommand("part", func(c *irc.Connection, e *irc.Event) {
		for _, admin := range Config.Admins {
			if e.Payload["sender"] == admin {
				c.Part(e.Params[0] + e.Params[1])
			}
		}
	})
	bot.AddCommand("echo", func(c *irc.Connection, e *irc.Event) {
		message := strings.Join(e.Params, " ")
		e.React(c, message)
	})
	bot.AddCommand("remind", remindCommandHandler)

	for {
		<-time.After(1 * time.Second)
	}
}
