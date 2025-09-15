import './App.css'
import {BrowserRouter, Routes, Route, Link, useParams} from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.css';
import {useEffect, useState} from "react";
import {Bulletin} from "./pages/Bulletin.jsx"
import {Inscriptions} from "./pages/Inscriptions.jsx"
import {Cours} from "./pages/Cours.jsx"
function Home() {
    return <h1>Home Page</h1>;
}

        function App() {

        return (
        <BrowserRouter>
        <div>
            <nav className="navbar navbar-expand-lg navbar-light bg-light">
                <div className="collapse navbar-collapse" id="navbarNav">
                    <ul className="navbar-nav">
                        <li className="nav-item">
                            <Link className="nav-link" to="/">Home</Link>
                        </li>
                        <li className="nav-item">
                            <Link className="nav-link" to="/cours">Cours</Link>
                        </li>
                        <li className="nav-item">
                            <Link className="nav-link" to="/notes">Notes</Link>
                        </li>
                        <li className="nav-item">
                            <Link className="nav-link" to="/inscriptions">Inscriptions</Link>
                        </li>
                    </ul>
                </div>
            </nav>
        </div>
    {/* Routes */}
        <Routes>
            <Route path="/" element={<Home/>}/>
            <Route path="/cours" element={<Cours/>}/>
            <Route path="/inscriptions" element={<Inscriptions/>}/>
            <Route path="/bulletin/:etudiant/:annee" element={<Bulletin/>}/>
        </Routes>

    </BrowserRouter>
)
}

export default App
