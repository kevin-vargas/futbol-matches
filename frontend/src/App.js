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
    const navigation = useNavigate();
    const username = useRef();
    const password = useRef();

    const handleLogin = () => {
        const user = username.current.value;
        const pass = password.current.value;

        if( !user ) return;
        if( !pass ) return;

        userService.login(user, pass).then((response) => {
                    setUser(response);
                    navigation("/principal")
            });
    }

    const closeSession = () =>{
        setUser(false);
        navigation("/");
    }

  return (
    <div className="App">
            <Header closeSession={closeSession} user={currentUser}/>
            <Routes>
                <Route path="/" element={<Login handleLogin={handleLogin} userName={username} password={password} />} />
                <Route path="/register" element={<Register />} />
                <Route path="/profile" element={<Profile user={currentUser}/>} />
                <Route path="/principal" element={<Principal />} />
                <Route path="/create-match" element={<CreateMatch user={currentUser} />} />
                <Route path="/statistics" element={<Statistics />} />
                <Route path="*" element={<Error />} />
            </Routes>
            <Footer />
    </div>
  );
}

export default App;
