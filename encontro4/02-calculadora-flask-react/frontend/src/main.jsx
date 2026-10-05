import { useState } from "react";
import { createRoot } from "react-dom/client";
import "./style.css";

function App() {
  const [a, setA] = useState(10);
  const [b, setB] = useState(2);
  const [operacao, setOperacao] = useState("somar");
  const [resultado, setResultado] = useState("Informe os valores e calcule.");

  async function calcular(event) {
    event.preventDefault();
    setResultado("Consultando a API pela rede interna...");

    try {
      const resposta = await fetch("/api/calcular", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ a, b, operacao }),
      });
      const dados = await resposta.json();
      setResultado(resposta.ok ? `${a} ${operacao} ${b} = ${dados.resultado}` : dados.erro);
    } catch (erro) {
      setResultado(`Não foi possível consultar a API: ${erro.message}`);
    }
  }

  return (
    <main>
      <p className="etiqueta">React → Nginx → Flask</p>
      <h1>Calculadora em Docker</h1>
      <p>O navegador fala com o frontend; o Nginx fala com a API pelo nome <code>api</code>.</p>
      <form onSubmit={calcular}>
        <label>Primeiro número<input type="number" value={a} onChange={(e) => setA(e.target.value)} required /></label>
        <label>Operação
          <select value={operacao} onChange={(e) => setOperacao(e.target.value)}>
            <option value="somar">somar</option><option value="subtrair">subtrair</option>
            <option value="multiplicar">multiplicar</option><option value="dividir">dividir</option>
          </select>
        </label>
        <label>Segundo número<input type="number" value={b} onChange={(e) => setB(e.target.value)} required /></label>
        <button>Calcular</button>
      </form>
      <output aria-live="polite">{resultado}</output>
    </main>
  );
}

createRoot(document.getElementById("root")).render(<App />);
