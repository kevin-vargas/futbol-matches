import React from 'react'
import {Button, TextField} from "@mui/material";
import {useNavigate} from "react-router-dom";
import Swal from 'sweetalert2'
import userService from "../services/userService";


const UserForm = (props) => {

    const navigation = useNavigate()

    const handleSubmit = (event) => {
        event.preventDefault()
        const user = {};

        user.username = event.target.username.value
        user.password = event.target.password.value
        user.name = event.target.name.value
        user.email = event.target.email.value
        user.phone = event.target.phone.value

        userService.signup(user).then(response => response.json()).then(jsonResponse => {
            if( jsonResponse.error ){
                Swal.fire({
                    title: jsonResponse.error,
                    icon: 'error',
                    confirmButtonText: 'Ok'
                }).then(r => console.log(r))

                return
            }
            else{
                Swal.fire({
                    title: "User was created!",
                    icon: 'success',
                    confirmButtonText: 'Ok'
                }).then(r => console.log(r))
                localStorage.setItem("token", jsonResponse.token)
                localStorage.setItem("username", user.username)
                navigation("/")
            }

        }).catch(error => {
            console.log(error)
        });
    }

    const handleCancel = () => {
        navigation("/")
    }

    return (
        <div className="container">
            <form autoComplete="off" onSubmit={handleSubmit}>
                <div className="row ">
                    <TextField  className="col-md-6"
                                label= "Username"
                                variant="outlined"
                                name="username"
                                required
                                sx={{marginRight: 10}} />

                    <TextField className="col-md-6"
                               label="Password"
                               variant="outlined"
                               name="password"
                               type="password"
                               required/>
                </div>
                <br/>


                <div className="row ">
                    <TextField className="col-md-6 "
                               label="Name"
                               variant="outlined"
                               name="name"
                               required
                               sx={{marginRight: 10}}/>


                    <TextField className="col-md-6" label="Email"
                               variant="outlined"
                               name="email"
                               required/>

                </div>
                <br/>

                <div className="row ">
                    <TextField className="col-md-6"
                               label="Phone"
                               variant="outlined"
                               name="phone"
                               required/>

                </div>


                <br />

                <br/>
                <div className="row">
                    <div className="col-md-6">
                        <Button type="submit"
                                fullWidth
                                variant="contained"
                                sx={{mt: 3, mb: 2}}
                        >
                            Create User
                        </Button>
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

export default UserForm;
