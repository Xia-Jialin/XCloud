apt-get update
docker pull rabbitmq:management
docker run -d --hostname fuyi-rabbit --name rabbitmq -e RABBITMQ_DEFAULT_USER=test -e RABBITMQ_DEFAULT_PASS=test -p 15672:15672 -p 5672:5672 rabbitmq:management
wget localhost:15672/cli/rabbitmqadmin
python3 rabbitmqadmin declare exchange name=dataServers type=fanout
python3 rabbitmqadmin declare exchange name=apiServers type=fanout

docker pull elasticsearch:6.7.0
docker run --name elasticsearch -d -e ES_JAVA_OPTS="-Xms512m -Xmx512m" -e "discovery.type=single-node" -p 9200:9200 -p 9300:9300 elasticsearch:6.7.0
curl -H "Content-Type: application/json"  localhost:9200/metadata -XPUT -d'{"mappings":{"objects":{"properties":{"name":{"type":"text","index":false},"version":{"type":"integer"},"size":{"type":"integer"},"hash":{"type":"text"}}}}}'