import React from 'react'
import CurrentMatches from '../match/currentMatches'

const Principal = (props) => {
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
