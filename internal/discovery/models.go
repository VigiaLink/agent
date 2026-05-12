package discovery

type Device struct {
    IP           string
    Manufacturer string
    Model        string
    RTSP         string
    ONVIF        bool
    Valid        bool
}