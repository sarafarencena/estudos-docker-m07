# Demonstração 2 — Calculadora React + Flask

Esta aplicação possui dois containers em execução:

- `frontend`: gera os arquivos React com Node.js e os serve com Nginx;
- `api`: prepara as dependências Python em um estágio de build e executa a API Flask em outro estágio.

O navegador acessa apenas o frontend em `http://localhost:8080`. Ao solicitar um cálculo, o React chama `/api/calcular`; o Nginx encaminha essa rota para `http://api:5000/calcular` pela rede interna `calculadora-net`.

## Executar

```bash
docker compose up --build -d
```

Abra <http://localhost:8080>. Para acompanhar a execução:

```bash
docker compose config
docker compose ps
docker compose logs -f
```

Para encerrar:

```bash
docker compose down
```

## Leitura orientada dos arquivos

### 1. Build do frontend

No `frontend/Dockerfile`, o estágio `build` usa Node.js, instala dependências e produz `/app/dist`. O estágio `runtime` usa Nginx e recebe somente `/app/dist`. Node.js, código-fonte e `node_modules` não são copiados para a imagem final.

### 2. Build da API

No `api/Dockerfile`, o estágio `builder` cria o ambiente virtual e instala as dependências. O estágio `runtime` recebe somente o ambiente virtual e o arquivo da aplicação. Ferramentas usadas no preparo ficam fora da imagem final.

### 3. Comunicação entre os containers

O `compose.yaml` conecta os dois serviços a `calculadora-net`. Dentro dessa rede, o nome do serviço `api` funciona como endereço da API. O Nginx conhece esse endereço em `frontend/nginx.conf`.

Não há uma instrução de “rede Docker” nos Dockerfiles porque ela depende do ambiente que executará os containers. Um mesmo Dockerfile pode ser usado com outra rede, em outro Compose ou com `docker run --network`.

## Critérios para discutir

1. Localize os dois `COPY --from=...` e descreva o artefato transferido em cada um.
2. Verifique que apenas `frontend` possui `ports:`; `api` usa `expose:` apenas como documentação da porta interna.
3. Explique por que o React usa a rota relativa `/api/calcular`, em vez de `http://localhost:5000`.
4. Interrompa apenas a API (`docker compose stop api`) e observe o comportamento do frontend e dos logs.
5. Reconstrua depois de uma alteração em `api/app.py` e discuta quais camadas podem ser reutilizadas pelo cache.
