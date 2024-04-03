package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analysticsChannel <-chan *slacker.CommandEvent) {
	for event := range analysticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6904819579380-6925193620096-e39BE5QOBwtXVrM8umG0BOUH")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06SBMTJ538-6915040349217-5fa03baed28a603abe2701738fa644d0404c9124b5f54af8def825ef1518518f")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	bot.Command("My year of birth is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year:= request.Param("year")
			yob,err := strconv.Atoi(year)
			if err != nil{
				println("error")
			}
			age:= 2024-yob
			r := fmt.Sprintf("Age is %d",age)
			response.Reply(r)
		},
	})
	
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}
