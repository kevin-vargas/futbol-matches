const config = {
    apiHost: `${process.env.REACT_APP_API_HOST}/`,
    endpoints: {
        signup: {
            path: 'users',
            requestOptions: {
                method: 'POST',
                headers: {'Content-Type': 'application/json'}
            }
        },
        login: {
            path: 'login',
            requestOptions: {
                method: 'POST',
                headers: {'Content-Type': 'application/json'}
            }
        },
        updateUser: {
            path: 'login',
            requestOptions: {
                method: 'POST',
                headers: {'Content-Type': 'application/json'}
            }
        },
        saveMatch: {
            path: 'matches',
            requestOptions: {
                method: 'POST',
                headers: {'Content-Type': 'application/json'}
            }
        },
        addPlayer: {
            requestOptions: {
                method: 'POST',
                headers: {'Content-Type': 'application/json'}
            }
        },
        getMatches: {
            path: 'matches',
            requestOptions: {
                method: 'GET',
                headers: {'Content-Type': 'application/json'}
            }
        },
        getMetrics: {
            path: 'metrics',
            requestOptions: {
                method: 'GET',
                headers: {'Content-Type': 'application/json'}
            }
        }
    }
}

export default config;
