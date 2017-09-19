# Builds sms.party executable for Linux and outputs as build/main-linux
echo "[sms.p] Building sms.party linux binary..."

# Enter src/ directory after resetting working directory to tools/
cd "${0%/*}"
cd ../src

# Run the build
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../build/main-linux

# Exit src/ directory into project root
cd ..

echo "[sms.p] Done!"
