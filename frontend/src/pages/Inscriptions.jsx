import {useEffect, useState} from "react";
import {Link} from "react-router-dom";
import {Button, Table} from "react-bootstrap";

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
                    <Table striped bordered hover responsive>
                        <thead className="bg-primary text-white">
                        <tr>
                            <th>Matricule</th>
                            <th>Nom</th>
                            <th>Prénom</th>
                            <th>Année</th>
                        </tr>
                        </thead>
                        <tbody>
                        {inscriptions.map((inscription) => (
                            <tr key={inscription.matricule+"_"+inscription.annee_etude}>
                                <td>{inscription.matricule}</td>
                                <td>{inscription.nom}</td>
                                <td>{inscription.prenom}</td>
                                <td>{inscription.annee_etude}</td>
                                <td>
                                    <Link className="nav-link" to={`/bulletin/${inscription.matricule}/${inscription.annee_etude}`}>
                                        <Button variant="info" size="sm">
                                            Voir détails
                                        </Button>
                                    </Link>
                                </td>
                            </tr>
                        ))}
                        </tbody>
                    </Table>
                    :
                    <div>Nothing</div>
                }
            </ul>
        </div>
    )
}