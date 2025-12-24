# Sistema de Monitorización

## Descripción
Sistema básico de monitorización para detectar fallos y eventos anómanos en la aplicación.

## Componentes

### 1. Logging Estructurado
- Formato: JSON
- Campos: timestamp, level, message, campos adicionales
- Niveles: INFO, ERROR

### 2. Script de Monitorización (monitor.sh)
Script bash que analiza los logs y detecta:
- Errores (level: ERROR)
- Número de requests
- Anomalías en los logs

### 3. Health Check Endpoint
- Endpoint: `/health`
- Respuesta: "OK" (HTTP 200)
- Permite monitorización externa

## Uso del monitor

```bash
# Ejecutar el monitor
./monitoring/monitor.sh

# Ver alertas generadas
cat /tmp/alerts.log
