import React from "react";
import ImageHome from "../../assets/go.gif";
import { useNavigate } from "react-router-dom";
export interface HomeInterface {}

const Home: React.FC<HomeInterface> = () => {
	const navigate = useNavigate();
  return (
    <div className="h-full flex flex-col items-center bg-[#fbfbfb] text-gray-400 ">
      <h2 className="my-14 font-bold text-4xl">Welcome to CRUD users by GOLANG and React</h2>
      <img className="max-w-3xl h-2/4" src={ImageHome} alt="Not found" />
      <button className="hover:text-gray-900 text-lg " onClick={() => navigate("/users")}>{`Crea un usuario->`}</button>
    </div>
  );
};

export default Home;
