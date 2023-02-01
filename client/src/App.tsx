import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { NotFound, UserPage ,Home } from "./pages";
import { Navbar } from "./components";


function App() {
  return (
    <Router>
      <div className=" h-screen text-white overflow-auto">
      <Navbar/>
      <Routes>
        <Route path="*" element={<NotFound/>} />
        <Route path="/" element={<Home/>} />
        <Route path="/users" element={<UserPage/>} />
        <Route path="/edit/:id" element={<UserPage/>} />
      </Routes>
      </div>
    </Router>
  );
}

export default App;
