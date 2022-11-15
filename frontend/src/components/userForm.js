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
                    confirmButtonText: 'Ok',
                    customClass: 'swal-height'
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
        <div className="container user-form">
            <form autoComplete="off" onSubmit={handleSubmit} className="form-padding">
                <div className="row ">
                    <div className="col-md-6">
                        <TextField  label= "Username"
                                    variant="outlined"
                                    name="username"
                                    required
                                    sx={{marginRight: 10, input: { color: 'white' }, label: { color: 'white' }}} />

                    </div>
                    <div className="col-md-6">
                        <TextField label="Password"
                                   variant="outlined"
                                   name="password"
                                   type="password"
                                   sx={{ input: { color: 'white' }, label: { color: 'white' } }}
                                   required/>
                    </div>


                </div>
                <br/>


                <div className="row ">
                    <div className="col-md-6">
                        <TextField label="Name"
                                   variant="outlined"
                                   name="name"
                                   required
                                   sx={{marginRight: 10, input: { color: 'white' }, label: { color: 'white' }}}/>

                    </div>
                    <div className="col-md-3">
                        <TextField variant="outlined"
                                   label="Email"
                                   name="email"
                                   type="email"
                                   required
                                   sx={{input: { color: 'white' }, label: { color: 'white' }}}/>
                    </div>
                    <div className="col-md-3">
                        <TextField label="Phone"
                                   variant="outlined"
                                   name="phone"
                                   required
                                   sx={{ input: { color: 'white' }, label: { color: 'white' }}}/>
                    </div>
                </div>
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
                    </div>
                    <div className="col-md-6">
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
