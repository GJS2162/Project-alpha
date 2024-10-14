import React, { useState, useEffect } from "react";
import axios from "axios";

const Bookings = () => {
    const [bookings, setBookings] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        axios
            .get("http://localhost:8080/api/v1/bookings")
            .then((response) => {
                setBookings(response.data);
                setLoading(false);
            })
            .catch((error) => {
                console.error("Error fetching bookings:", error);
            });
    }, []);

    if (loading) {
        return <div > Loading... < /div>;
    }

    return ( <
        div >
        <
        h1 > Bookings < /h1>{" "} <
        ul > { " " } {
            bookings.map((booking) => ( <
                li key = { booking.id } >
                <
                p > Booking ID: { booking.id } < /p>{" "} <
                p > Driver ID: { booking.driver_id } < /p>{" "} <
                p > Status: { booking.status } < /p>{" "} < /
                li >
            ))
        } { " " } <
        /ul>{" "} < /
        div >
    );
};

export default Bookings;