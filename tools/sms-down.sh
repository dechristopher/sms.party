# Kills active sms.party container that was brought up with ./sms-up
echo "[sms.p] Running devdown"

# Kill all local sms-p docker instances
docker kill -f sms-p

# Remove container from memory (keeps image)
docker rm -f sms-p

echo "[sms.p] Done!"
