import React, {useRef} from 'react'
import matchService from "../../services/matchService";
import {useNavigate} from "react-router-dom";

const CreateMatch = (props) => {

    const navigation = useNavigate();

    const user = props.user;
    const match = {};

    match.owner = useRef();
    match.date = useRef();
    match.place = useRef();
    match.time = useRef();
    match.format = useRef();
    match.maxplayers = useRef();
    match.price = useRef();

    const handleSaveMatch = (match) => {
        matchService.saveMatch(match).then(response => {
           if(response.status === 201) {
               console.log("Partido creado");
           }
           else {
               console.log("Error al crear el partido");
           }
            navigation("/principal")
        });
    }

    return (
        <section className="container">
            <form className="login-form rounded row">
                <div className="col-md-12">
                    <label className="form-label" htmlFor="owner">Owner</label>
                    <input type="email" id="owner" name="owner" className="form-control"
                           defaultValue={user.username} ref={match.owner}/>
                </div>

                <div className="col-md-12">
                    <label className="form-label" htmlFor="place">Place</label>
                    <input type="email" id="place" name="place" className="form-control"
                           defaultValue="" ref={match.place}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="date">Date</label>
                    <input type="email" id="date" name="date" className="form-control" ref={match.date}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="time">Hour</label>
                    <input type="email" id="time" name="time" className="form-control"
                           ref={match.time}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="matchType">Match type</label>
                    <input type="email" id="matchType" name="matchType" className="form-control"
                    ref={match.format}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="maxplayers">Max players</label>
                    <input type="email" id="maxplayers" name="maxplayers" className="form-control"
                    ref={match.maxplayers}/>
                </div>

                <div className="col-md-6">
                    <label className="form-label" htmlFor="price">Price per player</label>
                    <input type="email" id="price" name="price" className="form-control"
                           ref={match.price}/>
                </div>

                <div className="col-md-6"></div>
                <div className="col-md-12"><br /></div>
                <div className="col-md-6">
                    <button onClick={ e => navigation("/principal") } type="button" className="btn btn-danger" >Cancel</button>
                </div>
                <div className="col-md-6">
                    <button type="button" className="btn btn-success" onClick={ e => handleSaveMatch(match) } >Save</button>
                </div>
            </form>
        </section>
    );
}

export default CreateMatch;
