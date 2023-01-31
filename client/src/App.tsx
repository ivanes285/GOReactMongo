import { useState } from "react";
import axios from "axios";

function App() {

  const [name, setName] = useState<string>("");

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const response = await axios.post(import.meta.env.VITE_URLSERVER+"/users"); //usando variable de entorno en vite
    console.log(response.data);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  };

  return (
    <div>
      <form onSubmit={handleSubmit }>
        <input type="text" name="name" placeholder="Ingresa tu nombre" onChange={handleChange} />
        <button>Guardar</button>
      </form>
    </div>
  );
}

export default App;
