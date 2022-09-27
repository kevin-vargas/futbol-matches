import React from 'react'
import Menu from "../menu/menu";

const Header = (props) => {

    console.log("Header: ", props.user);

    return (
        <section className="header">
            <div className="row">
                <div className="col-md-6 text-align-left">
                    FUTBOL MATCHES
                </div>

                { props.user ?
                    <div className="col-md-6 text-align-right">
                        AVATAR
                    </div> : ''
                }

                { props.user ?
                    <div className="col-md-12">
                        <Menu handleCloseSession={props.closeSession} user={props.user}/>
                    </div>
                    : ''
                }

            </div>
            <hr />
        </section>
    )
}

export default Header;
