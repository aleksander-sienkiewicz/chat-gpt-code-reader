# chat gpt code reader
 reads code from a local file, sends it to chat gpt which returns the libraries used, then outputs to terminal, can also be used to ask anything and output to the file.
 
 CLI to build & run 
 

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air ~ % cd documents

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air documents % cd projectdev

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air projectdev % mkdir chat-gpt-code-reader

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air projectdev % cd chat-gpt-code-reader

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-code-reader % go mod init github.com/aleksander-sienkiewicz/chatt-gpt-code-reader

go: creating new go.mod: module github.com/aleksander-sienkiewicz/chatt-gpt-code-reader

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-code-reader % go mod tidy

go: finding module for package github.com/spf13/viper

go: finding module for package github.com/PullRequestInc/go-gpt3
go: found github.com/PullRequestInc/go-gpt3 in github.com/PullRequestInc/go-gpt3 v1.1.15

go: found github.com/spf13/viper in github.com/spf13/viper v1.15.0

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-code-reader % go run main.go

2023/05/01 18:20:09 [429:insufficient_quota] You exceeded your current quota, please check your plan and billing details.

exit status 1

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air chat-gpt-code-reader % 


