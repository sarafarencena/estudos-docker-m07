from flask import Flask, jsonify, request

app = Flask(__name__)


@app.post("/calcular")
def calcular():
    dados = request.get_json(silent=True) or {}
    operacao = dados.get("operacao")

    try:
        a = float(dados["a"])
        b = float(dados["b"])
    except (KeyError, TypeError, ValueError):
        return jsonify(erro="Informe dois números válidos em 'a' e 'b'."), 400

    operacoes = {
        "somar": lambda: a + b,
        "subtrair": lambda: a - b,
        "multiplicar": lambda: a * b,
        "dividir": lambda: a / b,
    }

    if operacao not in operacoes:
        return jsonify(erro="Operação inválida."), 400
    if operacao == "dividir" and b == 0:
        return jsonify(erro="Não é possível dividir por zero."), 400

    return jsonify(operacao=operacao, a=a, b=b, resultado=operacoes[operacao]())


@app.get("/health")
def health():
    return jsonify(status="ok", servico="calculadora-flask")
