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
	inputTime, err := strconv.ParseFloat(e.Params[0], 64)
	if err != nil {
		e.React(c, "Invalid time format!")
		return
	}
	sleepTime := time.Duration(inputTime) * time.Minute
	if sleepTime > time.Hour*24*7 {
		e.React(c, "Duration specified is not allowed!")
		return
	}
	time.Sleep(sleepTime * time.Minute)
	message := "This is your " + inputTimeS + " minute reminder " + strings.Join(e.Params[1:], " ")
	e.React(c, message)
}
