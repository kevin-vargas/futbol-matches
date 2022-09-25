import React from 'react'
import matchService from '../../services/matchService'
import Match from "./match";

const CurrentMatches = () => {
    const matches = matchService.getMatches();

    return(
        <div>
            {
                matches.map( (match, index) => {
                   return <Match match={match} key={index}/>
                })
            }
        </div>
    )
}

export default CurrentMatches;
