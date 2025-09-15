import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import {Badge, Card, Col, Container, Row, Table} from "react-bootstrap";

export function Bulletin() {
    const { etudiant, annee } = useParams();

    const [bulletin, setBulletin] = useState({
        liste_cours: []
    })
    const [isLoaded, setIsLoaded] = useState(false)


    useEffect(() => {
        fetch("http://localhost:8000/api/v1/etudiants/" + etudiant + "/annee/" + annee + "/bulletin")
            .then(res => res.json())
            .then(
                (result) => {
                    setBulletin(result.data);
                    setIsLoaded(true);
                })
    }, [])

    return (
        <Container className="my-5">
            {isLoaded ?
            <Row className="justify-content-center">
                <Col md={10}>
                    <Card className="shadow">
                        <Card.Header className="bg-primary text-white">
                            <h2 className="text-center mb-0">Résultats de l'examen</h2>
                        </Card.Header>
                        <Card.Body>
                            <Row className="mb-4">
                                <Col md={6}>
                                    <h4>Élève : {bulletin.prenom} {bulletin.nom}</h4>
                                    <p>Matricule : {bulletin.matricule}</p>
                                </Col>
                                <Col md={6} className="text-end">
                                    <h4>Année scolaire : {bulletin.annee == 1 ? bulletin.annee + "ère" : bulletin.annee+"ème"}</h4>
                                </Col>
                            </Row>

                            <Table striped bordered hover responsive>
                                <thead className="bg-light">
                                <tr>
                                    <th>Matière</th>
                                    <th>Note /20</th>
                                    <th>Credit</th>
                                    <th>titulaire</th>
                                </tr>
                                </thead>
                                <tbody>
                                {bulletin.liste_cours.map((cour) => (
                                    <tr key={cour.mnemonique}>
                                        <td>{cour.intitule}</td>
                                        <td>{cour.note}</td>
                                        <td>{cour.credit}</td>
                                        <td>{cour.titulaire}</td>
                                    </tr>
                                ))}
                                </tbody>
                            </Table>

                            <Row className="mt-4">
                                <Col md={6}>
                                    <Card className="text-center">
                                        <Card.Body>
                                            <Card.Title>Moyenne générale</Card.Title>
                                            <Card.Text className="fs-2 fw-bold">
                                                {bulletin.weighted_average}/20
                                            </Card.Text>
                                        </Card.Body>
                                    </Card>
                                </Col>
                                <Col md={6}>
                                    <Card className="text-center">
                                        <Card.Body>
                                            <Card.Title>Résultat</Card.Title>
                                            <Badge bg={bulletin.is_success ? "success" : "danger"} className="fs-4 p-2">
                                                {bulletin.is_success ? "Réussite" : "Echec"}
                                            </Badge>
                                        </Card.Body>
                                    </Card>
                                </Col>
                            </Row>
                        </Card.Body>
                    </Card>
                </Col>
            </Row>
                :
                <div>Loading ...</div> }
        </Container>

    )

}