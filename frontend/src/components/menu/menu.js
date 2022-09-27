import React from 'react'
import {NavLink, useNavigate} from "react-router-dom";

const Menu = (props) => {

    console.log("Menu: ", props.user);

    const navigation = useNavigate();

    const handleClickProfile = () => {
        navigation("/profile");
    }

    return (
        <section className="menu">
            <div className="row navbar ">
                <div className="col-md-2">
                   <NavLink to="/principal" style={{textDecoration: 'none'}}>Current Matches</NavLink>
                </div>
                <div className="col-md-2">
                    <NavLink to="/create-match" style={{textDecoration: 'none'}}>Create Match</NavLink>
                </div>
                { (props.user.role === 'admin') ?
                <div className="col-md-2">
                    <NavLink to="/statistics" style={{textDecoration: 'none'}}>Statistics</NavLink>
                </div>
                    :  ''
                }
                <div className="col-md-2">

                </div>
                <div className="col-md-2 text-align-right">
                    <button className="btn" onClick={ props.handleCloseSession }> Salir </button>
                </div>
                <div className="col-md-2 text-align-right">
                    <button className="btn" onClick={ handleClickProfile }> Mi Perfil </button>
                </div>
            </div>
        </section>
    );
}

export default Menu;
