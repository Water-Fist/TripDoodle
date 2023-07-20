import React, { MutableRefObject, useEffect, useRef } from "react";
import UseLocation from './UseLocation';

declare var kakao: any;

export const MyComponent = () => {
  const mapRef = useRef<HTMLElement | null>(null);
  const location = UseLocation();

  const initMap = () => {
    if (typeof location !== 'string' && location) {
      const container = document.getElementById('map');
      const options = {
        center: new kakao.maps.LatLng(location.latitude, location.longitude),
        level: 2
      };

      const map = new kakao.maps.Map(container as HTMLElement, options);
      (mapRef as MutableRefObject<any>).current = map;
    }
  };

  useEffect(() => {
    kakao.maps.load(() => initMap());
  }, [mapRef, location]);

  return (
      <div id="map" style={{ width: "500px", height: "400px" }}></div>
  );
};
