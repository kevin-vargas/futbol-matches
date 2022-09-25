import {matches} from "@testing-library/jest-dom/dist/utils";

const matchService = {

   matches: [{
        id: 'match 1',
        description: "match between friends",
        finished: false,
        date: '2022/09/30',
        time: '21:00',
        place: 'Club Ciudad de Buenos Aires, Av del Libertador',
        format: 8,
        price: 100,
        maxPlayers: 20,
        startingPlayers: [],
        substitutePlayer: []
    },
        {
            id: 'match 2',
            description: "match between friends",
            finished: false,
            date: '2022/09/30',
            time: '21:00',
            place: 'Club Ciudad de Buenos Aires, Av del Libertador',
            format: 8,
            price: 100,
            maxPlayers: 20,
            startingPlayers: [],
            substitutePlayer: []
        },
        {
            id: 'match 3',
            description: "match between friends",
            finished: false,
            date: '2022/09/30',
            time: '21:00',
            place: 'Club Ciudad de Buenos Aires, Av del Libertador',
            format: 8,
            price: 100,
            maxPlayers: 20,
            startingPlayers: [],
            substitutePlayer: []
        },
        {
            id: 'match 4',
            description: "match between friends",
            finished: false,
            date: '2022/09/30',
            time: '21:00',
            place: 'Club Ciudad de Buenos Aires, Av del Libertador',
            format: 8,
            price: 100,
            maxPlayers: 20,
            startingPlayers: [],
            substitutePlayer: []
        }],

    getMatches: () => {
        return matches;
    },
    saveMatch(match) {
       matches.push(match);
       console.log("match saved")
    }
}


export default matchService;
