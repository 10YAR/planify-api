#!/bin/bash

PATH=$PATH:/usr/local/go/bin
cd /home/planify
SHAONESUM=$(sha1sum api/database/01-tables.sql)
GITRES=$(git pull)
echo $GITRES
if [ "$GITRES" != "Already up to date." ];
then
  echo "Updates detected"
  SHATWOSUM=$(sha1sum api/database/01-tables.sql)
  if [ "$SHAONESUM" != "$SHATWOSUM" ];
  then

    if [ -f .env ]; then
      export $(echo $(cat .env | sed 's/#.*//g'| xargs) | envsubst)
    fi

    DB_HOST=${DB_HOST:0:len-1}
    DB_USER=${DB_USER:0:len-1}
    DB_PASSWORD=${DB_PASSWORD:0:len-1}
    DB_NAME=${DB_NAME:0:len-1}

    echo "Migration file changed, recreating database..."
    mysql -h$DB_HOST -u$DB_USER -p$DB_PASSWORD $DB_NAME < api/database/02-tables.down.sql
    mysql -h$DB_HOST -u$DB_USER -p$DB_PASSWORD $DB_NAME < api/database/01-tables.sql
    echo "Migration finished"
  fi

  echo "Building app..."
  screen -S planify -p 0 -X stuff "^C"
  screen -S planify -p 0 -X stuff "rm planifyApp && cd api && go build -v -o ../planifyApp . ^M"
  screen -S planify -p 0 -X stuff "cd ../ && chmod +x planifyApp ^M"
  screen -S planify -p 0 -X stuff "./planifyApp^M"
  sleep 3
  echo "Build finished"
else
  echo "App is up to date, exiting"
fi