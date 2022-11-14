import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import {Button} from "@mui/material";
import {useNavigate} from "react-router-dom";

export default function MatchCard(props) {
    const match = props.match
    const navigation = useNavigate();

    const handleClick = (event) => {

        navigation("/match/"+match._id )
    }

    return (
        <Card className="match-card" style={{backgroundColor: "transparent"}} >
            <CardContent>
                <Typography variant="h5" component="div">
                    { match.description }
                </Typography>
                <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
                   Created by: { match.owner } - Match Id: { match._id }
                </Typography>
                <Typography variant="h5" component="div">
                    Place: { match.place }
                </Typography>
                <Typography sx={{ mb: 1.5 }} color="text.secondary">
                    Date: { match.date.split("T")[0] }
                </Typography>
                <Typography sx={{ mb: 1.5 }} color="text.secondary">
                    Time: { match.time }
                </Typography>
                <Typography variant="body2">
                    Price per player: { match.price }
                </Typography>
                <Button variant="text" onClick={ handleClick}> Click to join! </Button>
            </CardContent>
        </Card>
    );
}
