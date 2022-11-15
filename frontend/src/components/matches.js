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
            jsonResponse.sort(function(a,b){
                return new Date(b.created_at) - new Date(a.created_at);
            });
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
        <div className="container">
            <div className="row">
                <div className="col-md-4">
                    <Button variant="contained" color="success" onClick={ handleCreateMatch }>
                        Create Match
                    </Button>
                </div>
                <div className="col-md-4">
                    {
                        (username === 'admin') &&
                        <Button variant="contained" color="success" onClick={ handleViewMetrics }>
                            View Metrics
                        </Button>
                    }
                </div>
                <div className="col-md-4">
                    <Button variant="contained" color="error" onClick={ handleLogout }>
                        Logout
                    </Button>
                </div>
            </div>
            <div className="row center">
                <div className="md-col-12 matches-list">
                    {
                        (matches.length > 0) ?
                            matches.map((match, index) => {
                                return <MatchCard match={match} key={index}/>
                            }) :
                            <strong className="match-list-empty">There is no matches created yet!</strong>
                    }
                </div>
            </div>
        </div>
    )
}

export default Matches;
