FROM node:latest as css
WORKDIR /app
COPY templates/ ./templates/
COPY tailwind.config.js .
COPY package.json .
COPY package-lock.json .
RUN npm install
RUN npx tailwindcss -i ./templates/main.css -o ./public/output.css --minify

FROM golang:latest
WORKDIR /app

COPY cmd/ ./cmd/
COPY pkg ./pkg/
COPY go.mod .
COPY go.sum .
COPY --from=css /app/public/output.css /app/public/output.css 
COPY --from=css /app/templates/ /app/templates/

ENV PORT=8080
EXPOSE ${PORT}

RUN go build -o /app/bin/server ./cmd/server/main.go
CMD [ "/app/bin/server" ]