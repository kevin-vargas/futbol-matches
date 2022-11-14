import React, {useState} from 'react'
import {useParams} from "react-router-dom";
import {useEffect} from 'react';
import matchService from "../services/matchService";
import MatchForm from "./matchForm";

const Match = (props) => {
    const [match, setMatch] = useState({})
    const { id } = useParams()

    console.log(id)

    useEffect(() => {
        matchService.getMatchById(id).then(response => response.json()).then(jsonResponse => {
            setMatch(jsonResponse)
        }).catch( error => {
            console.log(error)
        });
    }, [])

    return (
        <div>
            {
                (match._id) && <MatchForm match={match} />
            }

        </div>
    )
}

export default Match;
