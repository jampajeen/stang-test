#!/bin/bash
sleep 10

echo mongo_setup.sh time now: `date +"%T" `
mongo --host mongodb:27017 -u root -p password <<EOF
  db.createUser(
  {
    user: "root",
    pwd: "password",
    roles: [ { role: "readWrite", db: "test" } ]
  }
  );
EOF
