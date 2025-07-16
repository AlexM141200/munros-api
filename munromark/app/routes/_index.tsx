import type { MetaFunction, LoaderFunction } from "@remix-run/node";
import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { useState, useEffect } from "react";

// Dynamic import to avoid SSR issues
import { Suspense, lazy } from "react";

const MunroMapLazy = lazy(() => import("~/components/MunroMap"));

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

export const meta: MetaFunction = () => {
  return [
    { title: "Munro Mark - Interactive Map of Scottish Munros" },
    {
      name: "description",
      content: "Explore all the Munros in Scotland with our interactive map",
    },
  ];
};

export const loader: LoaderFunction = async () => {
  try {
    // Fetch munros from our Go API
    const response = await fetch("http://localhost:8080/api/munros");

    if (!response.ok) {
      throw new Error("Failed to fetch munros");
    }

    const munros = await response.json();
    return json({ munros });
  } catch (error) {
    console.error("Error fetching munros:", error);
    return json({ munros: [] });
  }
};

export default function Index() {
  const { munros } = useLoaderData<{ munros: Munro[] }>();
  const [selectedMunro, setSelectedMunro] = useState<Munro | null>(null);
  const [isClient, setIsClient] = useState(false);
  const [showHero, setShowHero] = useState(true);

  useEffect(() => {
    setIsClient(true);
  }, []);

  return (
    <div className="w-full h-screen bg-gray-50">
      <header className="bg-gradient-to-r from-blue-700 via-teal-600 to-green-500 text-white shadow-lg">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-4">
            <div className="flex items-center">
              <h1 className="text-2xl font-bold text-gray-900">Munro Mark</h1>
              <span className="ml-3 text-sm text-gray-500">
                Interactive Map of Scottish Munros
              </span>
            </div>
            <div className="flex items-center space-x-4">
              <div className="text-sm text-gray-600">
                {munros.length} Munros Available
              </div>
            </div>
          </div>
        </div>
      </header>

      <main className="h-[calc(100vh-80px)] relative">
        {isClient ? (
          <>
            {/* Decorative hero overlay */}
            {showHero && (
              <div className="absolute top-8 left-1/2 -translate-x-1/2 z-[1000] w-[90%] sm:w-[600px] bg-white/80 backdrop-blur-md rounded-xl shadow-xl p-6 flex flex-col items-center text-center gap-4">
                <button
                  onClick={() => setShowHero(false)}
                  aria-label="Close overlay"
                  className="absolute top-2 right-2 text-gray-600 hover:text-gray-800 text-2xl leading-none"
                >
                  &times;
                </button>
                <h2 className="text-3xl font-extrabold text-gray-900 drop-shadow-sm">
                  Discover Scotland
                  <span className="block text-4xl md:text-5xl bg-clip-text text-transparent bg-gradient-to-r from-blue-600 via-teal-500 to-green-500 tracking-tight">
                    Munros Map
                  </span>
                </h2>
                <p className="text-gray-700 max-w-prose text-sm sm:text-base">
                  Explore every Scottish Munro peak on an interactive map. Click
                  a summit to see height, routes, photos and more.
                </p>
                <div className="flex gap-3">
                  <a
                    href="#map"
                    className="px-5 py-2 rounded-lg text-white font-semibold bg-blue-600 hover:bg-blue-700 shadow"
                  >
                    Start Exploring
                  </a>
                  <a
                    href="https://en.wikipedia.org/wiki/Munro"
                    target="_blank"
                    rel="noreferrer"
                    className="px-5 py-2 rounded-lg font-semibold bg-white/70 border border-blue-600 text-blue-700 hover:bg-white"
                  >
                    What is a Munro?
                  </a>
                </div>
              </div>
            )}
            <div id="map" className="w-full h-full">
              <Suspense
                fallback={
                  <div className="w-full h-full flex items-center justify-center bg-gray-100">
                    <div className="text-center">
                      <div className="animate-spin rounded-full h-20 w-20 border-b-2 border-blue-500 mx-auto mb-4"></div>
                      <p className="text-gray-600">Loading map...</p>
                    </div>
                  </div>
                }
              >
                <MunroMapLazy
                  munros={munros}
                  onMunroClick={(munro: Munro) => setSelectedMunro(munro)}
                />
              </Suspense>
            </div>
          </>
        ) : (
          <div className="w-full h-full flex items-center justify-center bg-gray-100">
            <div className="text-center">
              <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-blue-500 mx-auto mb-4"></div>
              <p className="text-gray-600">Loading map...</p>
            </div>
          </div>
        )}
      </main>
    </div>
  );
}
