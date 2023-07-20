import { MutableRefObject, useEffect, useRef } from "react";

declare var kakao: any;

export const MyComponent = () => {
  const mapRef = useRef<HTMLElement | null>(null);

  const initMap = () => {
    const container = document.getElementById('map');
    const options = {
      center: new kakao.maps.LatLng(37.483034, 126.902435),
      level: 2
    };

    const map = new kakao.maps.Map(container as HTMLElement, options);
    (mapRef as MutableRefObject<any>).current = map;
  };

  useEffect(() => {
    kakao.maps.load(() => initMap());
  }, [mapRef]);

  return (
      <div id="map" style={{ width: "500px", height: "400px" }}></div>
  );
};
