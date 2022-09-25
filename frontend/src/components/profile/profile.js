import React from 'react';
import UserData from "../userData/userData";
import {useNavigate} from "react-router-dom";

const Profile = (props) => {
    const navigation = useNavigate();

    const handleSave = () => {
        navigation("/principal");
    }

    return (<UserData handleSave={handleSave} handleCancel={handleSave} user={props.user}/>);
}

export default Profile;
