import React, {useLayoutEffect} from 'react'
import CurrentMatches from '../match/currentMatches'
import {useNavigate} from "react-router-dom";

const Principal = (props) => {
    if( !localStorage.getItem("user") ){
        document.location.href = "/";
    }


    return (
        <section className="principal">
            <div className="row">
                <div className="col-md-9">
                    <CurrentMatches />
                </div>
            </div>
        </section>
    )
}

export default Principal;
