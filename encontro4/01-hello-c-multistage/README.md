# Demonstração 1 — Hello World em C

## O que demonstrar

O primeiro estágio usa uma imagem com compilador para gerar o binário `hello`. O estágio final parte de `scratch` e recebe somente esse binário por meio de `COPY --from=builder`.

```bash
docker build -t hello-c-multistage:1.0 .
docker run --rm hello-c-multistage:1.0
```

Para observar o estágio de compilação separadamente, construa-o como alvo:

```bash
docker build --target builder -t hello-c-builder:1.0 .
docker image ls hello-c-multistage hello-c-builder
```

## Perguntas para a turma

1. Qual arquivo atravessa a fronteira entre `builder` e `runtime`?
2. Por que `gcc` não existe na imagem final?
3. O que aconteceria se `COPY --from=builder /src /app` fosse usado no lugar de copiar apenas `/src/hello`?
4. Há mais de um container em execução neste exemplo? Por que, então, ele não precisa de rede?

O objetivo é separar visualmente **ambiente de construção** de **ambiente de execução**. Os estágios não são serviços e não se comunicam por rede: o Docker apenas copia o artefato solicitado durante o build.
