import './App.css';
import Login from "./components/login/login";
import Header from "./components/header/header";
import Footer from "./components/footer/footer";

import userService from "./services/userService"
import React, {useRef, useState} from "react";
import Principal from "./components/principal/principal";
import { Route, Routes} from "react-router-dom";
import { useNavigate } from "react-router-dom";
import Register from "./components/register/register";
import CreateMatch from "./components/match/createMatch";
import Error from "./components/error/error";
import Profile from "./components/profile/profile";
import Statistics from "./components/statistics/statistics";

function App() {
    const [currentUser, setUser ] = useState(undefined);
    const [token, setToken ] = useState({});

    const navigation = useNavigate();

    const username = useRef();
    const password = useRef();

    const handleLogin = () => {
        const user = username.current.value;
        const pass = password.current.value;

        if( !user ) return;
        if( !pass ) return;

        userService.login(user, pass).then((response) => {
            console.log(response.headers);
            return response.json();
        })
            .then( responseJson => {
                setUser(responseJson);
                navigation("/principal")
            }).catch(error => console.log(error));
    }

    const handleRegistration = ( user ) => {
        const userObj = {};
        userObj.name = user.name.current.value;
        userObj.lastname = user.lastname.current.value;
        userObj.username = user.username.current.value;
        userObj.email = user.email.current.value;
        userObj.phone = user.phone.current.value;
        userObj.password = user.password.current.value;

        userService.signup(userObj).then( response => {
            if( response.status === 201 ){
                console.log("Usuario Creado con exito!");
                setUser(userObj);
                localStorage.setItem("user", JSON.stringify(userObj));
                navigation("/principal")
            }
            else {
                console.log(`${response.status}: ${response.statusText}`);
                navigation("/");
            }
        }).catch( error => console.log("error: ", error));
    }

    const closeSession = () =>{
        setUser(false);
        localStorage.clear();
        navigation("/");
    }

  return (
    <div className="App">
            <Header closeSession={closeSession} user={currentUser}/>
            <Routes>
                <Route path="/" element={<Login handleLogin={handleLogin} userName={username} password={password} />} />
                <Route path="/register" element={<Register handleRegistration={handleRegistration}/>} />
                <Route path="/profile" element={<Profile user={currentUser}/>} />
                <Route path="/principal" element={<Principal user={currentUser}/>} />
                <Route path="/create-match" element={<CreateMatch user={currentUser} />} />
                <Route path="/statistics" element={<Statistics />} />
                <Route path="*" element={<Error />} />
            </Routes>
            <Footer />
    </div>
  );
}

export default App;
