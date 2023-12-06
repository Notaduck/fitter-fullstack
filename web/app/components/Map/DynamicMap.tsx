import { useEffect, FC, useState, useRef } from 'react';
import L from 'leaflet';
import * as ReactLeaflet from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import { Activity } from '../../../services/activity-services';

import MarkerIcon from 'leaflet/dist/images/marker-icon.png';
import MarkerIcon2x from 'leaflet/dist/images/marker-icon-2x.png';
import MarkerShadow from 'leaflet/dist/images/marker-shadow.png';

import styles from './Map.module.scss';
import { Marker, Polyline, Popup } from 'react-leaflet';

const { MapContainer } = ReactLeaflet;

type Props = {
  children: React.ReactNode[];
  className: string;

  width: number;
  height: number;
  activity: Activity;
};

type Point = {
  id: number;
  lat: number;
  long: number;
};


const findNearestPoint = (mouseLatLng, points) => {
  let nearestPoint = null;
  let minDistance = Number.MAX_VALUE;

  for (const point of points) {
    const distance = calculateDistance(mouseLatLng, { lat: point.lat, lng: point.long });

    if (distance < minDistance) {
      minDistance = distance;
      nearestPoint = point;
    }
  }

  return nearestPoint;
};

const calculateBounds = (routePoints: Point[]) => {
  if (!routePoints) return [];

  const latitudes = routePoints.map(({ lat }) => lat);
  const longitudes = routePoints.map(({ long }) => long);

  const minLat = Math.min(...latitudes);
  const maxLat = Math.max(...latitudes);
  const minLng = Math.min(...longitudes);
  const maxLng = Math.max(...longitudes);
  const padding = 10;

  return [
    [minLat, minLng],
    [maxLat, maxLng],
  ];
};

const calculateDistance = (point1: Point, point2: Point) => {
  const lat1 = point1.lat;
  const lng1 = point1.lng;
  const lat2 = point2.lat;
  const lng2 = point2.lng;

  const dLat = degToRad(lat2 - lat1);
  const dLng = degToRad(lng2 - lng1);

  const a =
    Math.sin(dLat / 2) * Math.sin(dLat / 2) +
    Math.cos(degToRad(lat1)) * Math.cos(degToRad(lat2)) * Math.sin(dLng / 2) * Math.sin(dLng / 2);

  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));

  // Radius of the Earth in kilometers
  const radius = 6371;

  const distance = radius * c;

  return distance;
};

const degToRad = (degrees) => {
  return degrees * (Math.PI / 180);
};

const Map: FC<Props> = ({ children, className, width, height, activity, ...rest }) => {
  let mapClassName = styles.map;

  if (className) {
    mapClassName = `${mapClassName} ${className}`;
  }

  useEffect(() => {
    (async function init() {
      delete L.Icon.Default.prototype._getIconUrl;
      L.Icon.Default.mergeOptions({
        iconRetinaUrl: MarkerIcon2x.src,
        iconUrl: MarkerIcon.src,
        shadowUrl: MarkerShadow.src,

      });
      L.canvas({ padding: 2.5, tolerance: 200 });

    })();
  }, []);

  const routePoints = activity?.records?.filter((e) => e.PositionLat !== 0 && e.PositionLong !== 0).map(({ PositionLat, PositionLong, ID }) => ({ id: ID, long: PositionLong, lat: PositionLat }));

  const [hoveredId, setHoveredId] = useState<number | null>(null);
  const [nearestPoint, setNearestPoint] = useState<Point | null>(null);
  const routePointsRef = useRef(routePoints);

  useEffect(() => {
    routePointsRef.current = routePoints;
  }, [routePoints]);



  const handleMouseOver = (e) => {
    const latlng = e.latlng;
    const nearestPoint = findNearestPoint(latlng, routePointsRef.current);
    setHoveredId(nearestPoint?.id);
    setNearestPoint(nearestPoint);
  };

  const bounds = calculateBounds(routePoints)
  console.log(bounds)
  const polylineRef = useRef(null)


  return (
    <MapContainer className={mapClassName} {...rest} >
      {routePoints && (
        <>
          <Polyline
            positions={routePoints.map(({ lat, long }) => [lat, long])}
            eventHandlers={{
              mouseover: handleMouseOver
            }}
          />
          {nearestPoint && (
            <Marker position={[nearestPoint.lat, nearestPoint.long]} />
          )}


        </>
      )}
      {bounds && calculateBounds(routePoints)?.length > 0 && (
        <ReactLeaflet.LayersControl position="topright">
          <ReactLeaflet.LayersControl.BaseLayer checked name="Street Map">
            <ReactLeaflet.TileLayer
              url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            />
          </ReactLeaflet.LayersControl.BaseLayer>
          <ReactLeaflet.LayersControl.Overlay name="Polyline">
            <ReactLeaflet.Rectangle bounds={bounds} />
          </ReactLeaflet.LayersControl.Overlay>
        </ReactLeaflet.LayersControl>
      )}

      {children(ReactLeaflet, L)}
    </MapContainer>
  );
};

export default Map;