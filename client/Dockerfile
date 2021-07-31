FROM node:16.5.0-alpine
LABEL maintainer="Ben Saufley<contact@bensaufley.com>"

ENV NODE_ENV development

WORKDIR /tmp
COPY package*.json ./
RUN npm install

WORKDIR /usr/src/graphql-preact-starter
RUN mv /tmp/package*.json /tmp/node_modules ./
