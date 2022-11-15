import React, {useState} from 'react'
import {Box, Button, Container, CssBaseline, Grid, Link, TextField} from "@mui/material";

import userService from "../services/userService";
import {useNavigate} from "react-router-dom";

function Intro(props){


    localStorage.clear()

    const [loginError, setLoginError ] = useState(undefined);

    const navigation = useNavigate();

    const handleSubmit = (event) => {
        event.preventDefault();
        const username = event.target.username.value;
        const password = event.target.password.value;

        if( !username || !password ) return;

        userService.login(username, password)
            .then((response) => {
                return response.json();
            }).then( responseJson => {
                if( responseJson.error ){
                    setLoginError(responseJson.error)
                    return
                }
                else{
                    localStorage.setItem("token", responseJson.token)
                    localStorage.setItem("username", username)
                    navigation("/matches")
                }

            }).catch(error => {
                console.log(error)
            });

    }

    return (
        <div className="Intro">
            <h1 className="intro-title">Futbol Matches</h1>
            <Container component="main" maxWidth="xs" className="with-background">
            <CssBaseline />
            <Box
                sx={{
                    marginTop: 8,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    color: 'withe',
                }}
            >

                <Box component="form" sx={{ mt: 1 }} onSubmit={handleSubmit}>
                    {
                        loginError != null && <p className="intro-error">{ loginError }</p>
                    }
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="username"
                        label="Username"
                        name="username"
                        autoComplete="Username"
                        autoFocus
                        sx={{ input: { color: 'white' }, label: { color: 'white' } }}
                    />
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        name="password"
                        label="Password"
                        type="password"
                        id="password"
                        autoComplete="current-password"
                        sx={{ input: { color: 'white' }, label: { color: 'white' } }}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                    >
                        Sign In
                    </Button>
                    <Grid container>
                        <Grid item>
                            <Link href="signup" variant="body2">
                                <span className="intro-link">Don't have an account? Sign Up</span>
                            </Link>
                        </Grid>
                    </Grid>
                </Box>
            </Box>
            </Container>
        </div>
    )
}

export default Intro;
