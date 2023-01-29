

function App() {
  const fechatData = async () => {
    const response = await fetch("http://127.0.0.1:3000/api/v1/users");
    const data = await response.json();
   console.log(data);
  };

  return (
    <div>
      <h1>Hello World</h1>
      <button onClick={() => fechatData()}>Obtener Datos</button>
    </div>
  );
}

export default App
