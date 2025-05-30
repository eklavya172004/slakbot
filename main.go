package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main(){

		if err := godotenv.Load(); err != nil {
		log.Println("[WARN] No .env file found. Using environment variables.")
	}

		botToken := os.Getenv("SLACK_BOT_TOKEN")
		appToken := os.Getenv("SLACK_APP_TOKEN")	

		bot := slacker.NewClient(botToken,appToken)
		
		go printCommandEvents(bot.CommandEvents())

		const botMentionName = "@age-bot"

		bot.Command("<year>",&slacker.CommandDefinition{
			Description: "Get the year of birth",
			Examples: []string{"My year of birth is 2004"},
			Handler: func(botctx slacker.BotContext,request slacker.Request,response slacker.ResponseWriter) {
			
			yearStr := request.Param("year")
			yob, err := strconv.Atoi(yearStr)
			if err != nil {
				response.Reply("Please enter a valid numeric year like 2004.")
				return
			}

			currentYear := time.Now().Year()

			if yob < 1900 || yob > currentYear {
				response.Reply(fmt.Sprintf("⚠️ %d is not a valid year of birth. Please enter a year between 1900 and %d.",yob, currentYear))
				return
			}

			age := currentYear - yob
			reply := fmt.Sprintf("You are %d years old.", age)
			response.Reply(reply)
				},
		})

		//The context package in Go is used to manage deadlines, timeouts, and cancellation signals across API boundaries and goroutines. It's especially helpful in concurrent or long-running operations.
		
		ctx , cancel := context.WithCancel(context.Background())
		defer cancel()

		err := bot.Listen(ctx)
		if err != nil {
			log.Fatal(err)
		}
}

func printCommandEvents(analyticschannel <-chan *slacker.CommandEvent){
	for event := range analyticschannel{
		fmt.Println("Command Events")
		fmt.Println("Command: ", event.Command)
		fmt.Println("Timestamp: ", event.Timestamp)
		fmt.Println("Event: ", event.Event)
		fmt.Println("Params: ", event.Parameters)
		fmt.Println()		
	}
}