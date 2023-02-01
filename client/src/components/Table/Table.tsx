import React, { useState } from "react";
import { IUser } from "../../interfaces";
import { useNavigate } from "react-router-dom";

export interface TableInterface {
  users: IUser[];
  handleDelete: (id: string) => void;
}

const Table: React.FC<TableInterface> = ({ users , handleDelete}) => {
  const navigate = useNavigate();
  let [isOpen, setIsOpen] = useState(true)
  return (
    <div className="w-7/12 text-lg  my-12">
      <table className="table-fixed w-full mx-auto">
        <thead>
          <tr className="h-16 border-y-2 border-slate-600 font-bold text-xl">
            <th>User</th>
            <th>Options</th>
          </tr>
        </thead>
        {users.map((user) => (
          <tbody key={user.id}>
            <tr className={`h-14 border-b-2 border-slate-600`}>
              <td>{user.name}</td>
              <td>
                <button className="bg-red-600 py-1 px-4 rounded-lg" onClick={() => handleDelete(user.id!)}>
                  Delete
                </button>
                {"  "}
                <button className="bg-green-500 py-1 px-4 rounded-lg"onClick={()=> navigate(`/edit/${user.id}`)}>Edit</button>
                {"  "}
                <button className="bg-violet-500 py-1 px-4 rounded-lg"onClick={()=>console.log("No hace nada XD")}>View</button>
              </td>
            </tr>
          </tbody>
        ))}
      </table>
    </div>
  );
};

export default Table;
