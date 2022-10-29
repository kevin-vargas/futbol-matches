import React from 'react'

const Match = (props) => {
    const remaining =  props.match.maxPlayers - (props.match.startingPlayers.length +
        props.match.substitutePlayer.length);

    return (
        <div className="row match-card rounded">
            <div className="col-md-3">Owner: {props.match.owner}</div>
            <div className="col-md-6">Place: {props.match.place}</div>
            <div className="col-md-3">Remaining Players: { remaining }</div>

            <div className="col-md-3">Date: {props.match.date.split("T")[0]}</div>
            <div className="col-md-3">Hour: {props.match.time}</div>
            <div className="col-md-3">Price: ${props.match.price}</div>
            <div className="col-md-3">
                <button type="button" className="btn btn-primary">Add me!</button>
            </div>
        </div>
    )
}

export default Match;
