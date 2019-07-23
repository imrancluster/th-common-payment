# booting up consul, redis, mysql containers
docker-compose up -d consul db redis

# building app
go build -v .

# setting KV, dependecy of app
curl --request PUT --data-binary @config.local.yml http://localhost:8500/v1/kv/th-common-payment

# start app
# THCOMMONP should be same of app.cons.AppName
export THCOMMONP_CONSUL_URL="127.0.0.1:8500"
export THCOMMONP_CONSUL_PATH="th-common-payment"
./th-common-payment serve
