# Munros API & Interactive Map

A comprehensive web application for exploring Scottish Munros with an interactive map interface. This project provides both a REST API for Munro data and a modern web interface built with Go, templ, HTMX, and Leaflet.

## Features

- 🏔️ **Interactive Map**: Explore all 282+ Scottish Munros on a detailed map
- 🔍 **Search & Filter**: Real-time search by name or SMC section
- 📊 **Comprehensive Data**: Height, grid references, classifications, and external links
- 🎯 **Custom Markers**: Mountain peak icons with detailed popups
- 📱 **Responsive Design**: Works on desktop and mobile devices
- 🚀 **Fast Performance**: Server-side rendering with templ templates
- 🔗 **External Links**: Direct access to maps, photos, and hill bagging resources

## Tech Stack

- **Backend**: Go 1.23+ with `http.ServeMux`
- **Templates**: [templ](https://templ.guide/) for type-safe HTML templates
- **Frontend**: HTMX for dynamic interactions
- **Styling**: Tailwind CSS for modern UI
- **Maps**: Leaflet.js for interactive mapping
- **Data**: CSV-based munro database

## Project Structure

```
munros-api/
├── src/
│   ├── api/           # API server setup
│   ├── cmd/           # Application entry point
│   ├── csv/           # CSV data handling
│   ├── handlers/      # HTTP handlers
│   ├── model/         # Data models
│   ├── routes/        # Route definitions
│   └── templates/     # templ templates
├── data/              # CSV data files
├── frontend/          # Static frontend files (legacy)
├── bin/               # Compiled binaries
└── scripts/           # Development scripts
```

## Quick Start

### Prerequisites

- Go 1.23 or higher
- templ CLI tool

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd munros-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Install templ CLI**
   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

4. **Run the server**
   ```bash
   ./run-server.sh
   ```

The application will be available at `http://localhost:8080`

## Development

### Development Server with Auto-Reload

For development with automatic reloading on file changes:

```bash
./dev-server.sh
```

This will:
- Watch for changes in Go and templ files
- Automatically regenerate templates
- Rebuild and restart the server
- Provide better development experience

### Manual Build Process

If you prefer to build manually:

```bash
# Generate templ files
templ generate

# Build the application
go build -o bin/munros-api ./src/cmd/main.go

# Run the server
./bin/munros-api
```

## API Endpoints

### Munros Data

- `GET /api/munros` - Get all munros with optional filtering
- `GET /api/munros/{id}` - Get specific munro by ID
- `GET /api/munros/csv` - Get munros in CSV format (legacy)
- `GET /api/munros/all` - Alias for /api/munros

### Query Parameters

- `classification` - Filter by classification (munro, top, other)
- `min_height` - Filter by minimum height in meters
- `section` - Filter by SMC section
- `search` - Search by name

### Example API Calls

```bash
# Get all munros
curl http://localhost:8080/api/munros

# Get munros over 1000m
curl "http://localhost:8080/api/munros?min_height=1000"

# Search for Ben Nevis
curl "http://localhost:8080/api/munros?search=Ben%20Nevis"

# Get specific munro
curl http://localhost:8080/api/munros/1
```

## Web Interface

### Main Features

- **Interactive Map**: Pan and zoom around Scotland
- **Search Bar**: Real-time filtering of visible munros
- **Munro Popups**: Click any marker for detailed information
- **Hero Overlay**: Welcome screen with quick start guide
- **Responsive Design**: Adapts to different screen sizes

### Data Display

Each munro popup includes:
- Name and height (meters and feet)
- Classification (Munro, Top, Other)
- SMC Section
- Grid reference
- Comments (if available)
- External links to maps and resources

## Data Source

The application uses the official Munro data from the Database of British and Irish Hills (DoBIH), including:
- All 282 official Munros
- Subsidiary tops
- Height data in meters and feet
- Grid references and coordinates
- Classifications and comments
- External resource links

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Generate templates (`templ generate`)
5. Test your changes
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## Development Notes

### Template Development

- Templates are located in `src/templates/`
- Use `templ generate` to compile templates to Go code
- Templates are type-safe and compiled at build time
- Follow templ syntax for component composition

### Adding New Features

1. Create new templ components in `src/templates/`
2. Add corresponding route handlers in `src/routes/`
3. Update the main router in `src/handlers/`
4. Test the changes using the development server

### Database Migration

The application currently uses CSV data. To migrate to a database:
1. Implement the `DataService` interface in `src/routes/routes.go`
2. Create database models in `src/model/`
3. Update the service initialization in `init()`

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Data provided by the Database of British and Irish Hills (DoBIH)
- Scottish Mountaineering Club (SMC) for section classifications
- OpenStreetMap for map tiles
- Leaflet.js for mapping functionality
- templ for modern Go templating