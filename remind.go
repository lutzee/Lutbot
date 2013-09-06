package main

import (
	"github.com/jdiez17/irc-go"
	"strconv"
	"strings"
	"time"
)

func remindCommandHandler(c *irc.Connection, e *irc.Event) {
	if len(e.Params) < 1 {
		e.React(c, "Not enough parameters! (syntax: .remind <time in minutes> <reason>)")
		return
	}
	inputTimeS := e.Params[0]
	inputTime, err := strconv.ParseInt(e.Params[0], 0, 64)
	if err != nil {
		e.React(c, "Invalid time format!")
		return
	}
	time.Sleep(time.Duration(inputTime) * time.Minute)
	message := "This is your " + inputTimeS + " minute reminder " + strings.Join(e.Params[1:], " ")
	e.React(c, message)
}
