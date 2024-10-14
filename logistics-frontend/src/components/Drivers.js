import React, { useState, useEffect } from "react";
import axios from "axios";

const Drivers = () => {
    const [drivers, setDrivers] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        axios
            .get("http://localhost:8080/api/v1/drivers")
            .then((response) => {
                setDrivers(response.data);
                setLoading(false);
            })
            .catch((error) => {
                console.error("Error fetching drivers:", error);
            });
    }, []);

    if (loading) {
        return <div > Loading... < /div>;
    }

    return ( <
        div >
        <
        h1 > Drivers < /h1>{" "} <
        ul > { " " } {
            drivers.map((driver) => ( <
                li key = { driver.id } >
                <
                p > Name: { driver.name } < /p>{" "} <
                p > Vehicle ID: { driver.vehicle_id } < /p>{" "} <
                p > Status: { driver.status } < /p>{" "} <
                /li>
            ))
        } { " " } <
        /ul>{" "} <
        /div>
    );
};

export default Drivers;