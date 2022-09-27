const config = {
    apiHost: 'http://localhost:8080/',
    endpoints: {
        signup: {
            path: 'signup',
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
        saveMatch: {
            path: 'matches',
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
        }
    }
}

export default config;
