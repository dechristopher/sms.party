echo "[sms.p] Checking your system for dev dependencies..."

gostring=$(which go)
if [[ $gostring == *"/"* ]];
then
  echo "[GOOD] Golang"
else
  echo "[FAIL] Golang"
  echo "- \"go\" command not found in your system path, please install go!"
fi

dockerstring=$(which docker)
if [[ $dockerstring == *"/"* ]];
then
  echo "[GOOD] Docker"
else
  echo "[FAIL] Docker"
  echo "- \"docker\" command not found in your system path, please install docker!"
fi

echo "[sms.p] Done!"
