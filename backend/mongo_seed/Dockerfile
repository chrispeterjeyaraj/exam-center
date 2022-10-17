FROM mongo:4
 # Will be set through Environment Files
 ARG MONGOURI

 ENV MONGOURI ${MONGOURI}

 COPY users.json /users.json

 CMD mongoimport --uri ${MONGOURI} --collection users --jsonArray --file /users.json


