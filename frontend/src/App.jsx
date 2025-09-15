import './App.css'
import {BrowserRouter, Routes, Route, Link, useParams} from "react-router-dom";
import 'bootstrap/dist/css/bootstrap.css';
import {useEffect, useState} from "react";
import {Bulletin} from "./pages/Bulletin.jsx"
import {Inscriptions} from "./pages/Inscriptions.jsx"
import {Cours} from "./pages/Cours.jsx"
function Home() {
    return (
        <div>
        <h1>Home Page</h1>
            <p>At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium
            voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi
            sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt
            mollitia animi, id est laborum et dolorum fuga.</p>
            <br/>
            <p>Et harum quidem rerum facilis est et
            expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil
            impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est,</p>
            <p>omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum
            necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae.
                Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores
            alias consequatur aut perferendis doloribus asperiores repellat.</p>
        </div>
);
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
