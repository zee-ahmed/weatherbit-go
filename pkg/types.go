package weatherbit

// BaseURL specifies Weatherbit 2.0 API
// ReqTimeout specifies how long to wait before assuming timeout
const (
	BaseURL    = "https://api.weatherbit.io/v2.0/"
	ReqTimeout = 10
)

// WbResponse holds entire response
type WbResponse struct {
	Data  []DataPoint `json:"data,omitempty"`
	Count int         `json:"count,omitempty"` // count of returned observations
}

// DataPoint holds nested metereological data from JSON object 'data'
// https://www.weatherbit.io/api/weather-current
// Default units for API response are metric (Celsius, m/s, mm)
type DataPoint struct {
	WeatherBlock             WeatherPoint `json:"weather,omitempty"`
	Latitude                 float64      `json:"lat,omitempty"` // degrees
	Longitude                float64      `json:"lon,omitempty"` // degrees
	CityName                 string       `json:"city_name,omitempty"`
	StateCode                string       `json:"state_code,omitempty"`
	CountryCode              string       `json:"country_code,omitempty"` // abbreviation
	PartOfDay                string       `json:"pod,omitempty"`          // "d" = day, "n" = night
	Temperature              float64      `json:"temp,omitempty"`
	TimeZone                 string       `json:"timezone,omitempty"` // local IANA Timezone
	Sunrise                  string       `json:"sunrise,omitempty"`  // HH:MM
	Sunset                   string       `json:"sunset,omitempty"`   // HH:MM
	LastObservationTime      string       `json:"ob_time,omitempty"`  // YYYY-MM-DD HH:MM
	LastObservationTimeStamp float64      `json:"ts,omitempty"`       // Unix timestamp - JSON marshalling errors if this is int64
	DateTime                 string       `json:"datetime,omitempty"` // current cycle hour (YYYY-MM-DD:HH)
	RelativeHumidity         float64      `json:"rh,omitempty"`       // percentage
	Pressure                 float64      `json:"pres,omitempty"`     // millibars
	WindDirection            float64      `json:"wind_dir,omitempty"` // degrees
	WindDirectionCardinal    string       `json:"wind_cdir,omitempty"`
	WindDirectionVerbal      string       `json:"wind_cdir_full,omitempty"`
	ApparentTemperature      float64      `json:"app_temp,omitempty"`
	SeaLevelPressure         float64      `json:"slp,omitempty"` // millibars
	Visibility               float64      `json:"vis,omitempty"` // KM
	DewPoint                 float64      `json:"dewpt,omitempty"`
	Precipitation            float64      `json:"precip,omitempty"`  // millimetres
	Station                  string       `json:"station,omitempty"` // source station ID
	WindSpeed                float64      `json:"wind_spd,omitempty"`
	// SolarElevationAngle is the altitude of the Sun (the angle between the horizon and the centre of the Sun's disc)
	SolarElevationAngle float64 `json:"elev_angle,omitempty"` // degrees
	// SolarHourAngle is positive during the morning, reduces to zero at solar noon and decreases further during the afternoon
	SolarHourAngle float64 `json:"h_angle,omitempty"` // degrees
	CloudCover     float64 `json:"clouds,omitempty"`  // percentage of sky covered by cloud
	UV             float64 `json:"uv,omitempty"`      // UV Index (0-11+)
	Dhi            float64 `json:"dhi,omitempty"`     // Diffuse Horizontal Irradiance

	// Additional fields for current (undocumented on Weatherbit.io API documentation?)
	Dni float64 `json:"dni,omitempty"` // Direct Normal Irradiance
	Ghi float64 `json:"ghi,omitempty"` // Global Horizontal Irradiance

	// Additional fields for history
	MaxWindSpd   float64 `json:"max_wind_spd,omitempty"`    // maximum 2 minute wind speed (m/s).
	MaxWindDir   float64 `json:"max_wind_dir,omitempty"`    // direction of maximum 2 minute wind gust (degrees)
	MaxWindSpdTs int64   `json:"max_wind_spd_ts,omitempty"` // time of maximum wind gust UTC (Unix Timestamp)
	WindGustSpd  float64 `json:"wind_gust_spd,omitempty"`
	MinTempTs    float64 `json:"min_temp_ts,omitempty"`
	MaxTempTs    float64 `json:"max_temp_ts,omitempty"`
	Snow         float64 `json:"snow,omitempty"`
	SnowDepth    float64 `json:"snow_depth,omitempty"`
	MaxDhi       float64 `json:"max_dhi,omitempty"`
	MaxTemp      float64 `json:"max_temp,omitempty"`
	MaxUv        float64 `json:"max_uv,omitempty"`
	TDhi         float64 `json:"t_dhi,omitempty"`
	MinTemp      float64 `json:"min_temp,omitempty"`

	// Additional fields for forecast
	TimestampLocal string  `json:"timestamp_local,omitempty"`
	TimestampUtc   string  `json:"timestamp_utc,omitempty"`
	Pop            float64 `json:"pop,omitempty"`
	CloudsLow      float64 `json:"clouds_low,omitempty"`
	CloudsMid      float64 `json:"clouds_mid,omitempty"`
	CloudsHi       float64 `json:"clouds_hi,omitempty"`
	Ozone          float64 `json:"ozone,omitempty"`
}

// WeatherPoint holds nested JSON object 'data.weather'
// See https://www.weatherbit.io/api/codes for details
type WeatherPoint struct {
	Icon string `json:"icon,omitempty"`
	// json field "Code" is sometimes understood to be float64, other times string. Generates Marshall error. Therefore declared as both below
	CodeString  string  `json:"code,omitempty"`
	CodeNumber  float64 `json:"code,omitempty"`
	Description string  `json:"description,omitempty"`
}

// Parameters - an exported struct which holds relevant information for each request
type Parameters struct {
	URL         string
	Apikey      string
	Temporality string
	Granularity string // "current", "forecast", "history"
	Lat, Lon    float64
	City        string
	State       string // "NC" or "North+Carolina"
	Country     string // "US"
	Cityid      float64
	Units       string // "M" (Metric), "S" (Scientific), "I" (Imperial) - Default "M"
	Marine      string // "f" (exclude offshore observations), "t" (include offshore observations) - Default "f"
	StartDate   string // YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
	EndDate     string // YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
}

// // parameters struct holds relevant information for each request
// type parameters struct {
// 	url         string
// 	apikey      string
// 	temporality string
// 	granularity string // "current", "forecast", "history"
// 	lat, lon    float64
// 	city        string
// 	state       string // "NC" or "North+Carolina"
// 	country     string // "US"
// 	cityid      float64
// 	units       string // "M" (Metric), "S" (Scientific), "I" (Imperial) - Default "M"
// 	marine      string // "f" (exclude offshore observations), "t" (include offshore observations) - Default "f"
// 	startDate   string // YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
// 	endDate     string // YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
// }
