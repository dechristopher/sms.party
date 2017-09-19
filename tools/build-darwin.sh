# Builds sms.party executable for darwin (OSX) and outputs as build/main-darwin
echo "[sms.p] Building sms.party darwin binary..."

# Enter src/ directory after resetting working directory to tools/
cd "${0%/*}"
cd ../src

# Run the build
CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o ../build/main-darwin

# Exit src/ directory into project root
cd ..

echo "[sms.p] Done!"
