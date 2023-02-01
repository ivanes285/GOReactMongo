import React from "react";
import { Link } from "react-router-dom";
export interface NavbarInterface {}

const Navbar: React.FC<NavbarInterface> = () => {
  
  return (
    <div className="bg-slate-800 flex justify-between px-20 py-4 text-xl items-center flex-wrap ">
      <Link to="/">
        <h1 className="text-2xl font-bold">React and Go server </h1>
      </Link>
      <nav className=" gap-6 flex flex-row flex-wrap">
        <Link className=" hover:border-b border-red-500 " to="/">
          Home
        </Link>
        <Link className=" hover:border-b border-red-500 " to="/users">
         Users
        </Link>
      </nav>
    </div>
  );
};

export default Navbar;
