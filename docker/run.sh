docker build -t wallet-api .
docker run -d -p 8003:8003 --name wallet-api wallet-api
