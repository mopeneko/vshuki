FROM node:12.16-alpine3.11

ENV PROJECT_WORKDIR /app/

WORKDIR ${PROJECT_WORKDIR}
COPY package.json yarn.lock ${PROJECT_WORKDIR}

RUN yarn install

COPY . .

EXPOSE 3000
ENV HOST 0.0.0.0

CMD [ "yarn", "dev" ]
