import React from "react";
import ImageNoFound from "../../assets/notfound.gif";
import { useNavigate } from "react-router-dom";
export interface NotFoundInterface {}
const NotFound: React.FC<NotFoundInterface> = () => {
	const navigate = useNavigate();
  return (
    <div className="h-full flex flex-col items-center bg-white text-gray-400 ">
      <h2 className="my-14 font-bold text-4xl">Page Not Found</h2>
      <img className="max-w-sm h-95  rounded-2xl  shadow-4xl" src={ImageNoFound} alt="Not found" />
	   <button className="hover:text-gray-900 text-lg" onClick={()=>navigate('/')}>{`Backt to Home->`}</button>
    </div>
  );
};

export default NotFound;
