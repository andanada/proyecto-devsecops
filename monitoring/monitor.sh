#!/bin/bash

# Script de monitorización básico para logs estructurados
# Simula detección de errores y envío de alertas

LOG_FILE="/tmp/app.log"
ALERT_FILE="/tmp/alerts.log"

echo "=== Monitor de logs iniciado ===" | tee -a $ALERT_FILE
echo "Timestamp: $(date)" | tee -a $ALERT_FILE

# Función para analizar logs
analyze_logs() {
    if [ -f "$LOG_FILE" ]; then
        echo "Analizando $LOG_FILE..." | tee -a $ALERT_FILE
        
        # Contar errores
        ERROR_COUNT=$(grep -c '"level":"ERROR"' "$LOG_FILE" 2>/dev/null || echo "0")
        
        if [ "$ERROR_COUNT" -gt 0 ]; then
            echo "ALERTA: Se detectaron $ERROR_COUNT errores" | tee -a $ALERT_FILE
            grep '"level":"ERROR"' "$LOG_FILE" | tee -a $ALERT_FILE
        else
            echo "INFO: No se detectaron errores" | tee -a $ALERT_FILE
        fi
        
        # Contar requests
        REQUEST_COUNT=$(grep -c '"message":"Request received"' "$LOG_FILE" 2>/dev/null || echo "0")
        echo "INFO: Total de requests: $REQUEST_COUNT" | tee -a $ALERT_FILE
        
    else
        echo "WARNING: Log file no encontrado: $LOG_FILE" | tee -a $ALERT_FILE
    fi
}

analyze_logs

echo "=== Monitor de logs finalizado ===" | tee -a $ALERT_FILE
