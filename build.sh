sudo docker build -t testapp . --progress=plain --no-cache
sudo docker run --rm --name testapp -p 3001:3001 -e APP_ID=helloapp1 testapp