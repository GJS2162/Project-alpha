import React from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Navbar from "./components/Navbar";
import Bookings from "./components/Bookings";
import Drivers from "./components/Drivers";

function App() {
    return ( <
        Router >
        <
        div >
        <
        Navbar / >
        <
        Switch >
        <
        Route path = "/"
        exact component = {
            () => < h1 > Welcome to the Logistics Platform < /h1>} /
            >
            <
            Route path = "/bookings"
            component = { Bookings }
            />{" "} <
            Route path = "/drivers"
            component = { Drivers }
            />{" "} <
            /Switch>{" "} <
            /div>{" "} <
            /Router>
        );
    }

    export default App;