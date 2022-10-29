import React from 'react'
import matchService from '../../services/matchService'
import Match from "./match";

class CurrentMatches extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            matches: []
        }
    }

    componentDidMount() {
        matchService.getMatches().then( response => response.json()).then( jsonResponse => {
            console.log("Response json: ", jsonResponse);
            this.setState({ matches: jsonResponse });
        });
    }

    render() {
        return(
            <div>
                {
                    (this.state.matches.length > 0) ?
                    this.state.matches.map( (match, index) => {
                        return <Match match={match} key={index}/>
                    }):
                        <div>
                            <strong>There is no matches created yet!</strong>
                        </div>
                }
            </div>
        )
    }
}

export default CurrentMatches;
