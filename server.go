// Pacote principal da aplicação em Go
package main

// Importação das bibliotecas necessárias
import (
   "fmt"
   "net/http"

   "github.com/prometheus/client_golang/prometheus"
   "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Definição de um contador Prometheus para contabilizar as requisições ao manipulador "/ping"
var pingCounter = prometheus.NewCounter(
   prometheus.CounterOpts{
       Name: "ping_request_count",
       Help: "No of requests handled by Ping handler",
   },
)

// Função que responde a requisições HTTP com "pong" e incrementa o contador
func ping(w http.ResponseWriter, req *http.Request) {
   pingCounter.Inc()
   fmt.Fprintf(w, "pong")
}

// Função principal que registra o contador Prometheus, configura manipuladores para as rotas "/ping" e "/metrics",
// e inicia o servidor HTTP na porta 8090
func main() {
   // Registro do contador Prometheus
   prometheus.MustRegister(pingCounter)

   // Configuração do manipulador para a rota "/ping"
   http.HandleFunc("/ping", ping)

   // Configuração do manipulador para a rota "/metrics" para expor métricas Prometheus
   http.Handle("/metrics", promhttp.Handler())

   // Inicia o servidor HTTP na porta 8090
   http.ListenAndServe(":8090", nil)
}
