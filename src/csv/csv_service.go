package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/AlexM141200/munros-api/src/model"
)

type CSVService struct {
	filePath string
}

func NewCSVService(filePath string) *CSVService {
	return &CSVService{
		filePath: filePath,
	}
}

// ConvertOSGridToLatLon converts Ordnance Survey grid coordinates to latitude and longitude
func ConvertOSGridToLatLon(easting, northing float64) (float64, float64) {
	// Constants for OSGB36 to WGS84 conversion
	a := 6377563.396        // Semi-major axis of Airy 1830
	b := 6356256.910        // Semi-minor axis of Airy 1830
	e2 := 0.00667054        // Eccentricity squared
	
	// False origin coordinates
	n0 := -100000.0         // Northing of false origin
	e0 := 400000.0          // Easting of false origin
	f0 := 0.9996012717      // Scale factor on central meridian
	lat0 := 0.855211333     // Latitude of false origin (49° in radians)
	lon0 := -0.034906585    // Longitude of false origin (-2° in radians)
	
	// Convert to radians and calculate
	n := (a - b) / (a + b)
	
	// Calculate latitude
	lat := lat0
	M := 0.0
	
	for i := 0; i < 10; i++ {
		M = b * f0 * (((1 + n + (5.0/4.0)*n*n + (5.0/4.0)*n*n*n) * (lat - lat0)) -
			((3*n + 3*n*n + (21.0/8.0)*n*n*n) * math.Sin(lat-lat0) * math.Cos(lat+lat0)) +
			(((15.0/8.0)*n*n + (15.0/8.0)*n*n*n) * math.Sin(2*(lat-lat0)) * math.Cos(2*(lat+lat0))) -
			((35.0/24.0)*n*n*n * math.Sin(3*(lat-lat0)) * math.Cos(3*(lat+lat0))))
		
		if math.Abs(northing-n0-M) < 0.01 {
			break
		}
		lat = lat + (northing-n0-M)/(a*f0)
	}
	
	// Calculate longitude
	nu := a * f0 / math.Sqrt(1-e2*math.Sin(lat)*math.Sin(lat))
	rho := a * f0 * (1 - e2) / math.Pow(1-e2*math.Sin(lat)*math.Sin(lat), 1.5)
	eta2 := nu/rho - 1
	
	VII := math.Tan(lat) / (2 * rho * nu)
	VIII := math.Tan(lat) / (24 * rho * math.Pow(nu, 3)) * (5 + 3*math.Pow(math.Tan(lat), 2) + eta2 - 9*math.Pow(math.Tan(lat), 2)*eta2)
	IX := math.Tan(lat) / (720 * rho * math.Pow(nu, 5)) * (61 + 90*math.Pow(math.Tan(lat), 2) + 45*math.Pow(math.Tan(lat), 4))
	
	X := 1.0 / (math.Cos(lat) * nu)
	XI := 1.0 / (math.Cos(lat) * 6 * math.Pow(nu, 3)) * (nu/rho + 2*math.Pow(math.Tan(lat), 2))
	XII := 1.0 / (math.Cos(lat) * 120 * math.Pow(nu, 5)) * (5 + 28*math.Pow(math.Tan(lat), 2) + 24*math.Pow(math.Tan(lat), 4))
	
	dE := easting - e0
	
	lat = lat - VII*dE*dE + VIII*math.Pow(dE, 4) - IX*math.Pow(dE, 6)
	lon := lon0 + X*dE - XI*math.Pow(dE, 3) + XII*math.Pow(dE, 5)
	
	// Convert to degrees
	lat = lat * 180 / math.Pi
	lon = lon * 180 / math.Pi
	
	return lat, lon
}

func (s *CSVService) ReadMunros() ([]model.Munro, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	// Read header
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV headers: %w", err)
	}

	var munros []model.Munro
	
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV record: %w", err)
		}

		// Parse the record
		munro, err := s.parseRecord(record, headers)
		if err != nil {
			continue // Skip malformed records
		}

		munros = append(munros, munro)
	}

	return munros, nil
}

func (s *CSVService) parseRecord(record []string, headers []string) (model.Munro, error) {
	munro := model.Munro{}

	// Helper function to safely get field value
	getField := func(fieldName string) string {
		for i, header := range headers {
			if strings.EqualFold(header, fieldName) && i < len(record) {
				return strings.TrimSpace(record[i])
			}
		}
		return ""
	}

	// Parse Running No
	if runningNo := getField("Running No"); runningNo != "" {
		if val, err := strconv.Atoi(runningNo); err == nil {
			munro.RunningNo = val
		}
	}

	// Parse DoBIH Number
	if dobihNum := getField("DoBIH Number"); dobihNum != "" {
		if val, err := strconv.Atoi(dobihNum); err == nil {
			munro.DoBIHNumber = val
		}
	}

	// Parse Name
	munro.Name = getField("Name")

	// Parse SMC Section
	munro.SMCSection = getField("SMC Section")

	// Parse RHB Section
	munro.RHBSection = getField("RHB Section")

	// Parse Height (m)
	if height := getField("Height (m)"); height != "" {
		if val, err := strconv.ParseFloat(height, 64); err == nil {
			munro.HeightM = val
		}
	}

	// Parse Height (ft) - handle quoted field names
	if height := getField("Height\n(ft)"); height != "" {
		if val, err := strconv.Atoi(height); err == nil {
			munro.HeightFt = val
		}
	}

	// Parse Map references
	munro.Map1to50k = getField("Map 1:50k")
	munro.Map1to25k = getField("Map 1:25k")
	munro.GridRef = getField("Grid Ref")
	munro.GridRefXY = getField("GridRefXY")

	// Parse coordinates
	if xCoord := getField("xcoord"); xCoord != "" {
		if val, err := strconv.ParseFloat(xCoord, 64); err == nil {
			munro.XCoord = val
		}
	}

	if yCoord := getField("ycoord"); yCoord != "" {
		if val, err := strconv.ParseFloat(yCoord, 64); err == nil {
			munro.YCoord = val
		}
	}

	// Convert OS Grid coordinates to Lat/Lon
	if munro.XCoord != 0 && munro.YCoord != 0 {
		lat, lon := ConvertOSGridToLatLon(munro.XCoord, munro.YCoord)
		munro.Latitude = lat
		munro.Longitude = lon
	}

	// Parse URLs
	munro.StreetmapURL = getField("Streetmap")
	munro.GeographURL = getField("Geograph")
	munro.HillBaggingURL = getField("Hill-bagging")

	// Parse Comments
	munro.Comments = getField("Comments")

	// Determine classification - check the 2021 column for MUN (Munro)
	if classification := getField("2021"); classification == "MUN" {
		munro.Classification = "Munro"
	} else if classification == "TOP" {
		munro.Classification = "Top"
	} else {
		munro.Classification = "Other"
	}

	return munro, nil
}
