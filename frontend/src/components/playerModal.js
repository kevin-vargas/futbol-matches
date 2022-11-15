import React, {useState} from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Modal from '@mui/material/Modal';
import {TextField} from "@mui/material";
import matchService from "../services/matchService";
import Swal from "sweetalert2";
import {useNavigate} from "react-router-dom";

const style = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
};

export default function JoinPlayerModal(props) {
    const [open, setOpen] = useState(false);
    const [phonePlayer, setPhonePlayer] = useState('');
    const [emailPlayer, setEmailPlayer] = useState('');
    const navigation = useNavigate();

    const handleOpen = () => {
        if(props.places == 0){
            Swal.fire({
                title: "The match has reached the maximum number of players.",
                icon: 'error',
                confirmButtonText: 'Ok'
            }).then(r => console.log(r))
        } else {
            setOpen(true);
        }
    }
    const handleClose = () => setOpen(false);

    const handleChange = (event) => {
        if(event.target.name === 'phonePlayer'){
            setPhonePlayer(event.target.value)
        }
        if(event.target.name === 'emailPlayer'){
            setEmailPlayer(event.target.value)
        }
    }

    const handleJoinMatch = (event) => {
        event.preventDefault()
            const player = {
                name: localStorage.getItem("username"),
                phone: phonePlayer,
                email: emailPlayer,
            }
            matchService.addPlayer(player, props.matchId).then(response => response.json()).then(jsonResponse => {
                if(jsonResponse.error){
                    Swal.fire({
                        title: jsonResponse.error,
                        icon: 'error',
                        confirmButtonText: 'Ok'
                    }).then(r => console.log(r))
                } else {
                    Swal.fire({
                        title: 'Player Added to the Match!',
                        icon: 'success',
                        confirmButtonText: 'Ok'
                    }).then(r => console.log(r))
                }
                setOpen(false)
                navigation('/match/'+props.matchId)
            });
    }

    return (
        <div>
            <Button variant="contained" onClick={handleOpen}>Join!</Button>
            <Modal
                open={open}
                onClose={handleClose}
                aria-labelledby="modal-modal-title"
                aria-describedby="modal-modal-description"
            >
                <Box sx={style}>
                    <TextField
                               label="Phone"
                               variant="outlined"
                               name="phonePlayer"
                               required
                               value={phonePlayer}
                               onChange={ handleChange }
                               sx={{marginRight: 10}}/>
                    <br />
                    <TextField
                               label="Email"
                               variant="outlined"
                               name="emailPlayer"
                               value={emailPlayer}
                               onChange={ handleChange }
                               required/>

                    <Button type="button"
                            fullWidth
                            variant="contained"
                            sx={{mt: 3, mb: 2}}
                    onClick={handleJoinMatch}>
                        Join!
                    </Button>
                </Box>
            </Modal>
        </div>
    );
}
