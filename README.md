# financial-chat

docker pull rabbitmq

docker run -d --hostname my-rabbit -p 5672:5672 -p 15672:15672 --name financial-rabbit rabbitmq