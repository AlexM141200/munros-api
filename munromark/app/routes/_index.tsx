import type { MetaFunction, LoaderFunction } from "@remix-run/node";
import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { useState, useEffect } from "react";

// Dynamic import to avoid SSR issues
let MunroMap: any = null;

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
    { name: "description", content: "Explore all the Munros in Scotland with our interactive map" },
  ];
};

export const loader: LoaderFunction = async () => {
  try {
    // Fetch munros from our Go API
    const response = await fetch('http://localhost:8080/api/munros');
    
    if (!response.ok) {
      throw new Error('Failed to fetch munros');
    }
    
    const munros = await response.json();
    return json({ munros });
  } catch (error) {
    console.error('Error fetching munros:', error);
    return json({ munros: [] });
  }
};

export default function Index() {
  const { munros } = useLoaderData<{ munros: Munro[] }>();
  const [selectedMunro, setSelectedMunro] = useState<Munro | null>(null);
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
    // Dynamically import the map component
    import("~/components/MunroMap").then((module) => {
      MunroMap = module.default;
    });
  }, []);

  return (
    <div className="w-full h-screen bg-gray-50">
      <header className="bg-white shadow-sm border-b border-gray-200">
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

      <main className="h-[calc(100vh-80px)]">
        {isClient && MunroMap ? (
          <MunroMap
            munros={munros}
            onMunroClick={(munro: Munro) => setSelectedMunro(munro)}
          />
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
