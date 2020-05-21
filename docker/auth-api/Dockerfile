FROM node:latest

WORKDIR /home/node/app
COPY auth-service/package*.json /home/node/app/
RUN npm install

CMD npm start
EXPOSE 9000
