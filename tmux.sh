#!/bin/bash

SESSIONNAME="chat-app"
CHAT_APP_PATH=~/dev/chat-app/
tmux has-session -t $SESSIONNAME &> /dev/null

if [ $? != 0 ]
 then
    tmux new-session -s $SESSIONNAME -c $CHAT_APP_PATH -n mysql -d
    tmux send-keys -t $SESSIONNAME:0 "docker-compose up mysql" C-m

    tmux new-window -t $SESSIONNAME -c $CHAT_APP_PATH -n rails
    tmux send-keys -t $SESSIONNAME:1 "docker-compose up rails_app" C-m

    tmux new-window -t $SESSIONNAME -c $CHAT_APP_PATH -n go
    tmux send-keys -t $SESSIONNAME:2 "docker-compose up go_app" C-m

    tmux new-window -t $SESSIONNAME -c $CHAT_APP_PATH -n rabbitmq
    tmux send-keys -t $SESSIONNAME:3 "docker-compose up rabbitmq" C-m

    tmux new-window -t $SESSIONNAME -c $CHAT_APP_PATH -n redis
    tmux send-keys -t $SESSIONNAME:4 "docker-compose up redis" C-m

    tmux new-window -t $SESSIONNAME -c $CHAT_APP_PATH -n elasticsearch
    tmux send-keys -t $SESSIONNAME:5 "docker-compose up elasticsearch" C-m
fi

tmux attach -t $SESSIONNAME
