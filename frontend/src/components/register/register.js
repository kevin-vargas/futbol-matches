import React from 'react'
import { useNavigate } from "react-router-dom";
import UserData from "../userData/userData";

const Register = (props) => {

    const navigation = useNavigate();

    const handleSave = () => {
        navigation("/principal");
    }

    const handleCancel = () => {
        navigation("/");
    }

    return(
        <UserData handleSave={ handleSave } handleCancel={ handleCancel } />
    )
}

export default Register;
