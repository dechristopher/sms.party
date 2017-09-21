# Spins up a new csms-b container for development and testing
echo "[sms.p] Running devup"

# Move to location of script
cd "${0%/*}"

# Removes old containers (if any)
docker rm -f csms/sms-p

# Build docker container
./docker-build.sh

# Move back to project root
cd ..

# Run docker container with args[1] being port
# EXAMPLE: ./sms-up.sh 3000
docker run -d --name sms-p -h sms-p -e PORT=$1 -p $1:$1 csms/sms-p

echo "[sms.p] Done!"
