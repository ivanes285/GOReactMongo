import axios from "axios";
import { useState } from "react";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { Notification } from "./components/Notification";

function App() {

  const [name, setName] = useState<string>("");

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    try {
      e.preventDefault();
      const response = await axios.post(import.meta.env.VITE_URLSERVER+"/users",{name}); //usando variable de entorno en vite
      toast.success(response.data.message);
    } catch (error: any) {
      toast.error(error.response.data.message);
    }
   
    
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  };

  return (
    <div>
      <form onSubmit={handleSubmit }>
        <input type="text" name="name" placeholder="Ingresa tu nombre" onChange={handleChange} />
        <button>Guardar</button>
        <Notification position={"top-right"} />
      </form>
    </div>
  );
}

export default App;
