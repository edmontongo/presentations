func (lh locationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    //...
    weather, err := darkskySvc.Forecast(darksky.Location{ Lat: lat, Long: lng })
    if err != nil {
        http.Error(w, "failed to fetch weather details", http.StatusInternalServerError)
        return
    }
}