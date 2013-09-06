package main

import (
	"fmt"
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
	sleepTime := inputTime * float64(time.Minute)
	if time.Duration(sleepTime) > time.Hour*24*7 {
		e.React(c, "Duration specified is not allowed!")
		return
	}
	fmt.Println(time.Duration(sleepTime))
	time.Sleep(time.Duration(sleepTime))
	message := "This is your " + inputTimeS + " minute reminder " + strings.Join(e.Params[1:], " ")
	e.React(c, message)
}
