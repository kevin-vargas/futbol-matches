import config from '../config/config'

const userService = {

    signup: (user) => {
        const requestOptions = config.endpoints.signup.requestOptions;
        requestOptions.body = JSON.stringify(user);
        const url = `${config.apiHost}${config.endpoints.signup.path}`;

        return fetch(url, requestOptions);
    },

    login: (username, password) => {
        const requestOptions = config.endpoints.login.requestOptions;
        requestOptions.body = JSON.stringify({username, password});
        const url = `${config.apiHost}${config.endpoints.login.path}`;

        return fetch(url, requestOptions);
    }
}

export default userService;
