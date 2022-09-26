import React, {useRef} from 'react'
import { useNavigate } from "react-router-dom";
import UserData from "../userData/userData";
import userService from "../../services/userService";

const Register = (props) => {

    const navigation = useNavigate();

    const user = {};
    user.name = useRef();
    user.lastname = useRef();
    user.username = useRef();
    user.password = useRef();
    user.email = useRef();
    user.phone = useRef();

    const handleSave = (user) => {
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
            }
            else {
                console.log(`${response.status}: ${response.statusText}`);
            }
            navigation("/");
        }).catch( error => console.log("error: ", error));
    }

    const handleCancel = () => {
        navigation("/");
    }

    return(
        <UserData handleSave={ handleSave } handleCancel={ handleCancel } user={user} />
    )
}

export default Register;
