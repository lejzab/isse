package networkdevice

import "net/http"

func GetNetworkDevice(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("network device"))
}

func DeleteNetworkDevice(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
