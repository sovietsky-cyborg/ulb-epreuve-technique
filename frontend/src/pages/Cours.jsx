import {useEffect, useState} from "react";
import {Button, Table} from "react-bootstrap";
import {Link} from "react-router-dom";

export function Cours() {

    const [cours, setCours] = useState([])
    const [isLoaded, setIsLoaded] = useState(false)

    useEffect(() => {
        fetch("http://localhost:8000/api/v1/liste-cours")
            .then(res => res.json() )
            .then(
                (result) => {
                    setCours(result.data);
                    setIsLoaded(true);
                }
            )
    }, []);

    return (
    <div>
        <h1>Cours</h1>;
        <ul className="list-group">
            {isLoaded ?
                <Table striped bordered hover responsive>
                    <thead className="bg-primary text-white">
                    <tr>
                        <th>Mnemonique</th>
                        <th>Intitule</th>
                        <th>Credit</th>
                        <th>Titulaire</th>
                    </tr>
                    </thead>
                    <tbody>
                    {cours.map((cour) => (
                        <tr key={cour.mnemonique}>
                            <td>{cour.mnemonique}</td>
                            <td>{cour.intitule}</td>
                            <td>{cour.credit}</td>
                            <td>{cour.titulaire}</td>
                        </tr>
                    ))}
                    </tbody>
                </Table>
                :
                <div>Nothing</div>
            }
        </ul>
    </div>
    );
}