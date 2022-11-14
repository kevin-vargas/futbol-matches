FROM node:alpine

WORKDIR /frontend

COPY package*.json .

COPY . .

RUN npm i

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]