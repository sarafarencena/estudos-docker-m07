# Desafio 5C — Calculadora com backend em Go

## Contexto

Uma versão funcional da calculadora já foi demonstrada em `../02-calculadora-flask-react/`. Agora a empresa decidiu reimplementar **somente a API** em Go. O frontend React é o mesmo apresentado antes: não copie, altere ou recrie seus arquivos neste desafio.

A equipe deve criar:

1. `api/Dockerfile`, usando build multi-stage para compilar o backend Go fornecido e executar somente o binário na imagem final;
2. `compose.yaml`, para subir o frontend já apresentado e a nova API como dois serviços conectados;
3. `registro-5c.md`, com a evidência do raciocínio e da validação.

Não há solução pronta para esses dois arquivos. A API Go completa já está em `api/`; a equipe deve partir dela e das evidências da demonstração anterior para criar a infraestrutura de build e execução.

## Requisitos de aceitação

| Requisito | Evidência esperada |
| --- | --- |
| A API é escrita em Go | `api/main.go` é compilado no estágio de build. |
| O Dockerfile é multi-stage | Há pelo menos um estágio de compilação e um estágio final sem ferramentas de compilação Go. |
| O frontend é o mesmo da demonstração 2 | O Compose usa o frontend existente em `../02-calculadora-flask-react/frontend`; não há frontend duplicado nesta pasta. |
| A interface funciona | Em `http://localhost:8081`, uma operação retorna o resultado da API Go. |
| A API não é exposta ao computador | Somente o frontend tem `ports:`. |
| Os serviços se comunicam | Ambos participam de uma rede declarada no Compose; Nginx resolve o serviço chamado `api`. |
| O projeto é demonstrável | `docker compose config`, `docker compose ps` e logs sustentam a explicação da equipe. |

## Contrato que a API Go deve preservar

O frontend chama `POST /api/calcular`. Como o Nginx remove o prefixo `/api/`, o servidor Go deve atender `POST /calcular` na porta `5000`.

Entrada JSON:

```json
{ "a": 10, "b": 2, "operacao": "dividir" }
```

Operações aceitas: `somar`, `subtrair`, `multiplicar` e `dividir`.

Saída de sucesso (status 200):

```json
{ "operacao": "dividir", "a": 10, "b": 2, "resultado": 5 }
```

Para operação inválida, valores inválidos ou divisão por zero, retorne JSON de erro com status 400. O endpoint `GET /health` deve retornar status 200 para facilitar a inspeção.

## Pistas técnicas, sem solução pronta

- O primeiro estágio do Dockerfile pode usar uma imagem oficial Go para executar a compilação.
- O estágio final deve receber apenas o binário compilado e os arquivos realmente necessários para executá-lo.
- `EXPOSE 5000` documenta a porta da imagem; ele não substitui `ports`, `expose` ou `networks` do Compose.
- A declaração da rede vai em `compose.yaml`. O Dockerfile não conecta containers a uma rede.
- O nome do serviço do backend precisa ser `api`, pois o Nginx do frontend usa esse nome.
- O contexto de build do frontend deve apontar para a pasta da demonstração anterior; o contexto da API é `./api`.

## Roteiro 5C

1. **Contextualizar:** descrevam o que precisa ser preservado para o frontend continuar funcionando.
2. **Consultar:** registrem uma pergunta delimitada sobre build multi-stage ou sobre a rede do Compose. Não peçam uma solução integral.
3. **Confrontar:** consultem documentação oficial e executem ao menos um teste que confirme ou refute a hipótese.
4. **Construir:** implementem os dois arquivos solicitados.
5. **Comprovar:** executem `docker compose up --build -d`, testem ao menos uma operação válida e uma entrada de erro, e registrem logs/configuração.

## Comandos de validação

```bash
docker compose config
docker compose up --build -d
docker compose ps
docker compose logs --tail=50
docker compose down
```
