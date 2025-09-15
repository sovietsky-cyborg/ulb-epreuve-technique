import {useEffect, useState} from "react";

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
                    console.log("result", result.data);
                }
            )
    }, []);

    return (
        <div>
            <h1>Cours</h1>
            <ul className="list-group">
                {isLoaded ?
                    cours.map(cour => (
                        <li className="list-group-item">{cour.intitule}</li>
                    ))
                    :
                    <div>Loading..</div>
                }
            </ul>
        </div>
    );
}