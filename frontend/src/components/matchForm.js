import React, {useState} from 'react'
import {Button, TextField} from "@mui/material";
import {useNavigate} from "react-router-dom";
import matchService from "../services/matchService";
import Swal from 'sweetalert2'
import {verifySession} from "../utils/utils";
import BasicModal from "./playerModal";


const MatchForm = (props) => {
    verifySession()
    const navigation = useNavigate()
    const creating = (!props.match) ? true : false
    const [match, setMatch] = useState(props.match || {})

    console.log(props.match)

    const handleSubmit = (event) => {
        event.preventDefault()
        const match = {}

        match.owner = localStorage.getItem("username")
        match.description = event.target.description.value
        match.date = event.target.date.value
        match.place = event.target.place.value
        match.time = event.target.time.value
        match.format = event.target.format.value
        match.maxplayers = event.target.maxplayers.value
        match.price = event.target.price.value

        matchService.saveMatch(match).then(response => {
            if(response.status === 201) {
                Swal.fire({
                    title: 'Match Created!',
                    icon: 'success',
                    confirmButtonText: 'Ok'
                }).then(r => console.log(r))

                navigation("/matches")
            }
            else {
                console.log("Error al crear el partido")
                navigation("/matches")
            }
        });
    }

    const handleCancel = () => {
        navigation("/matches")
    }

    const handleChange = (event) => {
        const matchChanged = {}
        matchChanged[event.target.name] = event.target.value
        const merged = {...match, ...matchChanged};
        setMatch(merged)

    }

    return (

        <div className="container">
            <form autoComplete="off" onSubmit={handleSubmit}>
                <div className="row ">
                    <TextField  className="col-md-6"
                        label= "Description"
                                variant="outlined"
                                name="description"
                                required
                                disabled={(!creating)}
                                value={match.description}
                                sx={{marginRight: 10}}
                                onChange={ handleChange }/>

                    <TextField className="col-md-6"
                               label="Place"
                               variant="outlined"
                               name="place"
                               required
                               value={match.place}
                               disabled={(!creating)}
                               onChange={ handleChange }/>
                </div>
                <br/>


                <div className="row ">
                        <TextField className="col-md-6 "
                        label="Date (yyyy-mm-dd)"
                                   variant="outlined"
                                   name="date"
                                   required
                                   value={(match.date)? match.date.split("T")[0] : ''}
                                   disabled={(!creating)}
                                   sx={{marginRight: 10}}
                                   onChange={ handleChange }/>


                    <TextField className="col-md-6" label="Time"
                               variant="outlined"
                               name="time"
                               required
                               value={match.time}
                               disabled={(!creating)}
                               onChange={ handleChange }/>

                </div>
                <br/>


                <div className="row ">
                        <TextField className="col-md-6"
                            label="Price per player"
                                   variant="outlined"
                                   name="price"
                                   required
                                   value={ (match.price) ? `${match.price}` : '' }
                                   disabled={(!creating)}
                                   sx={{marginRight: 10}}
                                   onChange={ handleChange }/>


                        <TextField className="col-md-6"
                            label="Format ( futbol 5, 7, 11, etc)"
                                   variant="outlined"
                                   name="format"
                                   required
                                   value={match.format}
                                   disabled={(!creating)}
                                   onChange={ handleChange }/>

                </div>


                <br />
                <div className="row">
                        <TextField className="col-md-6"
                            label="Max Players"
                                   variant="outlined"
                                   name="maxplayers"
                                   required
                                   value={match.maxPlayers}
                                   disabled={(!creating) ? true: false}
                                   sx={{marginRight: 10}}
                                   onChange={ handleChange }/>

                    {
                        (creating) ? '' :
                            <TextField className="col-md-6"
                                       label="Remaining Places"
                                       variant="outlined"
                                       name="remainingPlaces"
                                       required
                                       value={ (match && match.maxPlayers) ? (match.maxPlayers - match.startingPlayers.length - match.substitutePlayer.length) : ''}
                                       disabled={(match && match.maxPlayers)?true: false}
                                       onChange={ handleChange }
                            />
                    }
                </div>
                <br />

                    <div className={(creating)? "hide-div": ''}>
                            <span >Starting Players:
                            { (!match || !match._id) ? '' :

                                (match.startingPlayers.length <= 0 ) ? ' The List is empty' :
                                    match.startingPlayers.map( (player, index) => {
                                        return (
                                            <li key={index}>
                                                {player.name}
                                            </li>
                                        )
                                    })
                            }
                            </span>
                            <span >Substitute Players:
                            { ( !match || !match._id) ? '' :
                                (match.substitutePlayer.length <= 0 ) ? ' The List is empty' :
                                    match.substitutePlayer.map( (player, index) => {
                                        return (
                                            <li key={index}>
                                                {player.name}
                                            </li>
                                        )
                                    })
                            }</span>
                    </div>

                <br/>
                <div className="row">
                    <div className="col-md-6">
                        { (!creating && match && match._id) ? <BasicModal matchId={match._id} places={(match.maxPlayers - match.startingPlayers.length - match.substitutePlayer.length)}/> :
                            <Button type={ (!creating && match && match._id)? "button" : "submit"}
                                    fullWidth
                                    variant="contained"
                                    sx={{mt: 3, mb: 2}}
                            >
                                Create Match
                            </Button>
                        }

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
                </div>
            </form>
        </div>
    )
}

export default MatchForm;
