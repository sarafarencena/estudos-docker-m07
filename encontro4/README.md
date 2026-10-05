# Encontro 04 — Build multi-stage, aplicações web e Docker Compose

## Objetivo da instrução

Nesta instrução, a turma acompanha a evolução de uma imagem mínima até uma aplicação web com frontend e API. O foco não é apenas executar os comandos: é identificar **o que é produzido em cada estágio**, **o que atravessa os estágios de build** e **como os containers se comunicam em execução**.

> Regra central: multi-stage é um recurso do **Dockerfile**. O `compose.yaml` usa a imagem já construída para criar e conectar containers; ele não substitui os estágios `FROM ... AS ...` e `COPY --from=...`.

## Roteiro

| Etapa | Pasta | Pergunta que orienta a demonstração |
| --- | --- | --- |
| 1. Imagem mínima | `01-hello-c-multistage/` | Por que a imagem que executa não precisa ter compilador C? |
| 2. Calculadora web | `02-calculadora-flask-react/` | O que passa do build para a imagem final e o que passa entre os containers? |
| 3. Desafio 5C | `03-desafio-go-compose-5c/` | Como reconstruir o backend em Go sem alterar o frontend? |

## Dois fluxos diferentes

```text
BUILD (Dockerfile)                         EXECUÇÃO (Docker Compose)

builder ── COPY --from=builder ──> runtime frontend ── rede Compose ──> api
  artefato: binário ou dist/                 navegador -> :8080          :5000 interno
```

- No **build**, o Docker copia somente artefatos escolhidos de um estágio para outro. O estágio `builder` não vira um container em execução.
- Na **execução**, os containers já criados conversam por uma rede. No exemplo, o Nginx usa o nome do serviço `api` para alcançar a API.
- `EXPOSE` apenas documenta a porta da imagem. Ele não publica porta no computador e não cria rede.
- A rede é declarada em `compose.yaml`, em `networks:`. A publicação para o computador também é declarada ali, em `ports:`.

## Evidências esperadas

Ao final, cada equipe deve conseguir mostrar:

1. o trecho do Dockerfile que transfere o artefato entre estágios;
2. a imagem final sem compilador, Node.js ou ferramentas de build desnecessárias;
3. o frontend acessível no navegador e a API sem porta publicada diretamente;
4. o nome do serviço usado na comunicação interna;
5. `docker compose config`, `docker compose ps` e os logs como evidências da decisão tomada.

## Ciclo 5C durante a prática

Quando houver dúvida, registre-a como uma hipótese técnica — não como pedido de solução pronta:

1. **Contextualizar:** qual artefato ou qual conexão precisa existir?
2. **Consultar:** formule uma dúvida pequena, por exemplo: “`COPY --from` leva a imagem inteira ou somente o caminho indicado?”
3. **Confrontar:** compare a resposta com a documentação e com um build ou teste observado.
4. **Construir:** aplique a decisão no Dockerfile ou no Compose.
5. **Comprovar:** execute, registre o resultado e explique por que a decisão atende ao requisito.

Use IA como hipótese a ser verificada. A entrega deve demonstrar a autoria da equipe, a fonte consultada e uma evidência de execução.
