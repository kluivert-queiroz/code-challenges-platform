FROM node:24-alpine AS builder

WORKDIR /app

RUN corepack enable

COPY package.json yarn.lock ./

RUN yarn install

COPY . .

RUN yarn build

FROM node:24-alpine

WORKDIR /app

COPY --from=builder /app/dist ./dist

RUN yarn global add serve

EXPOSE 8080

CMD [ "serve", "-s", "dist", "-l", "8080" ]