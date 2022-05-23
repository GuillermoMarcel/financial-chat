# financial-chat
Financial chat application 
This is a challenge only.

Dependencies. 
This code is writen in go 1.17 so this is a requirement. follow instruction in golang website for instalation.
gcc is a dependecy, For windows recomended here https://jmeubank.github.io/tdm-gcc/
for linux :

```bash
apt-get install build-essential
```

This systems were design to work with a RabitMQ Message Broker
We reccomend to use the docker image for this.
Run this commands 
```bash
docker pull rabbitmq
docker run -d --hostname my-rabbit -p 5672:5672 -p 15672:15672 --name financial-rabbit rabbitmq
```


for development:
.vscode files are provided to just open and lunch with the build in debugger

for running:
locate under the application folder and runit with the -c configuration pointing to the config file

main.go files are located under the cmd folder.
To run the chat application run the next commands

```bash
cd .\cmd\financial-chat\
go run .\ -c ..\..\config\chat.json
```
and for runing the chat bot this command
```bash
cd .\cmd\chat-bot\
go run .\ -c ..\..\config\bot.json
```


change any necesary configuration on the ./config files


Author Guillermo Marcel
guille.marcel04@gmail.com