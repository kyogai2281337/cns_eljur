#FROM node:latest as build-stage
#
#WORKDIR /app
#
#RUN ls -a
#
#COPY ./frontend/cns_eljur/ ./
#
#RUN npm install
#
#COPY . .
#
#RUN npm run build

FROM nginx

#COPY --from=build-stage /app/dist /usr/share/nginx/html

COPY build/nginx/nginx.conf /etc/nginx/nginx.conf

COPY ssl/cert.crt /etc/nginx/cert.crt
COPY ssl/cert.key /etc/nginx/cert.key