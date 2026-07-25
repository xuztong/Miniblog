FROM node:22 AS builder

WORKDIR /app

COPY . .

RUN npm config set registry https://registry.npmmirror.com \
    && npm install \
    && npm run build

FROM nginx:1.25-alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
