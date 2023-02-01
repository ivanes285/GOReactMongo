import axios from "axios";
import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { Notification, Table } from "../../components";
import { IUser } from "../../interfaces";
export interface UserPageInterface {}

const UserPage: React.FC<UserPageInterface> = () => {
  const params = useParams();
  const navigate = useNavigate();
  const initialState = {
    name: "",
  };
  const [user, setUser] = useState<IUser>(initialState);
  const [users, setUsers] = useState<IUser[]>([]);

  const fetchUsers = async () => {
    try {
      const response = await axios.get(import.meta.env.VITE_URLSERVER + "/users");
      setUsers(response.data.users || []); //valido si no vienen datos
      setUser(initialState);
    } catch (error: any) {
      toast.error(error.response.data.message);
    }
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    try {
      e.preventDefault();
      if (params.id) {
        const response = await axios.put(import.meta.env.VITE_URLSERVER + `/users/${params.id}`, user); //usando variable de entorno en vite
        toast.success(response.data.message);
        navigate("/users");
      } else {
        await axios.post(import.meta.env.VITE_URLSERVER + "/users", user); //usando variable de entorno en vite
        toast.success("User created successfully");
      }
      setUser(initialState);
      fetchUsers();
    } catch (error: any) {
      toast.error(error.response.data.message);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUser({ ...user, [e.target.name]: e.target.value });
  };

  const handleDelete = async (id: string) => {
    try {
      const response = await axios.delete(import.meta.env.VITE_URLSERVER + `/users/${id}`);
      toast.success(response.data.message);
      fetchUsers();
    } catch (error: any) {
      toast.error(error.response.data.message);
    }
  };

  useEffect(() => {
    if (params.id) {
      const fetchUser = async () => {
        try {
          const response = await axios.get(import.meta.env.VITE_URLSERVER + `/users/${params.id}`);
          setUser(response.data.user);
        } catch (error: any) {
          toast.error(error.response.data.message);
        }
      };
      fetchUser();
    } else {
      fetchUsers();
    }
  }, [params.id]);

  return (
    <div className=" w-9/12 mt-10 mx-auto flex  flex-col text-center items-center text-white justify-center">
      <div className="flex flex-col ">
        <form
          onSubmit={handleSubmit}
          className="border-2 border-slate-600 rounded-md py-16 px-8  max-w-xl mx-auto flex flex-col gap-6 font-serif text-lg  items-center"
        >
          <span className="font-bold text-3xl">Formulario</span>
          <input
            className="py-2 px-6  rounded-md text-slate-900"
            type="text"
            name="name"
            value={user.name}
            placeholder="Ingresa tu nombre"
            onChange={handleChange}
          />
          <button className="bg-sky-800 w-full rounded-md py-2  mx-auto">{params.id ? "Editar" : "Guardar"}</button>
        </form>
      </div>
      <Table users={users} handleDelete={handleDelete} />
      <Notification position={"top-right"} />
    </div>
  );
};

export default UserPage;
