# Estágio de build: contém o compilador e existe somente enquanto a imagem é construída.
FROM alpine:3.20 AS builder

RUN apk add --no-cache build-base

WORKDIR /src
COPY hello.c .

# Binário estático: a imagem final não precisa de compilador nem de bibliotecas dinâmicas.
RUN gcc -static -Os -s -o hello hello.c

# Estágio de execução: recebe somente o resultado do build.
FROM scratch AS runtime

COPY --from=builder /src/hello /hello

ENTRYPOINT ["/hello"]
