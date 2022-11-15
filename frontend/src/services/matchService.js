import config from "../config/config";

const matchService = {
    getMatches: () => {
        const url = `${config.apiHost}${config.endpoints.getMatches.path}`
        return fetch(url);
    },

    getMatchById: (matchId) => {
        const url = `${config.apiHost}${config.endpoints.getMatches.path}/${matchId}`
        return fetch(url);
    },

    addPlayer: (player, matchId) => {
        const url = `${config.apiHost}${config.endpoints.getMatches.path}/${matchId}/player`
        const requestOptions = config.endpoints.addPlayer.requestOptions;
        requestOptions.body = JSON.stringify(player);
        return fetch(url, requestOptions);
    },

    saveMatch: (match) => {
        const matchObj = {};
        matchObj.owner = match.owner
        matchObj.description = match.description
        matchObj.place = match.place
        matchObj.date = new Date(match.date).toISOString()
        matchObj.time = match.time
        matchObj.price = parseInt(match.price)
        matchObj.format = parseInt(match.format)
        matchObj.maxPlayers = parseInt(match.maxplayers)

        const requestOptions = config.endpoints.saveMatch.requestOptions
        requestOptions.body = JSON.stringify(matchObj)

        const url = `${config.apiHost}${config.endpoints.saveMatch.path}`

        return fetch(url, requestOptions)
    }
}


export default matchService;
