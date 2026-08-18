package api
import("net/http";"github.com/kekelele996/offshore-rig-integrity-control-service/internal/domain")
func writeDecision(w http.ResponseWriter,d domain.ReleaseDecision){if !d.Allowed{http.Error(w,d.Reason,http.StatusConflict);return};w.WriteHeader(http.StatusNoContent)}
