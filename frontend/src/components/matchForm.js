import React, {useState} from 'react'
import {Button, TextField} from "@mui/material";
import {useNavigate} from "react-router-dom";
import matchService from "../services/matchService";
import Swal from 'sweetalert2'
import {verifySession} from "../utils/utils";
import JoinPlayerModal from "./playerModal";


const MatchForm = (props) => {
    verifySession()
    const navigation = useNavigate()
    const creating = (!props.match) ? true : false
    const [match, setMatch] = useState(props.match || {})

    console.log("match: ", props.match)

    const inputStyle = { marginRight: 10, input: { color: 'white' }, label: { color: 'white' },
        "& .MuiInputBase-input.Mui-disabled": {
            WebkitTextFillColor: 'white'
        },
        '& .MuiFormLabel-root.Mui-disabled': {
            color: 'white',
        }}

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

        console.log("SAVE MATCH: ", match)

        matchService.saveMatch(match).then(response => response.json()).then(response => {
            if (response.status === 201) {
                Swal.fire({
                    title: 'Match Created!',
                    icon: 'success',
                    confirmButtonText: 'Ok'
                }).then(r => console.log(r))

                navigation("/matches")
            } else {
                Swal.fire({
                    title: "Error creating match. Verify the fields and try again.",
                    icon: 'error',
                    confirmButtonText: 'Ok'
                }).then(r => console.log(r))
                console.log("Error creating match: ", response.error)
                navigation("/matches")
            }
        }).catch( error => console.log("ERROR: ", error));
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
        <div className="container match-form">
            <form autoComplete="off" onSubmit={handleSubmit} className="form-padding">
                <div className="row">
                    <div className="col-md-6">
                        <TextField
                            label="Description"
                            variant="outlined"
                            name="description"
                            required
                            disabled={(!creating)}
                            value={match.description}
                            sx={ inputStyle }
                            onChange={handleChange}/>
                    </div>
                    <div className="col-md-6">
                        <TextField
                            label="Place"
                            variant="outlined"
                            name="place"
                            required
                            value={match.place}
                            disabled={(!creating)}
                            sx={inputStyle}
                            onChange={handleChange}/>
                    </div>
                </div>
                <br/>
                <div className="row">
                    <div className="col-md-6">
                        <TextField
                            label={(creating)? '': "Date (mm-dd-yyyy)"}
                            variant="outlined"
                            name="date"
                            required
                            value={(match.date) ? match.date.split("T")[0] : ''}
                            disabled={(!creating)}
                            sx={inputStyle}
                            type="date"
                            onChange={handleChange}/>
                    </div>
                    <div className="col-md-6">
                        <TextField
                            label="Time"
                            variant="outlined"
                            name="time"
                            required
                            value={match.time}
                            disabled={(!creating)}
                            sx={inputStyle}
                            onChange={handleChange}/>
                    </div>
                </div>
                <br/>
                <div className="row">
                    <div className="col-md-6">
                        <TextField label="Price per player"
                                   variant="outlined"
                                   name="price"
                                   required
                                   value={(match.price) ? `${match.price}` : ''}
                                   disabled={(!creating)}
                                   sx={inputStyle}
                                   onChange={handleChange}
                                   type="number" step="0.01"/>
                    </div>
                    <div className="col-md-6">
                        <TextField label="Format (5, 7, 11, etc)"
                                   variant="outlined"
                                   name="format"
                                   required
                                   value={match.format}
                                   disabled={(!creating)}
                                   sx={inputStyle}
                                   type="number"
                                   onChange={handleChange}/>
                    </div>
                </div>
                <br />
                <div className="row">
                    <div className="col-md-6">
                        <TextField label="Max Players"
                                   variant="outlined"
                                   name="maxplayers"
                                   required
                                   value={match.maxPlayers}
                                   disabled={(!creating) ? true : false}
                                   sx={inputStyle}
                                   type="number"
                                   onChange={handleChange}/>
                    </div>
                    <div className="col-md-6">
                        {
                            (creating) ? '' :
                                <TextField className="col-md-6"
                                           label="Remaining Places"
                                           variant="outlined"
                                           name="remainingPlaces"
                                           required
                                           value={(match && match.maxPlayers) ? (match.maxPlayers - match.startingPlayers.length - match.substitutePlayer.length) : ''}
                                           disabled={(match && match.maxPlayers) ? true : false}
                                           sx={inputStyle}
                                           onChange={handleChange}
                                />
                        }
                    </div>
                </div>

                <br />
                <div className="row ">
                    <div className="col-md-6">
                        <div className={(creating) ? "hide-div match-form-starting-player-list" : 'match-form-starting-player-list'}>
                            <span className="match-form-player-list">Starting Players:
                                {(!match || !match._id) ? '' :

                                    (match.startingPlayers.length <= 0) ? ' The List is empty' :
                                        match.startingPlayers.map((player, index) => {
                                            return (
                                                <li key={index}>
                                                    {player.name}
                                                </li>
                                            )
                                        })
                                }
                            </span>
                        </div>
                    </div>
                    <div className="col-md-6">
                        <div className={(creating) ? "hide-div" : ''}>
                            <span className="match-form-player-list">Substitute Players:
                                {(!match || !match._id) ? '' :
                                    (match.substitutePlayer.length <= 0) ? ' The List is empty' :
                                        match.substitutePlayer.map((player, index) => {
                                            return (
                                                <li key={index}>
                                                    {player.name}
                                                </li>
                                            )
                                        })
                                }</span>
                        </div>
                    </div>
                </div>
                <br />
                <div className="row">
                    <div className="col-md-6">
                        {(!creating && match && match._id) ? <JoinPlayerModal matchId={match._id}
                                                                              places={(match.maxPlayers - match.startingPlayers.length - match.substitutePlayer.length)}/> :
                            <Button type={(!creating && match && match._id) ? "button" : "submit"}
                                    fullWidth
                                    variant="contained"
                            >
                                Create Match
                            </Button>
                        }
                    </div>
                    <div className="col-md-6">
                        <Button
                            fullWidth
                            variant="contained"
                            color="error"
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
