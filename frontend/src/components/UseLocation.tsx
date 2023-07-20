import React, { useEffect, useState } from "react";

const UseLocation = () => {
    const [location, setLocation] = useState<{latitude: number; longitude: number} | string>('');

    useEffect(() => {
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition(success, error);
        }

        function success(position: GeolocationPosition) {
            setLocation({
                latitude: position.coords.latitude,
                longitude: position.coords.longitude
            });
        }

        function error() {
            setLocation({
                latitude: 37.483034,
                longitude: 126.902435
            });
            console.log("위치 받기 실패");
        }
    }, []);

    return location;
};

export default UseLocation;
