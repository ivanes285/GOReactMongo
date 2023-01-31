import axios from "axios";
import { useEffect, useState } from "react";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { Notification } from "./components/Notification";


interface User {
  id?: string;
  name: string;
}

function App() {
  const initialState = {
    name: "",
  };

  const [user, setUser] = useState<User>(initialState);
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    try {
      e.preventDefault();
      await axios.post(import.meta.env.VITE_URLSERVER + "/users", user); //usando variable de entorno en vite
      setUser(initialState);
      toast.success("User created successfully");
    } catch (error: any) {
      toast.error(error.response.data.message);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUser({ ...user, [e.target.name]: e.target.value });
  };

  useEffect(() => {
    const fetchUsers = async () => {
      setLoading(true);
      const response = await axios.get(import.meta.env.VITE_URLSERVER + "/users");
      setUsers(response.data.users || []);  //valido si no vienen datos
      setLoading(false);
    };
    fetchUsers();
  }, [user]);

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input type="text" name="name" value={user.name} placeholder="Ingresa tu nombre" onChange={handleChange} />
        <button>Guardar</button>
        <Notification position={"top-right"} />
      </form>
      {!loading? users.length ? users.map((user) => <h2 key={user.id}>{user.name}</h2>) : <h2>No hay usuarios registrados</h2>: <h2>Cargando...</h2>}
    </div>
  );
}

export default App;
