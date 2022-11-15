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
        <div className="container match-card">
            <div className="row match-card-header">
                <div className="col-md-12">
                    { match.description }
                </div>
            </div>
            <hr />
            <div className="row match-card-header">
                <div className="col-md-12">
                    { match.place }
                </div>
            </div>
            <div className="row match-card-body">
                <div className="col-md-4">
                    { `${match.date.split("T")[0].split("-")[2]}-${match.date.split("T")[0].split("-")[1]}-${match.date.split("T")[0].split("-")[0]}` }
                </div>
                <div className="col-md-4">
                    { match.time } hs
                </div>
                <div className="col-md-4">
                    ${ match.price }
                </div>
            </div>
            <hr />
            <div className="row match-card-footer">
                <div className="col-md-4">
                    Owner: { match.owner }
                </div>
                <div className="col-md-8">
                    Id: { match._id }
                </div>
            </div>
            <div className="row ">
                <div className="col-md-12 ">
                    <Button variant="contained" onClick={ handleClick}> Click to join! </Button>
                </div>
            </div>
        </div>
    );
}
