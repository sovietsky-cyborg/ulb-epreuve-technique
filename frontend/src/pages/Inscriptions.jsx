import {useEffect, useState} from "react";
import {Link} from "react-router-dom";

export function Inscriptions() {
    const [inscriptions, setInscriptions] = useState({})
    const [isLoaded, setIsLoaded] = useState(false)


    useEffect(() => {
        fetch("http://localhost:8000/api/v1/liste-inscriptions")
            .then(res => res.json())
            .then(
                (result) => {
                    setInscriptions(result.data);
                    setIsLoaded(true);
                }
            )
    }, []);
    return (
        <div>
            <h1>Inscriptions</h1>;
            <ul className="list-group">
                {isLoaded ?
                    inscriptions.map(inscription => (
                        <li className="list-group-item">
                            <Link className="nav-link" to={`/bulletin/${inscription.matricule}/${inscription.annee_etude}`}>
                                {inscription.nom} {inscription.prenom} Année: {inscription.annee_etude}
                            </Link>
                        </li>
                    ))
                    :
                    <div>Nothing</div>
                }
            </ul>
        </div>
    )
}