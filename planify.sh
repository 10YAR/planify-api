#!/bin/bash
PATH=$PATH:/usr/local/go/bin
GITRES=$(cd /home/planify && git pull)
if [ "$GITRES" != "Déjà à jour." ];
then
  echo "Updates detected, building app..."
  screen -S planify -p 0 -X stuff "^C"
  screen -S planify -p 0 -X stuff "go build -v -o ./planifyApi api^M"
  screen -S planify -p 0 -X stuff "./planifyApi^M"
  echo "Build successful. App launched."
else
  echo "No updates, no need to build"
fi