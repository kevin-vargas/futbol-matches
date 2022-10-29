import config from "../config/config";

const matchService = {
    getMatches: () => {
        const url = `${config.apiHost}${config.endpoints.getMatches.path}`
        return fetch(url);
    },

    saveMatch: (match) => {
        console.log(match)
        const matchObj = {};
        matchObj.owner = match.owner.current.value;
        matchObj.place = match.place.current.value;
        matchObj.date = match.date.current.value+'T00:00:00Z';
        matchObj.time = match.time.current.value;
        matchObj.price = parseInt(match.price.current.value);
        matchObj.format = parseInt(match.format.current.value);
        matchObj.maxPlayers = parseInt(match.maxplayers.current.value);

        const requestOptions = config.endpoints.saveMatch.requestOptions;
        requestOptions.body = JSON.stringify(matchObj);

        const url = `${config.apiHost}${config.endpoints.saveMatch.path}`;

        return fetch(url, requestOptions);
    }
}


export default matchService;
