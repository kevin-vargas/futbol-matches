import React from 'react'
import Menu from "../menu/menu";

const Header = (props) => {

    console.log("Header: ", props.user);
    let user = {};

    if( !props.user ){
        user = JSON.parse(localStorage.getItem("user"));
    }
    else {
        user = props.user;
    }

    return (
        <section className="header">
            <div className="row">
                <div className="col-md-6 text-align-left">
                    FUTBOL MATCHES
                </div>

                { user ?
                    <div className="col-md-12">
                        <Menu handleCloseSession={props.closeSession} user={user}/>
                    </div>
                    : ''
                }

            </div>
            <hr />
        </section>
    )
}

export default Header;
