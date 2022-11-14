import React from 'react'
import './App.css';
import { Route, Routes} from "react-router-dom";
import Intro from "./components/intro";
import Matches from "./components/matches";
import MatchForm from "./components/matchForm";
import Match from "./components/match";
import UserForm from "./components/userForm";
import Metrics from "./components/metrics";

function App() {
  return (
        <div className="App">
          <Routes>
              <Route path="/" element={<Intro />} />
              <Route path="/matches" element={<Matches />} />
              <Route path="/match-form" element={<MatchForm />} />
              <Route path="/match/:id" element={<Match />} />
              <Route path="/signup" element={<UserForm />} />
              <Route path="/metrics" element={<Metrics />} />
          </Routes>
        </div>
  );
}

export default App;
