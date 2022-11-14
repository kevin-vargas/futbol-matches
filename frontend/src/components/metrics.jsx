import React, {useEffect, useState} from 'react'
import metricService from "../services/metricsService";
import {verifySession} from "../utils/utils";
import {Button} from "@mui/material";
import {useNavigate} from "react-router-dom";

const Metrics = () => {
    verifySession()

    const [matchesCreated, setMatchesCreated] = useState('')
    const [playersJoined, setPlayersJoined] = useState('')
    const interval = '2h'

    const navigation = useNavigate()

    const handleCancel = () => {
        navigation("/matches")
    }

    useEffect(() => {
        metricService.getMetric('created_matches', interval)
            .then(response => response.json())
            .then(jsonResponse => {
            console.log("Response json: ", jsonResponse);
                setMatchesCreated(jsonResponse.count);
        });

        metricService.getMetric('annotated_users', interval)
            .then(response => response.json())
            .then(jsonResponse => {
                console.log("Response json: ", jsonResponse);
                setPlayersJoined(jsonResponse.count);
            });
    }, [])

    return (
        <div>
            Matches created last 2 hours: { matchesCreated || '' }
            <br />
            Players Joined last 2 hours: { playersJoined || ''}
            <br />
            <Button
                fullWidth
                variant="contained"
                color="error"
                sx={{mt: 3, mb: 2}}
                onClick={handleCancel}
            >
                Cancel
            </Button>
        </div>
    )
}

export default Metrics;
