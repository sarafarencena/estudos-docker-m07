# Estágio de build: Node.js só é necessário para transformar o código React em arquivos estáticos.
FROM node:22-alpine AS build

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Estágio de execução: Nginx recebe somente os arquivos gerados em /app/dist.
FROM nginx:1.27-alpine AS runtime

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

# A rede dos containers é configurada no compose.yaml, não neste Dockerfile.
