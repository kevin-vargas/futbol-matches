import React, {useEffect, useState} from 'react'
import {verifySession} from "../utils/utils";
import matchService from "../services/matchService";
import MatchCard from "./matchCard";
import {Stack} from "@mui/material";
import Button from "@mui/material/Button";
import {useNavigate} from "react-router-dom";

const Matches = (props) => {

    verifySession()
    const username = localStorage.getItem("username")
    const [matches, setMatches] = useState([])
    const navigation = useNavigate();

    useEffect(() => {
        matchService.getMatches().then(response => response.json()).then(jsonResponse => {
            console.log("Response json: ", jsonResponse);
            setMatches(jsonResponse);
        });
    }, [])

    const handleLogout = () => {
        navigation("/")
    }

    const handleCreateMatch = () => {
        navigation("/match-form")
    }

    const handleViewMetrics = () => {
        navigation("/metrics")
    }

    return (
        <section className="row">
            <div className="md-col-12 matches-options">
                <Stack direction="row" spacing={50}>
                    <Button variant="contained" color="success" onClick={ handleCreateMatch }>
                        Create Match
                    </Button>

                    {
                        (username === 'admin') &&
                        <Button variant="contained" color="success" onClick={ handleViewMetrics }>
                            View Metrics
                        </Button>
                    }

                    <Button variant="contained" color="error" onClick={ handleLogout }>
                        Logout
                    </Button>
                </Stack>
            </div>

            <div className="md-col-12 matches-list">
                {
                    (matches.length > 0) ?
                        matches.map((match, index) => {
                            return <MatchCard match={match} key={index}/>
                        }) :
                        <strong className="match-list-empty">There is no matches created yet!</strong>
                }
            </div>
        </section>
    )
}

export default Matches;
