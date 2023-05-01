package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")          //our config file with api key in .env
	viper.ReadInConfig()                 //read the config file
	apiKey := viper.GetString("API_KEY") //read for api key, value assigned to it is api key
	if apiKey == "" {                    //if its empty
		panic("Missing API KEY") //panic missing api key
	}

	ctx := context.Background()      //use context
	client := gpt3.NewClient(apiKey) //crete client
	//cntxt done, client created

	//read input file, send it to the client along with context
	const inputFile = "./input_with_code.txt" //input file is filename
	fileBytes, err := os.ReadFile(inputFile)  //use os package to read that file, assign to fileBytes, or error if it doesnt work
	if err != nil {                           //handle error
		log.Fatalf("failed to read file: %v", err) //issue reading file
	}

	msgPrefix := "give me a short list of libraries that are used in the code\n```python\n"
	//this is the prompt the be used to instruct chatgpt, this is why this programs so versitile! <3
	msgSuffix := "\n```"
	msg := msgPrefix + string(fileBytes) + msgSuffix //construct msg , which has prefix, code, then suffix,

	outputBuilder := strings.Builder{} //build our output file
	//exact same code from last video
	err = client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{ //in prompt
			msg, //ur just sending the message, no need to type a string at all
		},
		MaxTokens:   gpt3.IntPtr(3000),  //assign a max tokens value,
		Temperature: gpt3.Float32Ptr(0), //Temperature is a parameter of OpenAI ChatGPT, GPT-3 and
		//GPT-4 models that governs the randomness and thus the creativity of the responses. It is always a number between 0 and 1.
	}, func(resp *gpt3.CompletionResponse) {
		outputBuilder.WriteString(resp.Choices[0].Text) //print out text
	})
	if err != nil { //handle error
		log.Fatalln(err)
	}
	output := strings.TrimSpace(outputBuilder.String()) //everything written in output builder

	const outputFile = "./output.txt"                           //create output file
	err = os.WriteFile(outputFile, []byte(output), os.ModePerm) //write output into output file
	if err != nil {                                             //if theres an error, print that its
		log.Fatalf("failed to read file: %v", err) //cuz u cant read file
	}
}
