#!/bin/bash
PATH=$PATH:/usr/local/go/bin
GITRES=$(cd /home/planify && git pull)

if [ "$GITRES" != "Déjà à jour." ]; then
  screen -S planify -p 0 -X stuff "^C"
  screen -S planify -p 0 -X stuff "go build -v -o ./planifyApi api && ./planifyApi^M"
fi