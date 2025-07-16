import { useEffect, useRef, useState } from "react";
import { MapContainer, TileLayer, Marker, Popup, useMap } from "react-leaflet";
import L from "leaflet";
import "leaflet/dist/leaflet.css";

// Fix for default markers in React-Leaflet
const DefaultIcon = L.icon({
  iconUrl:
    "https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon.png",
  iconRetinaUrl:
    "https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon-2x.png",
  shadowUrl:
    "https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-shadow.png",
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  shadowSize: [41, 41],
});

// Custom Munro icon - mountain peak
const MunroIcon = L.icon({
  iconUrl:
    "data:image/svg+xml;base64," +
    btoa(`
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24">
      <path fill="#2563eb" d="M12 2L3 22h18L12 2zm0 4.5L18.5 20h-13L12 6.5z"/>
      <circle cx="12" cy="8" r="1.5" fill="#ffffff"/>
    </svg>
  `),
  iconSize: [24, 24],
  iconAnchor: [12, 24],
  popupAnchor: [0, -24],
});

L.Marker.prototype.options.icon = DefaultIcon;

export interface Munro {
  running_no: number;
  dobih_number: number;
  name: string;
  smc_section: string;
  rhb_section: string;
  height_m: number;
  height_ft: number;
  map_1_50k: string;
  map_1_25k: string;
  grid_ref: string;
  grid_ref_xy: string;
  x_coord: number;
  y_coord: number;
  latitude: number;
  longitude: number;
  classification: string;
  comments: string;
  streetmap_url: string;
  geograph_url: string;
  hill_bagging_url: string;
}

interface MunroMapProps {
  munros: Munro[];
  onMunroClick?: (munro: Munro) => void;
}

function MapUpdater({ munros }: { munros: Munro[] }) {
  const map = useMap();

  useEffect(() => {
    if (munros.length > 0) {
      // Create bounds from munro coordinates
      const bounds = L.latLngBounds(
        munros.map((munro) => [munro.latitude, munro.longitude])
      );

      // Fit map to show all munros
      map.fitBounds(bounds, { padding: [20, 20] });
    }
  }, [munros, map]);

  return null;
}

export default function MunroMap({ munros, onMunroClick }: MunroMapProps) {
  const [selectedMunro, setSelectedMunro] = useState<Munro | null>(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [filteredMunros, setFilteredMunros] = useState<Munro[]>(munros);

  // Scotland bounds to restrict map view
  const scotlandBounds = L.latLngBounds(
    [54.6, -7.5], // Southwest corner (bottom-left)
    [60.9, -0.5] // Northeast corner (top-right)
  );

  useEffect(() => {
    if (searchTerm.trim() === "") {
      setFilteredMunros(munros);
    } else {
      const filtered = munros.filter(
        (munro) =>
          munro.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
          munro.smc_section.toLowerCase().includes(searchTerm.toLowerCase())
      );
      setFilteredMunros(filtered);
    }
  }, [searchTerm, munros]);

  const handleMunroClick = (munro: Munro) => {
    setSelectedMunro(munro);
    onMunroClick?.(munro);
  };

  const formatHeight = (heightM: number, heightFt: number) => {
    return `${heightM.toFixed(1)}m (${heightFt.toLocaleString()}ft)`;
  };

  return (
    <div className="w-full h-full relative">
      {/* Search Bar */}
      <div className="absolute top-4 left-4 z-[1000] bg-white p-4 rounded-lg shadow-lg">
        <div className="mb-4">
          <input
            type="text"
            placeholder="Search munros..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="w-64 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div className="text-sm text-gray-600">
          Showing {filteredMunros.length} of {munros.length} munros
        </div>
      </div>

      {/* Map */}
      <MapContainer
        center={[56.8, -4.2]} // Center of Scotland
        zoom={7}
        minZoom={6}
        maxZoom={14}
        maxBounds={scotlandBounds}
        maxBoundsViscosity={1.0}
        className="w-full h-full"
        style={{ height: "100%" }}
        worldCopyJump={false}
        zoomSnap={0.25}
      >
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        />

        <MapUpdater munros={filteredMunros} />

        {filteredMunros.map((munro) => (
          <Marker
            key={munro.running_no}
            position={[munro.latitude, munro.longitude]}
            icon={MunroIcon}
            eventHandlers={{
              click: () => handleMunroClick(munro),
            }}
          >
            <Popup>
              <div className="min-w-[280px] max-w-[400px]">
                <h3 className="text-lg font-bold text-gray-800 mb-2">
                  {munro.name}
                </h3>

                <div className="space-y-2 text-sm">
                  <div className="flex justify-between">
                    <span className="font-semibold text-gray-600">Height:</span>
                    <span className="text-gray-800">
                      {formatHeight(munro.height_m, munro.height_ft)}
                    </span>
                  </div>

                  <div className="flex justify-between">
                    <span className="font-semibold text-gray-600">
                      Classification:
                    </span>
                    <span
                      className={`px-2 py-1 rounded text-xs font-medium ${
                        munro.classification === "Munro"
                          ? "bg-blue-100 text-blue-800"
                          : "bg-gray-100 text-gray-800"
                      }`}
                    >
                      {munro.classification}
                    </span>
                  </div>

                  <div className="flex justify-between">
                    <span className="font-semibold text-gray-600">
                      SMC Section:
                    </span>
                    <span className="text-gray-800">{munro.smc_section}</span>
                  </div>

                  <div className="flex justify-between">
                    <span className="font-semibold text-gray-600">
                      Grid Reference:
                    </span>
                    <span className="text-gray-800 font-mono">
                      {munro.grid_ref}
                    </span>
                  </div>

                  {munro.comments && (
                    <div className="border-t pt-2">
                      <span className="font-semibold text-gray-600">
                        Comments:
                      </span>
                      <p className="text-gray-700 text-xs mt-1">
                        {munro.comments}
                      </p>
                    </div>
                  )}
                </div>

                <div className="flex gap-2 mt-3 pt-3 border-t">
                  {munro.streetmap_url && (
                    <a
                      href={munro.streetmap_url}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-xs bg-blue-500 text-white px-2 py-1 rounded hover:bg-blue-600"
                    >
                      Street Map
                    </a>
                  )}
                  {munro.geograph_url && (
                    <a
                      href={munro.geograph_url}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-xs bg-green-500 text-white px-2 py-1 rounded hover:bg-green-600"
                    >
                      Photos
                    </a>
                  )}
                  {munro.hill_bagging_url && (
                    <a
                      href={munro.hill_bagging_url}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-xs bg-purple-500 text-white px-2 py-1 rounded hover:bg-purple-600"
                    >
                      Hill Bagging
                    </a>
                  )}
                </div>
              </div>
            </Popup>
          </Marker>
        ))}
      </MapContainer>
    </div>
  );
}
