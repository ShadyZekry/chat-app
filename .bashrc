alias rails='docker-compose exec -it rails_app rails'
alias rake='docker-compose exec -it rails_app rake'
alias go='docker-compose exec -it go_app go'

alias redis='docker-compose exec -it redis redis-cli'
alias rds-flush='redis FLUSHALL'

alias rabbitmq='docker-compose exec -it rabbitmq'
alias rmq-purge='rabbitmq rabbitmqctl purge_queue $1'
rmq-get(){
	rabbitmq rabbitmqadmin get queue=$(echo $1 | tr -d [:space:])
}

