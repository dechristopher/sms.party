# Builds sms.party executable and packages it into a docker container
echo "[sms.p] Building sms.party container..."

# Move out of the tools/ dir and into the src/ dir
cd "${0%/*}"
cd ../src

# Build statically linked go app binary
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../build/main .

# Exit src/ dir into the project root
cd ..

# Remove old docker image
docker rmi -f csms/sms-p

# Build docker container
docker build -t csms/sms-p -f Dockerfile .

# List images
docker images

# Echo docker run instructions
echo "##############################################################################"
echo " docker run -d --name sms-p -h sms-p -e PORT=<port> -p <port>:<port> sms-p"
echo "##############################################################################"

echo "[sms.p] Done!"
